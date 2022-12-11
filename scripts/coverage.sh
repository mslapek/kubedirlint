#!/bin/sh

#
# Check coverage inside coverage.out file.
#
# Fails if a coverage <100% is detected.
#

set -e

COV=$(go tool cover -func=coverage.out)

echo "== COVERAGE =="
printf '%s\n\n' "${COV}"

if printf '%s\n' "${COV}" | grep -v '100.0*%' > /dev/null; then
    echo "some module has <100% coverage" >&2
    exit 1
fi
