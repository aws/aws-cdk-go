package interfacesawsdirectconnect


// A reference to a DirectConnectGateway resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   directConnectGatewayReference := &DirectConnectGatewayReference{
//   	DirectConnectGatewayArn: jsii.String("directConnectGatewayArn"),
//   }
//
type DirectConnectGatewayReference struct {
	// The DirectConnectGatewayArn of the DirectConnectGateway resource.
	DirectConnectGatewayArn *string `field:"required" json:"directConnectGatewayArn" yaml:"directConnectGatewayArn"`
}

