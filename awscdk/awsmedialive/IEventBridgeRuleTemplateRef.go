package awsmedialive

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventBridgeRuleTemplate.
// Experimental.
type IEventBridgeRuleTemplateRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEventBridgeRuleTemplateRef
type jsiiProxy_IEventBridgeRuleTemplateRef struct {
	internal.Type__constructsIConstruct
}

