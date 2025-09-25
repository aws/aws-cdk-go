package awsec2


// A reference to a NetworkInterface resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkInterfaceReference := &NetworkInterfaceReference{
//   	NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   }
//
type NetworkInterfaceReference struct {
	// The Id of the NetworkInterface resource.
	NetworkInterfaceId *string `field:"required" json:"networkInterfaceId" yaml:"networkInterfaceId"`
}

