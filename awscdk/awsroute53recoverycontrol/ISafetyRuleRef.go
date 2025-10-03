package awsroute53recoverycontrol

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53recoverycontrol/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SafetyRule.
// Experimental.
type ISafetyRuleRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISafetyRuleRef
type jsiiProxy_ISafetyRuleRef struct {
	internal.Type__constructsIConstruct
}

