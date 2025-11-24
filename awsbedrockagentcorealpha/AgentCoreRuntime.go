package awsbedrockagentcorealpha


// Bedrock AgentCore runtime environment for code execution Allowed values: PYTHON_3_10 | PYTHON_3_11 | PYTHON_3_12 | PYTHON_3_13.
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
//   	BucketName: codeBucket.BucketName,
//   	ObjectKey: jsii.String("deployment_package.zip"),
//   }, agentcore.AgentCoreRuntime_PYTHON_3_12, []*string{
//   	jsii.String("opentelemetry-instrument"),
//   	jsii.String("main.py"),
//   })
//
//   runtimeInstance := agentcore.NewRuntime(this, jsii.String("MyAgentRuntime"), &RuntimeProps{
//   	RuntimeName: jsii.String("myAgent"),
//   	AgentRuntimeArtifact: agentRuntimeArtifact,
//   })
//
// Experimental.
type AgentCoreRuntime string

const (
	// Python 3.10 runtime.
	// Experimental.
	AgentCoreRuntime_PYTHON_3_10 AgentCoreRuntime = "PYTHON_3_10"
	// Python 3.11 runtime.
	// Experimental.
	AgentCoreRuntime_PYTHON_3_11 AgentCoreRuntime = "PYTHON_3_11"
	// Python 3.12 runtime.
	// Experimental.
	AgentCoreRuntime_PYTHON_3_12 AgentCoreRuntime = "PYTHON_3_12"
	// Python 3.13 runtime.
	// Experimental.
	AgentCoreRuntime_PYTHON_3_13 AgentCoreRuntime = "PYTHON_3_13"
)

