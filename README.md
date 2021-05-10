# qrcode

Create a QRCode

## Build
- Build a binary into `bin/`
```shell
make build
```

- Build multi platform binaries
```shell
make xbuild
```

## Usage
```
./bin/qrcode-generator -o one.png -s 240 -c https://google.com -l high
```

### flags
| flag | description |
| --- | --- |
| -content, -c | QRCode content |
| -output, -o | The file name of output |
| -size, -s | The frame size of the output QRCode |
| -level, -l | The recovery level of the output QRCode |
