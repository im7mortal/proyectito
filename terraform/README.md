## Terraform task

### Encryption

The config `secret.json` is encrypted with `git-crypt`. I have the SOPS version in the `sops` branch.

### Test
#### Creation

```shell
./create.sh
```

#### Validation

```shell
./validate.sh
```

#### Cleaning

```shell
./destroy.sh
```