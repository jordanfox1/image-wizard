
############################################################	############################################################ 	############################################################ 

#K8s Helper commands
restart-pods:
	kubectl rollout restart deployment img-switch-api-depl
	kubectl rollout restart deployment img-switch-frontend-depl

deploy-pods:
	kubectl apply -f k8s/img-switch-api.yaml
	kubectl apply -f k8s/img-switch-frontend.yaml

dev-api:
	cd api/img-switch-api && go run main.go

dev-fe:
	cd frontend/img-switch-frontend && pnpm run dev