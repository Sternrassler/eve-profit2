#!/bin/bash

# EVE SDE Download Script - Fuzzwork SQLite
# Downloads the latest EVE Static Data Export in SQLite format

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
DATA_DIR="${SCRIPT_DIR}/data"
SDE_URL="https://www.fuzzwork.co.uk/dump/sqlite-latest.sqlite.bz2"
SDE_FILE="sqlite-latest.sqlite"
SDE_COMPRESSED="${SDE_FILE}.bz2"

echo "🚀 EVE Profit Calculator - SDE Download"
echo "========================================"

# Create data directory if it doesn't exist
mkdir -p "$DATA_DIR"
cd "$DATA_DIR"

# Check if SDE already exists and is recent
if [ -f "$SDE_FILE" ]; then
    FILE_AGE=$(find "$SDE_FILE" -mtime +7 2>/dev/null || echo "old")
    if [ "$FILE_AGE" != "old" ]; then
        echo "✅ SDE file exists and is less than 7 days old"
        echo "📁 File: $DATA_DIR/$SDE_FILE"
        echo "📊 Size: $(du -h "$SDE_FILE" | cut -f1)"
        echo ""
        echo "Use --force to download anyway"
        exit 0
    fi
fi

echo "📥 Downloading EVE SDE from Fuzzwork..."
echo "🔗 URL: $SDE_URL"

# Download the compressed SDE file
if command -v curl >/dev/null 2>&1; then
    curl -L -o "$SDE_COMPRESSED" "$SDE_URL"
elif command -v wget >/dev/null 2>&1; then
    wget -O "$SDE_COMPRESSED" "$SDE_URL"
else
    echo "❌ Error: Neither curl nor wget found"
    echo "Please install curl or wget to download the SDE"
    exit 1
fi

echo "📦 Downloaded: $(du -h "$SDE_COMPRESSED" | cut -f1)"

# Extract the SQLite database
echo "📤 Extracting SQLite database..."

if command -v bunzip2 >/dev/null 2>&1; then
    bunzip2 "$SDE_COMPRESSED"
elif command -v bzip2 >/dev/null 2>&1; then
    bzip2 -d "$SDE_COMPRESSED"
else
    echo "❌ Error: bzip2 not found"
    echo "Please install bzip2 to extract the SDE"
    exit 1
fi

# Verify the extracted file
if [ -f "$SDE_FILE" ]; then
    echo "✅ SDE successfully extracted!"
    echo "📁 File: $DATA_DIR/$SDE_FILE"
    echo "📊 Size: $(du -h "$SDE_FILE" | cut -f1)"
    
    # Basic SQLite verification
    if command -v sqlite3 >/dev/null 2>&1; then
        echo ""
        echo "🔍 Verifying SQLite database..."
        
        # Check if we can connect and run a basic query
        TABLES=$(sqlite3 "$SDE_FILE" ".tables" 2>/dev/null | wc -l)
        if [ "$TABLES" -gt 0 ]; then
            echo "✅ SQLite database is valid ($TABLES tables found)"
            
            # Show some basic stats
            echo ""
            echo "📈 SDE Statistics:"
            sqlite3 "$SDE_FILE" "SELECT 'Items: ' || COUNT(*) FROM invTypes WHERE published = 1;" 2>/dev/null || echo "Items: Unable to count"
            sqlite3 "$SDE_FILE" "SELECT 'Stations: ' || COUNT(*) FROM staStations;" 2>/dev/null || echo "Stations: Unable to count"
            sqlite3 "$SDE_FILE" "SELECT 'Systems: ' || COUNT(*) FROM mapSolarSystems;" 2>/dev/null || echo "Systems: Unable to count"
        else
            echo "⚠️  Warning: SQLite database may be corrupted"
        fi
    else
        echo "⚠️  Warning: sqlite3 not found - cannot verify database"
    fi
else
    echo "❌ Error: Extraction failed"
    exit 1
fi

echo ""
echo "🎉 SDE download complete!"
echo "💡 The SDE file is ready for use by the EVE Profit Calculator backend"
