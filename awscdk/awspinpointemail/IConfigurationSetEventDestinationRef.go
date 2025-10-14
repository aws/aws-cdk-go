package awspinpointemail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpointemail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConfigurationSetEventDestination.
// Experimental.
type IConfigurationSetEventDestinationRef interface {
	constructs.IConstruct
	// A reference to a ConfigurationSetEventDestination resource.
	// Experimental.
	ConfigurationSetEventDestinationRef() *ConfigurationSetEventDestinationReference
}

// The jsii proxy for IConfigurationSetEventDestinationRef
type jsiiProxy_IConfigurationSetEventDestinationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IConfigurationSetEventDestinationRef) ConfigurationSetEventDestinationRef() *ConfigurationSetEventDestinationReference {
	var returns *ConfigurationSetEventDestinationReference
	_jsii_.Get(
		j,
		"configurationSetEventDestinationRef",
		&returns,
	)
	return returns
}

