SHELL := /bin/bash

install:
	go install .\cmd\veil.go
test:
	veil get
	veil get something
	veil set key_without_value
	veil set twitter_api_key 12345678910
	veil get twitter_api_key