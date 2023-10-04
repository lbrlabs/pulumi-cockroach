# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetCockroachClusterResult',
    'AwaitableGetCockroachClusterResult',
    'get_cockroach_cluster',
    'get_cockroach_cluster_output',
]

@pulumi.output_type
class GetCockroachClusterResult:
    """
    A collection of values returned by getCockroachCluster.
    """
    def __init__(__self__, account_id=None, cloud_provider=None, cockroach_version=None, creator_id=None, dedicated=None, id=None, name=None, operation_status=None, parent_id=None, plan=None, regions=None, serverless=None, state=None, upgrade_status=None):
        if account_id and not isinstance(account_id, str):
            raise TypeError("Expected argument 'account_id' to be a str")
        pulumi.set(__self__, "account_id", account_id)
        if cloud_provider and not isinstance(cloud_provider, str):
            raise TypeError("Expected argument 'cloud_provider' to be a str")
        pulumi.set(__self__, "cloud_provider", cloud_provider)
        if cockroach_version and not isinstance(cockroach_version, str):
            raise TypeError("Expected argument 'cockroach_version' to be a str")
        pulumi.set(__self__, "cockroach_version", cockroach_version)
        if creator_id and not isinstance(creator_id, str):
            raise TypeError("Expected argument 'creator_id' to be a str")
        pulumi.set(__self__, "creator_id", creator_id)
        if dedicated and not isinstance(dedicated, dict):
            raise TypeError("Expected argument 'dedicated' to be a dict")
        pulumi.set(__self__, "dedicated", dedicated)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if operation_status and not isinstance(operation_status, str):
            raise TypeError("Expected argument 'operation_status' to be a str")
        pulumi.set(__self__, "operation_status", operation_status)
        if parent_id and not isinstance(parent_id, str):
            raise TypeError("Expected argument 'parent_id' to be a str")
        pulumi.set(__self__, "parent_id", parent_id)
        if plan and not isinstance(plan, str):
            raise TypeError("Expected argument 'plan' to be a str")
        pulumi.set(__self__, "plan", plan)
        if regions and not isinstance(regions, list):
            raise TypeError("Expected argument 'regions' to be a list")
        pulumi.set(__self__, "regions", regions)
        if serverless and not isinstance(serverless, dict):
            raise TypeError("Expected argument 'serverless' to be a dict")
        pulumi.set(__self__, "serverless", serverless)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if upgrade_status and not isinstance(upgrade_status, str):
            raise TypeError("Expected argument 'upgrade_status' to be a str")
        pulumi.set(__self__, "upgrade_status", upgrade_status)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> str:
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter(name="cloudProvider")
    def cloud_provider(self) -> str:
        return pulumi.get(self, "cloud_provider")

    @property
    @pulumi.getter(name="cockroachVersion")
    def cockroach_version(self) -> str:
        return pulumi.get(self, "cockroach_version")

    @property
    @pulumi.getter(name="creatorId")
    def creator_id(self) -> str:
        return pulumi.get(self, "creator_id")

    @property
    @pulumi.getter
    def dedicated(self) -> 'outputs.GetCockroachClusterDedicatedResult':
        return pulumi.get(self, "dedicated")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="operationStatus")
    def operation_status(self) -> str:
        return pulumi.get(self, "operation_status")

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> str:
        return pulumi.get(self, "parent_id")

    @property
    @pulumi.getter
    def plan(self) -> str:
        return pulumi.get(self, "plan")

    @property
    @pulumi.getter
    def regions(self) -> Sequence['outputs.GetCockroachClusterRegionResult']:
        return pulumi.get(self, "regions")

    @property
    @pulumi.getter
    def serverless(self) -> 'outputs.GetCockroachClusterServerlessResult':
        return pulumi.get(self, "serverless")

    @property
    @pulumi.getter
    def state(self) -> str:
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="upgradeStatus")
    def upgrade_status(self) -> str:
        return pulumi.get(self, "upgrade_status")


class AwaitableGetCockroachClusterResult(GetCockroachClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCockroachClusterResult(
            account_id=self.account_id,
            cloud_provider=self.cloud_provider,
            cockroach_version=self.cockroach_version,
            creator_id=self.creator_id,
            dedicated=self.dedicated,
            id=self.id,
            name=self.name,
            operation_status=self.operation_status,
            parent_id=self.parent_id,
            plan=self.plan,
            regions=self.regions,
            serverless=self.serverless,
            state=self.state,
            upgrade_status=self.upgrade_status)


def get_cockroach_cluster(id: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCockroachClusterResult:
    """
    CockroachDB Cloud cluster. Can be Dedicated or Serverless.
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('cockroach:index/getCockroachCluster:getCockroachCluster', __args__, opts=opts, typ=GetCockroachClusterResult).value

    return AwaitableGetCockroachClusterResult(
        account_id=__ret__.account_id,
        cloud_provider=__ret__.cloud_provider,
        cockroach_version=__ret__.cockroach_version,
        creator_id=__ret__.creator_id,
        dedicated=__ret__.dedicated,
        id=__ret__.id,
        name=__ret__.name,
        operation_status=__ret__.operation_status,
        parent_id=__ret__.parent_id,
        plan=__ret__.plan,
        regions=__ret__.regions,
        serverless=__ret__.serverless,
        state=__ret__.state,
        upgrade_status=__ret__.upgrade_status)


@_utilities.lift_output_func(get_cockroach_cluster)
def get_cockroach_cluster_output(id: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCockroachClusterResult]:
    """
    CockroachDB Cloud cluster. Can be Dedicated or Serverless.
    """
    ...
