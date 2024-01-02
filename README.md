# kubernetes-101

### Prerequisite
- minikube https://minikube.sigs.k8s.io/docs/start/
- kubernetes tools https://kubernetes.io/docs/tasks/tools/
- docker https://docs.docker.com/engine/install/

### Getting Started

```bash
# create namespace
kubectl create namespace kubernetes-101
kubectl apply -f deploy
# option if want to expose service to port 3000
kubectl port-forward svc/golang-api-service 3000:80
```

[![](https://mermaid.ink/img/pako:eNqNksluwyAQhl8F4Usr2amztIpIlVMulVqpao52DxjGDgoGC3AXJXn34kCU-NYL82uWj2GYA2aaAya4MbTbodePVakQYlKAcndFsJ_32QS9qMaAtVlLFW2Ao-fKrJHUlKOKSqoYGDTJ1iJkFTH780yLzixbHx-OFsyXYDAttkGgRvvyJqOdyGKMLPNQaPsqtMVkbx2YG1iIR5Ynd5pPi3fNQyEoHp5Brd1AjTpJhUK1kJIknPPUOqP3QJK6rqPOvgV3O7LoflKmpTYkyfN8NYLslzYi5rMnBo__ovjYmBKfEknXUpJUVTXGzK6YcOOVdBlDehnBRczSYRLDcVbz4VgMrd_Uhl8NQxm5Q2fRrnCKWzAtFdyvx2HIK7HbQQslJl5yavYlLtXJ59He6e2vYpg400OK-45TBxtB_fe1mNRUWu8FLpw2b2Hfzmt3-gPxyt2m?type=png)](https://mermaid.live/edit#pako:eNqNksluwyAQhl8F4Usr2amztIpIlVMulVqpao52DxjGDgoGC3AXJXn34kCU-NYL82uWj2GYA2aaAya4MbTbodePVakQYlKAcndFsJ_32QS9qMaAtVlLFW2Ao-fKrJHUlKOKSqoYGDTJ1iJkFTH780yLzixbHx-OFsyXYDAttkGgRvvyJqOdyGKMLPNQaPsqtMVkbx2YG1iIR5Ynd5pPi3fNQyEoHp5Brd1AjTpJhUK1kJIknPPUOqP3QJK6rqPOvgV3O7LoflKmpTYkyfN8NYLslzYi5rMnBo__ovjYmBKfEknXUpJUVTXGzK6YcOOVdBlDehnBRczSYRLDcVbz4VgMrd_Uhl8NQxm5Q2fRrnCKWzAtFdyvx2HIK7HbQQslJl5yavYlLtXJ59He6e2vYpg400OK-45TBxtB_fe1mNRUWu8FLpw2b2Hfzmt3-gPxyt2m)

```bash
# option if using ingress-nginx
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
minikube tunnel
curl http://adiatma.local
```

### References
- https://kubernetes.io/docs
