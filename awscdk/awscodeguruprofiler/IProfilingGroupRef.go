package awscodeguruprofiler

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodeguruprofiler/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ProfilingGroup.
// Experimental.
type IProfilingGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IProfilingGroupRef
type jsiiProxy_IProfilingGroupRef struct {
	internal.Type__constructsIConstruct
}

