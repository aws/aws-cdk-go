package awssecurityhub

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AutomationRule.
// Experimental.
type IAutomationRuleRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAutomationRuleRef
type jsiiProxy_IAutomationRuleRef struct {
	internal.Type__constructsIConstruct
}

