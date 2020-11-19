cd $1
docker-compose build
docker-compose push
kubectl apply -f manifests/deployment.yaml
kubectl delete deployments.apps $1-dep
kubectl apply -f manifests/deployment.yaml