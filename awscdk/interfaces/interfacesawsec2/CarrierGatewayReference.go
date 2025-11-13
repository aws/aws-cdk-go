package interfacesawsec2


// A reference to a CarrierGateway resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   carrierGatewayReference := &CarrierGatewayReference{
//   	CarrierGatewayId: jsii.String("carrierGatewayId"),
//   }
//
type CarrierGatewayReference struct {
	// The CarrierGatewayId of the CarrierGateway resource.
	CarrierGatewayId *string `field:"required" json:"carrierGatewayId" yaml:"carrierGatewayId"`
}

