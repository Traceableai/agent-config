BREAKING_CHECK_AGAINST := .git\#branch=main

lint-proto: ## Lints the proto files.
	buf lint

breaking-proto: ## Checks the proto files for breaking changes
	buf breaking --against $(BREAKING_CHECK_AGAINST)

generate-proto:
	@# Generates pb struct
	buf generate
	
	@# Generates pb loaders
	ROOT="$(PWD)/proto" \
	OUT_DIR="$(PWD)/gen/go" \
	OPT_MODULE="github.com/Traceableai/agent-config/gen/go" \
	ENV_PREFIX="TA_" \
	$(MAKE) -C ./tools/go-generator

	@echo "Tidy generated modules."
	@find $(PWD)/gen/go \( -name vendor -o -name '[._].*' -o -name node_modules \) -prune -o -name go.mod -print | sed 's:/go.mod::' | xargs -I {} bash -c 'cd {}; go mod tidy -go=1.21'

	@# Run gen/go load sanity tests
	cd $(PWD)/gen/go && go test ./...

generate-env-vars: ## Generates the ENV_VARS.md with all environment variables.
	docker build -t tools/env-vars-generator tools/env-vars-generator
	touch $(PWD)/ENV_VARS.md # makes sure this is created as a file and not as a directory
	docker run \
	-v $(PWD)/ENV_VARS.md:/usr/local/ENV_VARS.md \
	-v $(PWD)/proto/ai/traceable/agent/config/v1/config.proto:/usr/local/config.proto \
	tools/env-vars-generator \
	-o "/usr/local/ENV_VARS.md" -p "TA_" \
	/usr/local/config.proto

generate: generate-proto generate-env-vars

init-git-submodule:
	git submodule update --init --recursive
