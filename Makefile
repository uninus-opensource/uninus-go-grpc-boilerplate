MODULES      = $(shell cd extensions && ls -d */)
TMP_DIR     := $(shell mktemp -d)
UNAME       := $(shell uname)

# Default postgres migration settings
export POSTGRES_USER ?= root
export POSTGRES_PASS ?= rootpw
export POSTGRES_HOST ?= localhost
export POSTGRES_PORT ?= 15432
export POSTGRES_DATABASE ?= uninusdb
export POSTGRES_QUERYSTRING ?= sslmode=disable
export POSTGRES_MIGRATEUP ?= up
export POSTGRES_MIGRATEDOWN ?= down
export VERSION ?= 20230216134212

dep:
	GO111MODULE=on go mod download
	GO111MODULE=on go mod verify
	GO111MODULE=on go mod tidy

prepare:
	GO111MODULE=on go install github.com/golang/protobuf/protoc-gen-go@v1.3.2
	GO111MODULE=on go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v1.14.0
	GO111MODULE=on go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@v1.14.0
	GO111MODULE=on go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

run:
	go build -o userservice main.go
	./userservice -config=config/dev.conf

run-local:
	go build -o userservice main.go
	./userservice -config=config/local.conf

run-stage:
	go build -o chatservice main.go
	./chatservice -config=etc/config/stage.conf

docker-up:
	@docker-compose -f dev/docker-compose.yml up -d

docker-down:
	@docker-compose -f dev/docker-compose.yml down

bin:
	@mkdir -p bin

tool-migrate: bin
ifeq ($(UNAME), Linux)
	@curl -sSfL https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz | tar zxf - --directory /tmp \
	&& cp /tmp/migrate bin/
else ifeq ($(UNAME), Darwin)
	@curl -sSfL https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.darwin-amd64.tar.gz | tar zxf - --directory /tmp \
	&& cp /tmp/migrate bin/
else
	@echo "Your OS is not supported."
endif

migrate-up:
	@$(foreach module, $(MODULES), cp extensions/migrate/*.sql $(TMP_DIR);)
	@bin/migrate -source file://$(TMP_DIR) -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASS)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?$(POSTGRES_QUERYSTRING)" $(POSTGRES_MIGRATEUP)

migrate-down:
	@$(foreach module, $(MODULES), cp extensions/migrate/*.sql $(TMP_DIR);)
	@bin/migrate -source file://$(TMP_DIR) -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASS)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?$(POSTGRES_QUERYSTRING)" $(POSTGRES_MIGRATEDOWN)

migrate-force:
	@$(foreach module, $(MODULES), cp extensions/migrate/*.sql $(TMP_DIR);)
	@bin/migrate -source file://$(TMP_DIR) -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASS)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?$(POSTGRES_QUERYSTRING)" force $(VERSION)

test-coverage:
	go install gotest.tools/gotestsum@latest
	go env -w GOFLAGS=-mod=mod
	./coverage.sh

view-coverage:
	go tool cover -html=.\coverage.out
