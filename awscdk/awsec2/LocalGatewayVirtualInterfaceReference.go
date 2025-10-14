package awsec2


// A reference to a LocalGatewayVirtualInterface resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   localGatewayVirtualInterfaceReference := &LocalGatewayVirtualInterfaceReference{
//   	LocalGatewayVirtualInterfaceId: jsii.String("localGatewayVirtualInterfaceId"),
//   }
//
type LocalGatewayVirtualInterfaceReference struct {
	// The LocalGatewayVirtualInterfaceId of the LocalGatewayVirtualInterface resource.
	LocalGatewayVirtualInterfaceId *string `field:"required" json:"localGatewayVirtualInterfaceId" yaml:"localGatewayVirtualInterfaceId"`
}

