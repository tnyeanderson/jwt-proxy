#!/bin/bash

# Start the normal nginx process
/docker-entrypoint.sh nginx -g 'daemon off;' &

# Start the jwt-proxy process
/app/jwt-proxy &

# Wait for any process to exit
wait -n

# Exit with status of process that exited first
exit $?
