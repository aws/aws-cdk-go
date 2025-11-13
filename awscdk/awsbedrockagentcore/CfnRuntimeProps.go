package awsbedrockagentcore


// Properties for defining a `CfnRuntime`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRuntimeProps := &CfnRuntimeProps{
//   	AgentRuntimeArtifact: &AgentRuntimeArtifactProperty{
//   		CodeConfiguration: &CodeConfigurationProperty{
//   			Code: &CodeProperty{
//   				S3: &S3LocationProperty{
//   					Bucket: jsii.String("bucket"),
//   					Prefix: jsii.String("prefix"),
//
//   					// the properties below are optional
//   					VersionId: jsii.String("versionId"),
//   				},
//   			},
//   			EntryPoint: []*string{
//   				jsii.String("entryPoint"),
//   			},
//   			Runtime: jsii.String("runtime"),
//   		},
//   		ContainerConfiguration: &ContainerConfigurationProperty{
//   			ContainerUri: jsii.String("containerUri"),
//   		},
//   	},
//   	AgentRuntimeName: jsii.String("agentRuntimeName"),
//   	NetworkConfiguration: &NetworkConfigurationProperty{
//   		NetworkMode: jsii.String("networkMode"),
//
//   		// the properties below are optional
//   		NetworkModeConfig: &VpcConfigProperty{
//   			SecurityGroups: []*string{
//   				jsii.String("securityGroups"),
//   			},
//   			Subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	AuthorizerConfiguration: &AuthorizerConfigurationProperty{
//   		CustomJwtAuthorizer: &CustomJWTAuthorizerConfigurationProperty{
//   			DiscoveryUrl: jsii.String("discoveryUrl"),
//
//   			// the properties below are optional
//   			AllowedAudience: []*string{
//   				jsii.String("allowedAudience"),
//   			},
//   			AllowedClients: []*string{
//   				jsii.String("allowedClients"),
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	EnvironmentVariables: map[string]*string{
//   		"environmentVariablesKey": jsii.String("environmentVariables"),
//   	},
//   	LifecycleConfiguration: &LifecycleConfigurationProperty{
//   		IdleRuntimeSessionTimeout: jsii.Number(123),
//   		MaxLifetime: jsii.Number(123),
//   	},
//   	ProtocolConfiguration: jsii.String("protocolConfiguration"),
//   	RequestHeaderConfiguration: &RequestHeaderConfigurationProperty{
//   		RequestHeaderAllowlist: []*string{
//   			jsii.String("requestHeaderAllowlist"),
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-runtime.html
//
type CfnRuntimeProps struct {
	// The artifact of the agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-runtime.html#cfn-bedrockagentcore-runtime-agentruntimeartifact
	//
	AgentRuntimeArtifact interface{} `field:"required" json:"agentRuntimeArtifact" yaml:"agentRuntimeArtifact"`
	// The name of the AgentCore Runtime endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-runtime.html#cfn-bedrockagentcore-runtime-agentruntimename
	//
	AgentRuntimeName *string `field:"required" json:"agentRuntimeName" yaml:"agentRuntimeName"`
	// The network configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-runtime.html#cfn-bedrockagentcore-runtime-networkconfiguration
	//
	NetworkConfiguration interface{} `field:"required" json:"networkConfiguration" yaml:"networkConfiguration"`
	// The Amazon Resource Name (ARN) for for the role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-runtime.html#cfn-bedrockagentcore-runtime-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Represents inbound authorization configuration options used to authenticate incoming requests.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-runtime.html#cfn-bedrockagentcore-runtime-authorizerconfiguration
	//
	AuthorizerConfiguration interface{} `field:"optional" json:"authorizerConfiguration" yaml:"authorizerConfiguration"`
	// The agent runtime description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-runtime.html#cfn-bedrockagentcore-runtime-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The environment variables for the agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-runtime.html#cfn-bedrockagentcore-runtime-environmentvariables
	//
	EnvironmentVariables interface{} `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Configuration for managing the lifecycle of runtime sessions and resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-runtime.html#cfn-bedrockagentcore-runtime-lifecycleconfiguration
	//
	LifecycleConfiguration interface{} `field:"optional" json:"lifecycleConfiguration" yaml:"lifecycleConfiguration"`
	// The protocol configuration for an agent runtime.
	//
	// This structure defines how the agent runtime communicates with clients.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-runtime.html#cfn-bedrockagentcore-runtime-protocolconfiguration
	//
	ProtocolConfiguration *string `field:"optional" json:"protocolConfiguration" yaml:"protocolConfiguration"`
	// Configuration for HTTP request headers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-runtime.html#cfn-bedrockagentcore-runtime-requestheaderconfiguration
	//
	RequestHeaderConfiguration interface{} `field:"optional" json:"requestHeaderConfiguration" yaml:"requestHeaderConfiguration"`
	// The tags for the agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-runtime.html#cfn-bedrockagentcore-runtime-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

