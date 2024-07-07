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
cat <<EOF > /usr/local/bin/authorizedkeyscommand/<USERNAME>
/usr/local/bin/ssh-akc-github-auth --token <TOKEN HERE> --org <ORG HERE> --team <TEAM HERE>
EOF
chmod 100 /usr/local/bin/authorizedkeyscommand/*
exit
```

edit `/etc/ssh/sshd_config`
```
AuthorizedKeysCommand bash /usr/local/bin/authorizedkeyscommand/%u
AuthorizedKeysCommandUser root
```
