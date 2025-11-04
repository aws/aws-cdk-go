package awssecurityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConfigurationPolicy.
// Experimental.
type IConfigurationPolicyRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ConfigurationPolicy resource.
	// Experimental.
	ConfigurationPolicyRef() *ConfigurationPolicyReference
}

// The jsii proxy for IConfigurationPolicyRef
type jsiiProxy_IConfigurationPolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IConfigurationPolicyRef) ConfigurationPolicyRef() *ConfigurationPolicyReference {
	var returns *ConfigurationPolicyReference
	_jsii_.Get(
		j,
		"configurationPolicyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfigurationPolicyRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfigurationPolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

