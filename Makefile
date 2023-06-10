.PROXY: start
start: 
	FUNCTION_TARGET=SendGrindFunction go run cmd/main.go

.PHONY: deploy
deploy:
	gcloud functions deploy SendGrindFunction \
	--runtime go120 \
	--trigger-http

.PHONY: call
call:
	curl localhost:8080 -v