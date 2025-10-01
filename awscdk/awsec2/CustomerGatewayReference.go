package awsec2


// A reference to a CustomerGateway resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customerGatewayReference := &CustomerGatewayReference{
//   	CustomerGatewayId: jsii.String("customerGatewayId"),
//   }
//
type CustomerGatewayReference struct {
	// The CustomerGatewayId of the CustomerGateway resource.
	CustomerGatewayId *string `field:"required" json:"customerGatewayId" yaml:"customerGatewayId"`
}

