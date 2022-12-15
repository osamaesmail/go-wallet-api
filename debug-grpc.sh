#!/bin/sh
reflex -sr '\.go$' -- sh -c 'if pgrep dlv; then pkill dlv; fi && dlv debug  --headless --listen=:5001 --accept-multiclient --api-version=2 --log --continue --output=grpc-dev main.go -- grpc'