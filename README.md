# godeclutter
Declutters URLs in a lightning fast and flexible way, for improving input for web hacking automations such as crawlers and vulnerability scans. 

godeclutter is a very simple tool that will take a list of URLs, clean those URLs, and output the unique URLs that are present. This will reduce the number of requests you will have to make to your target website, and also filter URLs that are most likely uninteresting. 

# Features
godeclutter will perform the following steps on your URL host section:
- Clean http:// URLs pointing to the default SSL port (443) and vice-versa, since they are mostly CDN error pages.
- Clean port notation of URLs pointing to the default protocol ports, since those ports are already implied by the protocol scheme. (such as :443 and :80)
- Clean http:// URLs if a https:// to the same host and port is present, since 99,9% of those cases will just be a redirect to https://.
- Remove uninteresting media extensions such as png, jpg, css, etc. (This one will keep .js files since those are sometimes interesting)
- Sort query parameters
- Lowercase all schemes and hostnames, since upper-casing is irrelevant for those.
- Replace all lower-case URI encoding escapes to upper-case, to maintain a standard.
- Decode unnecessary escapes for characters that are not special in the URL context (i.e http://example.com/%41 ).
- Remove empty query strings (i.e http://example.com/? )
- Remove trailing slashes, this is rather aggressive but filters a lot of un-interesting duplicates on the majority of the cases. (i.e http://host/path/ -> http://host/path)
- Normalize dot segments, also rather aggressive but useful when working with dirty sources. (http://host/path/./a/b/../c -> http://host/path/a/c)


# Install

```bash
go install github.com/c3l3si4n/godeclutter@HEAD
```

# Basic Usage
You can send URLs by sending them to stdin.
```bash
arch ~>  cat test_urls.txt
https://1.1.1.1:443/
http://1.1.1.1/
http://1.1.1.1:80/
https://1.1.1.1:443/
https://1.1.1.1:80/
https://1.1.1.1:8443/
https://1.1.1.1:8443/
https://1.1.1.1:443/
https://1.1.1.1:8443/
https://1.1.1.1:8443/
https://1.1.1.1:443/?
http://1.1.1.1:443/?
https://1.1.1.1:80/?a
https://1.1.1.1:443/?
https://1.1.1.1:443/?
https://1.1.1.1:443/?1=1
https://1.1.1.1:443/?a=a&b=1
https://1.1.1.1:443/?a=a&b=1
https://1.1.1.1:443/a.js?a=a&b=1
https://1.1.1.1:443/fiqef.html?a=a&b=1
https://1.1.1.1:443/fmef.jpg?b=1&a=a
https://1.1.1.1:443/?b=1&a=a
https://1.1.1.1:443/a.js?
https://1.1.1.1:443/a.jpg
https://1.1.1.1:8443/
https://1.1.1.1:443/
arch ~>  cat test_urls.txt | godeclutter -b -c -p
https://1.1.1.1/
https://1.1.1.1:8443/
https://1.1.1.1/?1=1
https://1.1.1.1/?a=a&b=1
https://1.1.1.1/a.js?a=a&b=1
https://1.1.1.1/fiqef.html?a=a&b=1
https://1.1.1.1/a.js

arch ~> 
arch ~> 
```

# Arguments
```bash
$> ./godeclutter -h
Usage of ./godeclutter:
  -b	Blacklist Extensions - clean some uninteresting extensions. (default true)
  -c	Clean URLs - Aggressively clean/normalize URLs before outputting them.
  -p	Prefer HTTPS - If there's a https url present, don't print the http for it. (since it will probably just redirect to https)
```

# Aknowledgements
- **[@s0md3v](https://github.com/s0md3v)** for making **[uro](https://github.com/s0md3v/uro)**
- **[@PuerkitoBio](https://github.com/PuerkitoBio)** for making the amazing **[purell](https://github.com/PuerkitoBio/purell)** go library
- **[@ameenmaali](https://github.com/ameenmaali)** for making urldedupe which was one of the [first tools](https://github.com/ameenmaali/urldedupe) for doing that kind of work.
