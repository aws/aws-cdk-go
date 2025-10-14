package awssecurityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AutomationRule.
// Experimental.
type IAutomationRuleRef interface {
	constructs.IConstruct
	// A reference to a AutomationRule resource.
	// Experimental.
	AutomationRuleRef() *AutomationRuleReference
}

// The jsii proxy for IAutomationRuleRef
type jsiiProxy_IAutomationRuleRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAutomationRuleRef) AutomationRuleRef() *AutomationRuleReference {
	var returns *AutomationRuleReference
	_jsii_.Get(
		j,
		"automationRuleRef",
		&returns,
	)
	return returns
}

