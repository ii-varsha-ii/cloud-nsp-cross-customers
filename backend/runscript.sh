#!/usr/bin/env bash

HOST='http://localhost:8000'
echo "Hello! "
echo "$HOST/org/api/get"
ORGANIZATION=$(curl -s "$HOST/org/api/get")

echo "$ORGANIZATION"