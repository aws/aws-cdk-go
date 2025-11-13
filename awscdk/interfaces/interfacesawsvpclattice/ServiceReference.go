package interfacesawsvpclattice


// A reference to a Service resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceReference := &ServiceReference{
//   	ServiceArn: jsii.String("serviceArn"),
//   }
//
type ServiceReference struct {
	// The Arn of the Service resource.
	ServiceArn *string `field:"required" json:"serviceArn" yaml:"serviceArn"`
}

