package awsimagebuilder

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsimagebuilder/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DistributionConfiguration.
// Experimental.
type IDistributionConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDistributionConfigurationRef
type jsiiProxy_IDistributionConfigurationRef struct {
	internal.Type__constructsIConstruct
}

