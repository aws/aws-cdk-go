package awsapprunner

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapprunner/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AutoScalingConfiguration.
// Experimental.
type IAutoScalingConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAutoScalingConfigurationRef
type jsiiProxy_IAutoScalingConfigurationRef struct {
	internal.Type__constructsIConstruct
}

