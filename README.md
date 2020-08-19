![logo](https://raw.githubusercontent.com/sibis/producthunt-cli/master/logo.png "Product Hunt")
<a href="https://www.producthunt.com/posts/product-hunt-cli-2?utm_source=badge-featured&utm_medium=badge&utm_souce=badge-product-hunt-cli-2" target="_blank"><img src="https://api.producthunt.com/widgets/embed-image/v1/featured.svg?post_id=238170&theme=light" alt="Product Hunt CLI - Get the daily Product Hunt trending updates on your CLI | Product Hunt Embed" style="width: 250px; height: 54px;" width="250px" height="54px" /></a>

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
