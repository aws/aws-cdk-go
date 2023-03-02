package awsappmesh


// An object that represents a port mapping.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualGatewayPortMappingProperty := &virtualGatewayPortMappingProperty{
//   	port: jsii.Number(123),
//   	protocol: jsii.String("protocol"),
//   }
//
type CfnVirtualGateway_VirtualGatewayPortMappingProperty struct {
	// The port used for the port mapping.
	//
	// Specify one protocol.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The protocol used for the port mapping.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
}

