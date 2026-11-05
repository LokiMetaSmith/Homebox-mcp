#!/bin/bash
sed -i -e 's/os.Setenv/t.Setenv/g' homebox-mcp-server/main_test.go
