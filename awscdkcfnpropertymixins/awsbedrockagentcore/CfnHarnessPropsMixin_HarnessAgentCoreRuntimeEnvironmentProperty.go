package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   harnessAgentCoreRuntimeEnvironmentProperty := &HarnessAgentCoreRuntimeEnvironmentProperty{
//   	AgentRuntimeArn: jsii.String("agentRuntimeArn"),
//   	AgentRuntimeId: jsii.String("agentRuntimeId"),
//   	AgentRuntimeName: jsii.String("agentRuntimeName"),
//   	FilesystemConfigurations: []interface{}{
//   		&FilesystemConfigurationProperty{
//   			SessionStorage: &SessionStorageConfigurationProperty{
//   				MountPath: jsii.String("mountPath"),
//   			},
//   		},
//   	},
//   	LifecycleConfiguration: &LifecycleConfigurationProperty{
//   		IdleRuntimeSessionTimeout: jsii.Number(123),
//   		MaxLifetime: jsii.Number(123),
//   	},
//   	NetworkConfiguration: &NetworkConfigurationProperty{
//   		NetworkMode: jsii.String("networkMode"),
//   		NetworkModeConfig: &VpcConfigProperty{
//   			SecurityGroups: []*string{
//   				jsii.String("securityGroups"),
//   			},
//   			Subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessagentcoreruntimeenvironment.html
//
type CfnHarnessPropsMixin_HarnessAgentCoreRuntimeEnvironmentProperty struct {
	// The ARN of the underlying AgentCore Runtime.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessagentcoreruntimeenvironment.html#cfn-bedrockagentcore-harness-harnessagentcoreruntimeenvironment-agentruntimearn
	//
	AgentRuntimeArn *string `field:"optional" json:"agentRuntimeArn" yaml:"agentRuntimeArn"`
	// The ID of the underlying AgentCore Runtime.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessagentcoreruntimeenvironment.html#cfn-bedrockagentcore-harness-harnessagentcoreruntimeenvironment-agentruntimeid
	//
	AgentRuntimeId *string `field:"optional" json:"agentRuntimeId" yaml:"agentRuntimeId"`
	// The name of the underlying AgentCore Runtime.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessagentcoreruntimeenvironment.html#cfn-bedrockagentcore-harness-harnessagentcoreruntimeenvironment-agentruntimename
	//
	AgentRuntimeName *string `field:"optional" json:"agentRuntimeName" yaml:"agentRuntimeName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessagentcoreruntimeenvironment.html#cfn-bedrockagentcore-harness-harnessagentcoreruntimeenvironment-filesystemconfigurations
	//
	FilesystemConfigurations interface{} `field:"optional" json:"filesystemConfigurations" yaml:"filesystemConfigurations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessagentcoreruntimeenvironment.html#cfn-bedrockagentcore-harness-harnessagentcoreruntimeenvironment-lifecycleconfiguration
	//
	LifecycleConfiguration interface{} `field:"optional" json:"lifecycleConfiguration" yaml:"lifecycleConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessagentcoreruntimeenvironment.html#cfn-bedrockagentcore-harness-harnessagentcoreruntimeenvironment-networkconfiguration
	//
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
}

