# A Client for Bypassing Cloudflare

cf_forbidden is a Go package designed to make HTTP requests to websites protected by Cloudflare.

## Features

- Make HTTP requests to Cloudflare-protected websites.
- Set custom headers for each request.
- Configure UserAgent and JA3 fingerprinting.

## Installation

```bash
go get -u github.com/iarsham/cf_forbidden
```

## Usage

### 1. Create a Client:

```go
package main

import (
    cf "github.com/iarsham/cf_forbidden"
)

func main() {
    client, err := cf.New()
    if err != nil {
        // handle error
    }
}
```
### 2. Set Custom Headers (Optional):
```go
client.headers = cf.M{
    "X-Custom-Header": "value",
}
```


### 3. Set UserAgent and JA3 Fingerprint (Optional):
```go
client.agent = "My Custom User Agent"
client.ja3 = "your_ja3_fingerprint"
```


### 4. Make Requests:

#### GET request:

```go
response, err := client.Get("https://target.com", cf.M{"Authorization": "Bearer your_token"})
if err != nil {
    // handle error
}
// Access response body and headers
body := response.Body
headers := response.Headers
```

#### POST request:

```go
response, err := client.Post("https://target.com/api/endpoint", cf.M{"Content-Type": "application/json"}, {"data": "your_data"} )
if err != nil {
// handle error
}

// Access response body and headers
body := response.Body
headers := response.Headers
```

**Note:** Refer to the documentation of the CycleTLS library [here](https://github.com/Danny-Dasilva/CycleTLS) for more
advanced configuration options.

## Contributing

We welcome contributions to this project! Please see the `CONTRIBUTING.md` file (if available) for details.

## License

This project is licensed under the MIT License.

**Disclaimer:** This project is provided for educational purposes only. Use it responsibly and ethically. The authors
are not responsible for any misuse of this tool.
You can copy this Markdown text into a .md file for your project. Let me know if you need any further modifications!