package awsnimblestudio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsnimblestudio/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Studio.
// Experimental.
type IStudioRef interface {
	constructs.IConstruct
	// A reference to a Studio resource.
	// Experimental.
	StudioRef() *StudioReference
}

// The jsii proxy for IStudioRef
type jsiiProxy_IStudioRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IStudioRef) StudioRef() *StudioReference {
	var returns *StudioReference
	_jsii_.Get(
		j,
		"studioRef",
		&returns,
	)
	return returns
}

