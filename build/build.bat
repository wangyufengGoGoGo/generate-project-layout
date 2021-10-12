set GOOS=windows
set GOARCH=amd64
set ENABLED=0

go build ..\cmd\generate_porject_layout\server.go

move server.exe layout.exe