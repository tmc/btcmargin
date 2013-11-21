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
$ btcmargin -key=80ed07b3-8cfc-5b27-a4dg-dbf2cef7d1da -secret=j6jfbMc2blT2x14QfgmZKUezFmEKbXIJXAJn5Hqz/lNWmAwWAUOP+O/mNYFph3IN7NFJ/XRmLkXEywjk0rcExA== -value=0.0048 -amount=5000
Made $3718104.45 15492101.88% - currently at $743.625690 USD (spent:$24.00, made:$3718128.45)
Made $3718104.45 15492101.88% - currently at $743.625690 USD (spent:$24.00, made:$3718128.45)
```

License: ISC

