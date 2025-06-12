package awsec2


// Properties for defining a `CfnNetworkInterfacePermission`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNetworkInterfacePermissionProps := &CfnNetworkInterfacePermissionProps{
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   	Permission: jsii.String("permission"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterfacepermission.html
//
type CfnNetworkInterfacePermissionProps struct {
	// The AWS account ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterfacepermission.html#cfn-ec2-networkinterfacepermission-awsaccountid
	//
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// The ID of the network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterfacepermission.html#cfn-ec2-networkinterfacepermission-networkinterfaceid
	//
	NetworkInterfaceId *string `field:"required" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// The type of permission to grant: `INSTANCE-ATTACH` or `EIP-ASSOCIATE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterfacepermission.html#cfn-ec2-networkinterfacepermission-permission
	//
	Permission *string `field:"required" json:"permission" yaml:"permission"`
}

