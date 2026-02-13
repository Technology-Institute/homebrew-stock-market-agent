#!/bin/bash

# Guide to push to GitHub with token authentication

echo "========================================"
echo "GitHub Push Authentication Guide"
echo "========================================"
echo ""

# Load GITHUB_TOKEN from .env if it exists
if [ -f .env ]; then
    echo "Loading GITHUB_TOKEN from .env..."
    export $(grep -v '^#' .env | grep GITHUB_TOKEN | xargs)
fi

# Check if GITHUB_TOKEN is set
if [ -z "$GITHUB_TOKEN" ]; then
    echo "⚠️  GITHUB_TOKEN is not set!"
    echo ""
    echo "Add it to your .env file:"
    echo "  echo 'GITHUB_TOKEN=ghp_yourTokenHere' >> .env"
    echo ""
    echo "Or set it as environment variable:"
    echo "  export GITHUB_TOKEN='ghp_yourTokenHere'"
    echo ""
    exit 1
fi

echo "✓ GITHUB_TOKEN is set"
echo ""
echo "Updating git remote to use token..."
echo ""

# Remove existing github remote and add with token
git remote remove github 2>/dev/null || true
git remote add github "https://$GITHUB_TOKEN@github.com/Technology-Institute/homebrew-stock-market-agent.git"


echo "✓ Remote configured with token authentication"
echo ""
echo "Now you can push:"
echo "  git push github main"
echo "  git push github --tags"
echo ""
