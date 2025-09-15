package awsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a JobDefinition.
// Experimental.
type IJobDefinitionRef interface {
	constructs.IConstruct
	// A reference to a JobDefinition resource.
	// Experimental.
	JobDefinitionRef() *JobDefinitionReference
}

// The jsii proxy for IJobDefinitionRef
type jsiiProxy_IJobDefinitionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IJobDefinitionRef) JobDefinitionRef() *JobDefinitionReference {
	var returns *JobDefinitionReference
	_jsii_.Get(
		j,
		"jobDefinitionRef",
		&returns,
	)
	return returns
}

