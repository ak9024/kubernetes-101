# kubernetes-101

### Prerequisite
- kubernetes tools https://kubernetes.io/docs/tasks/tools/
- docker https://docs.docker.com/engine/install/

### Getting Started

```bash
cd golang-api
docker build -t golang-api:v1 -f Dockerfile .
cd ..
# create namespace
kubectl create namespace kubernetes-101
kubectl apply -f deploy/golang-api.yaml
```

```bash
# enable metrics api
kubectl apply -f deploy/compoonents.yaml
```

```bash
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
curl http://adiatma.local/api
```

### References
- https://kubernetes.io/docs/concepts/workloads/controllers/deployment/
- https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/
- https://kubernetes.io/docs/concepts/configuration/secret/
- https://kubernetes.io/docs/tasks/inject-data-application/define-environment-variable-container/
- https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/
- https://kubernetes.github.io/ingress-nginx/deploy/#quick-start
