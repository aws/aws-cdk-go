package awsec2


// Properties for defining a `CfnNetworkInterfacePermission`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNetworkInterfacePermissionProps := &cfnNetworkInterfacePermissionProps{
//   	awsAccountId: jsii.String("awsAccountId"),
//   	networkInterfaceId: jsii.String("networkInterfaceId"),
//   	permission: jsii.String("permission"),
//   }
//
type CfnNetworkInterfacePermissionProps struct {
	// The AWS account ID.
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// The ID of the network interface.
	NetworkInterfaceId *string `field:"required" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// The type of permission to grant: `INSTANCE-ATTACH` or `EIP-ASSOCIATE` .
	Permission *string `field:"required" json:"permission" yaml:"permission"`
}

