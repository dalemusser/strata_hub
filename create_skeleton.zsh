#!/usr/bin/env zsh
#
# create_skeleton.zsh — scaffold for strata_hub (gowebcore edition)
# Usage:  zsh create_skeleton.zsh
# Safe to re‑run; existing directories stay intact.

set -e          # exit on first error
set -u          # error on undefined variables
set -o pipefail # propagate failure through pipes

# -------------------------------------------------------------------
# Directory list — add/remove feature folders to taste
# -------------------------------------------------------------------
typeset -a DIRS=(
  cmd/strata_hub           # main() / CLI entry
  internal/config          # config structs
  internal/handler         # shared handler base
  internal/routes          # RegisterAllRoutes
  internal/templates       # embedded HTML / HTMx
  internal/db              # Mongo / DocumentDB helpers
  internal/about           # feature: /about
  internal/auth            # feature: /auth/*
  internal/dashboard       # feature: /dashboard/*
  assets                   # Tailwind / HTMx build output
)

# -------------------------------------------------------------------
# Create directories and keep them under version control
# -------------------------------------------------------------------
for d in "${DIRS[@]}"; do
  mkdir -p "$d"
  touch "$d"/.gitkeep
done

print -P "%F{green}✅  Project skeleton created.%f"
