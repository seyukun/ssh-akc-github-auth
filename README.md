# ssh-akc-github-auth
The software retrieves ssh pub keys from the GitHub Org team.

# Usage
edit `/etc/ssh/sshd_config`
```
AuthorizedKeysCommand /usr/local/bin/ssh-akc-github-auth --token <TOKEN HERE> --org <ORG HERE> --team <TEAM HERE>
AuthorizedKeysCommandUser root
```
