package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The execution status of an online evaluation configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   executionStatus := awscdk.Aws_bedrockagentcore.NewExecutionStatus(jsii.String("value"))
//
type ExecutionStatus interface {
	// The string value of the execution status.
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


func NewExecutionStatus(value *string) ExecutionStatus {
	_init_.Initialize()

	if err := validateNewExecutionStatusParameters(value); err != nil {
		panic(err)
	}
	j := jsiiProxy_ExecutionStatus{}

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrockagentcore.ExecutionStatus",
		[]interface{}{value},
		&j,
	)

	return &j
}

func NewExecutionStatus_Override(e ExecutionStatus, value *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrockagentcore.ExecutionStatus",
		[]interface{}{value},
		e,
	)
}

func ExecutionStatus_DISABLED() ExecutionStatus {
	_init_.Initialize()
	var returns ExecutionStatus
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.ExecutionStatus",
		"DISABLED",
		&returns,
	)
	return returns
}

func ExecutionStatus_ENABLED() ExecutionStatus {
	_init_.Initialize()
	var returns ExecutionStatus
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.ExecutionStatus",
		"ENABLED",
		&returns,
	)
	return returns
}

