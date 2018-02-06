# consul-dns
Authoritative DNS server that plug into consul

DNS shouldn't be hard, as a former sysadmin i like the convenience of RFC1035 zone files (eg.: the ones bind9 uses), so the base idea is use the standards that are already in place.

Go is the choosen language, principally for this fantastic [dns libray](https://github.com/miekg/dns)

In addition, supporting RFC2136+RFC2845 make the the system pluggable to [terraform](https://www.terraform.io/docs/providers/dns/index.html) or old school nsupdate

NOTE: I am still new to go and i need some time to figure some things out.

## Usage
1. Install default policy for anonymous token on consul ACL (to allow dns queries)
```
node "" {
  policy = "read"
}

service "" {
  policy = "read"
}
```

2. Clone the repo in you gopath

3. Install dependencies
```
go get ./...
```

4. Build the binary
```
go build cmd/consul-dns/main.go
```

5. Run
./consul-dns -c config.json

Example config:
```
{
  "ConsulConfig": {
    "AuthToken": "anonymous",
    "HttpAddress": "http://127.0.0.1:8500",
    "DnsAddress": "127.0.0.1:8600"
  },
  "SystemConfig": {
    "DnsPort": "53",
    "HttpPort": "4367",
    "TsigKey": "v41HAYWrgX88krtc7x/X1Q=="
  }
}
```

You can also use env variables, look in the config package

6. Add your bind zone files under a "dns" folder in consul KV, remember to name the sub-keys in fqdn format (eg.: with the dot at the end)

## Modules description
### QUERY (3 scenarios)
1. The tld is .consul

    The request is proxied as-is to the DNS interface of consul
2. The request contains .node. or .service.

    The request is masquerated as .consul TLD, sent to consul DNS interface and the responses are converted to original TLD
3. The requested zone (or a higher one) is present in the K/V
    The zone is parsed and if a matching record is found (or a wildcard) it's returned in the resposes (CNAME also resolves the associated A record to the google servers)

### UPDATE
1. If the domain has a tsig keys, verify it
2. Update the zone accordingly

## Feature list
| Name                                 | Status | Module |
| ------------------------------------ |:------:| :-----:|
| Query masquerading                   | alpha,untested    | QUERY  |
| RFC1035 zone parsing from consul K/V | alpha,untested    | QUERY  |
| RFC2136 zone update                  | alpha,untested    | UPDATE |
| RFC2845 tsig verification            | alpha,untested    | UPDATE |
| DDNS like update                     | todo   | REST   |
| TSIG key generation                  | alpha,untested   | REST   |
| query caching                        | todo   | --     |
| zone formatter                       | todo   | --     |
| phrometeus metrics                   | todo   | --     |

## Visual rappresentation
![consul-dns diagram](https://github.com/ennetech/consul-dns/raw/master/docs/diagram.png "consul-dns")