pb:
	goctl  rpc protoc distribute.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=./app/manager

sql:
	goctl model mysql ddl --src=./doc/sql/convert_job.sql --dir=app/manager/internal/model
	goctl model mysql ddl --src=./doc/sql/convert_config.sql --dir=app/manager/internal/model