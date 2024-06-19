#!/usr/bin/env bash

set -euxo pipefail

time ./prefetch run -dir foo -prefetch=false -hit=false | grep Mapped:
time ./prefetch run -dir foo -prefetch=true -hit=false | grep Mapped:
time ./prefetch run -dir foo -prefetch=false -hit=true | grep Mapped:
time ./prefetch run -dir foo -prefetch=true -hit=true | grep Mapped: