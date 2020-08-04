# Brandfolder Code-Challenge-Golang

### Requirements
- Docker Compose

### How to get started
Given the need to pull some docker containers, it'd be beneficial to start that process:
```
docker-compose up
```


This repository represents an application that accepts a list of files via a 
```
POST /download 
[
  {
    "filename": "fname.png",
    "url": "https://url.com/to/file.png",
    "gcsUrl": "gs://bucketname/path/to/file.png"  # Optional
  }.
  ...
]
```

The repo in it's entirely contains a few components (loadgen-server, main.go, dependency-server). Here is a Flow Diagram that depicts the lifecycle of the repo.

![Flow Diagram](https://www.websequencediagrams.com/cgi-bin/cdraw?lz=dGl0bGUgY29kZS1jaGFsbGVuZ2UtZ29sYW5nIGFwcGxpY2F0aW9uIGZsb3cgZGlhZ3JhbQoKcGFydGljaXBhbnQgbG9hZGdlbi1zZXJ2ZXIADg1uZ2lueAAgDQBUFQATD250ZW50CgoASg4tPisAgRAVOiByZXF1ZXN0IGxpc3Qgb2YgZmlsZXMKbG9vcAAGBQogAIFIFgBMBQBmBTogZmV0Y2gAJAoAfAUtLT4tAFsXc3RyZWFtAFgGZW5kCgBKFi0-LQCCDQ4AMAl6aXA&s=default)


The responsibility of the application under question is contained within the `main.go` application. It's responsibility is to decode the request (list of files), compile them into a zip, and stream them to the client.

This service has a mixture of end-users, and it's importance and footprint has grown. We're starting to hear that there are errors in the functionality delivered to end-users, and would like help debugging and correcting some of the errors.

In the midst of this work, please make an effort to make logical improvements to the service and up-level our practices. Do your best to describe your intentions and coach us as team members on how we should think about making changes.

### Goal of this exercise
1. Help us identify the effectiveness/failures of this service
   1. Suggestion of how this could be quantified/monitored
1. Make changes to reduce the failure rate
1. Suggest and possibly implement structural changes to the code for better maintainability
1. Discuss how we could operate this service more effectively in production
   1. What specific things could we do such that we aren't caught off guard by a client bug report

Open main.go for assessing code footpring and editing

### Commands to help
```
docker-compose logs -f nginx
docker-compose logs -f code-challenge-golang
```
