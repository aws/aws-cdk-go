package awsbedrockagentcore


// Configuration for permissions associated with a capacity provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   permissionsConfigurationProperty := &PermissionsConfigurationProperty{
//   	CapacityProviderOperatorRoleArn: jsii.String("capacityProviderOperatorRoleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-permissionsconfiguration.html
//
type CfnCapacityProviderPropsMixin_PermissionsConfigurationProperty struct {
	// The ARN of the IAM role that operators use to manage the capacity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-capacityprovider-permissionsconfiguration.html#cfn-bedrockagentcore-capacityprovider-permissionsconfiguration-capacityprovideroperatorrolearn
	//
	CapacityProviderOperatorRoleArn *string `field:"optional" json:"capacityProviderOperatorRoleArn" yaml:"capacityProviderOperatorRoleArn"`
}

