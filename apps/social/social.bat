@echo off

call goctl rpc protoc ./apps/social/rpc/social.proto --go_out=./apps/social/rpc/ --go-grpc_out=./apps/social/rpc/ --zrpc_out=./apps/social/rpc/

call goctl model mysql ddl -src="./deploy/sql/social.sql" -dir="./apps/social/models/" -c

call goctl api go -api ./apps/social/api/social.api -dir ./apps/social/api -style gozero