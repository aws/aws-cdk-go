package awspinpointemail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpointemail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConfigurationSetEventDestination.
// Experimental.
type IConfigurationSetEventDestinationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ConfigurationSetEventDestination resource.
	// Experimental.
	ConfigurationSetEventDestinationRef() *ConfigurationSetEventDestinationReference
}

// The jsii proxy for IConfigurationSetEventDestinationRef
type jsiiProxy_IConfigurationSetEventDestinationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IConfigurationSetEventDestinationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfigurationSetEventDestinationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

