package interfacesawssecurityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AutomationRuleV2.
// Experimental.
type IAutomationRuleV2Ref interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AutomationRuleV2 resource.
	// Experimental.
	AutomationRuleV2Ref() *AutomationRuleV2Reference
}

// The jsii proxy for IAutomationRuleV2Ref
type jsiiProxy_IAutomationRuleV2Ref struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IAutomationRuleV2Ref) AutomationRuleV2Ref() *AutomationRuleV2Reference {
	var returns *AutomationRuleV2Reference
	_jsii_.Get(
		j,
		"automationRuleV2Ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAutomationRuleV2Ref) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAutomationRuleV2Ref) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

