package awsiotsitewise


// A reference to a Gateway resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayReference := &GatewayReference{
//   	GatewayId: jsii.String("gatewayId"),
//   }
//
type GatewayReference struct {
	// The GatewayId of the Gateway resource.
	GatewayId *string `field:"required" json:"gatewayId" yaml:"gatewayId"`
}

