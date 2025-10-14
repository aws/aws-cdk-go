package awsrekognition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsrekognition/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Project.
// Experimental.
type IProjectRef interface {
	constructs.IConstruct
	// A reference to a Project resource.
	// Experimental.
	ProjectRef() *ProjectReference
}

// The jsii proxy for IProjectRef
type jsiiProxy_IProjectRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IProjectRef) ProjectRef() *ProjectReference {
	var returns *ProjectReference
	_jsii_.Get(
		j,
		"projectRef",
		&returns,
	)
	return returns
}

