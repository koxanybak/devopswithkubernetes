cd $1
docker-compose build
docker-compose push
kubectl delete deployments.apps $1-dep
kubectl apply -f manifests/deployment.yaml