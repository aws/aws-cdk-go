package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Log types for AgentCore Runtime observability.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   logType := bedrock_agentcore_alpha.LogType_Of(jsii.String("value"))
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type LogType interface {
	// The log type value.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Value() *string
}

// The jsii proxy struct for LogType
type jsiiProxy_LogType struct {
	_ byte // padding
}

func (j *jsiiProxy_LogType) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// A custom log type value.
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func LogType_Of(value *string) LogType {
	_init_.Initialize()

	if err := validateLogType_OfParameters(value); err != nil {
		panic(err)
	}
	var returns LogType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.LogType",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func LogType_APPLICATION_LOGS() LogType {
	_init_.Initialize()
	var returns LogType
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.LogType",
		"APPLICATION_LOGS",
		&returns,
	)
	return returns
}

func LogType_USAGE_LOGS() LogType {
	_init_.Initialize()
	var returns LogType
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.LogType",
		"USAGE_LOGS",
		&returns,
	)
	return returns
}

