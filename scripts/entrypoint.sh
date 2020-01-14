#!/usr/bin/env sh
WORK_DIR=${WORK_DIR:-$GOPATH/src}
cd "${WORK_DIR}" || exit

exec "$@"
