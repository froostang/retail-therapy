tidy:
	@find . -name "go.mod" -execdir go mod tidy \;

start-api:
	cd api && make build && ./build/api-service