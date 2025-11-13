package awsbedrockagentcore


// The artifact of the agent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   agentRuntimeArtifactProperty := &AgentRuntimeArtifactProperty{
//   	CodeConfiguration: &CodeConfigurationProperty{
//   		Code: &CodeProperty{
//   			S3: &S3LocationProperty{
//   				Bucket: jsii.String("bucket"),
//   				Prefix: jsii.String("prefix"),
//
//   				// the properties below are optional
//   				VersionId: jsii.String("versionId"),
//   			},
//   		},
//   		EntryPoint: []*string{
//   			jsii.String("entryPoint"),
//   		},
//   		Runtime: jsii.String("runtime"),
//   	},
//   	ContainerConfiguration: &ContainerConfigurationProperty{
//   		ContainerUri: jsii.String("containerUri"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-agentruntimeartifact.html
//
type CfnRuntime_AgentRuntimeArtifactProperty struct {
	// Representation of a code configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-agentruntimeartifact.html#cfn-bedrockagentcore-runtime-agentruntimeartifact-codeconfiguration
	//
	CodeConfiguration interface{} `field:"optional" json:"codeConfiguration" yaml:"codeConfiguration"`
	// Representation of a container configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-agentruntimeartifact.html#cfn-bedrockagentcore-runtime-agentruntimeartifact-containerconfiguration
	//
	ContainerConfiguration interface{} `field:"optional" json:"containerConfiguration" yaml:"containerConfiguration"`
}

