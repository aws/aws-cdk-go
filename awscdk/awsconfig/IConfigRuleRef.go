package awsconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConfigRule.
// Experimental.
type IConfigRuleRef interface {
	constructs.IConstruct
	// A reference to a ConfigRule resource.
	// Experimental.
	ConfigRuleRef() *ConfigRuleReference
}

// The jsii proxy for IConfigRuleRef
type jsiiProxy_IConfigRuleRef struct {
	internal.Type__constructsIConstruct
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

