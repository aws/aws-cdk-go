package awsssmquicksetup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssmquicksetup/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConfigurationManager.
// Experimental.
type IConfigurationManagerRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ConfigurationManager resource.
	// Experimental.
	ConfigurationManagerRef() *ConfigurationManagerReference
}

// The jsii proxy for IConfigurationManagerRef
type jsiiProxy_IConfigurationManagerRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IConfigurationManagerRef) ConfigurationManagerRef() *ConfigurationManagerReference {
	var returns *ConfigurationManagerReference
	_jsii_.Get(
		j,
		"configurationManagerRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfigurationManagerRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfigurationManagerRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

