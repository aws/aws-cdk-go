package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for creating a Bedrock Agent Core Runtime resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var agentRuntimeArtifact AgentRuntimeArtifact
//   var loggingDestination LoggingDestination
//   var logType LogType
//   var role Role
//   var runtimeAuthorizerConfiguration RuntimeAuthorizerConfiguration
//   var runtimeNetworkConfiguration RuntimeNetworkConfiguration
//
//   runtimeProps := &RuntimeProps{
//   	AgentRuntimeArtifact: agentRuntimeArtifact,
//
//   	// the properties below are optional
//   	AuthorizerConfiguration: runtimeAuthorizerConfiguration,
//   	Description: jsii.String("description"),
//   	EnvironmentVariables: map[string]*string{
//   		"environmentVariablesKey": jsii.String("environmentVariables"),
//   	},
//   	ExecutionRole: role,
//   	LifecycleConfiguration: &LifecycleConfiguration{
//   		IdleRuntimeSessionTimeout: cdk.Duration_Minutes(jsii.Number(30)),
//   		MaxLifetime: cdk.Duration_*Minutes(jsii.Number(30)),
//   	},
//   	LoggingConfigs: []LoggingConfig{
//   		&LoggingConfig{
//   			Destination: loggingDestination,
//   			LogType: logType,
//   		},
//   	},
//   	NetworkConfiguration: runtimeNetworkConfiguration,
//   	ProtocolConfiguration: bedrock_agentcore_alpha.ProtocolType_MCP,
//   	RequestHeaderConfiguration: &RequestHeaderConfiguration{
//   		AllowlistedHeaders: []*string{
//   			jsii.String("allowlistedHeaders"),
//   		},
//   	},
//   	RuntimeName: jsii.String("runtimeName"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	TracingEnabled: jsii.Boolean(false),
//   }
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
	// Logging configuration for the runtime.
	//
	// Allows sending APPLICATION_LOGS and USAGE_LOGS to CloudWatch Logs, S3, or Kinesis Data Firehose.
	// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/observability.html
	//
	// Default: - No logging configured.
	//
	// Experimental.
	LoggingConfigs *[]*LoggingConfig `field:"optional" json:"loggingConfigs" yaml:"loggingConfigs"`
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
	// Whether to enable X-Ray tracing for this runtime.
	//
	// When enabled, traces will be delivered to AWS X-Ray.
	// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/observability.html
	//
	// Default: false.
	//
	// Experimental.
	TracingEnabled *bool `field:"optional" json:"tracingEnabled" yaml:"tracingEnabled"`
}

