DEFAULT_GOAL = help
.SILENT: 


build-windows:
	env GOOS=windows GOARCH=amd64 go build -o build/copyfast cmd/copyfast.go


build-all: build-windows

test: build-windows
	docker build --force-rm -t locals/copyfast .
	docker run locals/copyfast

run:
	go run cmd/copyfast.go
	

help: #Pour générer automatiquement l'aide ## Display all commands available
    @grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


git: 
	git add -A .
	git commit -m "Auto Commit"
	git push 