package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   containerConfigurationProperty := &ContainerConfigurationProperty{
//   	ContainerUri: jsii.String("containerUri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-containerconfiguration.html
//
type CfnHarnessPropsMixin_ContainerConfigurationProperty struct {
	// The ECR URI of the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-containerconfiguration.html#cfn-bedrockagentcore-harness-containerconfiguration-containeruri
	//
	ContainerUri *string `field:"optional" json:"containerUri" yaml:"containerUri"`
}

