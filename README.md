# nmap2json

```
Usage: nmap2json [OPTION]... [FILE]...
Convert nmap XML FILE(s) to JSON

With no FILE, or when FILE is -, read standard input.

  -o file
        Write output to file instead of standard output
  -p    Pretty-print JSON output

Examples:
  nmap -sn -oX network_data.xml 192.168.2.0/24 && nmap2json network_data.xml
  nmap -sn -oX - 192.168.2.0/24 | nmap2json
```
