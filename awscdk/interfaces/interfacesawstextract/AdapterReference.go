package interfacesawstextract


// A reference to a Adapter resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   adapterReference := &AdapterReference{
//   	AdapterArn: jsii.String("adapterArn"),
//   }
//
type AdapterReference struct {
	// The Arn of the Adapter resource.
	AdapterArn *string `field:"required" json:"adapterArn" yaml:"adapterArn"`
}

