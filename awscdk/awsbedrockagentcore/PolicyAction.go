package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The action a policy statement applies to.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   awscdk.PolicyAction_Any() // action
//   awscdk.PolicyAction_One(jsii.String("AgentCore::Action::GetGateway")) // action == AgentCore::Action::"GetGateway"
//   awscdk.PolicyAction_AnyOf([]*string{
//   	jsii.String("AgentCore::Action::GetGateway"),
//   	jsii.String("AgentCore::Action::ListGateways"),
//   })
//
type PolicyAction interface {
}

// The jsii proxy struct for PolicyAction
type jsiiProxy_PolicyAction struct {
	_ byte // padding
}

// Any action.
func PolicyAction_Any() PolicyAction {
	_init_.Initialize()

	var returns PolicyAction

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyAction",
		"any",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Any one of the given actions.
func PolicyAction_AnyOf(actions *[]*string) PolicyAction {
	_init_.Initialize()

	if err := validatePolicyAction_AnyOfParameters(actions); err != nil {
		panic(err)
	}
	var returns PolicyAction

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyAction",
		"anyOf",
		[]interface{}{actions},
		&returns,
	)

	return returns
}

// One specific action.
func PolicyAction_One(action *string) PolicyAction {
	_init_.Initialize()

	if err := validatePolicyAction_OneParameters(action); err != nil {
		panic(err)
	}
	var returns PolicyAction

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyAction",
		"one",
		[]interface{}{action},
		&returns,
	)

	return returns
}

