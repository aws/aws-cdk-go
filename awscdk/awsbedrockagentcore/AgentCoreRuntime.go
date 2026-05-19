package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Bedrock AgentCore runtime environment for code execution Allowed values: PYTHON_3_10 | PYTHON_3_11 | PYTHON_3_12 | PYTHON_3_13 | PYTHON_3_14 | NODE_22.
//
// Example:
//   // S3 bucket containing the agent core
//   codeBucket := s3.NewBucket(this, jsii.String("AgentCode"), &BucketProps{
//   	BucketName: jsii.String("my-code-bucket"),
//   	RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
//   })
//
//   // the bucket above needs to contain the agent code
//
//   agentRuntimeArtifact := agentcore.AgentRuntimeArtifact_FromS3(&Location{
//   	BucketName: codeBucket.bucketName,
//   	ObjectKey: jsii.String("deployment_package.zip"),
//   }, agentcore.AgentCoreRuntime_PYTHON_3_12(), []*string{
//   	jsii.String("opentelemetry-instrument"),
//   	jsii.String("main.py"),
//   })
//
//   runtimeInstance := agentcore.NewRuntime(this, jsii.String("MyAgentRuntime"), &RuntimeProps{
//   	RuntimeName: jsii.String("myAgent"),
//   	AgentRuntimeArtifact: agentRuntimeArtifact,
//   })
//
type AgentCoreRuntime interface {
	// The runtime string value.
	Value() *string
	// Returns the runtime string value.
	ToString() *string
}

// The jsii proxy struct for AgentCoreRuntime
type jsiiProxy_AgentCoreRuntime struct {
	_ byte // padding
}

func (j *jsiiProxy_AgentCoreRuntime) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Create a custom runtime value for runtimes not yet defined in this enum.
func AgentCoreRuntime_Of(value *string) AgentCoreRuntime {
	_init_.Initialize()

	if err := validateAgentCoreRuntime_OfParameters(value); err != nil {
		panic(err)
	}
	var returns AgentCoreRuntime

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.AgentCoreRuntime",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func AgentCoreRuntime_NODE_22() AgentCoreRuntime {
	_init_.Initialize()
	var returns AgentCoreRuntime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.AgentCoreRuntime",
		"NODE_22",
		&returns,
	)
	return returns
}

func AgentCoreRuntime_PYTHON_3_10() AgentCoreRuntime {
	_init_.Initialize()
	var returns AgentCoreRuntime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.AgentCoreRuntime",
		"PYTHON_3_10",
		&returns,
	)
	return returns
}

func AgentCoreRuntime_PYTHON_3_11() AgentCoreRuntime {
	_init_.Initialize()
	var returns AgentCoreRuntime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.AgentCoreRuntime",
		"PYTHON_3_11",
		&returns,
	)
	return returns
}

func AgentCoreRuntime_PYTHON_3_12() AgentCoreRuntime {
	_init_.Initialize()
	var returns AgentCoreRuntime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.AgentCoreRuntime",
		"PYTHON_3_12",
		&returns,
	)
	return returns
}

func AgentCoreRuntime_PYTHON_3_13() AgentCoreRuntime {
	_init_.Initialize()
	var returns AgentCoreRuntime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.AgentCoreRuntime",
		"PYTHON_3_13",
		&returns,
	)
	return returns
}

func AgentCoreRuntime_PYTHON_3_14() AgentCoreRuntime {
	_init_.Initialize()
	var returns AgentCoreRuntime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.AgentCoreRuntime",
		"PYTHON_3_14",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AgentCoreRuntime) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

