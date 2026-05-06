package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
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
//   	Definition: fmt.Sprintf("\n    permit(\n      principal,\n      action,\n      resource == AgentCore::Gateway::\"%v\"\n    );\n  ", gateway.GatewayArn),
//   	Description: jsii.String("Allow all actions on specific gateway (development)"),
//   	ValidationMode: agentcore.PolicyValidationMode_IGNORE_ALL_FINDINGS(),
//   })
//
//   // you can add multiple policies to the policy engine
//   policyEngine.AddPolicy(jsii.String("SpecificToolPolicy"), &AddPolicyOptions{
//   	Definition: fmt.Sprintf("\n    permit(\n      principal is AgentCore::OAuthUser,\n      action == AgentCore::Action::\"WeatherTool__get_forecast\",\n      resource == AgentCore::Gateway::\"%v\"\n    );\n  ", gateway.*GatewayArn),
//   	Description: jsii.String("Allow specific weather tool access"),
//   	ValidationMode: agentcore.PolicyValidationMode_FAIL_ON_ANY_FINDINGS(),
//   })
//
// Experimental.
type PolicyEngineMode interface {
	// The string value of the policy engine mode.
	// Experimental.
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


// Experimental.
func NewPolicyEngineMode(value *string) PolicyEngineMode {
	_init_.Initialize()

	if err := validateNewPolicyEngineModeParameters(value); err != nil {
		panic(err)
	}
	j := jsiiProxy_PolicyEngineMode{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.PolicyEngineMode",
		[]interface{}{value},
		&j,
	)

	return &j
}

// Experimental.
func NewPolicyEngineMode_Override(p PolicyEngineMode, value *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.PolicyEngineMode",
		[]interface{}{value},
		p,
	)
}

func PolicyEngineMode_ENFORCE() PolicyEngineMode {
	_init_.Initialize()
	var returns PolicyEngineMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.PolicyEngineMode",
		"ENFORCE",
		&returns,
	)
	return returns
}

func PolicyEngineMode_LOG_ONLY() PolicyEngineMode {
	_init_.Initialize()
	var returns PolicyEngineMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.PolicyEngineMode",
		"LOG_ONLY",
		&returns,
	)
	return returns
}

