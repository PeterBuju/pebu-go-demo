# Overview

This application is a small Go application demo.

It hosts an API endpoint that returns the result from the [Maxmind GeoIp demo](https://www.maxmind.com/en/geoip-demo) that looks up particular IPs.

## Layout

Project layout based on these [Golang standards project](https://github.com/golang-standards/project-layout/tree/master).

## Getting started
See the [Getting Started](./wiki/getting-started/getting-started.md) section to run the application.

## Swagger
The API specs are defined in Swagger. To update the Swagger docs run:
```
swag init --generalInfo cmd/api/main.go
```
