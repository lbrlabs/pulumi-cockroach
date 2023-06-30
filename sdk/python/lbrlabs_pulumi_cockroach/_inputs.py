# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'ClusterDedicatedArgs',
    'ClusterRegionArgs',
    'ClusterServerlessArgs',
    'ClusterServerlessUsageLimitsArgs',
    'CmekAdditionalRegionArgs',
    'CmekRegionArgs',
    'CmekRegionKeyArgs',
    'LogExportConfigGroupArgs',
    'PrivateEndpointServicesServiceArgs',
    'PrivateEndpointServicesServiceAwsArgs',
    'UserRoleGrantsRoleArgs',
]

@pulumi.input_type
class ClusterDedicatedArgs:
    def __init__(__self__, *,
                 disk_iops: Optional[pulumi.Input[int]] = None,
                 machine_type: Optional[pulumi.Input[str]] = None,
                 memory_gib: Optional[pulumi.Input[float]] = None,
                 num_virtual_cpus: Optional[pulumi.Input[int]] = None,
                 private_network_visibility: Optional[pulumi.Input[bool]] = None,
                 storage_gib: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[bool] private_network_visibility: Set to true to assign private IP addresses to nodes. Required for CMEK and other advanced networking features.
        """
        if disk_iops is not None:
            pulumi.set(__self__, "disk_iops", disk_iops)
        if machine_type is not None:
            pulumi.set(__self__, "machine_type", machine_type)
        if memory_gib is not None:
            pulumi.set(__self__, "memory_gib", memory_gib)
        if num_virtual_cpus is not None:
            pulumi.set(__self__, "num_virtual_cpus", num_virtual_cpus)
        if private_network_visibility is not None:
            pulumi.set(__self__, "private_network_visibility", private_network_visibility)
        if storage_gib is not None:
            pulumi.set(__self__, "storage_gib", storage_gib)

    @property
    @pulumi.getter(name="diskIops")
    def disk_iops(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "disk_iops")

    @disk_iops.setter
    def disk_iops(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "disk_iops", value)

    @property
    @pulumi.getter(name="machineType")
    def machine_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "machine_type")

    @machine_type.setter
    def machine_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "machine_type", value)

    @property
    @pulumi.getter(name="memoryGib")
    def memory_gib(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "memory_gib")

    @memory_gib.setter
    def memory_gib(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "memory_gib", value)

    @property
    @pulumi.getter(name="numVirtualCpus")
    def num_virtual_cpus(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "num_virtual_cpus")

    @num_virtual_cpus.setter
    def num_virtual_cpus(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "num_virtual_cpus", value)

    @property
    @pulumi.getter(name="privateNetworkVisibility")
    def private_network_visibility(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to true to assign private IP addresses to nodes. Required for CMEK and other advanced networking features.
        """
        return pulumi.get(self, "private_network_visibility")

    @private_network_visibility.setter
    def private_network_visibility(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "private_network_visibility", value)

    @property
    @pulumi.getter(name="storageGib")
    def storage_gib(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "storage_gib")

    @storage_gib.setter
    def storage_gib(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "storage_gib", value)


@pulumi.input_type
class ClusterRegionArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 node_count: Optional[pulumi.Input[int]] = None,
                 primary: Optional[pulumi.Input[bool]] = None,
                 sql_dns: Optional[pulumi.Input[str]] = None,
                 ui_dns: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Name of cluster
        :param pulumi.Input[bool] primary: Set to true to mark this region as the primary for a Serverless cluster. Exactly one region must be primary. Dedicated clusters expect to have no primary region.
        """
        pulumi.set(__self__, "name", name)
        if node_count is not None:
            pulumi.set(__self__, "node_count", node_count)
        if primary is not None:
            pulumi.set(__self__, "primary", primary)
        if sql_dns is not None:
            pulumi.set(__self__, "sql_dns", sql_dns)
        if ui_dns is not None:
            pulumi.set(__self__, "ui_dns", ui_dns)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Name of cluster
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nodeCount")
    def node_count(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "node_count")

    @node_count.setter
    def node_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "node_count", value)

    @property
    @pulumi.getter
    def primary(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to true to mark this region as the primary for a Serverless cluster. Exactly one region must be primary. Dedicated clusters expect to have no primary region.
        """
        return pulumi.get(self, "primary")

    @primary.setter
    def primary(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "primary", value)

    @property
    @pulumi.getter(name="sqlDns")
    def sql_dns(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "sql_dns")

    @sql_dns.setter
    def sql_dns(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sql_dns", value)

    @property
    @pulumi.getter(name="uiDns")
    def ui_dns(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ui_dns")

    @ui_dns.setter
    def ui_dns(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ui_dns", value)


@pulumi.input_type
class ClusterServerlessArgs:
    def __init__(__self__, *,
                 routing_id: Optional[pulumi.Input[str]] = None,
                 spend_limit: Optional[pulumi.Input[int]] = None,
                 usage_limits: Optional[pulumi.Input['ClusterServerlessUsageLimitsArgs']] = None):
        """
        :param pulumi.Input[int] spend_limit: Spend limit in US cents.
        """
        if routing_id is not None:
            pulumi.set(__self__, "routing_id", routing_id)
        if spend_limit is not None:
            pulumi.set(__self__, "spend_limit", spend_limit)
        if usage_limits is not None:
            pulumi.set(__self__, "usage_limits", usage_limits)

    @property
    @pulumi.getter(name="routingId")
    def routing_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "routing_id")

    @routing_id.setter
    def routing_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "routing_id", value)

    @property
    @pulumi.getter(name="spendLimit")
    def spend_limit(self) -> Optional[pulumi.Input[int]]:
        """
        Spend limit in US cents.
        """
        return pulumi.get(self, "spend_limit")

    @spend_limit.setter
    def spend_limit(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "spend_limit", value)

    @property
    @pulumi.getter(name="usageLimits")
    def usage_limits(self) -> Optional[pulumi.Input['ClusterServerlessUsageLimitsArgs']]:
        return pulumi.get(self, "usage_limits")

    @usage_limits.setter
    def usage_limits(self, value: Optional[pulumi.Input['ClusterServerlessUsageLimitsArgs']]):
        pulumi.set(self, "usage_limits", value)


@pulumi.input_type
class ClusterServerlessUsageLimitsArgs:
    def __init__(__self__, *,
                 request_unit_limit: pulumi.Input[int],
                 storage_mib_limit: pulumi.Input[int]):
        pulumi.set(__self__, "request_unit_limit", request_unit_limit)
        pulumi.set(__self__, "storage_mib_limit", storage_mib_limit)

    @property
    @pulumi.getter(name="requestUnitLimit")
    def request_unit_limit(self) -> pulumi.Input[int]:
        return pulumi.get(self, "request_unit_limit")

    @request_unit_limit.setter
    def request_unit_limit(self, value: pulumi.Input[int]):
        pulumi.set(self, "request_unit_limit", value)

    @property
    @pulumi.getter(name="storageMibLimit")
    def storage_mib_limit(self) -> pulumi.Input[int]:
        return pulumi.get(self, "storage_mib_limit")

    @storage_mib_limit.setter
    def storage_mib_limit(self, value: pulumi.Input[int]):
        pulumi.set(self, "storage_mib_limit", value)


@pulumi.input_type
class CmekAdditionalRegionArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 node_count: Optional[pulumi.Input[int]] = None,
                 primary: Optional[pulumi.Input[bool]] = None,
                 sql_dns: Optional[pulumi.Input[str]] = None,
                 ui_dns: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[bool] primary: Set to true to mark this region as the primary for a Serverless cluster. Exactly one region must be primary. Dedicated clusters expect to have no primary region.
        """
        pulumi.set(__self__, "name", name)
        if node_count is not None:
            pulumi.set(__self__, "node_count", node_count)
        if primary is not None:
            pulumi.set(__self__, "primary", primary)
        if sql_dns is not None:
            pulumi.set(__self__, "sql_dns", sql_dns)
        if ui_dns is not None:
            pulumi.set(__self__, "ui_dns", ui_dns)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nodeCount")
    def node_count(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "node_count")

    @node_count.setter
    def node_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "node_count", value)

    @property
    @pulumi.getter
    def primary(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to true to mark this region as the primary for a Serverless cluster. Exactly one region must be primary. Dedicated clusters expect to have no primary region.
        """
        return pulumi.get(self, "primary")

    @primary.setter
    def primary(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "primary", value)

    @property
    @pulumi.getter(name="sqlDns")
    def sql_dns(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "sql_dns")

    @sql_dns.setter
    def sql_dns(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sql_dns", value)

    @property
    @pulumi.getter(name="uiDns")
    def ui_dns(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ui_dns")

    @ui_dns.setter
    def ui_dns(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ui_dns", value)


@pulumi.input_type
class CmekRegionArgs:
    def __init__(__self__, *,
                 key: pulumi.Input['CmekRegionKeyArgs'],
                 region: pulumi.Input[str],
                 status: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] status: Aggregated status of the cluster's encryption key(s)
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "region", region)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input['CmekRegionKeyArgs']:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input['CmekRegionKeyArgs']):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Aggregated status of the cluster's encryption key(s)
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class CmekRegionKeyArgs:
    def __init__(__self__, *,
                 auth_principal: pulumi.Input[str],
                 type: pulumi.Input[str],
                 uri: pulumi.Input[str],
                 created_at: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None,
                 user_message: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] status: Aggregated status of the cluster's encryption key(s)
        """
        pulumi.set(__self__, "auth_principal", auth_principal)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "uri", uri)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)
        if user_message is not None:
            pulumi.set(__self__, "user_message", user_message)

    @property
    @pulumi.getter(name="authPrincipal")
    def auth_principal(self) -> pulumi.Input[str]:
        return pulumi.get(self, "auth_principal")

    @auth_principal.setter
    def auth_principal(self, value: pulumi.Input[str]):
        pulumi.set(self, "auth_principal", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def uri(self) -> pulumi.Input[str]:
        return pulumi.get(self, "uri")

    @uri.setter
    def uri(self, value: pulumi.Input[str]):
        pulumi.set(self, "uri", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Aggregated status of the cluster's encryption key(s)
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)

    @property
    @pulumi.getter(name="userMessage")
    def user_message(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "user_message")

    @user_message.setter
    def user_message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_message", value)


@pulumi.input_type
class LogExportConfigGroupArgs:
    def __init__(__self__, *,
                 channels: pulumi.Input[Sequence[pulumi.Input[str]]],
                 log_name: pulumi.Input[str],
                 min_level: Optional[pulumi.Input[str]] = None,
                 redact: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] channels: A list of CRDB log channels to include in this group
        :param pulumi.Input[str] log_name: The name of the group, reflected in the log sink
        :param pulumi.Input[str] min_level: The minimum log level to filter to this log group
        :param pulumi.Input[bool] redact: Governs whether this log group should aggregate redacted logs if unset
        """
        pulumi.set(__self__, "channels", channels)
        pulumi.set(__self__, "log_name", log_name)
        if min_level is not None:
            pulumi.set(__self__, "min_level", min_level)
        if redact is not None:
            pulumi.set(__self__, "redact", redact)

    @property
    @pulumi.getter
    def channels(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A list of CRDB log channels to include in this group
        """
        return pulumi.get(self, "channels")

    @channels.setter
    def channels(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "channels", value)

    @property
    @pulumi.getter(name="logName")
    def log_name(self) -> pulumi.Input[str]:
        """
        The name of the group, reflected in the log sink
        """
        return pulumi.get(self, "log_name")

    @log_name.setter
    def log_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "log_name", value)

    @property
    @pulumi.getter(name="minLevel")
    def min_level(self) -> Optional[pulumi.Input[str]]:
        """
        The minimum log level to filter to this log group
        """
        return pulumi.get(self, "min_level")

    @min_level.setter
    def min_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "min_level", value)

    @property
    @pulumi.getter
    def redact(self) -> Optional[pulumi.Input[bool]]:
        """
        Governs whether this log group should aggregate redacted logs if unset
        """
        return pulumi.get(self, "redact")

    @redact.setter
    def redact(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "redact", value)


@pulumi.input_type
class PrivateEndpointServicesServiceArgs:
    def __init__(__self__, *,
                 aws: Optional[pulumi.Input['PrivateEndpointServicesServiceAwsArgs']] = None,
                 cloud_provider: Optional[pulumi.Input[str]] = None,
                 region_name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        if aws is not None:
            pulumi.set(__self__, "aws", aws)
        if cloud_provider is not None:
            pulumi.set(__self__, "cloud_provider", cloud_provider)
        if region_name is not None:
            pulumi.set(__self__, "region_name", region_name)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def aws(self) -> Optional[pulumi.Input['PrivateEndpointServicesServiceAwsArgs']]:
        return pulumi.get(self, "aws")

    @aws.setter
    def aws(self, value: Optional[pulumi.Input['PrivateEndpointServicesServiceAwsArgs']]):
        pulumi.set(self, "aws", value)

    @property
    @pulumi.getter(name="cloudProvider")
    def cloud_provider(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "cloud_provider")

    @cloud_provider.setter
    def cloud_provider(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cloud_provider", value)

    @property
    @pulumi.getter(name="regionName")
    def region_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region_name")

    @region_name.setter
    def region_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region_name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class PrivateEndpointServicesServiceAwsArgs:
    def __init__(__self__, *,
                 availability_zone_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None):
        if availability_zone_ids is not None:
            pulumi.set(__self__, "availability_zone_ids", availability_zone_ids)
        if service_id is not None:
            pulumi.set(__self__, "service_id", service_id)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="availabilityZoneIds")
    def availability_zone_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "availability_zone_ids")

    @availability_zone_ids.setter
    def availability_zone_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "availability_zone_ids", value)

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "service_id")

    @service_id.setter
    def service_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_id", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)


@pulumi.input_type
class UserRoleGrantsRoleArgs:
    def __init__(__self__, *,
                 resource_type: pulumi.Input[str],
                 role_name: pulumi.Input[str],
                 resource_id: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "resource_type", resource_type)
        pulumi.set(__self__, "role_name", role_name)
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_type", value)

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "role_name")

    @role_name.setter
    def role_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_name", value)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_id", value)

