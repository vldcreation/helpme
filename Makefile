.PHONY: gen_mem_password_example gen_secure_password_example find_go_pkg_strings_example build install clean pull

build_dir=./bin
helpme=$(build_dir)/helpme

# Default values
BRANCH ?= master

# Repository mapping
REPO_USER := vldcreation
REPO_NAME := helpme-package

.DEFAULT_GOAL := all

all: $(helpme) gen_mem_password_example gen_secure_password_example find_go_pkg_strings_example

GO_FILES := $(shell git ls-files '*.go')
.PHONY: $(GO_FILES)

$(helpme): $(GO_FILES)
	@echo "Building..."
	@go build -o $(build_dir)/helpme .

gen_mem_password_example: $(helpme)
	@echo "Generating memorable password example..."
	@$(helpme) generate-password --qty=1 --len=3 --type=2

gen_secure_password_example: $(helpme)
	@echo "Generating secure password example..."
	@$(helpme) generate-password --qty=1 --len=16 --type=4

find_go_pkg_strings_example: $(helpme)
	@echo "Finding Go packages example 'strings.Join'"
	@$(helpme) find -l=go -p=strings -s -e "Join"

install: $(helpme)
	@echo "Installing..."
	@go install .

clean:
	@echo "Cleaning..."
	@rm -rf $(build_dir)

pull:
	@if [ "$(r)" = "pkg" ]; then \
		REPO_NAME="helpme-package"; \
	elif [ "$(r)" = "src" ]; then \
		REPO_NAME="go-ressources"; \
	else \
		echo "Invalid repository flag. Use r=pkg or r=src"; \
		exit 1; \
	fi; \
	if [ ! -z "$(b)" ]; then \
		BRANCH="$(b)"; \
	fi; \
	echo "Pulling from $(REPO_USER)/$(REPO_NAME) branch: $(BRANCH)"; \
	$(build_dir)/helpme pull -u=$(REPO_USER) -r=$(REPO_NAME) -b=$(BRANCH)
	