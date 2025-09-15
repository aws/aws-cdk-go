package awssecurityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AutomationRuleV2.
// Experimental.
type IAutomationRuleV2Ref interface {
	constructs.IConstruct
	// A reference to a AutomationRuleV2 resource.
	// Experimental.
	AutomationRuleV2Ref() *AutomationRuleV2Reference
}

// The jsii proxy for IAutomationRuleV2Ref
type jsiiProxy_IAutomationRuleV2Ref struct {
	internal.Type__constructsIConstruct
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

