#!/bin/bash

# Function to detect profile file
detect_profile() {
    if [[ -f "$HOME/.zshrc" ]]; then
        echo "$HOME/.zshrc"
    elif [[ -f "$HOME/.bash_profile" ]]; then
        echo "$HOME/.bash_profile"
    elif [[ -f "$HOME/.bashrc" ]]; then
        echo "$HOME/.bashrc"
    else
        echo "No profile file found."
        exit 1
    fi
}

# Detecting profile file
profile_file=$(detect_profile)

# Path to be added, assuming pwa is a variable holding path
pwa="$( pwd )"  # Change this path to your executable path

# Adding PATH to user profile
echo "export PATH=\"$pwa:\$PATH\"" >> "$profile_file"
echo "PATH updated in $profile_file"

echo "Please run source .zshrc/.bash_profice/.bashrc before running CEMQ commands."
echo "CEMQ installation and setup complete."
