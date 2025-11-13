package interfacesawsec2


// A reference to a InstanceConnectEndpoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceConnectEndpointReference := &InstanceConnectEndpointReference{
//   	InstanceConnectEndpointId: jsii.String("instanceConnectEndpointId"),
//   }
//
type InstanceConnectEndpointReference struct {
	// The Id of the InstanceConnectEndpoint resource.
	InstanceConnectEndpointId *string `field:"required" json:"instanceConnectEndpointId" yaml:"instanceConnectEndpointId"`
}

