package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The enforcement mode for a policy engine associated with a gateway.
//
// Example:
//   // Create a Policy engine
//   policyEngine := agentcore.NewPolicyEngine(this, jsii.String("MyPolicyEngine"), &PolicyEngineProps{
//   	PolicyEngineName: jsii.String("my_policy_engine"),
//   	Description: jsii.String("Policy engine for access control"),
//   })
//
//   gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-gateway"),
//   	PolicyEngineConfiguration: &GatewayPolicyEngineConfig{
//   		PolicyEngine: policyEngine,
//   		Mode: agentcore.PolicyEngineMode_ENFORCE(),
//   	},
//   })
//
//   // Add policy to policy engine
//   policyEngine.AddPolicy(jsii.String("AllowAllActions"), &AddPolicyOptions{
//   	Statement: agentcore.PolicyStatement_FromCedar(fmt.Sprintf("\n    permit(\n      principal,\n      action,\n      resource == AgentCore::Gateway::\"%v\"\n    );\n  ", gateway.GatewayArn)),
//   	Description: jsii.String("Allow all actions on specific gateway (development)"),
//   	ValidationMode: agentcore.PolicyValidationMode_IGNORE_ALL_FINDINGS(),
//   })
//
//   // you can add multiple policies to the policy engine
//   policyEngine.AddPolicy(jsii.String("SpecificToolPolicy"), &AddPolicyOptions{
//   	Statement: agentcore.PolicyStatement_*FromCedar(fmt.Sprintf("\n    permit(\n      principal is AgentCore::OAuthUser,\n      action == AgentCore::Action::\"WeatherTool__get_forecast\",\n      resource == AgentCore::Gateway::\"%v\"\n    );\n  ", gateway.*GatewayArn)),
//   	Description: jsii.String("Allow specific weather tool access"),
//   	ValidationMode: agentcore.PolicyValidationMode_FAIL_ON_ANY_FINDINGS(),
//   })
//
type PolicyEngineMode interface {
	// The string value of the policy engine mode.
	Value() *string
}

// The jsii proxy struct for PolicyEngineMode
type jsiiProxy_PolicyEngineMode struct {
	_ byte // padding
}

func (j *jsiiProxy_PolicyEngineMode) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// A policy engine mode that this version of the CDK does not model.
//
// Prefer the static members above. Use this when the service has added a mode
// that the CDK has no member for yet: the CloudFormation registry schemas that
// validate the synthesized template ship with `aws-cdk-lib` and refresh on
// release, while the members above are added by hand, so a released CDK can
// accept a mode before it models one.
func PolicyEngineMode_Of(value *string) PolicyEngineMode {
	_init_.Initialize()

	if err := validatePolicyEngineMode_OfParameters(value); err != nil {
		panic(err)
	}
	var returns PolicyEngineMode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyEngineMode",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func PolicyEngineMode_ENFORCE() PolicyEngineMode {
	_init_.Initialize()
	var returns PolicyEngineMode
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyEngineMode",
		"ENFORCE",
		&returns,
	)
	return returns
}

func PolicyEngineMode_LOG_ONLY() PolicyEngineMode {
	_init_.Initialize()
	var returns PolicyEngineMode
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyEngineMode",
		"LOG_ONLY",
		&returns,
	)
	return returns
}

