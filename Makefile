all:
	go build -o plugin-sample main.go
	go build -buildmode=plugin -o writers/plugins/en.so writers/plugins/en/en.go
	go build -buildmode=plugin -o writers/plugins/de.so writers/plugins/de/de.go