```bash=
docker buildx create --name grpc-builder --driver=docker-container

sudo docker buildx build \
    --tag nu1lspaxe/grpc-go:server \
    --file server/Dockerfile \
    --platform linux/x86_64 \
    --builder grpc-builder \
    --load .

sudo docker buildx build \
    --tag nu1lspaxe/grpc-go:client \
    --file client/Dockerfile \
    --build-arg SERVER_ADDR="dns:///todo-server.default.svc.cluster.local:50051" \
    --platform linux/x86_64 \
    --builder grpc-builder
    --load .

kind create cluster --config k8s/kind.yaml
# kind delete cluster

kubectl apply -f k8s/server.yaml
# kubectl get pods

kubectl apply -f k8s/client.yaml
```