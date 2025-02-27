.PHONY: gen_mem_password_example gen_secure_password_example find_go_pkg_strings_example build install

build_dir=./bin
helpme=./bin/helpme

gen_mem_password_example:
	@echo "Generating memorable password example..."
	@$(helpme) generate-password --qty=1 --len=3 --type=3
gen_secure_password_example:
	@echo "Generating secure password example..."
	@$(helpme) generate-password --qty=1 --len=16 --type=4
find_go_pkg_strings_example:
	@echo "Finding Go packages example 'strings.Join'"
	@$(helpme) find -l=go -p=strings -s -e "Join"

# dev step 
build:
	@echo "Building..."
	@go build -o $(build_dir) .
install:
	@echo "Installing..."
	@go install .
