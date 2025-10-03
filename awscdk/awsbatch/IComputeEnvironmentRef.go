package awsbatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ComputeEnvironment.
// Experimental.
type IComputeEnvironmentRef interface {
	constructs.IConstruct
}

// The jsii proxy for IComputeEnvironmentRef
type jsiiProxy_IComputeEnvironmentRef struct {
	internal.Type__constructsIConstruct
}

