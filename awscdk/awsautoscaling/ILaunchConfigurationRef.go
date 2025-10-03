package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LaunchConfiguration.
// Experimental.
type ILaunchConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for ILaunchConfigurationRef
type jsiiProxy_ILaunchConfigurationRef struct {
	internal.Type__constructsIConstruct
}

