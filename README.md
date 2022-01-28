# GoSub
Subdomain finder based on crt.sh

This will simply make an HTTP request to crt.sh passing the URL you provided, parse the results and print the subdomains it found.

This is my first non-useless code in go. I just started learning, so it might not be the best =)

# Usage 

`Gosub <url>`

```
$ GoSub tesla.com
[+] Subdomains found:
[...]
owerhub.energy.tesla.com
autobidder-preprd.powerhub.energy.tesla.com
autobidder.powerhub.energy.tesla.com
autodiscover.tesla.com
avatars.github-it.tesla.com
awsbtest.tesla.com
[...]
factory-berlin.tesla.com
feedback.tesla.com
forums.tesla.com
gist.github-it.tesla.com
github-it.tesla.com
github.tesla.com
gridlogic-eng.energy.tesla.com
gridlogic-eng.powerhub.energy.tesla.com
gridlogic.energy.tesla.com
gridlogic.powerhub.energy.tesla.com
manager.courses.tesla.com
[...]
smarttax.tesla.com
solarbonds.tesla.com
sso-dec.tesla.com
sso-dev.tesla.com
sso-sig.tesla.com
sso.tesla.com
toolbox.tesla.com
tx.tesla.com
ug.tesla.com
uploads.github-it.tesla.com
view.emails.tesla.com
vpn1.tesla.com
vpn2.tesla.com
[...]
```

# Install
```bash
go get github.com/Nickguitar/GoSub
```

or

```bash
git clone https://github.com/Nickguitar/GoSub/
cd GoSub
go install
```
