# XRepository
Provides SDKs to interact with the XRepository API.

## Usage
Not published yet.

## Prerequisites
* Go
* make
* go-swagger

## Installation
* `git clone git@github.com:diggi-foundation/xrepository.git`
* `cd xrepsoitory`
* `make pull`

## Serve UI
* `make serve`

## Generate Client
* `make generate`

## Limitations
This repository just downloads the openapi spec from https://www.xrepository.de/api/swagger and generates the respective clients. The downloaded spec is written in swagger 2.0 and has several documentation as well as validation issues.

## Features
- [x] downloads openapi spec via command
- [x] generates go client
- [ ] replaces xrep-wildfly02.init.de/192.168.212.84 with www.xrepository.de/api
- [ ] publishes itself via Github Actions
- [ ] uses semantic release
- [ ] fixes error with x_ö_v_bibliothek 
