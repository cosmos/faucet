## Deployment Makefile for AWS networks

# Binary for the faucet
BINARY ?= ../faucet

# Frontend tar.gz-d
DISTTGZ ?= ../frontend/dist.tgz

# Testnet name used in --chain-id
TESTNET_NAME ?= gaia-6002

# Application name used in AWS. Use only alphanumerical characters.
APP_NAME ?= gaia6002

# Recaptcha secret key is mandatory
ifndef SECRETKEY
$(error Please provide Recaptcha SECRETKEY in environment variables)
endif

# gaiacli binary path
ifndef GAIACLIBINARY
$(error GAIACLIBINARY not set)
endif

# Seedwords for the key to use
ifndef SEEDWORDS
$(error SEEDWORDS not set)
endif

build:
	@echo "Start with make app-start-aws in the cosmos-sdk networks folder."
	cd remote/ansible && ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook -i inventory/ec2.py -l "tag_Environment_$(APP_NAME)" -u centos -b -e BINARY=$(BINARY) -e GAIACLIBINARY="$(GAIACLIBINARY)" -e DISTTGZ="$(DISTTGZ)" -e SECRETKEY="$(SECRETKEY)" -e SEEDWORDS="$(SEEDWORDS)" -e TESTNET_NAME="$(SEEDWORDS)" deploy-faucet.yml

