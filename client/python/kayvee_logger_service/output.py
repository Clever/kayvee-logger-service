import os
import logging
logging.basicConfig(level=logging.INFO)

import discovery

from .api_client import ApiClient
from .apis.default_api import DefaultApi
from .rest import ApiException


class Output:

  """
  A Kayvee logger output that writes to the kayvee-logger-service.

  :Example:

    import kayvee_logger_service.output as kv_output
    import logger

    kv_logger = logger.Logger("kayvee-logger-service-test", output=kv_output.Output())
    kv_logger.info("sample-log-title", dict(msg="This gets logged to the kayvee-logger-service."))

  """

  def __init__(self, kv_service_default_api=None):
    """
    :param kv_service_default_api:
      an instance of kayvee_logger_service.apis.default_api.DefaultApi; uses a
      client created via discovery by default.
    """

    self.logger = logging.getLogger('kayvee-logger-service-output')

    if (kv_service_default_api is None):
      try:
        kv_service_host = discovery.host_port('kayvee-logger-service', 'http')
        kv_service_default_api = DefaultApi(ApiClient(kv_service_host))
      except discovery.MissingEnvironmentVariableError as e:
        self.logger.warning('Kayvee Logger Service discovery failed: %s', e)

    self.default_api = kv_service_default_api

  def write(self, log_string):
    if self.default_api is None:
      return

    try:
      self.default_api.log(log_string, callback=self._noOpCallback)
    except ApiException as e:
      self.logger.warning('Unable to log to kayvee-logger-service: %s', e)

  def _noOpCallback(self, response):
    """ Used to facilitate fire-and-forget logging. """
    return
