.PHONY: run

backend: $(wildcard backend/*)
	go build -C backend -o ../build/

frontend: $(wildcard frontend/*)
	cd frontend/; npm run build

build: backend frontend

run: build
	@cd build; ./myservices
