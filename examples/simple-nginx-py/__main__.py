# Example inspired by:
# https://www.digitalocean.com/community/tutorials/how-to-set-up-an-nginx-ingress-on-digitalocean-kubernetes-using-helm
# Note that, as described in that article, you will need to configure DNS for hw1/hw2.your_domain_name.

from grpc import server
import pulumi
from pulumi_kubernetes.apps.v1 import Deployment
from pulumi_kubernetes.core.v1 import Service
from pulumi_kubernetes.networking.v1 import Ingress
from pulumi_kubernetes_ingress_nginx import (
    IngressController,
    ControllerArgs,
    ControllerHostPortArgs,
    ControllerPublishServiceArgs,
    ControllerServiceArgs,
)

# Install the NGINX ingress controller to our cluster. The controller
# consists of a Pod and a Service. Install it and configure the controller
# to publish the load balancer IP address on each Ingress so that
# applications can depend on the IP address of the load balancer if needed.
ctrl = IngressController(
    "myctrl",
    controller=ControllerArgs(
        host_port=ControllerHostPortArgs(
            enabled=True,
        ),
        publish_service=ControllerPublishServiceArgs(
            enabled=True,
        ),
        service=ControllerServiceArgs(type="NodePort", external_traffic_policy="Local"),
    ),
)

# Now let's deploy two applications which are identical except for the
# names. We will later configure the ingress to direct traffic to them,
# one domain name per application instance.
apps = []
app_base = "hello-k8s"
app_names = [f"{app_base}-first", f"{app_base}-second"]
for app_name in app_names:
    app_svc = Service(
        f"{app_name}-svc",
        metadata={"name": app_name},
        spec={
            "type": "ClusterIP",
            "ports": [{"port": 80, "targetPort": 8080}],
            "selector": {"app": app_name},
        },
    )
    app_dep = Deployment(
        f"{app_name}-dep",
        metadata={"name": app_name},
        spec={
            "replicas": 3,
            "selector": {
                "matchLabels": {"app": app_name},
            },
            "template": {
                "metadata": {
                    "labels": {"app": app_name},
                },
                "spec": {
                    "containers": [
                        {
                            "name": app_name,
                            "image": "paulbouwer/hello-kubernetes:1.8",
                            "ports": [{"containerPort": 8080}],
                            "env": [{"name": "MESSAGE", "value": "Hello K8s!"}],
                        }
                    ],
                },
            },
        },
    )
    apps.append(app_svc.status)

# Next, expose the app using an Ingress.
app_ingress = Ingress(
    f"{app_base}-ingress",
    metadata={
        "name": "hello-k8s-ingress",
        "annotations": {
            "kubernetes.io/ingress.class": "nginx",
        },
    },
    spec={
        "rules": [
            {
                # Replace this with your own domain!
                "host": "myservicea.foo.org",
                "http": {
                    "paths": [
                        {
                            "pathType": "Prefix",
                            "path": "/",
                            "backend": {
                                "service": {
                                    "name": app_names[0],
                                    "port": {"number": 80},
                                },
                            },
                        }
                    ],
                },
            },
            {
                # Replace this with your own domain!
                "host": "myserviceb.foo.org",
                "http": {
                    "paths": [
                        {
                            "pathType": "Prefix",
                            "path": "/",
                            "backend": {
                                "service": {
                                    "name": app_names[1],
                                    "port": {"number": 80},
                                },
                            },
                        }
                    ],
                },
            },
        ],
    },
)

pulumi.export("app_statuses", apps)
pulumi.export("controller_status", ctrl.status)
