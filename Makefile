proto:
	protoc pkg/**/pb/*.proto --go-grpc_out=. --go_out=.

server:
	go run ./cmd/main.go

kubernetes-init-databases:
	kubectl apply -f ./kubernetes/postgres-jobs.yaml

kubernetes-start:
	kubectl apply -f ./kubernetes/postgres-config-map.yaml
	kubectl apply -f ./kubernetes/postgres-db-pv.yaml
	kubectl apply -f ./kubernetes/postgres-db-pvc.yaml
	kubectl apply -f ./kubernetes/postgres-db-deployment.yaml
	kubectl apply -f ./kubernetes/postgres-db-service.yaml
	kubectl apply -f ./kubernetes/xm-gateway-deployment.yaml
	kubectl apply -f ./kubernetes/xm-gateway-srv.yaml
	kubectl apply -f ./kubernetes/xm-auth-deployment.yaml
	kubectl apply -f ./kubernetes/xm-auth-srv.yaml
	kubectl apply -f ./kubernetes/xm-company-deployment.yaml
	kubectl apply -f ./kubernetes/xm-company-srv.yaml

kubernetes-stop:
	kubectl delete -f ./kubernetes/postgres-config-map.yaml
	# kubectl delete -f ./kubernetes/postgres-db-pv.yaml
	# kubectl delete -f ./kubernetes/postgres-db-pvc.yaml
	kubectl delete -f ./kubernetes/postgres-db-deployment.yaml
	kubectl delete -f ./kubernetes/postgres-db-service.yaml
	kubectl delete -f ./kubernetes/xm-gateway-deployment.yaml
	kubectl delete -f ./kubernetes/xm-gateway-srv.yaml
	kubectl delete -f ./kubernetes/xm-auth-deployment.yaml
	kubectl delete -f ./kubernetes/xm-auth-srv.yaml
	kubectl delete -f ./kubernetes/xm-company-deployment.yaml
	kubectl delete -f ./kubernetes/xm-company-srv.yaml