# kayvee-logger-service
A service for recording client-side events and errors via the Clever logging pipeline.

*NOTE: Currently in development.*

## Clients

### Python

#### Installation:

The auto-generated Python clinet can be installed via pip.

Add the following to your requirements.txt file:
```
git+https://<clever-drone Github token>@github.com/Clever/kayvee-logger-service.git@<version_tag>#egg=kayvee-logger-service&subdirectory=client/python
```

The Github token can be found in [dev-passwords](https://github.com/Clever/clever-ops/tree/master/credentials).

#### Usage:

TODO

## Development

### Making API definition changes

- Make any desired API changes in `kayvee-logger-service.yaml`.
- To update the auto-generated server and client code with your changes:

    ```sh
    make codegen
    ```

#### Requirements

##### go-swagger

`go-swagger` is required to generate swagger server code.
To install or update:
```sh
go get -u github.com/go-swagger/go-swagger/cmd/swagger
```

##### Java Runtime Environment
v1.7 or greater of the Java Runtime Environment is required to generate swagger client code.

First, check your current version of the JRE:
```sh
java -version
```

If there is no appropriate version of JRE installed, you can install v1.7 by running:
```sh
sudo apt-get install openjdk-7-jre
```

### Adding a new client

TODO

## Deployment

TODO
