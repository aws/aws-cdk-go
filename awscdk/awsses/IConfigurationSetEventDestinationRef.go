package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConfigurationSetEventDestination.
// Experimental.
type IConfigurationSetEventDestinationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IConfigurationSetEventDestinationRef
type jsiiProxy_IConfigurationSetEventDestinationRef struct {
	internal.Type__constructsIConstruct
}

