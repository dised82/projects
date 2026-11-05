import sys
import json
import re
import urllib.request

"""
this script is build to be called by a go script to implement the mit 6.824 webcrawlwer.
it will bw called by a go script who will manage the ouput data
"""

def fetch_links(url):
    with urllib.request.urlopen(url, timeout=10) as resp:
        html = resp.read().decode("utf-8", errors="ignore")
    # naive extraction, use BeautifulSoup if you want something robust
    return re.findall(r'href=["\'](https?://[^"\']+)["\']', html)

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("usage: crawler.py <url>", file=sys.stderr)
        sys.exit(1)
    links = fetch_links(sys.argv[1])
    print(json.dumps(links))   # ONLY the JSON goes to stdout
