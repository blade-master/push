#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/pushcore"
exec "$CURDIR/bin/pushcore"
