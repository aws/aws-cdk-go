package awsimagebuilder

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsimagebuilder/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InfrastructureConfiguration.
// Experimental.
type IInfrastructureConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IInfrastructureConfigurationRef
type jsiiProxy_IInfrastructureConfigurationRef struct {
	internal.Type__constructsIConstruct
}

