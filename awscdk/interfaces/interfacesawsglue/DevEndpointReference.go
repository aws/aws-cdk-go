package interfacesawsglue


// A reference to a DevEndpoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   devEndpointReference := &DevEndpointReference{
//   	EndpointName: jsii.String("endpointName"),
//   }
//
type DevEndpointReference struct {
	// The EndpointName of the DevEndpoint resource.
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
}

