# alfred - Deployment Assistant
alfred is able to facilitiate deployments. He does not like the capital A.

## Getting Started

### First time setup
```sh
make install
make setup
# configure .alfred/config.yml to your liking
```

### Configuring and deploying
```sh
# Replace {{ENV}} with the environment you are
# building a deploy config for (qa, for example)
ENV={{ENV}} make repos
# configure .alfred/{{ENV}}.yml to your liking

# Confirm the output is to your liking
ENV={{ENV}} make verify

# Create release
ENV={{ENV}} make release

## Development and Testing
```sh
make test
```