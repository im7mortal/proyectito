#!/bin/bash

tmp_tfplan="tfplan.local"
flags="--terragrunt-working-dir=envs/dev/aws-secret-manager"

terragrunt plan "${flags}" -out="${tmp_tfplan}" &&
terragrunt apply "${flags}" "${tmp_tfplan}"
