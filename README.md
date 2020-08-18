![logo](https://raw.githubusercontent.com/sibis/producthunt-cli/master/logo.png "Product Hunt")
## **Product hunt CLI**

A CLI application to access the current trending products and view its information, without the need to open the browser. If you happen to use the producthunt everyday and this will help you fetch the product lists on your CLI. 

## Installation 

#### macOS
```bash
 brew install sibis/ph/ph
```

### Debian/Ubuntu Linux
Install and upgrade:

1. Download the `.deb` file from the [releases page][];
2. Install the downloaded file: `sudo apt install ./ph_*_amd64.deb`

```bash
 wget https://github.com/sibis/producthunt-cli/releases/download/v0.1.0/ph_0.1.0_amd64.deb
 sudo apt install ./ph_0.1.0_amd64.deb
```

### Usage
- To authenticate Producthunt from your account
```bash
ph signin
```

- Fetch the current trending products
```bash
ph list
```


[releases page]: https://github.com/sibis/producthunt-cli/releases/latest
