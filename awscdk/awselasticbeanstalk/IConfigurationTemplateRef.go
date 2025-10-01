package awselasticbeanstalk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticbeanstalk/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConfigurationTemplate.
// Experimental.
type IConfigurationTemplateRef interface {
	constructs.IConstruct
	// A reference to a ConfigurationTemplate resource.
	// Experimental.
	ConfigurationTemplateRef() *ConfigurationTemplateReference
}

// The jsii proxy for IConfigurationTemplateRef
type jsiiProxy_IConfigurationTemplateRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IConfigurationTemplateRef) ConfigurationTemplateRef() *ConfigurationTemplateReference {
	var returns *ConfigurationTemplateReference
	_jsii_.Get(
		j,
		"configurationTemplateRef",
		&returns,
	)
	return returns
}

