package interfacesawsses


// A reference to a MultiRegionEndpoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multiRegionEndpointReference := &MultiRegionEndpointReference{
//   	EndpointName: jsii.String("endpointName"),
//   }
//
type MultiRegionEndpointReference struct {
	// The EndpointName of the MultiRegionEndpoint resource.
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
}

