cd client/ && npm run build && cd ..
docker-compose build
docker-compose push
kubectl delete deployments.apps todo-dep
kubectl apply -f manifests/deployment.yaml