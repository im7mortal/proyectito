#!/bin/bash

aws secretsmanager list-secrets | jq '.SecretList[] | .Name'
ARN1=$(aws secretsmanager list-secrets | jq '.SecretList[] | .ARN'| head -n1 | sed -e 's/"//g')
aws secretsmanager get-secret-value --secret-id $ARN1 | jq '.SecretString'
