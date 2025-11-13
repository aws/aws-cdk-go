package interfacesawsneptunegraph


// A reference to a PrivateGraphEndpoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   privateGraphEndpointReference := &PrivateGraphEndpointReference{
//   	PrivateGraphEndpointIdentifier: jsii.String("privateGraphEndpointIdentifier"),
//   }
//
type PrivateGraphEndpointReference struct {
	// The PrivateGraphEndpointIdentifier of the PrivateGraphEndpoint resource.
	PrivateGraphEndpointIdentifier *string `field:"required" json:"privateGraphEndpointIdentifier" yaml:"privateGraphEndpointIdentifier"`
}

