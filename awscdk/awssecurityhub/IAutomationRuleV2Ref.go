package awssecurityhub

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AutomationRuleV2.
// Experimental.
type IAutomationRuleV2Ref interface {
	constructs.IConstruct
}

// The jsii proxy for IAutomationRuleV2Ref
type jsiiProxy_IAutomationRuleV2Ref struct {
	internal.Type__constructsIConstruct
}

