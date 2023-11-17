
############################################################	############################################################ 	############################################################ 

#K8s Helper commands
restart-pods:
	kubectl rollout restart deployment image-wizard-api-depl
	kubectl rollout restart deployment image-wizard-frontend-depl

deploy-pods:
	kubectl apply -f api/image-wizard-api/k8s/deployment.yaml
	kubectl apply -f frontend/image-wizard-frontend/k8s/deployment.yaml
