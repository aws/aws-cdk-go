package awsec2


// A reference to a LocalGatewayVirtualInterfaceGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   localGatewayVirtualInterfaceGroupReference := &LocalGatewayVirtualInterfaceGroupReference{
//   	LocalGatewayVirtualInterfaceGroupArn: jsii.String("localGatewayVirtualInterfaceGroupArn"),
//   	LocalGatewayVirtualInterfaceGroupId: jsii.String("localGatewayVirtualInterfaceGroupId"),
//   }
//
type LocalGatewayVirtualInterfaceGroupReference struct {
	// The ARN of the LocalGatewayVirtualInterfaceGroup resource.
	LocalGatewayVirtualInterfaceGroupArn *string `field:"required" json:"localGatewayVirtualInterfaceGroupArn" yaml:"localGatewayVirtualInterfaceGroupArn"`
	// The LocalGatewayVirtualInterfaceGroupId of the LocalGatewayVirtualInterfaceGroup resource.
	LocalGatewayVirtualInterfaceGroupId *string `field:"required" json:"localGatewayVirtualInterfaceGroupId" yaml:"localGatewayVirtualInterfaceGroupId"`
}

