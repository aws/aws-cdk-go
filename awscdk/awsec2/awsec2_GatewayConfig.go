package awsec2


// Pair represents a gateway created by NAT Provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayConfig := &GatewayConfig{
//   	Az: jsii.String("az"),
//   	GatewayId: jsii.String("gatewayId"),
//   }
//
// Experimental.
type GatewayConfig struct {
	// Availability Zone.
	// Experimental.
	Az *string `field:"required" json:"az" yaml:"az"`
	// Identity of gateway spawned by the provider.
	// Experimental.
	GatewayId *string `field:"required" json:"gatewayId" yaml:"gatewayId"`
}

