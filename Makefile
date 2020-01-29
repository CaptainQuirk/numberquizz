COMMANDS := $(MAKEFILE_LIST)
build: ## Compiles the project and installs a binary in the bin subdirectory
	@cd generator; go build; cd - > /dev/null
	@go install .
test:
	@go test ./...
help:
	@echo "\033[91mNumber quizz\n============\033[0m\n\nCommand line application that quizzes for different representation of a number. For educational purposes."
	@echo "\n\033[35mAvailable commands\033[0m\n"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(COMMANDS) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo "\n\033[35mShortcuts\033[0m\n"
	@grep -P '^[a-zA-Z_-]+: (?!##)' $(COMMANDS) | awk '{printf "\033[36m%-30s\033[0m \n", $$1; for (i = 2; i<= NF; i++) print " - " $$i}'
