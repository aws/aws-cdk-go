package interfacesawsec2


// A reference to a ClientVpnRoute resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clientVpnRouteReference := &ClientVpnRouteReference{
//   	ClientVpnRouteId: jsii.String("clientVpnRouteId"),
//   }
//
type ClientVpnRouteReference struct {
	// The Id of the ClientVpnRoute resource.
	ClientVpnRouteId *string `field:"required" json:"clientVpnRouteId" yaml:"clientVpnRouteId"`
}

