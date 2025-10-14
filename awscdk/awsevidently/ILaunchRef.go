package awsevidently

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevidently/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Launch.
// Experimental.
type ILaunchRef interface {
	constructs.IConstruct
	// A reference to a Launch resource.
	// Experimental.
	LaunchRef() *LaunchReference
}

// The jsii proxy for ILaunchRef
type jsiiProxy_ILaunchRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ILaunchRef) LaunchRef() *LaunchReference {
	var returns *LaunchReference
	_jsii_.Get(
		j,
		"launchRef",
		&returns,
	)
	return returns
}

