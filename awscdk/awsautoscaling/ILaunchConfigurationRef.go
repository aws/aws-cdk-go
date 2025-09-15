package awsautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LaunchConfiguration.
// Experimental.
type ILaunchConfigurationRef interface {
	constructs.IConstruct
	// A reference to a LaunchConfiguration resource.
	// Experimental.
	LaunchConfigurationRef() *LaunchConfigurationReference
}

// The jsii proxy for ILaunchConfigurationRef
type jsiiProxy_ILaunchConfigurationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ILaunchConfigurationRef) LaunchConfigurationRef() *LaunchConfigurationReference {
	var returns *LaunchConfigurationReference
	_jsii_.Get(
		j,
		"launchConfigurationRef",
		&returns,
	)
	return returns
}

