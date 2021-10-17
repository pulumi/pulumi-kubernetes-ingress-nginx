// Example inspired by:
// https://www.digitalocean.com/community/tutorials/how-to-set-up-an-nginx-ingress-on-digitalocean-kubernetes-using-helm
// Note that, as described in that article, you will need to configure DNS for hw1/hw2.your_domain_name.

import * as k8s from "@pulumi/kubernetes";
import * as nginx from "@pulumi/kubernetes-ingress-nginx";

// Install the NGINX ingress controller to our cluster. The controller
// consists of a Pod and a Service. Install it and configure the controller
// to publish the load balancer IP address on each Ingress so that
// applications can depend on the IP address of the load balancer if needed.
const ctrl = new nginx.IngressController("myctrl", {
    controller: {
        publishService: {
            enabled: true,
        },
    },
});

// Now let's deploy two applications which are identical except for the
// names. We will later configure the ingress to direct traffic to them,
// one domain name per application instance.
const apps = [];
const appBase = "hello-k8s";
const appNames = [ `${appBase}-first`, `${appBase}-second` ];
for (const appName of appNames) {
    const appSvc = new k8s.core.v1.Service(`${appName}-svc`, {
        metadata: { name: appName },
        spec: {
            type: "ClusterIP",
            ports: [{ port: 80, targetPort: 8080 }],
            selector: { app: appName },
        },
    });
    const appDep = new k8s.apps.v1.Deployment(`${appName}-dep`, {
        metadata: { name: appName },
        spec: {
            replicas: 3,
            selector: {
                matchLabels: { app: appName },
            },
            template: {
                metadata: {
                    labels: { app: appName },
                },
                spec: {
                    containers: [{
                        name: appName,
                        image: "paulbouwer/hello-kubernetes:1.8",
                        ports: [{ containerPort: 8080 }],
                        env: [{ name: "MESSAGE", value: "Hello K8s!" }],
                    }],
                },
            },
        },
    });
    apps.push(appSvc.status);
}

// Next, expose the app using an Ingress.
const appIngress = new k8s.networking.v1.Ingress(`${appBase}-ingress`, {
    metadata: {
        name: "hello-k8s-ingress",
        annotations: {
            "kubernetes.io/ingress.class": "nginx",
        },
    },
    spec: {
        rules: [
            {
                // Replace this with your own domain!
                host: "myservicea.foo.org",
                http: {
                    paths: [{
                        pathType: "Prefix",
                        path: "/",
                        backend: {
                            service: {
                                name: appNames[0],
                                port: { number: 80 },
                            },
                        },
                    }],
                },
            },
            {
                // Replace this with your own domain!
                host: "myserviceb.foo.org",
                http: {
                    paths: [{
                        pathType: "Prefix",
                        path: "/",
                        backend: {
                            service: {
                                name: appNames[1],
                                port: { number: 80 },
                            },
                        },
                    }],
                },
            },
        ],
    },
});

export const appStatuses = apps;
export const controllerStatus = ctrl.status;
