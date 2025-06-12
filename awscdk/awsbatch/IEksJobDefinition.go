package awsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A JobDefinition that uses Eks orchestration.
type IEksJobDefinition interface {
	IJobDefinition
	// The container this Job Definition will run.
	Container() EksContainerDefinition
	// The DNS Policy of the pod used by this Job Definition.
	// See: https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pod-s-dns-policy
	//
	// Default: `DnsPolicy.CLUSTER_FIRST`
	//
	DnsPolicy() DnsPolicy
	// The name of the service account that's used to run the container.
	//
	// service accounts are Kubernetes method of identification and authentication,
	// roughly analogous to IAM users.
	// See: https://docs.aws.amazon.com/eks/latest/userguide/associate-service-account-role.html
	//
	// Default: - the default service account of the container.
	//
	ServiceAccount() *string
	// If specified, the Pod used by this Job Definition will use the host's network IP address.
	//
	// Otherwise, the Kubernetes pod networking model is enabled.
	// Most AWS Batch workloads are egress-only and don't require the overhead of IP allocation for each pod for incoming connections.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/#pod-networking
	//
	// Default: true.
	//
	UseHostNetwork() *bool
}

// The jsii proxy for IEksJobDefinition
type jsiiProxy_IEksJobDefinition struct {
	jsiiProxy_IJobDefinition
}

func (j *jsiiProxy_IEksJobDefinition) Container() EksContainerDefinition {
	var returns EksContainerDefinition
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEksJobDefinition) DnsPolicy() DnsPolicy {
	var returns DnsPolicy
	_jsii_.Get(
		j,
		"dnsPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEksJobDefinition) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEksJobDefinition) UseHostNetwork() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"useHostNetwork",
		&returns,
	)
	return returns
}

