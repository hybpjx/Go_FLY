SET CGO_ENABLED=O
SET GOOS=linux
SET GOARCH=amd64
echo now the CGO_ENABLED:
    go env CGO_ENABLED

echo now the GoOS:
    go env GOOS
echo now the GOARCH:
    go env GOARCH
go build main.go
SET CGO_ENABLED=1
SET GOOS=windows
SET GOARCH=amd64

