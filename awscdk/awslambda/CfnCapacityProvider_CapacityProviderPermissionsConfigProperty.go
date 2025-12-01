package awslambda


// IAM permissions configuration for the capacity provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityProviderPermissionsConfigProperty := &CapacityProviderPermissionsConfigProperty{
//   	CapacityProviderOperatorRoleArn: jsii.String("capacityProviderOperatorRoleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-capacityprovider-capacityproviderpermissionsconfig.html
//
type CfnCapacityProvider_CapacityProviderPermissionsConfigProperty struct {
	// The ARN of the IAM role that Lambda assumes to manage the capacity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-capacityprovider-capacityproviderpermissionsconfig.html#cfn-lambda-capacityprovider-capacityproviderpermissionsconfig-capacityprovideroperatorrolearn
	//
	CapacityProviderOperatorRoleArn *string `field:"required" json:"capacityProviderOperatorRoleArn" yaml:"capacityProviderOperatorRoleArn"`
}

