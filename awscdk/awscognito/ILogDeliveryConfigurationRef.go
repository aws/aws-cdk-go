package awscognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LogDeliveryConfiguration.
// Experimental.
type ILogDeliveryConfigurationRef interface {
	constructs.IConstruct
	// A reference to a LogDeliveryConfiguration resource.
	// Experimental.
	LogDeliveryConfigurationRef() *LogDeliveryConfigurationReference
}

// The jsii proxy for ILogDeliveryConfigurationRef
type jsiiProxy_ILogDeliveryConfigurationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ILogDeliveryConfigurationRef) LogDeliveryConfigurationRef() *LogDeliveryConfigurationReference {
	var returns *LogDeliveryConfigurationReference
	_jsii_.Get(
		j,
		"logDeliveryConfigurationRef",
		&returns,
	)
	return returns
}

