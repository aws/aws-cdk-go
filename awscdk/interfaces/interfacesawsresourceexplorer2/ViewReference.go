package interfacesawsresourceexplorer2


// A reference to a View resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   viewReference := &ViewReference{
//   	ViewArn: jsii.String("viewArn"),
//   }
//
type ViewReference struct {
	// The ViewArn of the View resource.
	ViewArn *string `field:"required" json:"viewArn" yaml:"viewArn"`
}

