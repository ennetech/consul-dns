# consul-dns
Authoritative DNS server that plug into consul

DNS shouldn't be hard, as a former sysadmin i like the convenience of RFC1035 zone files (eg.: the ones bind9 uses), so the base idea is use the standards that are already in place.

Go is the choosen language, principally for this fantastic [dns libray](https://github.com/miekg/dns), i will start publishing the code as soon it takes an human readable shape (spaghetty POC anyone?), i am still new to go and i need some time to figure some things out.

In addition, supporting RFC2136+RFC2845 make the the system pluggable to [terraform](https://www.terraform.io/docs/providers/dns/index.html) or old school nsupdate

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
| Query masquerading                   | POC    | QUERY  |
| RFC1035 zone parsing from consul K/V | POC    | QUERY  |
| RFC2136 zone update                  | POC    | UPDATE |
| RFC2845 tsig verification            | POC    | UPDATE |
| DDNS like update                     | todo   | REST   |
| TSIG key generation                  | todo   | REST   |
| query caching                        | todo   | --     |
| zone formatter                       | todo   | --     |
| phrometeus metrics                   | todo   | --     |

## Visual rappresentation
![consul-dns diagram](https://github.com/ennetech/consul-dns/raw/master/docs/diagram.png "consul-dns")