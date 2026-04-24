package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Log types for AgentCore Runtime observability.
//
// Example:
//   repository := ecr.NewRepository(this, jsii.String("TestRepository"), &RepositoryProps{
//   	RepositoryName: jsii.String("test-agent-runtime"),
//   })
//
//   agentRuntimeArtifact := agentcore.AgentRuntimeArtifact_FromEcrRepository(repository, jsii.String("v1.0.0"))
//
//   // Create logging destinations
//   logGroup := logs.NewLogGroup(this, jsii.String("RuntimeLogGroup"))
//   logBucket := s3.NewBucket(this, jsii.String("RuntimeLogBucket"))
//   firehoseStream := firehose.NewDeliveryStream(this, jsii.String("RuntimeLogStream"), &DeliveryStreamProps{
//   	Destination: firehose.NewS3Bucket(logBucket),
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
//   		&LoggingConfig{
//   			LogType: agentcore.LogType_APPLICATION_LOGS(),
//   			Destination: agentcore.LoggingDestination_S3(logBucket),
//   		},
//   		&LoggingConfig{
//   			LogType: agentcore.LogType_APPLICATION_LOGS(),
//   			Destination: agentcore.LoggingDestination_Firehose(firehoseStream),
//   		},
//   	},
//   })
//
// Experimental.
type LogType interface {
	// The log type value.
	// Experimental.
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
// Experimental.
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

