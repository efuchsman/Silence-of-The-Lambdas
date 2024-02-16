#!/bin/bash

# Function to check if a command is installed
check_command() {
    if ! command -v "$1" &> /dev/null; then
        return 1  # Command not found
    fi
}

# Check if AWS SAM CLI is installed
if ! check_command "sam"; then
    echo "AWS SAM CLI not found. Installing..."
    brew install aws-sam-cli
fi

if ! check_command "java"; then
    echo "Java not found. Installing..."
    brew install openjdk@11
fi

# Set DynamoDB Local installation directory
dynamodb_local_dir="/usr/local/Caskroom/dynamodb-local/latest"

# Check if DynamoDB Local is installed
if [ ! -d "$dynamodb_local_dir" ] || [ ! -f "$dynamodb_local_dir/DynamoDBLocal.jar" ]; then
    echo "DynamoDB Local not found. Downloading..."
    mkdir -p "$dynamodb_local_dir"
    curl -O https://s3.us-west-2.amazonaws.com/dynamodb-local/dynamodb_local_latest.tar.gz
    tar -zxvf dynamodb_local_latest.tar.gz -C "$dynamodb_local_dir"
    rm dynamodb_local_latest.tar.gz
fi

echo "AWS SAM CLI, Java, and DynamoDB Local are installed."
