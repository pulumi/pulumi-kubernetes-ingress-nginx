# Pulumi NGINX Ingress Controller Component

This repo contains the Pulumi NGINX Ingress Controller component for Kubernetes. This ingress controller
uses NGINX as a reverse proxy and load balancer.

This component wraps [the official NGINX Ingress Controller](https://github.com/kubernetes/ingress-nginx),
and offers a Pulumi-friendly and strongly-typed way to manage ingress controller installations.

After installing this component to your cluster, you can use it by adding the
`kubernetes.io/ingress.class: nginx` annotation to your `Ingress` resources.

## To Use

To use this component, first install the Pulumi Package:

Afterwards, import the library and instantiate it within your Pulumi program:

## Configuration

This component supports all of the configuration options of the [official Helm chart](
https://github.com/kubernetes/ingress-nginx/tree/main/charts/ingress-nginx), except that these
are strongly typed so you will get IDE support and static error checking.

For complete details, refer to the Pulumi Package details within the Pulumi Registry.
