package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StudioLifecycleConfig.
// Experimental.
type IStudioLifecycleConfigRef interface {
	constructs.IConstruct
	// A reference to a StudioLifecycleConfig resource.
	// Experimental.
	StudioLifecycleConfigRef() *StudioLifecycleConfigReference
}

// The jsii proxy for IStudioLifecycleConfigRef
type jsiiProxy_IStudioLifecycleConfigRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IStudioLifecycleConfigRef) StudioLifecycleConfigRef() *StudioLifecycleConfigReference {
	var returns *StudioLifecycleConfigReference
	_jsii_.Get(
		j,
		"studioLifecycleConfigRef",
		&returns,
	)
	return returns
}

