package awsroute53recoverycontrol

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53recoverycontrol/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SafetyRule.
// Experimental.
type ISafetyRuleRef interface {
	constructs.IConstruct
	// A reference to a SafetyRule resource.
	// Experimental.
	SafetyRuleRef() *SafetyRuleReference
}

// The jsii proxy for ISafetyRuleRef
type jsiiProxy_ISafetyRuleRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISafetyRuleRef) SafetyRuleRef() *SafetyRuleReference {
	var returns *SafetyRuleReference
	_jsii_.Get(
		j,
		"safetyRuleRef",
		&returns,
	)
	return returns
}

