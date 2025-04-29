$ports = @(8081, 8082, 8083)
foreach ($port in $ports) {
    Start-Process powershell -ArgumentList "-NoExit", "-Command `"`$env:PORT='$port'; go run server.go`""
}
