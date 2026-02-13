# Homebrew Installation Guide

This guide covers both GitHub and GitLab for hosting your Homebrew tap.

## Setup for GitHub (Technology-Institute)

### Step 1: Create GitHub Personal Access Token

1. Go to https://github.com/settings/tokens/new
2. Give it a descriptive name like "GoReleaser - Stock Market Agent"
3. Set expiration (or select "No expiration" if you prefer)
4. Select the following scopes:
   - ✅ `repo` (Full control of private repositories)
   - ✅ `write:packages` (Upload packages to GitHub Package Registry)
5. Click "Generate token"
6. **Copy the token immediately** (you won't be able to see it again!)

### Step 2: Set the Token as Environment Variable

```bash
export GITHUB_TOKEN="ghp_your_token_here"
```

To make it permanent, add to your `~/.bashrc` or `~/.zshrc`:
```bash
echo 'export GITHUB_TOKEN="ghp_your_token_here"' >> ~/.bashrc
source ~/.bashrc
```

### Step 3: Initialize the Homebrew Tap Repository

The tap repository is already created at:
https://github.com/Technology-Institute/homebrew-stock-market-agent

Initialize it with a Formula directory:
```bash
git clone https://github.com/Technology-Institute/homebrew-stock-market-agent.git
cd homebrew-stock-market-agent
mkdir -p Formula
echo "# Homebrew Tap for Stock Market Agent" > README.md
git add .
git commit -m "Initialize tap repository"
git push origin main
```

### Step 4: Create a Release

```bash
cd /path/to/stock-market-agent

# Make sure you're on the right remote
git remote add github https://github.com/Technology-Institute/stock-market-agent.git

# Push your code
git push github main

# Create and push a tag
git tag -a v1.0.0 -m "First release"
git push github v1.0.0

# Run GoReleaser (this will automatically update the homebrew tap)
goreleaser release --clean
```

### Step 5: Users Can Install

```bash
brew tap Technology-Institute/stock-market-agent
brew install stock-market-agent
```

## Using GitLab (Your Configuration)

### Creating Your Homebrew Tap on GitLab

1. **Create a new GitLab repository** for your Homebrew tap:
   ```bash
   # Create at: https://gitlab.com/technology_institute/homebrew-stock-market-agent
   # Note: The name must start with "homebrew-" for Homebrew to recognize it
   ```

2. **Set up the tap repository** (do this once):
   ```bash
   # Clone your tap repository
   git clone https://gitlab.com/technology_institute/homebrew-stock-market-agent.git
   cd homebrew-stock-market-agent
   mkdir -p Formula
   echo "# Homebrew Tap for Stock Market Agent" > README.md
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

4. **Create a release in your stock-market-agent repository**:
   ```bash
   cd /path/to/stock-market-agent
   
   # Tag your release
   git tag -a v1.0.0 -m "First release"
   git push origin v1.0.0
   
   # Run GoReleaser (this will create the Homebrew formula and push it to your tap repo)
   goreleaser release --clean
   ```

5. **Users can now install via Homebrew**:
   ```bash
   brew tap technology_institute/stock-market-agent https://gitlab.com/technology_institute/homebrew-stock-market-agent.git
   brew install stock-market-agent
   ```

   Or the shorter form after tapping:
   ```bash
   brew tap technology_institute/stock-market-agent https://gitlab.com/technology_institute/homebrew-stock-market-agent.git
   brew install technology_institute/stock-market-agent/stock-market-agent
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
