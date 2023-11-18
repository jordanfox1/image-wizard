
############################################################	############################################################ 	############################################################ 

#K8s Helper commands
restart-pods:
	kubectl rollout restart deployment image-wizard-api-depl
	kubectl rollout restart deployment image-wizard-frontend-depl

deploy-pods:
	kubectl apply -f k8s/image-wizard-api.yaml
	kubectl apply -f k8s/image-wizard-frontend.yaml
