#!/bin/sh
set -e

OUTPUT_DIR="${OUTPUT_DIR:-/output}"

find /plugins -name "go.mod" | while read gomod; do
  plugin_dir=$(dirname "$gomod")
  plugin_name=$(basename "$plugin_dir")
  echo "--- Compiling: $plugin_name ---"
  PLUGIN_SOURCE_PATH="$plugin_dir" /build.sh "${plugin_name}.so" "" linux arm64
  find "$plugin_dir" -name "*.so" -exec cp {} "$OUTPUT_DIR/" \;
done

echo "Compilation complete. Output:"
ls -la "$OUTPUT_DIR/"
