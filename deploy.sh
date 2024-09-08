#!/bin/sh

. ./env.sh


# This script is used to deploy the binary file to a server and
# restart the corresponding service.
#
# Args:
#   None

# Build the binary
GOOS=linux GOARCH=amd64 go build -o bin/main cmd/*.go

# Get password from rsync_pass
sshpass -f ./rsync_pass

# Deploy the binary file from local machine to the server
rsync --archive --compress --delete bin/main ${SSH_URL}:${SERVICE_PATH}/bin/

# Restart the service on the server
ssh "${SSH_URL}" "
    sudo systemctl restart ${SERVICE_NAME}
"

# Inform the user that the script has finished executing
if [ $? -eq 0 ]; then
    echo "Binary has been deployed successfully."
else
    echo "Binary deployment failed."
    exit 1
fi
