set CGO_ENABLED=0 
set GOARCH=amd64 
set GOOS=linux 
go build ./src/core.go
mv core chembox-core-linux-amd64
echo "done"