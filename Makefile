.PHONY: build install clean test release

# Build the binary
build:
	go build -o stock-market-agent .

# Install dependencies
deps:
	go mod download
	go mod tidy

# Clean build artifacts
clean:
	rm -f stock-market-agent
	rm -rf dist/

# Test the application
test:
	go test -v ./...

# Run the agent
run: build
	./stock-market-agent

# Install locally
install: build
	mkdir -p ~/.local/bin
	cp stock-market-agent ~/.local/bin/
	mkdir -p ~/.config/stock-market-agent
	cp config.yaml ~/.config/stock-market-agent/
	cp .sample-env ~/.config/stock-market-agent/
	@echo "Installed to ~/.local/bin/stock-market-agent"
	@echo "Config files in ~/.config/stock-market-agent/"
	@echo "Copy .sample-env to ~/.config/stock-market-agent/.env and configure"

# Create a release (requires goreleaser)
release:
	goreleaser release --snapshot --clean

# Create a release with tag
release-tag:
	goreleaser release --clean
