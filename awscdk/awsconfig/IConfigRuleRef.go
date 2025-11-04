package awsconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConfigRule.
// Experimental.
type IConfigRuleRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ConfigRule resource.
	// Experimental.
	ConfigRuleRef() *ConfigRuleReference
}

// The jsii proxy for IConfigRuleRef
type jsiiProxy_IConfigRuleRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IConfigRuleRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

