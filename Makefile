proto:
	protoc pkg/**/pb/*.proto --go-grpc_out=. --go_out=.

server:
	go run cmd/main.go

kubernetes-start:
	kubectl create -f ./build/postgres-secret.yaml
	kubectl apply -f ./build/postgres-db-pv.yaml
	kubectl apply -f ./build/postgres-db-pvc.yaml
	kubectl apply -f ./build/postgres-db-deployment.yaml
	kubectl apply -f ./build/postgres-db-service.yaml
	kubectl apply -f ./build/xm-gateway-deployment.yaml
	kubectl apply -f ./build/xm-gateway-srv.yaml
	kubectl apply -f ./build/xm-auth-deployment.yaml
	kubectl apply -f ./build/xm-auth-srv.yaml
	kubectl apply -f ./build/xm-company-deployment.yaml
	kubectl apply -f ./build/xm-company-srv.yaml

kubernetes-stop:
	kubectl delete -f ./build/postgres-secret.yaml
	# kubectl delete -f ./build/postgres-db-pv.yaml
	# kubectl delete -f ./build/postgres-db-pvc.yaml
	kubectl delete -f ./build/postgres-db-deployment.yaml
	kubectl delete -f ./build/postgres-db-service.yaml
	kubectl delete -f ./build/xm-gateway-deployment.yaml
	kubectl delete -f ./build/xm-gateway-srv.yaml
	kubectl delete -f ./build/xm-auth-deployment.yaml
	kubectl delete -f ./build/xm-auth-srv.yaml
	kubectl delete -f ./build/xm-company-deployment.yaml
	kubectl delete -f ./build/xm-company-srv.yaml