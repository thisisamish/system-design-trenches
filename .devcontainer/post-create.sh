#!/usr/bin/env bash
set -euo pipefail

# make and graphviz support the trenches; curl installs Claude Code.
if ! command -v make >/dev/null || ! command -v dot >/dev/null || ! command -v curl >/dev/null; then
  sudo apt-get update
  sudo apt-get install -y --no-install-recommends make graphviz curl
fi

if [ ! -x "$HOME/.local/bin/claude" ]; then
  curl -fsSL https://claude.ai/install.sh | bash
fi

go version
docker --version || echo "docker not ready yet; re-open a terminal in a moment"
