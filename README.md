**To run server**
cmd : go run server/main.go

**To run client**
cmd : go run client/main.go

**To run test files**
cmd : go test ./...

**Pre-requistes:**

Go need to be installed in your system.

install mockGen to generate mock files and
protoc to compile *.proto files

*cmd to compile proto :*
protoc \                                                  
    --go_out=./proto \
    --go_opt=paths=source_relative \
    --go-grpc_out=./proto \
    --go-grpc_opt=paths=source_relative \
    ./proto/blog.proto

*cmd to generate mock file*
mockgen -source=SOURCE_FILE_PATH -destination=DESTINATION_FILE_PATH -package=PACKAGE_NAME

export above mockGen and protoc to path 

to install required packages -> 
*go mod tidy*
