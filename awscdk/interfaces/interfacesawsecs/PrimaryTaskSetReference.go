package interfacesawsecs


// A reference to a PrimaryTaskSet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   primaryTaskSetReference := &PrimaryTaskSetReference{
//   	Cluster: jsii.String("cluster"),
//   	Service: jsii.String("service"),
//   }
//
type PrimaryTaskSetReference struct {
	// The Cluster of the PrimaryTaskSet resource.
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
	// The Service of the PrimaryTaskSet resource.
	Service *string `field:"required" json:"service" yaml:"service"`
}

