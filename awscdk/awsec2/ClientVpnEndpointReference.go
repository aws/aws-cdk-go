package awsec2


// A reference to a ClientVpnEndpoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clientVpnEndpointReference := &ClientVpnEndpointReference{
//   	ClientVpnEndpointId: jsii.String("clientVpnEndpointId"),
//   }
//
type ClientVpnEndpointReference struct {
	// The Id of the ClientVpnEndpoint resource.
	ClientVpnEndpointId *string `field:"required" json:"clientVpnEndpointId" yaml:"clientVpnEndpointId"`
}

