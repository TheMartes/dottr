MIN_MAKE_VERSION := 3.81

# Min version
ifneq ($(MIN_MAKE_VERSION),$(firstword $(sort $(MAKE_VERSION) $(MIN_MAKE_VERSION))))
	$(error GNU Make $(MIN_MAKE_VERSION) or higher required)
endif

.DEFAULT_GOAL:=help

##@ Test Files
.PHONY: test-files delete create-config
init: ## Create test-files for dottr
	./.scripts/init_dottr.sh
cleanup: ## Cleanup after dev
	./.scripts/delete_dottr.sh

##@ install
.PHONY: install
install: ## Start only daemon, assuming youre infrastructure is running already (use "args=" to supply custom arguments)
	go install .

##@ Help
.PHONY: help
help: ## Show all available commands (you are looking at it)
	@awk 'BEGIN {FS = ":.*##"; printf "Usage: make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
