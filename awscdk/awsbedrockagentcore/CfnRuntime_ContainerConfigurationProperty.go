package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerConfigurationProperty := &ContainerConfigurationProperty{
//   	ContainerUri: jsii.String("containerUri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-containerconfiguration.html
//
type CfnRuntime_ContainerConfigurationProperty struct {
	// The ECR URI of the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-containerconfiguration.html#cfn-bedrockagentcore-runtime-containerconfiguration-containeruri
	//
	ContainerUri *string `field:"required" json:"containerUri" yaml:"containerUri"`
}

