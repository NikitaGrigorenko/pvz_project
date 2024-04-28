ifeq ($(POSTGRES_SETUP_TEST),)
	POSTGRES_SETUP_TEST := user=test password=test dbname=test host=localhost port=5432 sslmode=disable
endif

ifeq ($(POSTGRES_SETUP_TEST_TEST),)
	POSTGRES_SETUP_TEST_TEST := user=test_test password=test_test dbname=test_test host=localhost port=5433 sslmode=disable
endif

INTERNAL_PKG_PATH=$(CURDIR)/internal/pkg
MOCKGEN_TAG=v1.6.0
MIGRATION_FOLDER=$(INTERNAL_PKG_PATH)/db/migrations

.PHONY: migration-create
migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

.PHONY: test-migration-up
test-migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" up

.PHONY: test-migration-down
test-migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" down

.PHONY: migration-create-test
migration-create-test:
	goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

.PHONY: test-migration-up-test
test-migration-up-test:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST_TEST)" up

.PHONY: test-migration-down-test
test-migration-down-test:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST_TEST)" down

.PHONY: .generate-mockgen-deps
.generate-mockgen-deps:
ifeq ($(wildcard $(MOCKGEN_BIN)),)
	@GOBIN=$(LOCAL_BIN) go install github.com/golang/mock/mockgen@$(MOCKGEN_TAG)
endif

.PHONY: .generate-mockgen
.generate-mockgen:
	PATH="$(LOCAL_BIN):$$PATH" go generate -x -run=mockgen ./...

.test:
	$(info Running tests...)
	go test -count=1 -cover ./...

test_unit: .test

.test_int:
	$(info Running tests...)
	go test -count=1 -tags integration ./tests

test_integration: .test_int

.PHONY: run-test-env
run-test-env:
	docker-compose -f docker-compose-test.yaml up -d

.PHONY: stop-test-env
stop-test-env:
	docker-compose -f docker-compose-test.yaml down

.PHONY: run-main-bd
run-main-bd:
	docker-compose -f docker-compose.yaml up -d

.PHONY: stop-main-bd
stop-main-bd:
	docker-compose -f docker-compose.yaml down

build:
	docker-compose -f docker-compose-kafka.yaml build

up-kafka:
	docker-compose -f docker-compose-kafka.yaml up -d zookeeper kafka1 kafka2 kafka3

down-kafka:
	docker-compose -f docker-compose-kafka.yaml down

proto-gen:
	rm -rf internal/pkg/pb
	mkdir -p internal/pkg/pb

	protoc \
    		--go_out=internal/pkg/pb \
    		--go-grpc_out=internal/pkg/pb \
    		api/devtask/**/**/*.proto

up-grpc:
	docker-compose -f docker-compose-grpc.yaml up -d

down-grpc:
	docker-compose -f docker-compose-grpc.yaml down
