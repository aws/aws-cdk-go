package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ManagedPolicy.
// Experimental.
type IManagedPolicyRef interface {
	constructs.IConstruct
	// A reference to a ManagedPolicy resource.
	// Experimental.
	ManagedPolicyRef() *ManagedPolicyReference
}

// The jsii proxy for IManagedPolicyRef
type jsiiProxy_IManagedPolicyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IManagedPolicyRef) ManagedPolicyRef() *ManagedPolicyReference {
	var returns *ManagedPolicyReference
	_jsii_.Get(
		j,
		"managedPolicyRef",
		&returns,
	)
	return returns
}

