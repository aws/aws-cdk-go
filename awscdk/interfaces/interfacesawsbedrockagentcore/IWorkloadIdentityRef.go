package interfacesawsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbedrockagentcore/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WorkloadIdentity.
// Experimental.
type IWorkloadIdentityRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a WorkloadIdentity resource.
	// Experimental.
	WorkloadIdentityRef() *WorkloadIdentityReference
}

// The jsii proxy for IWorkloadIdentityRef
type jsiiProxy_IWorkloadIdentityRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IWorkloadIdentityRef) WorkloadIdentityRef() *WorkloadIdentityReference {
	var returns *WorkloadIdentityReference
	_jsii_.Get(
		j,
		"workloadIdentityRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkloadIdentityRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkloadIdentityRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

