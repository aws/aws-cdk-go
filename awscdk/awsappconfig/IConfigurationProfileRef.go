package awsappconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConfigurationProfile.
// Experimental.
type IConfigurationProfileRef interface {
	constructs.IConstruct
	// A reference to a ConfigurationProfile resource.
	// Experimental.
	ConfigurationProfileRef() *ConfigurationProfileReference
}

// The jsii proxy for IConfigurationProfileRef
type jsiiProxy_IConfigurationProfileRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IConfigurationProfileRef) ConfigurationProfileRef() *ConfigurationProfileReference {
	var returns *ConfigurationProfileReference
	_jsii_.Get(
		j,
		"configurationProfileRef",
		&returns,
	)
	return returns
}

