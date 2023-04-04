# webgl-01

build : 
GOOS=js GOARCH=wasm go build -o  ../../assets/json.wasm 

run : 
cd cmd/server/
go run main.go