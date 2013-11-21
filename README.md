# btcmargin is a command line utility that shows you your bitcoin returns

Utilizes the MtGox streaming API


## Installation

btcmargin is written in go so you need install go if you haven't already.

```sh
$ go get github.com/tmc/btcmargin
```

## Running

You need an API key from MtGox to run btcmargin. It's best if this key only has Get Info privileges. The keys (and amounts) shown are fake.

```sh
$ btcmargin -value=5.24 -amount=100 -key=80ed07b3-8cfc-5b27-a4dg-dbf2cef7d1da -secret=j6jfbMc2blT2x14QfgmZKUezFmEKbXIJXAJn5Hqz/lNWmAwWAUOP+O/mNYFph3IN7NFJ/XRmLkXEywjk0rcExA==
$74476.00 - 14212.98% - currently trading at $750.000000 USD (spent:$524.00, made:$75000.00)
$74476.00 - 14212.98% - currently trading at $750.000000 USD (spent:$524.00, made:$75000.00)
$74201.00 - 14160.50% - currently trading at $747.250020 USD (spent:$524.00, made:$74725.00)
$74201.00 - 14160.50% - currently trading at $747.250020 USD (spent:$524.00, made:$74725.00)
$74201.00 - 14160.50% - currently trading at $747.250030 USD (spent:$524.00, made:$74725.00)
$74726.00 - 14260.69% - currently trading at $752.500020 USD (spent:$524.00, made:$75250.00)
```

License: ISC

