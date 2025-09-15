package awsdatazone

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatazone/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ProjectProfile.
// Experimental.
type IProjectProfileRef interface {
	constructs.IConstruct
	// A reference to a ProjectProfile resource.
	// Experimental.
	ProjectProfileRef() *ProjectProfileReference
}

// The jsii proxy for IProjectProfileRef
type jsiiProxy_IProjectProfileRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IProjectProfileRef) ProjectProfileRef() *ProjectProfileReference {
	var returns *ProjectProfileReference
	_jsii_.Get(
		j,
		"projectProfileRef",
		&returns,
	)
	return returns
}

