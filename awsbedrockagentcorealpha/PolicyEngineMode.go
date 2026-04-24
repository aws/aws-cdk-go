package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The enforcement mode for a policy engine associated with a gateway.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   policyEngineMode := bedrock_agentcore_alpha.NewPolicyEngineMode(jsii.String("value"))
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

