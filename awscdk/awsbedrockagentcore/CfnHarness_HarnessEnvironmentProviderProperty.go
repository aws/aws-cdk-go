package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   harnessEnvironmentProviderProperty := &HarnessEnvironmentProviderProperty{
//   	AgentCoreRuntimeEnvironment: &HarnessAgentCoreRuntimeEnvironmentProperty{
//   		AgentRuntimeArn: jsii.String("agentRuntimeArn"),
//   		AgentRuntimeId: jsii.String("agentRuntimeId"),
//   		AgentRuntimeName: jsii.String("agentRuntimeName"),
//   		FilesystemConfigurations: []interface{}{
//   			&FilesystemConfigurationProperty{
//   				SessionStorage: &SessionStorageConfigurationProperty{
//   					MountPath: jsii.String("mountPath"),
//   				},
//   			},
//   		},
//   		LifecycleConfiguration: &LifecycleConfigurationProperty{
//   			IdleRuntimeSessionTimeout: jsii.Number(123),
//   			MaxLifetime: jsii.Number(123),
//   		},
//   		NetworkConfiguration: &NetworkConfigurationProperty{
//   			NetworkMode: jsii.String("networkMode"),
//
//   			// the properties below are optional
//   			NetworkModeConfig: &VpcConfigProperty{
//   				SecurityGroups: []*string{
//   					jsii.String("securityGroups"),
//   				},
//   				Subnets: []*string{
//   					jsii.String("subnets"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessenvironmentprovider.html
//
type CfnHarness_HarnessEnvironmentProviderProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessenvironmentprovider.html#cfn-bedrockagentcore-harness-harnessenvironmentprovider-agentcoreruntimeenvironment
	//
	AgentCoreRuntimeEnvironment interface{} `field:"optional" json:"agentCoreRuntimeEnvironment" yaml:"agentCoreRuntimeEnvironment"`
}

