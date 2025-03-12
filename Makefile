.PHONY: gen_mem_password_example gen_secure_password_example find_go_pkg_strings_example build install clean

build_dir=./bin
helpme=$(build_dir)/helpme

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

.PHONY: pull

# Default values
BRANCH ?= master

# Repository mapping
REPO_USER := vldcreation
REPO_NAME := helpme-package

# Parse repository type from command line
ifneq (,$(findstring pkg,$(r)))
	REPO_NAME := helpme-package
else ifneq (,$(findstring src,$(r)))
	REPO_NAME := go-ressources
else
	$(error Invalid repository flag. Use r=pkg or r=src)
endif

# Override branch if specified
ifneq (,$(b))
	BRANCH := $(b)
endif

pull:
	@echo "Pulling from $(REPO_USER)/$(REPO_NAME) branch: $(BRANCH)"
	@$(build_dir)/helpme pull -u=$(REPO_USER) -r=$(REPO_NAME) -b=$(BRANCH)
	