# coding: utf-8

import os
import pkg_resources
import sys

from pip.req import parse_requirements
from setuptools import setup, find_packages

NAME = "kayvee_logger_service"
VERSION = "1.0.0"

here = os.path.abspath(os.path.dirname(__file__))

reqs = './requirements.txt'
if len(sys.argv) > 1 and sys.argv[1] in ['develop', 'test']:
  reqs = './requirements-dev.txt'

pr_kwargs = {}
if pkg_resources.get_distribution("pip").version >= '6.0':
  pr_kwargs = {"session": False}

install_reqs = parse_requirements(os.path.join(here, reqs), **pr_kwargs)

setup(
    name=NAME,
    version=VERSION,
    description="Kayvee Logger Service",
    author_email="tech-notify@getclever.com",
    url="https://github.com/Clever/kayvee-logger-service",
    keywords=["Swagger", "Kayvee Logger Service"],
    install_requires=[str(ir.req) for ir in install_reqs],
    packages=find_packages(),
    long_description="""\
    Logs kayvee events.
    """
)
