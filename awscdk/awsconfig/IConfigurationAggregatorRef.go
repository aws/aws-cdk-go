package awsconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConfigurationAggregator.
// Experimental.
type IConfigurationAggregatorRef interface {
	constructs.IConstruct
}

// The jsii proxy for IConfigurationAggregatorRef
type jsiiProxy_IConfigurationAggregatorRef struct {
	internal.Type__constructsIConstruct
}

