# Simple NGINX Example

This example demonstrates how to deploy a simple NGINX server using Pulumi and Kubernetes.

## Prerequisites

- [Pulumi](https://www.pulumi.com/docs/get-started/install/)
- [Go](https://golang.org/doc/install)
- [Kubernetes](https://kubernetes.io/docs/tasks/tools/install-kubectl/)

## Getting Started

1. Clone the repository:

    ```sh
    git clone https://github.com/pulumi/pulumi-kubernetes-ingress-nginx.git
    cd pulumi-kubernetes-ingress-nginx/examples/simple-nginx-go
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

3. Modify the `go.mod` file to add the following line:

    ```go
    replace github.com/pulumi/pulumi-kubernetes-ingress-nginx/sdk => ../../sdk/
    ```

4. Deploy the NGINX server:

    ```sh
    pulumi up
    ```

5. To destroy the resources when you're done:

    ```sh
    pulumi destroy
    ```

## Testing

To test the deployment, follow these steps:

1. Verify that the NGINX server is running:

    ```sh
    kubectl get pods
    ```

2. Get the external IP address of the NGINX service:

    ```sh
    kubectl get svc
    ```

3. Open a web browser and navigate to the external IP address. You should see the default NGINX welcome page.

## Cleanup

To clean up the resources created by this example, run:

```sh
pulumi destroy
```
