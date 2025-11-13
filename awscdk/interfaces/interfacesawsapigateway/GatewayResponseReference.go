package interfacesawsapigateway


// A reference to a GatewayResponse resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayResponseReference := &GatewayResponseReference{
//   	GatewayResponseId: jsii.String("gatewayResponseId"),
//   }
//
type GatewayResponseReference struct {
	// The Id of the GatewayResponse resource.
	GatewayResponseId *string `field:"required" json:"gatewayResponseId" yaml:"gatewayResponseId"`
}

