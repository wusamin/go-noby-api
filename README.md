# go-noby-api
go-noby-api is wrapper library "noby api(cotogo)".
> https://www.cotogoto.ai/

## Install
```
go get -u github.com/wusamin/go-noby-api
```

## Usage
```
// initialize NobyRequest and set parameter.
req := NobyRequest{}
req.Appkey = "api key"
req.Text = "Hello"
req.Persona = Normal

// "Call" returns API response. It is a bound struct.
res, err := Call(&req)
```
