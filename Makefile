test: docker_build_test
	docker-compose up -d
	docker-compose exec -T http go test ./...
	docker-compsoe down

unit-test:
	go test `go list ./... | grep -v internal/test/unit`

e2e-test:
	go test `go list ./... | grep -v internal/test/e2e`

docker_build_test:
	docker build . -t service_test --target=test

docker_run:
	docker run --publish 8080:8080 service