## Invite-list

This is a tiny cmd for my **Intercom** application.

### Install and run
Install the binary
```bash
$ make install
```
check if it is installed
```bash
$ customers -h
Usage:
  customers [OPTIONS]

Application Options:
      --office=              Office location {SF, Chicago, SurryHills, Dublin, London} (default: Dublin)
  -d, --max-invite-distance= The largest distance (km) to invite customer to the office (default: 100)
  -i, --input=               File with customer records (default: all-customers.txt)
  -o, --output=              Output file
      --loglevel=            Logging level for all subsystems {trace, debug, info, error, critical} (default: info)

Help Options:
  -h, --help                 Show this help message

```
and run it
```bash
$ customers
2019-03-10 02:02:58.134 [INF] CUST: reading customers
2019-03-10 02:02:58.135 [INF] CUST: inviting customers
2019-03-10 02:02:58.135 [INF] CUST: sorting customers
2019-03-10 02:02:58.135 [INF] CUST: writing customers


        ID: 4           Name: Ian Kehoe
        ID: 5           Name: Nora Dempsey
        ID: 6           Name: Theresa Enright
        ID: 8           Name: Eoin Ahearn
        ID: 11          Name: Richard Finnegan
        ID: 12          Name: Christina McArdle
        ID: 13          Name: Olive Ahearn
        ID: 15          Name: Michael Ahearn
        ID: 17          Name: Patricia Cahill
        ID: 23          Name: Eoin Gallagher
        ID: 24          Name: Rose Enright
        ID: 26          Name: Stephen McArdle
        ID: 29          Name: Oliver Ahearn
        ID: 30          Name: Nick Enright
        ID: 31          Name: Alan Behan
        ID: 39          Name: Lisa Ahearn


```

### Greetings
Hello team)

### Approach
I decided to not overthink this task, so the code is simple and
not optimized.<br />
Optimization makes sense only if there are constraints such as:
- number of customers
- runtime limit
- memory limit
- output format
- reading input from a network
- compressed input file
- many input files

### Stuff to add
- CI/CD scripts
- Benchmarks
- Platform & architecture supporting releases
