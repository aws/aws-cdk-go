package interfacesawselasticbeanstalk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawselasticbeanstalk/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConfigurationTemplate.
// Experimental.
type IConfigurationTemplateRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ConfigurationTemplate resource.
	// Experimental.
	ConfigurationTemplateRef() *ConfigurationTemplateReference
}

// The jsii proxy for IConfigurationTemplateRef
type jsiiProxy_IConfigurationTemplateRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_IConfigurationTemplateRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfigurationTemplateRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

