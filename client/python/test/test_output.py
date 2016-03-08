import json
import logger
import unittest

from kayvee_logger_service.output import Output
from mock import Mock


class TestOutput(unittest.TestCase):

  def setUp(self):
    self.mock_api_client = Mock()

  def test_write(self):
    error = 'An unknown error occurred.'

    kv_logger = logger.Logger('kv-service-output-test', output=Output(self.mock_api_client))
    kv_logger.error('test-error', data=dict(msg=error))

    call_args = self.mock_api_client.log.call_args
    self.assertIsNotNone(call_args)

    expected = dict(source='kv-service-output-test', level='error', title='test-error', msg=error)
    actual = json.loads(call_args[0][0].strip())
    self.assertEqual(actual, expected)
