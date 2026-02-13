# Homebrew Installation Guide

## Creating Your Homebrew Tap

1. **Create a new GitHub repository** for your Homebrew tap:
   ```bash
   # The repository MUST be named: homebrew-tap
   # GitHub URL will be: https://github.com/bjoern/homebrew-tap
   ```

2. **Set up the tap repository** (do this once):
   ```bash
   # Clone your tap repository
   git clone https://github.com/bjoern/homebrew-tap.git
   cd homebrew-tap
   mkdir -p Formula
   # Create initial commit
   git add .
   git commit -m "Initial commit"
   git push origin main
   ```

3. **Configure GoReleaser** with GitHub token:
   ```bash
   # Create a GitHub Personal Access Token at:
   # https://github.com/settings/tokens/new
   # Scopes needed: repo, write:packages
   
   export GITHUB_TOKEN="your_github_token_here"
   ```

4. **Create a release in your go-agent repository**:
   ```bash
   cd /path/to/go-agent
   
   # Tag your release
   git tag -a v1.0.0 -m "First release"
   git push origin v1.0.0
   
   # Run GoReleaser
   goreleaser release --clean
   ```

5. **Users can now install via Homebrew**:
   ```bash
   brew tap bjoern/tap
   brew install stock-market-agent
   ```

## Testing Locally Before Release

Test the release process without publishing:
```bash
# Install goreleaser if not already installed
brew install goreleaser

# Create a snapshot release (doesn't publish)
goreleaser release --snapshot --clean

# Check the dist/ folder for generated artifacts
ls -la dist/
```

## Manual Installation from Release

If you don't want to use Homebrew:

1. Download the appropriate archive from GitHub releases
2. Extract it:
   ```bash
   tar -xzf stock-market-agent_1.0.0_linux_amd64.tar.gz
   ```
3. Move to your PATH:
   ```bash
   sudo mv stock-market-agent /usr/local/bin/
   ```

## Configuration After Installation

After installing via Homebrew, config files are located at:
- `/usr/local/etc/stock-market-agent/config.yaml` (Intel Mac)
- `/opt/homebrew/etc/stock-market-agent/config.yaml` (Apple Silicon Mac)
- `~/.linuxbrew/etc/stock-market-agent/config.yaml` (Linux)

Copy and configure:
```bash
cp $(brew --prefix)/etc/stock-market-agent/.sample-env ~/.stock-market-agent.env
# Edit with your credentials
nano ~/.stock-market-agent.env
```
