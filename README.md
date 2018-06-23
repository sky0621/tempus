# tempus

## docker

https://docs.docker.com/install/linux/docker-ce/ubuntu/

$ sudo docker version

Client:
 Version:      18.03.1-ce
 API version:  1.37
 Go version:   go1.9.5
 Git commit:   9ee9f40
 Built:        Thu Apr 26 07:17:38 2018
 OS/Arch:      linux/amd64
 Experimental: false
 Orchestrator: swarm

Server:
 Engine:
  Version:      18.03.1-ce
  API version:  1.37 (minimum version 1.12)
  Go version:   go1.9.5
  Git commit:   9ee9f40
  Built:        Thu Apr 26 07:15:45 2018
  OS/Arch:      linux/amd64
  Experimental: false

## docker build

$ sudo docker build -t sky0621/tempus:v0.1 .

## docker run

$ sudo docker run sky0621/tempus:v0.1

## GCP Cloud SDK

https://cloud.google.com/sdk/docs/

https://cloud.google.com/sdk/docs/components

## gcloud and kubectl setup

$ gcloud init

$ gcloud auth list

$ gcloud info

$ gcloud components list

$ gcloud components update

$ gcloud components install kubectl

$ gcloud container clusters get-credentials {my cluster name} --zone=asia-northeast1-a

## to gcr

$ sudo docker tag sky0621/tempus:v0.1 gcr.io/$PROJECT_ID/tempus:v0.1

$ gcloud docker -- push gcr.io/$PROJECT_ID/tempus:v0.1

$ gsutil ls -R gs://artifacts.*

## deploy

$ kubectl create -f tempus-deployment.yaml 
deployment "tempus" created

$ kubectl delete deployment tempus
deployment "tempus" deleted
