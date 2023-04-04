# webgl-01

build : 
GOOS=js GOARCH=wasm go build -o  ../../assets/json.wasm 

run : 
cd cmd/server/
go run main.go


apt-get install libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev mesa-common-dev libxxf86vm-dev