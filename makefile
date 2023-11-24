
############################################################	############################################################ 	############################################################ 

#K8s Helper commands
restart-pods:
	kubectl rollout restart deployment img-switch-api-depl
	kubectl rollout restart deployment img-switch-frontend-depl

deploy-pods:
	kubectl apply -f k8s/img-switch-api.yaml
	kubectl apply -f k8s/img-switch-frontend.yaml

dev-deploy-pods:
	kubectl apply -f k8s/dev/img-switch-api.yaml
	kubectl apply -f k8s/dev/img-switch-frontend.yaml
