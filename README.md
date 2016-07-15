# kayvee-logger-service
A service for recording client-side events and errors via the Clever logging pipeline.

*NOTE: Currently in development.*

## Clients

### Python

#### Installation:

##### pip

Assuming kayvee-logger-service v1.0.0 is being installed:

```sh
pip install git+https://github.com/Clever/kayvee-logger-service.git@v1.0.0
```

##### setup.py

```python
from setuptools import setup

# Assuming kayvee-logger-service v1.0.0 is being installed:
setup(

    # ...

    install_requires=['kayvee-logger-service==1.0.0'],
    dependency_links=[
      'https://github.com/Clever/kayvee-logger-service/tarball/v1.0.0#egg=kayvee-logger-service-1.0.0'
    ],

    # ...

)
```

#### Usage:

```python
import kayvee_logger_service.output as kv_output
import logger

kv_logger = logger.Logger("kayvee-logger-service-test", output=kv_output.Output())
kv_logger.info("sample-log-title", dict(msg="This will get logged to the kayvee-logger-service."))
```

#### Environment Variables:

The following environment variables are required to enable `kayvee-logger-service` discovery:

```
SERVICE_KAYVEE_LOGGER_SERVICE_HTTP_HOST
SERVICE_KAYVEE_LOGGER_SERVICE_HTTP_PORT
```

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

`kayvee-logger-service` is deployed using ark:

```sh
ark start kayvee-logger-service -e system
```

You can also deploy a new version of the service by running the `deploy` Make target:

```sh
make deploy
```
