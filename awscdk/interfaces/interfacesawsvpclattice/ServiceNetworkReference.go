package interfacesawsvpclattice


// A reference to a ServiceNetwork resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceNetworkReference := &ServiceNetworkReference{
//   	ServiceNetworkArn: jsii.String("serviceNetworkArn"),
//   }
//
type ServiceNetworkReference struct {
	// The Arn of the ServiceNetwork resource.
	ServiceNetworkArn *string `field:"required" json:"serviceNetworkArn" yaml:"serviceNetworkArn"`
}

