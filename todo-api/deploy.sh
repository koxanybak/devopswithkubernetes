docker-compose build
docker-compose push
kubectl delete deployments.apps todo-api-dep
kubectl apply -f manifests/deployment.yaml