#!/bin/sh
reflex -sr '\.go$' -- sh -c 'if pgrep dlv; then pkill dlv; fi && dlv debug  --headless --listen=:5002 --accept-multiclient --api-version=2 --log --continue --output=dev main.go -- grpc'