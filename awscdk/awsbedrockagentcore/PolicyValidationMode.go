package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Validation mode for Cedar policy definitions.
//
// Example:
//   var policyEngine PolicyEngine
//   var gateway Gateway
//
//
//   // Allow access unless the user is suspended
//   policyWithUnless := agentcore.NewPolicy(this, jsii.String("UnlessPolicy"), &PolicyProps{
//   	PolicyEngine: policyEngine,
//   	PolicyName: jsii.String("unless_suspended"),
//   	Statement: agentcore.NewPolicyStatement(&PolicyStatementProps{
//   		Effect: agentcore.PolicyEffect_PERMIT,
//   		Principal: agentcore.PolicyPrincipal_EntityType(jsii.String("AgentCore::OAuthUser")),
//   		Action: agentcore.PolicyAction_Any(),
//   		Resource: agentcore.PolicyResource_Instance(jsii.String("AgentCore::Gateway"), gateway.GatewayArn),
//   		Unless: []PolicyCondition{
//   			agentcore.PolicyCondition_BooleanEquals(agentcore.PolicyAttribute_Principal(jsii.String("suspended")), jsii.Boolean(true)),
//   		},
//   	}),
//   	Description: jsii.String("Allow all actions unless user is suspended"),
//   	ValidationMode: agentcore.PolicyValidationMode_FAIL_ON_ANY_FINDINGS(),
//   })
//
type PolicyValidationMode interface {
	// The string value of the validation mode.
	Value() *string
}

// The jsii proxy struct for PolicyValidationMode
type jsiiProxy_PolicyValidationMode struct {
	_ byte // padding
}

func (j *jsiiProxy_PolicyValidationMode) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// A validation mode that this version of the CDK does not model.
//
// Prefer the static members above. Use this when the service has added a mode
// that the CDK has no member for yet: the CloudFormation registry schemas that
// validate the synthesized template ship with `aws-cdk-lib` and refresh on
// release, while the members above are added by hand, so a released CDK can
// accept a mode before it models one.
func PolicyValidationMode_Of(value *string) PolicyValidationMode {
	_init_.Initialize()

	if err := validatePolicyValidationMode_OfParameters(value); err != nil {
		panic(err)
	}
	var returns PolicyValidationMode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyValidationMode",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func PolicyValidationMode_FAIL_ON_ANY_FINDINGS() PolicyValidationMode {
	_init_.Initialize()
	var returns PolicyValidationMode
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyValidationMode",
		"FAIL_ON_ANY_FINDINGS",
		&returns,
	)
	return returns
}

func PolicyValidationMode_IGNORE_ALL_FINDINGS() PolicyValidationMode {
	_init_.Initialize()
	var returns PolicyValidationMode
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyValidationMode",
		"IGNORE_ALL_FINDINGS",
		&returns,
	)
	return returns
}

