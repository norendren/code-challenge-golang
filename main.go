package main

import (
	"archive/zip"
	"cloud.google.com/go/storage"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Asset struct {
	Filename string
	Url      string
	GcsUrl   string
}

var client = &http.Client{}

func main() {
	http.HandleFunc("/download", func(w http.ResponseWriter, req *http.Request) {
		assetManifest := make([]Asset, 0)
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		json.Unmarshal(body, &assetManifest)

		w.Header()["Content-Type"] = []string{"application/zip"}
		w.Header()["Content-Disposition"] = []string{"attachment; filename=\"brandfolder_assets.zip\""}
		w.WriteHeader(200) // Status Code pushes headers so we can stream data

		writer := zip.NewWriter(w)
		for _, asset := range assetManifest {
			func() (err error) {
				zipWriter, err := writer.CreateHeader(
					&zip.FileHeader{
						Name: asset.Filename,
					},
				)

				var fileReader io.Reader
				if asset.GcsUrl != "" {
					url, err := url.Parse(asset.GcsUrl)
					client, err := storage.NewClient(context.Background())
					fileReader, err = client.Bucket(url.Host).Object(url.Path).NewReader(context.Background())
					if err != nil {
						return err
					}
				} else {
					req, err := http.NewRequest("GET", asset.Url, nil)
					assetResp, err := client.Do(req)
					if err != nil {
						return err
					}
					defer assetResp.Body.Close()
					fileReader = assetResp.Body
				}

				io.Copy(zipWriter, fileReader)
				return
			}()
		}
		writer.Close()
	})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil))
}
