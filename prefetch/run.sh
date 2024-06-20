#!/usr/bin/env bash

set -euo pipefail

function run() {
    echo === ./prefetch run -dir foo -prefetch="$1" -hit="$2"
    local NAME="run_prefetch_$1_hit_$2"
    systemctl stop "$NAME".slice
    systemd-run --scope --slice "$NAME" -p IOAccounting=true time ./prefetch run -dir foo -prefetch="$1" -hit="$2"
    mkdir -p "$NAME"
    cp -r "/sys/fs/cgroup/$NAME.slice/." "$NAME" 2>/dev/null || true
    cat "$NAME/io.stat"
    echo ===
}

rm -rf ./run_*/
run false false
run true false
run false true
run true true
