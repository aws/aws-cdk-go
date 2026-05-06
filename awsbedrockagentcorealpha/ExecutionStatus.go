package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The execution status of an online evaluation configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   executionStatus := bedrock_agentcore_alpha.NewExecutionStatus(jsii.String("value"))
//
// Experimental.
type ExecutionStatus interface {
	// The string value of the execution status.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for ExecutionStatus
type jsiiProxy_ExecutionStatus struct {
	_ byte // padding
}

func (j *jsiiProxy_ExecutionStatus) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func NewExecutionStatus(value *string) ExecutionStatus {
	_init_.Initialize()

	if err := validateNewExecutionStatusParameters(value); err != nil {
		panic(err)
	}
	j := jsiiProxy_ExecutionStatus{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ExecutionStatus",
		[]interface{}{value},
		&j,
	)

	return &j
}

// Experimental.
func NewExecutionStatus_Override(e ExecutionStatus, value *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ExecutionStatus",
		[]interface{}{value},
		e,
	)
}

func ExecutionStatus_DISABLED() ExecutionStatus {
	_init_.Initialize()
	var returns ExecutionStatus
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ExecutionStatus",
		"DISABLED",
		&returns,
	)
	return returns
}

func ExecutionStatus_ENABLED() ExecutionStatus {
	_init_.Initialize()
	var returns ExecutionStatus
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ExecutionStatus",
		"ENABLED",
		&returns,
	)
	return returns
}

