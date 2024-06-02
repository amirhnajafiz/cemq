#!/bin/bash

# Function to check if current Go version is greater than or equal to required version
check_go_version() {
    local min_version="1.22.0"
    local current_version=$(go version | awk '{print $3}' | cut -d 'o' -f2)

    if [ "$(printf '%s\n' "$min_version" "$current_version" | sort -V | head -n1)" = "$min_version" ]; then
        return 0  # True, current version is greater or equal
    else
        return 1  # False, current version is less
    fi
}

# Checking Go version
if check_go_version; then
    echo "Go version is sufficient."
    exit 0
else
    echo "Go version is insufficient. Please upgrade to Go 1.22.0 or later."
    exit 1
fi
