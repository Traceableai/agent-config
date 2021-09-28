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
	$(MAKE) -C ./tools/hypertrace/agent-config/tools/go-generator

	@echo "Tidy generated modules."
	@find $(PWD)/gen/go \( -name vendor -o -name '[._].*' -o -name node_modules \) -prune -o -name go.mod -print | sed 's:/go.mod::' | xargs -I {} bash -c 'cd {}; go mod tidy -go=1.15'

generate-env-vars: init-git-submodule ## Generates the ENV_VARS.md with all environment variables.
	docker build -t hypertrace/agent-config/env-vars-generator tools/hypertrace/agent-config/tools/env-vars-generator
	touch $(PWD)/ENV_VARS.md # makes sure this is created as a file and not as a directory
	docker run \
	-v $(PWD)/ENV_VARS.md:/usr/local/ENV_VARS.md \
	-v $(PWD)/proto/ai/traceable/agent/config/v1/config.proto:/usr/local/config.proto \
	hypertrace/agent-config/env-vars-generator \
	-o "/usr/local/ENV_VARS.md" -p "TA_" \
	/usr/local/config.proto

init-git-submodule:
	git submodule update --init --recursive
