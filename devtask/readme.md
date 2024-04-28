# Pickup Point Management Utility 

This tool is the manager for Pickup Point which managing how to take orders from courier, return orders to courier, give orders to clients and more. 


## Build instructions

Prerequisites:
- Go 1.21
- Grafana
- Prometheus
- BloomRPC (or other tool to make grpc requests)
- Offset Explorer 3 (or other tool to interact with kafka)

```bash
# Navigate to the cloned Git repository folder with the source code of the tool
cd homework-ozon
```

## Sample Usage

The syntax for a sample invocation of this tool is as follows:

```bash
go run ./cmd/console-app/main.go take_order --order_id=5 --client_id=4 --date_exp=2024-03-29 --weight=40 --cost=32 --package=2
```

Above, `go run` invokes an automatic pipeline for the tool. 
For such a sample run, the following commands are specified:
- `take_order` takes an order from courier. Command has a list of parameters:
  - `--order_id` id of an order to be taken
  - `--client_id` id of a client owns an order
  - `--date_exp` date of expiration in pickup point
  - `--weight` weight of the order in kilograms
  - `--cost` initial cost of the order (without package)
  - `--package` type of the package (0-packet, 1-box, 2-tape)
  
  sample invocation `go run ./cmd/console-app/main.go take_order --order_id=5 --client_id=4 --date_exp=2024-03-29 --weight=40 --cost=32 --package=2`

- `return_order` returns an order to courier. Command has a list of parameters:
  - `--order_id` id of an order to be returned

   sample invocation `go run ./cmd/console-app/main.go return_order --order_id=4`

- `give_order` gives slice of orders to the client. Command has a list of parameters:
  - `--slice` slice of ids to be given

  sample invocation `go run ./cmd/console-app/main.go give_order --slice=7,12,11`

- `list_orders` shows orders of the client. Command has a list of parameters:
  - `--client_id` id of a client owns an orders
  - `--last_n` number of last orders to be swown
  - `--inpp` if true -> shows only orders which are in pickup point

  sample invocation `go run ./cmd/console-app/main.go list_orders --client_id=3 --last_n=3 --inpp=false`

- `client_refund` refunds an order from the client to Pickup Point. Command has a list of parameters:
  - `--order_id` id of an order to be refunded
  - `--client_id` id of a client owns an order

  sample invocation `go run ./cmd/console-app/main.go client_refund  --client_id=3 --order_id=2`

- `list_refunds` shows 5 refunds on each page. Command has a list of parameters:
  - `--page_number` the number of page to be shown

  sample invocation `go run ./cmd/console-app/main.go list_refunds --page_number=1`

- `pvz` runs an interactive menu to work with pickup points.

  sample invocation `go run ./cmd/console-app/main.go pvz`

- `help` shows a list of commands which can be used with their description

  sample invocation `go run ./cmd/console-app/main.go help`

## Sample Usage with HTTP server

```bash
# At first, we need to run database by docker
make run-main-bd

# Next we need to migrate our database
make test-migration-up

# Also, we need to build docker-compose file for kafka
make build

# And run all kafka brokers
make up-kafka

# To run http server application use
go run ./cmd/http-app/main.go
```

Now the following **curl** commands can be used:

- `curl -i -u nikita:1234 -X POST localhost:9000/pvz --header "Content-Type: application/json" --data '{"name": "pvzInnopolis", "address": "Innopolis", "contact": "+79221593814"}'` to add new pvz information into database
- `curl -i -u nikita:1234 -X PUT localhost:9000/pvz/1 --header "Content-Type: application/json" --data '{"name": "pvzMoscow", "address": "Moscow", "contact": "+792219923232"}'` to update existing information about pvz by id in database
- `curl -i -u nikita:1234 -X GET http://localhost:9000/pvz/1` to get information about one pvz by id from database
- `curl -i -u nikita:1234 -X GET http://localhost:9000/pvz` to get information about all pvz from database
- `curl -i -u nikita:1234 -X DELETE http://localhost:9000/pvz/1` to remove information about one pvz by id from database

Also, secure connection can be established by using 9001 port and https

- `curl -i -u nikita:1234 -X GET https://localhost:9001/pvz/1` try to establish secure connection (now it gives an error because of self-signed certificates)

Examples of incorrect requests

- `curl -i -u nikita:12345 -X POST localhost:9000/pvz --header "Content-Type: application/json" --data '{"name": "pvzInnopolis", "address": "Innopolis", "contact": "+79221593814"}'` should return `401 Unauthorized` because of incorrect password
- If you try to access a pvz row which is not in database you will get `404 Not Found`

When you are done with application

```bash
# You can clean all migration if you want
make test-migration-down

# To stop kafka brokers
make down-kafka

# Next down the docker database
make stop-main-bd
```

## Sample Usage with GRPC server

```bash
# At first, we need to run database by docker
make run-main-bd

# Next we need to migrate our database
make test-migration-up

# Also, we need to build docker-compose file for kafka
make build

# And run all kafka brokers
make up-kafka

# Also, you need to up grpc server
make up-grpc

# To run http server application use
go run ./cmd/http-app/main.go

# To run grpc server application use
go run ./cmd/grpc-app/main.go
```

Now the application can be used by BloomRPC app

When you are done with application

```bash
# You can clean all migration if you want
make test-migration-down

# To stop kafka brokers
make down-kafka

# Also, you need to stop grpc server
make down-grpc

# Next down the docker database
make stop-main-bd
```

# Testing
To test app by integration tests at first we need to run testing environment:

```bash
# At first, we need to run test database by docker
make run-test-env

# Next we need to migrate our test database
make test-migration-up-test

# Also, we need to build docker-compose file for kafka
make build

# And run all kafka brokers
make up-kafka

# Also, you need to up grpc server
make up-grpc

# To test app by using integration tests use:
make test_integration
```

Also, if you want to run unit test you need to run

```bash
make test_unit
```

When you are done with testing:
```bash
# You need to clean all migration to test database
make test-migration-down-test

# To stop kafka brokers
make down-kafka

# Also, you need to stop grpc server
make down-grpc

# Next down the docker test database
make stop-test-env
```

# Metrics
To see all metrics from the application:

```bash
# Custom metrics
curl -i -u nikita:1234 -X GET http://localhost:9000/metrics

# Standard metrics
curl -i -u nikita:1234 -X GET http://localhost:9090/metrics
```

# Tracing
To see tracing of the app in grpc mode open in the browser:

`localhost:16686`


# Technologies used:

- Golang
- Postgresql
- Redis
- Grafana
- Prometheus
- JaegerUI
- Kafka
- Rest
- Grpc
- Docker