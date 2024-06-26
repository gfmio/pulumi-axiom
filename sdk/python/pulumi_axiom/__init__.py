# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .dataset import *
from .get_dataset import *
from .get_monitor import *
from .get_notifier import *
from .get_user import *
from .monitor import *
from .notifier import *
from .provider import *
from .user import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_axiom.config as __config
    config = __config
else:
    config = _utilities.lazy_import('pulumi_axiom.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "axiom",
  "mod": "index/dataset",
  "fqn": "pulumi_axiom",
  "classes": {
   "axiom:index/dataset:Dataset": "Dataset"
  }
 },
 {
  "pkg": "axiom",
  "mod": "index/monitor",
  "fqn": "pulumi_axiom",
  "classes": {
   "axiom:index/monitor:Monitor": "Monitor"
  }
 },
 {
  "pkg": "axiom",
  "mod": "index/notifier",
  "fqn": "pulumi_axiom",
  "classes": {
   "axiom:index/notifier:Notifier": "Notifier"
  }
 },
 {
  "pkg": "axiom",
  "mod": "index/user",
  "fqn": "pulumi_axiom",
  "classes": {
   "axiom:index/user:User": "User"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "axiom",
  "token": "pulumi:providers:axiom",
  "fqn": "pulumi_axiom",
  "class": "Provider"
 }
]
"""
)
