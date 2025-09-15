package awsmsk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmsk/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Configuration.
// Experimental.
type IConfigurationRef interface {
	constructs.IConstruct
	// A reference to a Configuration resource.
	// Experimental.
	ConfigurationRef() *ConfigurationReference
}

// The jsii proxy for IConfigurationRef
type jsiiProxy_IConfigurationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IConfigurationRef) ConfigurationRef() *ConfigurationReference {
	var returns *ConfigurationReference
	_jsii_.Get(
		j,
		"configurationRef",
		&returns,
	)
	return returns
}

