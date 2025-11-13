package interfacesawsconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConfigRule.
// Experimental.
type IConfigRuleRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ConfigRule resource.
	// Experimental.
	ConfigRuleRef() *ConfigRuleReference
}

// The jsii proxy for IConfigRuleRef
type jsiiProxy_IConfigRuleRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IConfigRuleRef) ConfigRuleRef() *ConfigRuleReference {
	var returns *ConfigRuleReference
	_jsii_.Get(
		j,
		"configRuleRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfigRuleRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfigRuleRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

