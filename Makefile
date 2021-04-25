all:
	go build -o plugin-sample main.go
	go build -ldflags="-s -w" -buildmode=plugin -o writers/plugins/en.so writers/plugins/en/en.go
	go build -buildmode=plugin -o writers/plugins/de.so writers/plugins/de/de.go

compress:
	chmod +x ./writers/plugins/en.so
	upx -9 -k ./writers/plugins/en.so