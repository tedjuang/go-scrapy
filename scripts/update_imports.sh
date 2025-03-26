#!/bin/bash

# Update import paths in Go files
find . -name "*.go" -type f -exec sed -i '' 's|github.com/yourusername/scrapy/pkg/|github.com/yourusername/scrapy/internal/|g' {} \;

echo "Import paths updated." 