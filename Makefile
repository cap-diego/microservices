swagger:
	swagger generate spec -o ./swagger.yaml --scan-models

api: go_client
	swagger generate client -f ../swagger.yaml -A microservices
