package awsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SchedulingPolicy.
// Experimental.
type ISchedulingPolicyRef interface {
	constructs.IConstruct
	// A reference to a SchedulingPolicy resource.
	// Experimental.
	SchedulingPolicyRef() *SchedulingPolicyReference
}

// The jsii proxy for ISchedulingPolicyRef
type jsiiProxy_ISchedulingPolicyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISchedulingPolicyRef) SchedulingPolicyRef() *SchedulingPolicyReference {
	var returns *SchedulingPolicyReference
	_jsii_.Get(
		j,
		"schedulingPolicyRef",
		&returns,
	)
	return returns
}

