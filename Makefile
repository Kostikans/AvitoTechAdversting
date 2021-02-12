.PHONY:
gen:
	 GO111MODULE=on  swagger generate spec -o ./api/swagger/swagger.yaml --scan-models
run:
	go run cmd/main.go

start:
	sudo APP_VERSION=latest docker-compose up

upload:
	sudo docker build -t kostikan/avito_advertising:latest -f ./Dockerfile .
	sudo docker push kostikan/avito_advertising:latest
	sudo APP_VERSION=latest docker-compose up

pull:
	sudo docker pull kostikan/avito_advertising:latest

tests:
	go test -coverprofile=coverage1.out -coverpkg=./... -cover ./... && cat coverage1.out | grep -v  easyjson | grep -v mocks | grep -v main  > cover.out &&go tool cover -func=cover.out
