package awsconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConfigurationAggregator.
// Experimental.
type IConfigurationAggregatorRef interface {
	constructs.IConstruct
	// A reference to a ConfigurationAggregator resource.
	// Experimental.
	ConfigurationAggregatorRef() *ConfigurationAggregatorReference
}

// The jsii proxy for IConfigurationAggregatorRef
type jsiiProxy_IConfigurationAggregatorRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IConfigurationAggregatorRef) ConfigurationAggregatorRef() *ConfigurationAggregatorReference {
	var returns *ConfigurationAggregatorReference
	_jsii_.Get(
		j,
		"configurationAggregatorRef",
		&returns,
	)
	return returns
}

