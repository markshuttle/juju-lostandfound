[![Juju logo](doc/juju-logo.png?raw=true)](https://juju.is/)

A model-driven Operator Lifecycle Manager.

An "operator" is a reusable operations code package that drives the
configuration and operation of a particular workload. By encapsulating ops
code as a reusable package, the operator pattern moves beyond traditional
config management to allow much more agile operations for complex cloud
workloads.

Shared, open source operators take **infra as code to the next level** with
community-driven ops and integration code.

Juju is a **universal OLM** that supports operators on bare metal, virtual
machines or cloud instances, as well as on Kubernetes. Embrace the operator
pattern across container and legacy estate.

**Juju excels at application integration**. Instead of simply focusing on
lifecycle management, the Juju OLM provides a rich application graph model
that tells operators how to integrate with one another. This dramatically
simplifies the operations of large deployments.

Composable operators enable very rich scenarios to be constructed out of
simpler operators that do one thing, and do it well.

## Open Operator Collection

A world's largest collection of composable operators. The community
emphasizes quality, collaboration and consistency. Publish your own operator
and share integration code for other operators to connect to your
application.

## Pure Python operators

The Python Operator Framework makes it easy to write an operator. The
framework handles all the details of communication between integrated
operators, so you can focus on your own application lifecycle management.

Code sharing between operator publishers is simplified making it much faster
to collaborate on distributed systems involving components from many
different publishers and upstreams.

With the Python Operator Framework, your operator is a Python event handler.
Lifecycle management, configuration and integration are all events delivered
to your charm by the framework.

## Multi cloud and hybrid operations across ARM and x86 infrastructure

THe Juju OLM supports AWS, Azure, Google, Oracle, OpenStack, VMware and bare
metal machines, as well as any conformant Kubernetes cluster. Integrate
operators across clouds, and across machines and containers, just as easily.

Juju operators support multiple architectures. Connect applications on ARM
with applications on x86 to take advantage of architecture-specific
optimisations.

## Get started

Our community hangs out at the [Charmhub
discourse](https://discourse.jujucharms.com/) which serves as a combination
mailing list and web forum. Keep up with the news and get a feel for
operator engineering and usage there.

Get started with the [install instructions](https://juju.is/docs/installing)
and [try the
tutorials](https://juju.is/docs/tutorials). All you need is a small K8s
cluster, or an Ubuntu machine or VM to run MicroK8s.

Read the [documentation](https://juju.is/docs) for a comprehensive reference
of commands and usage.

## Contributing

Follow our [code and contribution guidelines](CONTRIBUTING.md) to learn how
to make code changes. File bugs in
[Launchpad](https://bugs.launchpad.net/juju/+filebug) or ask questions on
our [Freenode IRC channel](https://webchat.freenode.net/#juju).

