package main

import (
	"archive/zip"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sony/sonyflake"
)

type Asset struct {
	Filename string
	Url      string
	GcsUrl   string
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// Opportunity for dynamic log levels here
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	s, err := NewServer()
	if err != nil {
		log.Fatal().Err(err).Msg("unable to start http server")
	}

	http.HandleFunc("/download", s.Download)
	log.Info().Str("port", os.Getenv("PORT")).Msg("Starting server.")
	log.Fatal().Msg(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil).Error())
}

type Server struct {
	IDGen *sonyflake.Sonyflake
}

func NewServer() (*Server, error) {
	return &Server{
		IDGen: sonyflake.NewSonyflake(sonyflake.Settings{
			StartTime: time.Now(),
		}),
	}, nil
}

func (s *Server) Download(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
	// Set up transaction ID for each incoming request
	transactionID, err := s.IDGen.NextID()
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	subLog := log.With().Uint64("transactionID", transactionID).Logger()

	assetManifest := make([]Asset, 0)
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		subLog.Error().Err(err).Msg("unable to read request body")
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &assetManifest)
	if err != nil {
		subLog.Error().Err(err).Msg("unable to unmarshal asset manifest")
		http.Error(w, "invalid manifest format", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", `attachment; filename="brandfolder_assets.zip"`)

	writer := zip.NewWriter(w)
	defer writer.Close()
	for _, asset := range assetManifest {
		if err := s.zipAsset(writer, asset); err != nil {
			subLog.Error().Err(err).Msg("failed to zip asset")
			return
		}
	}
	subLog.Debug().Msg("Asset Manifest processed successfully")
}

func (s *Server) zipAsset(writer *zip.Writer, asset Asset) (err error) {
	zipWriter, err := writer.CreateHeader(
		&zip.FileHeader{
			Name: asset.Filename,
		},
	)
	if err != nil {
		return err
	}
	var fileReader io.Reader
	if asset.GcsUrl != "" {
		url, err := url.Parse(asset.GcsUrl)
		if err != nil {
			return err
		}
		c, err := storage.NewClient(context.Background())
		if err != nil {
			return err
		}
		fileReader, err = c.Bucket(url.Host).Object(url.Path).NewReader(context.Background())
		if err != nil {
			return err
		}
	} else {
		assetResp, err := http.Get(asset.Url)
		if err != nil {
			return err
		} else if assetResp.StatusCode != 200 {
			return errors.New(fmt.Sprintf("unable to retrieve asset at given url=%v", asset.Url))
		}
		defer assetResp.Body.Close()
		fileReader = assetResp.Body
	}

	_, err = io.Copy(zipWriter, fileReader)
	if err != nil {
		return err
	}
	return
}
