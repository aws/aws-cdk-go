package awsnimblestudio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsnimblestudio/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LaunchProfile.
// Experimental.
type ILaunchProfileRef interface {
	constructs.IConstruct
	// A reference to a LaunchProfile resource.
	// Experimental.
	LaunchProfileRef() *LaunchProfileReference
}

// The jsii proxy for ILaunchProfileRef
type jsiiProxy_ILaunchProfileRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ILaunchProfileRef) LaunchProfileRef() *LaunchProfileReference {
	var returns *LaunchProfileReference
	_jsii_.Get(
		j,
		"launchProfileRef",
		&returns,
	)
	return returns
}

