package awsbatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SchedulingPolicy.
// Experimental.
type ISchedulingPolicyRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISchedulingPolicyRef
type jsiiProxy_ISchedulingPolicyRef struct {
	internal.Type__constructsIConstruct
}

