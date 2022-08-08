# DSRP
Dead Simple Reverse Proxy.

## Usage
```SH
$ dsrp -h
Usage of dsrp:
  -crt string
        [REQ] certificate file
  -if string
        [OPT] the interface to listen on (default "127.0.0.1")
  -key string
        [REQ] key file
  -port uint
        [OPT] the port to listen on (default 443)
  -url string
        [REQ] the target url to forward requests to

dsrp -crt ./server.crt -key ./server.key -port 8443 -url http://localhost:9876
```

## Create a self signed certificate
There are probably better ways, but it does the trick for a self
hosted nextcloud deployment.

```SH
OUT="$(ip route | grep default | cut -d ' ' -f 3)"
PORT='443'
openssl req -x509 -newkey rsa:4096 -sha256 -days 3650 -nodes \
  -keyout server.key -out server.crt -subj "/CN=$OUT:$PORT" \
  -addext "subjectAltName=DNS:$OUT,DNS:localhost,IP:127.0.0.1"
```
