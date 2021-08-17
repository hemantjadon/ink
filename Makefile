# Path of current Makefile
ROOT_DIR := $(abspath $(dir $(abspath $(lastword $(MAKEFILE_LIST)))))

# All directories with go.mod files related to ink library. Used for building, testing and linting.
ALL_GO_MOD_DIRS := $(shell find $(ROOT_DIR) -type f -name 'go.mod' -exec dirname {} \; | sort)

# go toolchain.
GO := $(or $(shell which go), $(error "unable to find go toolchain in $$PATH"))

# tparse for pretty printing test output.
TPARSE := $(or $(shell which tparse), )

# golangci-lint for linting the project.
GOLANGCI_LINT := $(or $(shell which golangci-lint), $(error "unable to find golangci-lint in $$PATH"))

define get_module
	$(GO) get -v -t -d ./...
endef

define test_module
	$(if $(TPARSE),	\
		$(GO) test -cover -race -count 1 -timeout 30s -json ./... | $(TPARSE), \
		$(GO) test -cover -race -count 1 -timeout 30s ./... \
	)
endef

define test_module_verbose
	$(if $(TPARSE), \
		$(GO) test -cover -race -count 1 -timeout 30s -v -json ./... | $(TPARSE) -all, \
		$(GO) test -cover -race -count 1 -timeout 30s -v ./... \
	)
endef

define lint_module
	$(GOLANGCI_LINT) run ./...
endef

define print_header
	echo "-----------------------------------------------------------------"; \
	echo "$(1)"; \
	echo "-----------------------------------------------------------------"; \
	echo ""
endef

define print_footer
	echo "-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-"; \
	echo""
endef

get:
	@for DIR in $(ALL_GO_MOD_DIRS); do \
		$(call print_header,Get mod $$DIR); \
		cd $$DIR && $(call get_module); cd $(ROOT_DIR); \
		$(call print_footer); \
	done

test:
	@for DIR in $(ALL_GO_MOD_DIRS); do \
		$(call print_header,Test mod $$DIR); \
		cd $$DIR && $(call test_module); cd $(ROOT_DIR); \
		$(call print_footer); \
	done

test-verbose:
	@for DIR in $(ALL_GO_MOD_DIRS); do \
		$(call print_header,Test mod $$DIR); \
		cd $$DIR && $(call test_module_verbose); cd $(ROOT_DIR); \
		$(call print_footer); \
	done

lint:
	@for DIR in $(ALL_GO_MOD_DIRS); do \
		$(call print_header,Lint mod $$DIR); \
		cd $$DIR && $(call lint_module); cd $(ROOT_DIR); \
        $(call print_footer); \
	done
