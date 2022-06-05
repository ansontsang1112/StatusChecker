# Internet Status Checker
***
A simple Go based online services checker.

## How to achieve the status check and what it is?
***
We check the service status when HTTP request comes and fetch the services at the devices you've configurated to see if the devices is online. 
The result will finally respond as JSON. You can simply access though any API fetching technology.

## Deployment
***
This application can simply deploy on either **PM2** or **Docker** platform.

_Details Belows:_
* Go Version: 1.18rc1

## Current Features
***
* "Ping" response

## Future Update
* Support Schedule and save status data at redis / sqlite
* Support Discord Webhook notification function
* Support "cli" based edit on "pingDevices.json"

## Configurations
***
### config.json
```json
{
  "httpConfig": {
    "port": 8080
  }
}
```
### pingDevices.json
```json
{
  "devices": [
    {
      "id": "cfdns",
      "host": "1.1.1.1",
      "port": 53,
      "protocol": "udp",
      "description": "Cloudflare DNS"
    }
  ]
}
```

## Responses
***
```json
{
  "cfdns":{
    "Ping":{
      "id":"cfdns",
      "host":"1.1.1.1",
      "port":53,
      "protocol":"udp",
      "description":"Cloudflare DNS"
    },
    "Status":true,
    "Err":"N/A"
  }
}
```
