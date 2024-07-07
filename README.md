# ssh-akc-github-auth
The software retrieves ssh pub keys from the GitHub Org team.

# Usage
```bash
sudo -i
curl -sLO https://github.com/seyukun/ssh-akc-github-auth/releases/download/latest/ssh-akc-github-auth
chmod +x ssh-akc-github-auth
install ssh-akc-github-auth /usr/local/bin/
rm ssh-akc-github-auth
mkdir -p /usr/local/bin/authorizedkeyscommand

# PER USER
cat <<EOF > /usr/local/bin/authorizedkeyscommand/<LOGIN>
/usr/local/bin/ssh-akc-github-auth --token <TOKEN HERE> --org <ORG HERE> --team <TEAM HERE>
EOF

chmod 100 /usr/local/bin/authorizedkeyscommand/*
exit
```

edit `/etc/ssh/sshd_config`
```
AuthorizedKeysCommand /bin/sh -c '/usr/local/bin/authorizedkeyscommand/%u 2> /dev/null'
AuthorizedKeysCommandUser root
```
