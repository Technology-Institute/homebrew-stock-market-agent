# Homebrew Installation Guide

This guide covers both GitHub and GitLab for hosting your Homebrew tap.

## Option 1: Using GitHub

### Creating Your Homebrew Tap

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

4. **Edit `goreleaser.yaml`** and uncomment the GitHub brews section

5. **Create a release in your go-agent repository**:
   ```bash
   cd /path/to/go-agent
   
   # Tag your release
   git tag -a v1.0.0 -m "First release"
   git push origin v1.0.0
   
   # Run GoReleaser
   goreleaser release --clean
   ```

6. **Users can now install via Homebrew**:
   ```bash
   brew tap bjoern/tap
   brew install stock-market-agent
   ```

## Option 2: Using GitLab

### Creating Your Homebrew Tap on GitLab

1. **Create a new GitLab repository** for your Homebrew tap:
   ```bash
   # The repository MUST be named: homebrew-tap
   # GitLab URL will be: https://gitlab.com/bjoern/homebrew-tap
   ```

2. **Set up the tap repository** (do this once):
   ```bash
   # Clone your tap repository
   git clone https://gitlab.com/bjoern/homebrew-tap.git
   cd homebrew-tap
   mkdir -p Formula
   # Create initial commit
   git add .
   git commit -m "Initial commit"
   git push origin main
   ```

3. **Configure GoReleaser** with GitLab token:
   ```bash
   # Create a GitLab Personal Access Token at:
   # https://gitlab.com/-/profile/personal_access_tokens
   # Scopes needed: api, write_repository
   
   export GITLAB_TOKEN="your_gitlab_token_here"
   ```

4. **Edit `goreleaser.yaml`**:
   - Uncomment the `gitlab_urls` section
   - Uncomment the GitLab brews section
   - Comment out or remove the GitHub brews section

5. **Create a release in your go-agent repository**:
   ```bash
   cd /path/to/go-agent
   
   # Add GitLab as remote if needed
   git remote add gitlab https://gitlab.com/bjoern/go-agent.git
   
   # Tag your release
   git tag -a v1.0.0 -m "First release"
   git push gitlab v1.0.0
   
   # Run GoReleaser
   goreleaser release --clean
   ```

6. **Users can now install via Homebrew**:
   ```bash
   brew tap bjoern/tap https://gitlab.com/bjoern/homebrew-tap.git
   brew install stock-market-agent
   ```

## Testing Locally Before Release

Test the release process without publishing:
```bash
# Install goreleaser if not already installed
brew install goreleaser  # macOS
# or download from goreleaser.com for Linux

# Create a snapshot release (doesn't publish)
goreleaser release --snapshot --clean

# Check the dist/ folder for generated artifacts
ls -la dist/
```

## Manual Installation from Release

If you don't want to use Homebrew:

### From GitHub Releases
1. Go to https://github.com/bjoern/go-agent/releases
2. Download the appropriate archive
3. Extract and install:
   ```bash
   tar -xzf stock-market-agent_1.0.0_linux_amd64.tar.gz
   sudo mv stock-market-agent /usr/local/bin/
   ```

### From GitLab Releases
1. Go to https://gitlab.com/bjoern/go-agent/-/releases
2. Download the appropriate archive
3. Extract and install:
   ```bash
   tar -xzf stock-market-agent_1.0.0_linux_amd64.tar.gz
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

## Notes

- **Repository naming**: The tap repository MUST be named `homebrew-tap` for Homebrew to recognize it
- **Token permissions**: Make sure your token has sufficient permissions to push to the tap repository
- **GitLab URLs**: If using a self-hosted GitLab instance, update the `gitlab_urls` section in `goreleaser.yaml` with your instance URL
