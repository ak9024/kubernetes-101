# kubernetes-101

### Prerequisite
- kubernetes tools https://kubernetes.io/docs/tasks/tools/
- docker https://docs.docker.com/engine/install/

### Getting Started

```
cd golang-api
docker build -t golang-api:v1 -f Dockerfile .
cd ..
kubectl apply -f deploy/golang-api.yaml
```

### References
- https://kubernetes.io/docs/concepts/workloads/controllers/deployment/
- https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/
- https://kubernetes.io/docs/concepts/configuration/secret/
- https://kubernetes.io/docs/tasks/inject-data-application/define-environment-variable-container/
- https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/
