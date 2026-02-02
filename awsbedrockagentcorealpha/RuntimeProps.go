package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for creating a Bedrock Agent Core Runtime resource.
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
type RuntimeProps struct {
	// The artifact configuration for the agent runtime Contains the container configuration with ECR URI.
	// Experimental.
	AgentRuntimeArtifact AgentRuntimeArtifact `field:"required" json:"agentRuntimeArtifact" yaml:"agentRuntimeArtifact"`
	// Authorizer configuration for the agent runtime Use RuntimeAuthorizerConfiguration static methods to create the configuration.
	// Default: - RuntimeAuthorizerConfiguration.iam() (IAM authentication)
	//
	// Experimental.
	AuthorizerConfiguration RuntimeAuthorizerConfiguration `field:"optional" json:"authorizerConfiguration" yaml:"authorizerConfiguration"`
	// Optional description for the agent runtime.
	// Default: - No description
	// Length Minimum: 1 , Maximum: 1200.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Environment variables for the agent runtime - Maximum 50 environment variables - Key: Must be 1-100 characters, start with letter or underscore, contain only letters, numbers, and underscores - Value: Must be 0-2048 characters (per CloudFormation specification).
	// Default: - No environment variables.
	//
	// Experimental.
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// The IAM role that provides permissions for the agent runtime If not provided, a role will be created automatically.
	// Default: - A new role will be created.
	//
	// Experimental.
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The life cycle configuration for the AgentCore Runtime.
	// Default: - No lifecycle configuration.
	//
	// Experimental.
	LifecycleConfiguration *LifecycleConfiguration `field:"optional" json:"lifecycleConfiguration" yaml:"lifecycleConfiguration"`
	// Network configuration for the agent runtime.
	// Default: - RuntimeNetworkConfiguration.usingPublicNetwork()
	//
	// Experimental.
	NetworkConfiguration RuntimeNetworkConfiguration `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// Protocol configuration for the agent runtime.
	// Default: - ProtocolType.HTTP
	//
	// Experimental.
	ProtocolConfiguration ProtocolType `field:"optional" json:"protocolConfiguration" yaml:"protocolConfiguration"`
	// Configuration for HTTP request headers that will be passed through to the runtime.
	// Default: - No request headers configured.
	//
	// Experimental.
	RequestHeaderConfiguration *RequestHeaderConfiguration `field:"optional" json:"requestHeaderConfiguration" yaml:"requestHeaderConfiguration"`
	// The name of the agent runtime Valid characters are a-z, A-Z, 0-9, _ (underscore) Must start with a letter and can be up to 48 characters long Pattern: ^[a-zA-Z][a-zA-Z0-9_]{0,47}$.
	// Default: - auto generate.
	//
	// Experimental.
	RuntimeName *string `field:"optional" json:"runtimeName" yaml:"runtimeName"`
	// Tags for the agent runtime A list of key:value pairs of tags to apply to this Runtime resource.
	// Default: {} - no tags.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

