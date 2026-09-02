package awsbedrockagentcore


// Configuration for a capacity provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityProviderConfigurationProperty := &CapacityProviderConfigurationProperty{
//   	CapacityProviderArn: jsii.String("capacityProviderArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-capacityproviderconfiguration.html
//
type CfnRuntime_CapacityProviderConfigurationProperty struct {
	// ARN of the capacity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-capacityproviderconfiguration.html#cfn-bedrockagentcore-runtime-capacityproviderconfiguration-capacityproviderarn
	//
	CapacityProviderArn *string `field:"required" json:"capacityProviderArn" yaml:"capacityProviderArn"`
}

