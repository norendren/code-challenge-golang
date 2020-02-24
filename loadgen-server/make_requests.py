import time
import requests
import logging

logger = logging.getLogger('make_requests')


request_datas = [
    [
        {
            "filename": "brandfolder-icon-favicon (2)-1299x1222-59ebf79.png",
            "url": "http://localhost:8090/?statusCode=200"
        },
    ],
    [
        {
            "filename": "badfile",
            "url": "https://s3.amazonaws.com/nont-existant-bucket/file; \n'.nothere"
        },
    ],
    [
        {
            "filename": "50x",
            "url": "http://dependent-server:8090?duration=3s&statusCode=503"
        },
    ],
    [
        {
            "filename": "bflogo.png",
            "url": "https://storage.googleapis.com/bf-boulder-staging/pvf0r1-7tf5g8-3urlmq/v/5495330/original/loader.gif",
            "gcsUrl": "gs://bf-boulder-staging/pvf0r1-7tf5g8-3urlmq/v/5495330/original/loader.gif"
        },
    ],
    [
        {
            "filename": "img.png",
            "url": "http://localhost:8090/?statusCode=200"
        }, {
            "filename": "img.png",
            "url": "http://localhost:8090/?statusCode=200"
        }, {
            "filename": "img.png",
            "url": "http://localhost:8090/?statusCode=200"
        }, {
            "filename": "img.png",
            "url": "http://localhost:8090/?statusCode=200"
        }, {
            "filename": "img.png",
            "url": "http://localhost:8090/?statusCode=200"
        },
    ]
]

time.sleep(20)
while True:
    for request_data in request_datas:
        try:
            response = requests.post('http://nginx:80/download', json=request_data)
            #logger.info(request_data[0]['filename'] + " " + str(len(request_data)) + " " + str(response.elapsed.total_seconds()) + " seconds " + str(len(response.content)) + " bytes")
        except KeyboardInterrupt as exc:
            raise exc
        except:
            pass
