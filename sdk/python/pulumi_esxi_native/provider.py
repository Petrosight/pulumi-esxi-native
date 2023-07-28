# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 host: pulumi.Input[str],
                 password: pulumi.Input[str],
                 ssh_port: Optional[pulumi.Input[str]] = None,
                 ssl_port: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[str] host: ESXi Host Name config
        :param pulumi.Input[str] password: ESXi Password config
        :param pulumi.Input[str] ssh_port: ESXi Host SSH Port config
        :param pulumi.Input[str] ssl_port: ESXi Host SSL Port config
        :param pulumi.Input[str] username: ESXi Username config
        """
        pulumi.set(__self__, "host", host)
        pulumi.set(__self__, "password", password)
        if ssh_port is None:
            ssh_port = '22'
        if ssh_port is not None:
            pulumi.set(__self__, "ssh_port", ssh_port)
        if ssl_port is None:
            ssl_port = '443'
        if ssl_port is not None:
            pulumi.set(__self__, "ssl_port", ssl_port)
        if username is None:
            username = 'root'
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def host(self) -> pulumi.Input[str]:
        """
        ESXi Host Name config
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: pulumi.Input[str]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Input[str]:
        """
        ESXi Password config
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: pulumi.Input[str]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="sshPort")
    def ssh_port(self) -> Optional[pulumi.Input[str]]:
        """
        ESXi Host SSH Port config
        """
        return pulumi.get(self, "ssh_port")

    @ssh_port.setter
    def ssh_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssh_port", value)

    @property
    @pulumi.getter(name="sslPort")
    def ssl_port(self) -> Optional[pulumi.Input[str]]:
        """
        ESXi Host SSL Port config
        """
        return pulumi.get(self, "ssl_port")

    @ssl_port.setter
    def ssl_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_port", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        ESXi Username config
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 ssh_port: Optional[pulumi.Input[str]] = None,
                 ssl_port: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The provider type for the ESXi native package. By default, resources use package-wide configuration settings, however an explicit `Provider` instance may be created and passed during resource construction to achieve fine-grained programmatic control over provider settings. See the [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] host: ESXi Host Name config
        :param pulumi.Input[str] password: ESXi Password config
        :param pulumi.Input[str] ssh_port: ESXi Host SSH Port config
        :param pulumi.Input[str] ssl_port: ESXi Host SSL Port config
        :param pulumi.Input[str] username: ESXi Username config
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProviderArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the ESXi native package. By default, resources use package-wide configuration settings, however an explicit `Provider` instance may be created and passed during resource construction to achieve fine-grained programmatic control over provider settings. See the [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 ssh_port: Optional[pulumi.Input[str]] = None,
                 ssl_port: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderArgs.__new__(ProviderArgs)

            if host is None and not opts.urn:
                raise TypeError("Missing required property 'host'")
            __props__.__dict__["host"] = host
            if password is None and not opts.urn:
                raise TypeError("Missing required property 'password'")
            __props__.__dict__["password"] = password
            if ssh_port is None:
                ssh_port = '22'
            __props__.__dict__["ssh_port"] = ssh_port
            if ssl_port is None:
                ssl_port = '443'
            __props__.__dict__["ssl_port"] = ssl_port
            if username is None:
                username = 'root'
            __props__.__dict__["username"] = username
        super(Provider, __self__).__init__(
            'esxi-native',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter
    def host(self) -> pulumi.Output[str]:
        """
        ESXi Host Name config
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[str]:
        """
        ESXi Password config
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="sshPort")
    def ssh_port(self) -> pulumi.Output[Optional[str]]:
        """
        ESXi Host SSH Port config
        """
        return pulumi.get(self, "ssh_port")

    @property
    @pulumi.getter(name="sslPort")
    def ssl_port(self) -> pulumi.Output[Optional[str]]:
        """
        ESXi Host SSL Port config
        """
        return pulumi.get(self, "ssl_port")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[Optional[str]]:
        """
        ESXi Username config
        """
        return pulumi.get(self, "username")

