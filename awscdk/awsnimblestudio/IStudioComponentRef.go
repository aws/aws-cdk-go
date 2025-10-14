package awsnimblestudio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsnimblestudio/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StudioComponent.
// Experimental.
type IStudioComponentRef interface {
	constructs.IConstruct
	// A reference to a StudioComponent resource.
	// Experimental.
	StudioComponentRef() *StudioComponentReference
}

// The jsii proxy for IStudioComponentRef
type jsiiProxy_IStudioComponentRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IStudioComponentRef) StudioComponentRef() *StudioComponentReference {
	var returns *StudioComponentReference
	_jsii_.Get(
		j,
		"studioComponentRef",
		&returns,
	)
	return returns
}

