package awsmedialive

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventBridgeRuleTemplateGroup.
// Experimental.
type IEventBridgeRuleTemplateGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEventBridgeRuleTemplateGroupRef
type jsiiProxy_IEventBridgeRuleTemplateGroupRef struct {
	internal.Type__constructsIConstruct
}

