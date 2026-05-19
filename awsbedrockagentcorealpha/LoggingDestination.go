package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Represents a logging destination for AgentCore Runtime.
//
// Use the static factory methods to create instances:
// - `LoggingDestination.cloudWatchLogs(logGroup)` - Send logs to CloudWatch Logs
// - `LoggingDestination.s3(bucket)` - Send logs to S3
// - `LoggingDestination.firehose(stream)` - Send logs to Kinesis Data Firehose
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var logGroup LogGroup
//
//   loggingDestination := bedrock_agentcore_alpha.LoggingDestination_CloudWatchLogs(logGroup)
//
// Experimental.
type LoggingDestination interface {
}

// The jsii proxy struct for LoggingDestination
type jsiiProxy_LoggingDestination struct {
	_ byte // padding
}

// Experimental.
func NewLoggingDestination_Override(l LoggingDestination) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.LoggingDestination",
		nil, // no parameters
		l,
	)
}

// Create a logging destination that sends logs to a CloudWatch Log Group.
// Experimental.
func LoggingDestination_CloudWatchLogs(logGroup awslogs.ILogGroup) LoggingDestination {
	_init_.Initialize()

	if err := validateLoggingDestination_CloudWatchLogsParameters(logGroup); err != nil {
		panic(err)
	}
	var returns LoggingDestination

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.LoggingDestination",
		"cloudWatchLogs",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

// Create a logging destination that sends logs to a Kinesis Data Firehose delivery stream.
// Experimental.
func LoggingDestination_Firehose(stream awskinesisfirehose.IDeliveryStream) LoggingDestination {
	_init_.Initialize()

	if err := validateLoggingDestination_FirehoseParameters(stream); err != nil {
		panic(err)
	}
	var returns LoggingDestination

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.LoggingDestination",
		"firehose",
		[]interface{}{stream},
		&returns,
	)

	return returns
}

// Create a logging destination that sends logs to an S3 bucket.
// Experimental.
func LoggingDestination_S3(bucket awss3.IBucket) LoggingDestination {
	_init_.Initialize()

	if err := validateLoggingDestination_S3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns LoggingDestination

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.LoggingDestination",
		"s3",
		[]interface{}{bucket},
		&returns,
	)

	return returns
}

