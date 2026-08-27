package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
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
//   repository := ecr.NewRepository(this, jsii.String("TestRepository"), &RepositoryProps{
//   	RepositoryName: jsii.String("test-agent-runtime"),
//   })
//
//   agentRuntimeArtifact := agentcore.AgentRuntimeArtifact_FromEcrRepository(repository, jsii.String("v1.0.0"))
//
//   // Use a /aws/vendedlogs/ log group for same-account delivery without explicit resource policy
//   logGroup := logs.NewLogGroup(this, jsii.String("RuntimeLogGroup"), &LogGroupProps{
//   	LogGroupName: jsii.String("/aws/vendedlogs/bedrock-agentcore/my-runtime"),
//   })
//
//   agentcore.NewRuntime(this, jsii.String("test-runtime"), &RuntimeProps{
//   	RuntimeName: jsii.String("test_runtime"),
//   	AgentRuntimeArtifact: agentRuntimeArtifact,
//   	TracingEnabled: jsii.Boolean(true),
//   	LoggingConfigs: []LoggingConfig{
//   		&LoggingConfig{
//   			LogType: agentcore.LogType_APPLICATION_LOGS(),
//   			Destination: agentcore.LoggingDestination_CloudWatchLogs(logGroup),
//   		},
//   	},
//   	ManageDeliveryResourcePolicy: jsii.Boolean(false),
//   })
//
type LoggingDestination interface {
}

// The jsii proxy struct for LoggingDestination
type jsiiProxy_LoggingDestination struct {
	_ byte // padding
}

func NewLoggingDestination_Override(l LoggingDestination) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrockagentcore.LoggingDestination",
		nil, // no parameters
		l,
	)
}

// Create a logging destination that sends logs to a CloudWatch Log Group.
func LoggingDestination_CloudWatchLogs(logGroup awslogs.ILogGroup) LoggingDestination {
	_init_.Initialize()

	if err := validateLoggingDestination_CloudWatchLogsParameters(logGroup); err != nil {
		panic(err)
	}
	var returns LoggingDestination

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.LoggingDestination",
		"cloudWatchLogs",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

// Create a logging destination that sends logs to a Kinesis Data Firehose delivery stream.
func LoggingDestination_Firehose(stream awskinesisfirehose.IDeliveryStream) LoggingDestination {
	_init_.Initialize()

	if err := validateLoggingDestination_FirehoseParameters(stream); err != nil {
		panic(err)
	}
	var returns LoggingDestination

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.LoggingDestination",
		"firehose",
		[]interface{}{stream},
		&returns,
	)

	return returns
}

// Create a logging destination that sends logs to an S3 bucket.
func LoggingDestination_S3(bucket awss3.IBucket) LoggingDestination {
	_init_.Initialize()

	if err := validateLoggingDestination_S3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns LoggingDestination

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.LoggingDestination",
		"s3",
		[]interface{}{bucket},
		&returns,
	)

	return returns
}

