# kubernetes-101

### Prerequisite
- kubernetes tools https://kubernetes.io/docs/tasks/tools/
- docker https://docs.docker.com/engine/install/

### Overview

[![](https://mermaid.ink/img/pako:eNqNksluwyAQhl8FkUsrmTRbq4hUOeVSqZWq5mj3gGHsoGCwAHdRkncvDkSJb70wv2b5GIY5YG4EYIpry9odev1YFRohriRof5dH-3lPxuhF1xacIw3TrAaBnku7RsowgUqmmOZg0ZisZczKU_bnmZachKyPD0cH9ktymObbKFBtQnlNWCtJitHlJBa6roxtcdU5D_YGFuOJFcitEdP83YgbXMBMEgi0iM9izm2gQq1iUqNKKkVHQojMeWv2QEdVVSVNvqXwO7pofzJulLF0FGCrAWS_dAkxnz1xePwXJcSGlPS0RLqW0lFZlkPM7IqJN15Jl7Fkl5FcxCzrJ9MfZzXvj0Xf-k1t_OU4lIE7dpbsCme4AdswKcK6HPq8AvsdNFBgGqRgdl_gQp9CHuu82f5qjqm3HWS4awXzsJEsfGeDacWUC14Q0hv7FvfvvIanPx6T4qc?type=png)](https://mermaid.live/edit#pako:eNqNksluwyAQhl8FkUsrmTRbq4hUOeVSqZWq5mj3gGHsoGCwAHdRkncvDkSJb70wv2b5GIY5YG4EYIpry9odev1YFRohriRof5dH-3lPxuhF1xacIw3TrAaBnku7RsowgUqmmOZg0ZisZczKU_bnmZachKyPD0cH9ktymObbKFBtQnlNWCtJitHlJBa6roxtcdU5D_YGFuOJFcitEdP83YgbXMBMEgi0iM9izm2gQq1iUqNKKkVHQojMeWv2QEdVVSVNvqXwO7pofzJulLF0FGCrAWS_dAkxnz1xePwXJcSGlPS0RLqW0lFZlkPM7IqJN15Jl7Fkl5FcxCzrJ9MfZzXvj0Xf-k1t_OU4lIE7dpbsCme4AdswKcK6HPq8AvsdNFBgGqRgdl_gQp9CHuu82f5qjqm3HWS4awXzsJEsfGeDacWUC14Q0hv7FvfvvIanPx6T4qc)

### Getting Started

```bash
# create namespace
kubectl create namespace kubernetes-101
kubectl apply -f deploy
```

```bash
# enable metrics-server for CPU/Memory based for HPA
# https://github.com/kubernetes-sigs/metrics-server
kubectl apply -f kube-system/compoonents.yaml
```

```bash
# https://github.com/kubernetes/ingress-nginx
# install ingress-nginx with helm
helm upgrade --install ingress-nginx ingress-nginx \
  --repo https://kubernetes.github.io/ingress-nginx \
  --namespace ingress-nginx --create-namespace

# install ingress-nginx without helm
kubectl create namespace ingress-nginx
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.8.2/deploy/static/provider/aws/deploy.yaml
```

```bash
# vi /etc/hosts
# 127.0.0.1 adiatma.local
curl http://adiatma.local
```

### References
- https://kubernetes.io/docs/concepts/workloads/controllers/deployment/
- https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/
- https://kubernetes.io/docs/concepts/configuration/secret/
- https://kubernetes.io/docs/tasks/inject-data-application/define-environment-variable-container/
- https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/
- https://kubernetes.github.io/ingress-nginx/deploy/#quick-start
