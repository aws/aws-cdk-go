package awscodeguruprofiler

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodeguruprofiler/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ProfilingGroup.
// Experimental.
type IProfilingGroupRef interface {
	constructs.IConstruct
	// A reference to a ProfilingGroup resource.
	// Experimental.
	ProfilingGroupRef() *ProfilingGroupReference
}

// The jsii proxy for IProfilingGroupRef
type jsiiProxy_IProfilingGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IProfilingGroupRef) ProfilingGroupRef() *ProfilingGroupReference {
	var returns *ProfilingGroupReference
	_jsii_.Get(
		j,
		"profilingGroupRef",
		&returns,
	)
	return returns
}

