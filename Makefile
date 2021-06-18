.DEFAUT_GOAL = help
.SILENT:

ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
$(eval $(ARGS):;@:)

build-linux: ##build de la version linux
	env GOOS=linux GOARCH=amd64 go build -o ./build/linux/copyfast cmd/copyfast.go 

build-darwin:  ##build de la version darwin
	env GOOS=darwin GOARCH=amd64 go build -o ./build/darwin/copyfast cmd/copyfast.go 

build-windows: ##build de la version windows
	env GOOS=windows GOARCH=amd64 go build -o ./build/windows/copyfast cmd/copyfast.go 

build-all: build-linux build-darwin build-windows ##build de toutes les versions

test: build-linux	##test la version linux sur un container
	docker build --force-rm -t locals/copyfast .
	docker run --rm -ti -v $$(pwd)/build/linux/copyfast:/bin/copyfast locals/copyfast /bin/copyfast

run: ## run sans compilation du porjet GO
	go run cmd/copyfast.go $(ARGS)

help: #pour générer automatiquement l'aide ## Display all commands available
	@grep -E '^[a-zA-Z_-]+:.*? ## .*$$ $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

git:
	git add -A .
	git commit -m "Auto Commit"
	git push