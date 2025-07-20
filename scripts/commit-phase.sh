#!/bin/bash

# Git Commit & Push Script für Entwicklungsphasen
# Usage: ./commit-phase.sh <phase-number> <phase-name> [description]

set -e

if [ "$#" -lt 2 ]; then
    echo "Usage: $0 <phase-number> <phase-name> [description]"
    echo "Example: $0 2 'SDE Client Implementation' 'SQLite integration for Items and Stations'"
    exit 1
fi

PHASE_NUMBER="$1"
PHASE_NAME="$2"
DESCRIPTION="${3:-''}"

echo "🚀 Git Commit & Push - Phase $PHASE_NUMBER"
echo "=========================================="

# Check if we're in a git repository
if [ ! -d ".git" ]; then
    echo "❌ Error: Not in a git repository"
    exit 1
fi

# Add all files (except ignored ones)
echo "📁 Adding files to git..."
git add .

# Check if there are changes to commit
if git diff --staged --quiet; then
    echo "ℹ️  No changes to commit"
    exit 0
fi

# Show what will be committed
echo ""
echo "📋 Files to be committed:"
git diff --staged --name-status

echo ""
read -p "Continue with commit? (y/N): " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "❌ Commit cancelled"
    exit 1
fi

# Create commit message
COMMIT_MSG="✅ Phase $PHASE_NUMBER: $PHASE_NAME"

if [ -n "$DESCRIPTION" ]; then
    COMMIT_MSG="$COMMIT_MSG

📋 Description:
$DESCRIPTION

📅 Phase $PHASE_NUMBER completed $(date '+%Y-%m-%d %H:%M')"
fi

# Commit changes
echo "💾 Creating commit..."
git commit -m "$COMMIT_MSG"

# Push to remote
echo "🚀 Pushing to remote..."
git push origin main

echo ""
echo "✅ Phase $PHASE_NUMBER successfully committed and pushed!"
echo "🔗 Repository: https://github.com/Sternrassler/eve-profit2"
echo ""

# Show commit info
echo "📊 Commit Info:"
git log -1 --oneline
echo ""

echo "🎯 Ready for next phase!"
