.PHONY: gen_mem_password_example gen_secure_password_example
gen_mem_password_example:
	@echo "Generating memorable password example..."
	@go run main.go generate-password --qty=1 --len=3 --type=3
gen_secure_password_example:
	@echo "Generating secure password example..."
	@go run main.go generate-password --qty=1 --len=16 --type=4