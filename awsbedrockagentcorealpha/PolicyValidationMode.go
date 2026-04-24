package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Validation mode for Cedar policy definitions.
//
// Example:
//   gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-gateway"),
//   })
//
//   policyEngine := agentcore.NewPolicyEngine(this, jsii.String("MyPolicyEngine"), &PolicyEngineProps{
//   	PolicyEngineName: jsii.String("my_policy_engine"),
//   })
//
//   allowAllPolicy := agentcore.NewPolicy(this, jsii.String("AllowAllPolicy"), &PolicyProps{
//   	PolicyEngine: policyEngine,
//   	PolicyName: jsii.String("allow_all"),
//   	Statement: agentcore.PolicyStatement_Permit().ForAllPrincipals().OnAllActions().OnResource(jsii.String("AgentCore::Gateway"), gateway.GatewayArn),
//   	Description: jsii.String("Allow all actions on specific gateway (development only)"),
//   	ValidationMode: agentcore.PolicyValidationMode_IGNORE_ALL_FINDINGS(),
//   })
//
// Experimental.
type PolicyValidationMode interface {
	// The string value of the validation mode.
	// Experimental.
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


// Experimental.
func NewPolicyValidationMode(value *string) PolicyValidationMode {
	_init_.Initialize()

	if err := validateNewPolicyValidationModeParameters(value); err != nil {
		panic(err)
	}
	j := jsiiProxy_PolicyValidationMode{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.PolicyValidationMode",
		[]interface{}{value},
		&j,
	)

	return &j
}

// Experimental.
func NewPolicyValidationMode_Override(p PolicyValidationMode, value *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.PolicyValidationMode",
		[]interface{}{value},
		p,
	)
}

func PolicyValidationMode_FAIL_ON_ANY_FINDINGS() PolicyValidationMode {
	_init_.Initialize()
	var returns PolicyValidationMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.PolicyValidationMode",
		"FAIL_ON_ANY_FINDINGS",
		&returns,
	)
	return returns
}

func PolicyValidationMode_IGNORE_ALL_FINDINGS() PolicyValidationMode {
	_init_.Initialize()
	var returns PolicyValidationMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.PolicyValidationMode",
		"IGNORE_ALL_FINDINGS",
		&returns,
	)
	return returns
}

