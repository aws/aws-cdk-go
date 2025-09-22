package awsiotwireless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TaskDefinition.
// Experimental.
type ITaskDefinitionRef interface {
	constructs.IConstruct
	// A reference to a TaskDefinition resource.
	// Experimental.
	TaskDefinitionRef() *TaskDefinitionReference
}

// The jsii proxy for ITaskDefinitionRef
type jsiiProxy_ITaskDefinitionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITaskDefinitionRef) TaskDefinitionRef() *TaskDefinitionReference {
	var returns *TaskDefinitionReference
	_jsii_.Get(
		j,
		"taskDefinitionRef",
		&returns,
	)
	return returns
}

