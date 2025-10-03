package awsapprunner

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapprunner/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ObservabilityConfiguration.
// Experimental.
type IObservabilityConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IObservabilityConfigurationRef
type jsiiProxy_IObservabilityConfigurationRef struct {
	internal.Type__constructsIConstruct
}

