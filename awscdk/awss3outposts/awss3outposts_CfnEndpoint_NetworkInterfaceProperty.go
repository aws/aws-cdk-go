package awss3outposts


// The container for the network interface.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkInterfaceProperty := &networkInterfaceProperty{
//   	networkInterfaceId: jsii.String("networkInterfaceId"),
//   }
//
type CfnEndpoint_NetworkInterfaceProperty struct {
	// The ID for the network interface.
	NetworkInterfaceId *string `field:"required" json:"networkInterfaceId" yaml:"networkInterfaceId"`
}

