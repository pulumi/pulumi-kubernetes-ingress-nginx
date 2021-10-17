# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *
import pulumi_kubernetes

__all__ = ['IngressControllerArgs', 'IngressController']

@pulumi.input_type
class IngressControllerArgs:
    def __init__(__self__, *,
                 controller: Optional[pulumi.Input['ControllerArgs']] = None,
                 default_backend: Optional[pulumi.Input['ControllerDefaultBackendArgs']] = None,
                 dh_param: Optional[pulumi.Input[str]] = None,
                 fullname_override: Optional[pulumi.Input[str]] = None,
                 helm_options: Optional[pulumi.Input['ReleaseArgs']] = None,
                 image_pull_secrets: Optional[pulumi.Input[Sequence[pulumi.Input['pulumi_kubernetes.core.v1.LocalObjectReferenceArgs']]]] = None,
                 name_override: Optional[pulumi.Input[str]] = None,
                 pod_security_policy: Optional[pulumi.Input['ControllerPodSecurityPolicyArgs']] = None,
                 rbac: Optional[pulumi.Input['ControllerRBACArgs']] = None,
                 revision_history_limit: Optional[pulumi.Input[int]] = None,
                 service_account: Optional[pulumi.Input['ControllerServiceAccountArgs']] = None,
                 tcp: Optional[pulumi.Input[Mapping[str, pulumi.Input[Mapping[str, pulumi.Input[str]]]]]] = None,
                 udp: Optional[pulumi.Input[Mapping[str, pulumi.Input[Mapping[str, pulumi.Input[str]]]]]] = None):
        """
        The set of arguments for constructing a IngressController resource.
        :param pulumi.Input['ControllerDefaultBackendArgs'] default_backend: Default 404 backend.
        :param pulumi.Input[str] dh_param: A base64ed Diffie-Hellman parameter. This can be generated with: openssl dhparam 4096 2> /dev/null | base64 Ref: https://github.com/kubernetes/ingress-nginx/tree/main/docs/examples/customization/ssl-dh-param.
        :param pulumi.Input[str] fullname_override: Overrides for generated resource names.
        :param pulumi.Input['ReleaseArgs'] helm_options: HelmOptions is an escape hatch that lets the end user control any aspect of the Helm deployment. This exposes the entirety of the underlying Helm Release component args.
        :param pulumi.Input[Sequence[pulumi.Input['pulumi_kubernetes.core.v1.LocalObjectReferenceArgs']]] image_pull_secrets: Optional array of imagePullSecrets containing private registry credentials Ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/.
        :param pulumi.Input[str] name_override: Overrides for generated resource names.
        :param pulumi.Input['ControllerPodSecurityPolicyArgs'] pod_security_policy: If true, create & use Pod Security Policy resources https://kubernetes.io/docs/concepts/policy/pod-security-policy/
        :param pulumi.Input['ControllerRBACArgs'] rbac: Enable RBAC as per https://github.com/kubernetes/ingress-nginx/blob/main/docs/deploy/rbac.md and https://github.com/kubernetes/ingress-nginx/issues/266
        :param pulumi.Input[int] revision_history_limit: Rollback limit.
        :param pulumi.Input[Mapping[str, pulumi.Input[Mapping[str, pulumi.Input[str]]]]] tcp: TCP service key:value pairs Ref: https://github.com/kubernetes/ingress-nginx/blob/main/docs/user-guide/exposing-tcp-udp-services.md.
        :param pulumi.Input[Mapping[str, pulumi.Input[Mapping[str, pulumi.Input[str]]]]] udp: UDP service key:value pairs Ref: https://github.com/kubernetes/ingress-nginx/blob/main/docs/user-guide/exposing-tcp-udp-services.md.
        """
        if controller is not None:
            pulumi.set(__self__, "controller", controller)
        if default_backend is not None:
            pulumi.set(__self__, "default_backend", default_backend)
        if dh_param is not None:
            pulumi.set(__self__, "dh_param", dh_param)
        if fullname_override is not None:
            pulumi.set(__self__, "fullname_override", fullname_override)
        if helm_options is not None:
            pulumi.set(__self__, "helm_options", helm_options)
        if image_pull_secrets is not None:
            pulumi.set(__self__, "image_pull_secrets", image_pull_secrets)
        if name_override is not None:
            pulumi.set(__self__, "name_override", name_override)
        if pod_security_policy is not None:
            pulumi.set(__self__, "pod_security_policy", pod_security_policy)
        if rbac is not None:
            pulumi.set(__self__, "rbac", rbac)
        if revision_history_limit is not None:
            pulumi.set(__self__, "revision_history_limit", revision_history_limit)
        if service_account is not None:
            pulumi.set(__self__, "service_account", service_account)
        if tcp is not None:
            pulumi.set(__self__, "tcp", tcp)
        if udp is not None:
            pulumi.set(__self__, "udp", udp)

    @property
    @pulumi.getter
    def controller(self) -> Optional[pulumi.Input['ControllerArgs']]:
        return pulumi.get(self, "controller")

    @controller.setter
    def controller(self, value: Optional[pulumi.Input['ControllerArgs']]):
        pulumi.set(self, "controller", value)

    @property
    @pulumi.getter(name="defaultBackend")
    def default_backend(self) -> Optional[pulumi.Input['ControllerDefaultBackendArgs']]:
        """
        Default 404 backend.
        """
        return pulumi.get(self, "default_backend")

    @default_backend.setter
    def default_backend(self, value: Optional[pulumi.Input['ControllerDefaultBackendArgs']]):
        pulumi.set(self, "default_backend", value)

    @property
    @pulumi.getter(name="dhParam")
    def dh_param(self) -> Optional[pulumi.Input[str]]:
        """
        A base64ed Diffie-Hellman parameter. This can be generated with: openssl dhparam 4096 2> /dev/null | base64 Ref: https://github.com/kubernetes/ingress-nginx/tree/main/docs/examples/customization/ssl-dh-param.
        """
        return pulumi.get(self, "dh_param")

    @dh_param.setter
    def dh_param(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dh_param", value)

    @property
    @pulumi.getter(name="fullnameOverride")
    def fullname_override(self) -> Optional[pulumi.Input[str]]:
        """
        Overrides for generated resource names.
        """
        return pulumi.get(self, "fullname_override")

    @fullname_override.setter
    def fullname_override(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fullname_override", value)

    @property
    @pulumi.getter(name="helmOptions")
    def helm_options(self) -> Optional[pulumi.Input['ReleaseArgs']]:
        """
        HelmOptions is an escape hatch that lets the end user control any aspect of the Helm deployment. This exposes the entirety of the underlying Helm Release component args.
        """
        return pulumi.get(self, "helm_options")

    @helm_options.setter
    def helm_options(self, value: Optional[pulumi.Input['ReleaseArgs']]):
        pulumi.set(self, "helm_options", value)

    @property
    @pulumi.getter(name="imagePullSecrets")
    def image_pull_secrets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['pulumi_kubernetes.core.v1.LocalObjectReferenceArgs']]]]:
        """
        Optional array of imagePullSecrets containing private registry credentials Ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/.
        """
        return pulumi.get(self, "image_pull_secrets")

    @image_pull_secrets.setter
    def image_pull_secrets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['pulumi_kubernetes.core.v1.LocalObjectReferenceArgs']]]]):
        pulumi.set(self, "image_pull_secrets", value)

    @property
    @pulumi.getter(name="nameOverride")
    def name_override(self) -> Optional[pulumi.Input[str]]:
        """
        Overrides for generated resource names.
        """
        return pulumi.get(self, "name_override")

    @name_override.setter
    def name_override(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name_override", value)

    @property
    @pulumi.getter(name="podSecurityPolicy")
    def pod_security_policy(self) -> Optional[pulumi.Input['ControllerPodSecurityPolicyArgs']]:
        """
        If true, create & use Pod Security Policy resources https://kubernetes.io/docs/concepts/policy/pod-security-policy/
        """
        return pulumi.get(self, "pod_security_policy")

    @pod_security_policy.setter
    def pod_security_policy(self, value: Optional[pulumi.Input['ControllerPodSecurityPolicyArgs']]):
        pulumi.set(self, "pod_security_policy", value)

    @property
    @pulumi.getter
    def rbac(self) -> Optional[pulumi.Input['ControllerRBACArgs']]:
        """
        Enable RBAC as per https://github.com/kubernetes/ingress-nginx/blob/main/docs/deploy/rbac.md and https://github.com/kubernetes/ingress-nginx/issues/266
        """
        return pulumi.get(self, "rbac")

    @rbac.setter
    def rbac(self, value: Optional[pulumi.Input['ControllerRBACArgs']]):
        pulumi.set(self, "rbac", value)

    @property
    @pulumi.getter(name="revisionHistoryLimit")
    def revision_history_limit(self) -> Optional[pulumi.Input[int]]:
        """
        Rollback limit.
        """
        return pulumi.get(self, "revision_history_limit")

    @revision_history_limit.setter
    def revision_history_limit(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "revision_history_limit", value)

    @property
    @pulumi.getter(name="serviceAccount")
    def service_account(self) -> Optional[pulumi.Input['ControllerServiceAccountArgs']]:
        return pulumi.get(self, "service_account")

    @service_account.setter
    def service_account(self, value: Optional[pulumi.Input['ControllerServiceAccountArgs']]):
        pulumi.set(self, "service_account", value)

    @property
    @pulumi.getter
    def tcp(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[Mapping[str, pulumi.Input[str]]]]]]:
        """
        TCP service key:value pairs Ref: https://github.com/kubernetes/ingress-nginx/blob/main/docs/user-guide/exposing-tcp-udp-services.md.
        """
        return pulumi.get(self, "tcp")

    @tcp.setter
    def tcp(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[Mapping[str, pulumi.Input[str]]]]]]):
        pulumi.set(self, "tcp", value)

    @property
    @pulumi.getter
    def udp(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[Mapping[str, pulumi.Input[str]]]]]]:
        """
        UDP service key:value pairs Ref: https://github.com/kubernetes/ingress-nginx/blob/main/docs/user-guide/exposing-tcp-udp-services.md.
        """
        return pulumi.get(self, "udp")

    @udp.setter
    def udp(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[Mapping[str, pulumi.Input[str]]]]]]):
        pulumi.set(self, "udp", value)


