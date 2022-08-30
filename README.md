# godeclutter
Declutters URLs in a flexible way, for improving input for web hacking automations such as crawlers and vulnerability scans. 


# Install
```bash
go install github.com/c3l3si4n/godeclutter@latest
```

# Basic Usage
You can send URLs by sending them to stdin.
```bash
$> cat test_urls.txt | godeclutter -b -c -p
https://1.1.1.1/
https://1.1.1.1:8443/
https://1.1.1.1/?1=1
https://1.1.1.1/?a=a&b=1
https://1.1.1.1/a.js?a=a&b=1
https://1.1.1.1/fiqef.html?a=a&b=1
https://1.1.1.1/a.js
```

# Arguments
```bash
$> ./godeclutter -h
Usage of ./godeclutter:
  -b	Blacklist Extensions - clean some uninteresting extensions. (default true)
  -c	Clean URLs - Aggressively clean/normalize URLs before outputting them.
  -p	Prefer HTTPS - If there's a https url present, don't print the http for it. (since it will probably just redirect to https)
```