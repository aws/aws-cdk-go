package awscognito

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LogDeliveryConfiguration.
// Experimental.
type ILogDeliveryConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for ILogDeliveryConfigurationRef
type jsiiProxy_ILogDeliveryConfigurationRef struct {
	internal.Type__constructsIConstruct
}

