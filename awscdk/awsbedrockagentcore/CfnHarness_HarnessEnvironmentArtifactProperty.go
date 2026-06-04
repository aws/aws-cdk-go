package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   harnessEnvironmentArtifactProperty := &HarnessEnvironmentArtifactProperty{
//   	ContainerConfiguration: &ContainerConfigurationProperty{
//   		ContainerUri: jsii.String("containerUri"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessenvironmentartifact.html
//
type CfnHarness_HarnessEnvironmentArtifactProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessenvironmentartifact.html#cfn-bedrockagentcore-harness-harnessenvironmentartifact-containerconfiguration
	//
	ContainerConfiguration interface{} `field:"optional" json:"containerConfiguration" yaml:"containerConfiguration"`
}

