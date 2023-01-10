# Provider Linode

**###############################################################################################################**

**NOTE: THIS IS A PRE-RELEASE REPO AND SHOULD NOT BE DEPLOYED TO PRODUCTION AT THIS TIME UNTIL OFFICIALLY RELEASED**

**###############################################################################################################**

- Crossplane Website: <https://www.crossplane.io>
- Linode Website: <https://www.linode.com>
- Documentation: <https://marketplace.upbound.io/providers/linode/provider-linode>


## Abount Crossplane

<img src="https://github.com/crossplane/artwork/blob/420102818bc649730cc97de5b4ed8178e9333eb5/logo/icon.svg" height="200px" width="300px">

Crossplane is an open source Kubernetes add-on that transforms your cluster into a universal control plane. Crossplane enables platform teams to assemble infrastructure from multiple vendors, and expose higher level self-service APIs for application teams to consume, without having to write any code.

Crossplane extends your Kubernetes cluster to support orchestrating any infrastructure or managed service. Compose Crossplane’s granular resources into higher level abstractions that can be versioned, managed, deployed and consumed using your favorite tools and existing processes. Install Crossplane into any Kubernetes cluster to get started.

Crossplane is a Cloud Native Compute Foundation project.

## About Linode / Akamai Cloud Computing

<img src="https://www.linode.com/wp-content/themes/linode-website-theme/images/linode-akamai-logo.svg?ver=1663187393" height="200px" width="300px">

Akamai Cloud Computing based on Linode accelerates innovation with scalable, simple, affordable, and accessible Linux cloud solutions and services. Our products, services, and people give developers and enterprises the flexibility, support, and trust they need to build, deploy, secure, and scale applications more easily and cost-effectively from cloud to edge on the world’s most distributed network.

## Maintainers

This provider plugin is maintained by Linode.

## Requirements

- [Crossplane](https://docs.crossplane.io/v1.10/getting-started/install-configure/) 0.10.0+
- [Go](https://golang.org/doc/install) 1.18.0 or higher (to build the provider plugin)

## Using the provider

`provider-linode` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/upbound/upjet) code
generation tools and exposes XRM-conformant managed resources for the
Linode API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/linode/provider-linode):
```
up ctp provider install linode/provider-linode:v0.1.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-linode
spec:
  package: linode/provider-linode:v0.1.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/linode/provider-linode).

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/linode/provider-linode/issues).
