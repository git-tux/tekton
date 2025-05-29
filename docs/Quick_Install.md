# Steps to quickly setup a ci/cd pipeline in kubernetes for a go app

Please make sure that you have kubectl installed and configured with a test 
kubernetes cluster. All the following commands have been tested locally using
[KinD](https://kind.sigs.k8s.io/)

First, please download the Github Personal Access Token that I have already 
shared with you. This is a test github account and the shared pat has limited
access and will be expired after a few days

```
export GITHUB_TOKEN=<Github Token>
```


## Create Kubernetes namespace for apps and create a deployment
```
kubectl create namespace apps
kubectl -n apps apply -f K8s_manifests/Deployment.yaml
```

## Create cicd namespace and add registry secrets and serviceaccount

```
kubectl create namespace cicd
cat  K8s_manifests/Buildah_docker_push_sa.yaml | sed s/GITHUB_TOKEN/$GITHUB_TOKEN/ |
kubectl apply -f -

```

## Install tekton and tekton tasks for git-clone and buildah
```
kubectl apply -f https://storage.googleapis.com/tekton-releases/pipeline/latest/release.yaml

kubectl -n cicd apply -f https://raw.githubusercontent.com/tektoncd/catalog/refs/heads/main/task/git-clone/0.9/git-clone.yaml

kubectl -n cicd apply -f https://api.hub.tekton.dev/v1/resource/tekton/task/buildah/0.9/raw
```

## Create a pipeline for building and deploying a go app
```
kubectl -n cicd apply -f pipelines/Build_and_Push.yaml
```
## Allow build-bot serviceaccount to deploy on Kubernetes

```
kubectl apply -f K8s_manifests/rbac.yaml
```

## Create ta task for k8s deployment which will be called from our pipeline
```
kubectl -n cicd apply -f tasks/restart_deployment.yaml
```

## Run the pipeline
```
kubectl -n cicd apply -f pipeline_runs/build_and_push_go_server.yml
```


