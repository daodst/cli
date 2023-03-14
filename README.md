# It mainly includes wallet creation and viewing functions

## Compile command
`
go mod tidy
`

`go build
`

## method

### 1.create wallet
`fm-cli create`

### 2.View mnemonics
`fm-cli import --seed "absurd mixture rural brick snow stumble gadget walk noodle pause like run"`


### 3.View privateKey
`fm-cli import --pk 12fc22d015c4f7e015eb15668e4c35a6303a987b0cba53c00d5c2c357443c94f`

### 4.transfer
`fm-cli transfer 0B9A99EDA4B4C12323C4C1EB1210E98E8FE22FE1AE3543CACD044825CC9A793B  dex10436g09cfa2nz2k6yphzj9qs9u93qtvvkkftur  dex1va8aaeystat4twpy70ns7235pxwczg0698twv4  100tt,100fm`

### 5.query tx
`fm-cli tx 7CCF00A97EACBBF73D72E2F9BD5D5CA13F63761753C535170E991FBA5E0C1E67 --node tcp://3.35.42.193:26657`

### 6.query balance
`fm-cli balance  dex1cnty9am46c97s8w9e54xaxztl9g68u4f8njjeu --node tcp://3.35.42.193:26657`