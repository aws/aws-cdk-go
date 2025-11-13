package interfacesawsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Job.
// Experimental.
type IJobRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Job resource.
	// Experimental.
	JobRef() *JobReference
}

// The jsii proxy for IJobRef
type jsiiProxy_IJobRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IJobRef) JobRef() *JobReference {
	var returns *JobReference
	_jsii_.Get(
		j,
		"jobRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJobRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJobRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

