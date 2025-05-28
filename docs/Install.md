# Tekton CI/CD Pipeline

WARNING!!! : This file is outdated. Please read the [Quick Install](docs/Quick_Install.md)
In this document I'm going to create step by step a sample Tekton pipelines which builds a go based web server Docker image
and uploads it to Github Packages. For the build step, I will use the kaniko tool which builds docker images inside of pod containers without the need of docker.

## Installation

I will start by installing the latest releases of tekton pipeline and tekton triggers.

Install tekton pipeline

```bash
kubectl apply -f https://storage.googleapis.com/tekton-releases/pipeline/latest/release.yaml
```

Install tekton triggers

```bash
kubectl apply --filename https://storage.googleapis.com/tekton-releases/triggers/latest/release.yaml
```

```bash
kubectl apply -f https://storage.googleapis.com/tekton-releases/triggers/latest/interceptors.yaml

```

After the installation, I'm checking the deployed pods in the tekton-pipelines namespace:

```bash
kubectl get pods -n tekton-pipelines

NAME                                                READY   STATUS    RESTARTS      AGE
tekton-events-controller-58b9f95676-hc2d4           1/1     Running   1 (25h ago)   3d5h
tekton-pipelines-controller-7d7fc8479b-pqgbp        1/1     Running   1 (25h ago)   3d5h
tekton-pipelines-webhook-5984b98f88-59rv6           1/1     Running   5 (25h ago)   3d5h
tekton-triggers-controller-784d896c7f-79xmj         1/1     Running   0             6m53s
tekton-triggers-core-interceptors-54d9f764b-z2jx9   1/1     Running   0             45s
tekton-triggers-webhook-666c844478-nb6ck            1/1     Running   0             6m53s
```


## Configuration

First I'm going to install some useful tasks:

- [git-clone task](https://raw.githubusercontent.com/tektoncd/catalog/refs/heads/main/task/git-clone/0.9/git-clone.yaml)
- [buildah](https://api.hub.tekton.dev/v1/resource/tekton/task/buildah/0.9/raw)

Git-clone task will be used for cloning git repositories while buildah will be used for building docker images. 

```bash
kubectl apply -f https://raw.githubusercontent.com/tektoncd/catalog/refs/heads/main/task/git-clone/0.9/git-clone.yaml

task.tekton.dev/git-clone created

kubectl apply -f https://api.hub.tekton.dev/v1/resource/tekton/task/buildah/0.9/raw

task.tekton.dev/buildah created
```

Verifying that the tasks are created:

```bash
kubectl get tasks
NAME        AGE
git-clone   4d22h
buildah     4d16h
```


Next, I will create a pipeline that downloads a git repo using the following arguments:

 **repo-url**: The URL of the repo that will be downloaded
 **branch**: The branch that will be downloaded ( default main )

```
kubectl apply -f Clone_git_repo_pipeline.yaml
```

A pipeline named 'clone-repo' has been created:

```bash
kubectl get pipelines
NAME         AGE
clone-repo   7m13s
```

Since we need to clone and push on ghrc.io registry, I will create a secret in kubernetes which will include the authentication token for ghcr.io. Then I will create a serviceaccount which will be used from buildah tasks and I will assign the secret in this account

```yaml
---
# K8s_manifests/Buildah_docker_push.yaml

apiVersion: v1
kind: Secret
type: kubernetes.io/basic-auth
metadata:
  name: ghcr-credentials
  annotations:
    tekton.dev/docker-0: https://ghcr.io
type: kubernetes.io/basic-auth
stringData:
  username: git-tux
  password: <github token>

---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: build-bot
secrets:
  - name: ghcr-credentials
```

Now we are ready to create a pipeline that clones a repo, builds docker image using buildah and uploads the image to my ghcr.io account.

```yaml


```

Install the Pipeline and run it by creating a PipelineRun:

```
kubectl apply -f Build_and_push_Task.yaml
pipeline.tekton.dev/pipeline-git-buildah created

kubectl create -f pipeline_runs/build_and_push_go_server.yml
pipelinerun.tekton.dev/goserver-pipelinerun-8ff98 created