class IngressController(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 controller: Optional[pulumi.Input[pulumi.InputType['ControllerArgs']]] = None,
                 default_backend: Optional[pulumi.Input[pulumi.InputType['ControllerDefaultBackendArgs']]] = None,
                 dh_param: Optional[pulumi.Input[str]] = None,
                 fullname_override: Optional[pulumi.Input[str]] = None,
                 helm_options: Optional[pulumi.Input[pulumi.InputType['ReleaseArgs']]] = None,
                 image_pull_secrets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['pulumi_kubernetes.core.v1.LocalObjectReferenceArgs']]]]] = None,
                 name_override: Optional[pulumi.Input[str]] = None,
                 pod_security_policy: Optional[pulumi.Input[pulumi.InputType['ControllerPodSecurityPolicyArgs']]] = None,
                 rbac: Optional[pulumi.Input[pulumi.InputType['ControllerRBACArgs']]] = None,
                 revision_history_limit: Optional[pulumi.Input[int]] = None,
                 service_account: Optional[pulumi.Input[pulumi.InputType['ControllerServiceAccountArgs']]] = None,
                 tcp: Optional[pulumi.Input[Mapping[str, pulumi.Input[Mapping[str, pulumi.Input[str]]]]]] = None,
                 udp: Optional[pulumi.Input[Mapping[str, pulumi.Input[Mapping[str, pulumi.Input[str]]]]]] = None,
                 __props__=None):
        """
        Ingress controller for Kubernetes using NGINX as a reverse proxy and load balancer

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ControllerDefaultBackendArgs']] default_backend: Default 404 backend.
        :param pulumi.Input[str] dh_param: A base64ed Diffie-Hellman parameter. This can be generated with: openssl dhparam 4096 2> /dev/null | base64 Ref: https://github.com/kubernetes/ingress-nginx/tree/main/docs/examples/customization/ssl-dh-param.
        :param pulumi.Input[str] fullname_override: Overrides for generated resource names.
        :param pulumi.Input[pulumi.InputType['ReleaseArgs']] helm_options: HelmOptions is an escape hatch that lets the end user control any aspect of the Helm deployment. This exposes the entirety of the underlying Helm Release component args.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['pulumi_kubernetes.core.v1.LocalObjectReferenceArgs']]]] image_pull_secrets: Optional array of imagePullSecrets containing private registry credentials Ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/.
        :param pulumi.Input[str] name_override: Overrides for generated resource names.
        :param pulumi.Input[pulumi.InputType['ControllerPodSecurityPolicyArgs']] pod_security_policy: If true, create & use Pod Security Policy resources https://kubernetes.io/docs/concepts/policy/pod-security-policy/
        :param pulumi.Input[pulumi.InputType['ControllerRBACArgs']] rbac: Enable RBAC as per https://github.com/kubernetes/ingress-nginx/blob/main/docs/deploy/rbac.md and https://github.com/kubernetes/ingress-nginx/issues/266
        :param pulumi.Input[int] revision_history_limit: Rollback limit.
        :param pulumi.Input[Mapping[str, pulumi.Input[Mapping[str, pulumi.Input[str]]]]] tcp: TCP service key:value pairs Ref: https://github.com/kubernetes/ingress-nginx/blob/main/docs/user-guide/exposing-tcp-udp-services.md.
        :param pulumi.Input[Mapping[str, pulumi.Input[Mapping[str, pulumi.Input[str]]]]] udp: UDP service key:value pairs Ref: https://github.com/kubernetes/ingress-nginx/blob/main/docs/user-guide/exposing-tcp-udp-services.md.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[IngressControllerArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Ingress controller for Kubernetes using NGINX as a reverse proxy and load balancer

        :param str resource_name: The name of the resource.
        :param IngressControllerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IngressControllerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 controller: Optional[pulumi.Input[pulumi.InputType['ControllerArgs']]] = None,
                 default_backend: Optional[pulumi.Input[pulumi.InputType['ControllerDefaultBackendArgs']]] = None,
                 dh_param: Optional[pulumi.Input[str]] = None,
                 fullname_override: Optional[pulumi.Input[str]] = None,
                 helm_options: Optional[pulumi.Input[pulumi.InputType['ReleaseArgs']]] = None,
                 image_pull_secrets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['pulumi_kubernetes.core.v1.LocalObjectReferenceArgs']]]]] = None,
                 name_override: Optional[pulumi.Input[str]] = None,
                 pod_security_policy: Optional[pulumi.Input[pulumi.InputType['ControllerPodSecurityPolicyArgs']]] = None,
                 rbac: Optional[pulumi.Input[pulumi.InputType['ControllerRBACArgs']]] = None,
                 revision_history_limit: Optional[pulumi.Input[int]] = None,
                 service_account: Optional[pulumi.Input[pulumi.InputType['ControllerServiceAccountArgs']]] = None,
                 tcp: Optional[pulumi.Input[Mapping[str, pulumi.Input[Mapping[str, pulumi.Input[str]]]]]] = None,
                 udp: Optional[pulumi.Input[Mapping[str, pulumi.Input[Mapping[str, pulumi.Input[str]]]]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is not None:
            raise ValueError('ComponentResource classes do not support opts.id')
        else:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IngressControllerArgs.__new__(IngressControllerArgs)

            __props__.__dict__["controller"] = controller
            __props__.__dict__["default_backend"] = default_backend
            __props__.__dict__["dh_param"] = dh_param
            __props__.__dict__["fullname_override"] = fullname_override
            __props__.__dict__["helm_options"] = helm_options
            __props__.__dict__["image_pull_secrets"] = image_pull_secrets
            __props__.__dict__["name_override"] = name_override
            __props__.__dict__["pod_security_policy"] = pod_security_policy
            __props__.__dict__["rbac"] = rbac
            __props__.__dict__["revision_history_limit"] = revision_history_limit
            __props__.__dict__["service_account"] = service_account
            __props__.__dict__["tcp"] = tcp
            __props__.__dict__["udp"] = udp
            __props__.__dict__["status"] = None
        super(IngressController, __self__).__init__(
            'chart-ingress-nginx:index:IngressController',
            resource_name,
            __props__,
            opts,
            remote=True)

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output['outputs.ReleaseStatus']:
        """
        Detailed information about the status of the underlying Helm deployment.
        """
        return pulumi.get(self, "status")
