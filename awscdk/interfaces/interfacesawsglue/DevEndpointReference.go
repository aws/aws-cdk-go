package interfacesawsglue


// A reference to a DevEndpoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   devEndpointReference := &DevEndpointReference{
//   	DevEndpointId: jsii.String("devEndpointId"),
//   }
//
type DevEndpointReference struct {
	// The Id of the DevEndpoint resource.
	DevEndpointId *string `field:"required" json:"devEndpointId" yaml:"devEndpointId"`
}

