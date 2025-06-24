# Setup the Platform
setup:
	@echo "Setting up the Platform"
	@echo "Generating the .env file"
	cp sample.env .env
	cp sample.env test.env
	@echo ".env file generated successfully update values in them"


# Migrate Database
migrate:
	@echo "Migrating the Database"
	sh scripts/migrate.sh
	@echo "Database Migrated Successfully"
# Run pretest script
pretest:
	sh scripts/test-helper.sh
	
migrate-test:
	@echo "Running migrations for tests..."
	sh scripts/migrate-tests.sh

# Run Tests
test-cover: migrate-test
	go test `go list ./... | grep -v cmd` -coverprofile=/tmp/coverage.out -coverpkg=./...
	go tool cover -html=/tmp/coverage.out
	
.PHONY: tests

tests:
	go test ./tests/integration/user_test.go ./tests/integration/resturent_test.go ./tests/integration/rating_test.go ./tests/integration/menu_card_test.go -v

	# Generate API documentation
doc:
	@echo "Generating swagger docs..."
	swag fmt --exclude ./internal/domain
	swag init --parseDependency --parseInternal -g internal/http/api/resturent_api.go -ot go,yaml -o internal/http/swagger

wire:
	cd internal/dependency/ && wire && cd ../..
