package interfacesawscases


// A reference to a Layout resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   layoutReference := &LayoutReference{
//   	LayoutArn: jsii.String("layoutArn"),
//   }
//
type LayoutReference struct {
	// The LayoutArn of the Layout resource.
	LayoutArn *string `field:"required" json:"layoutArn" yaml:"layoutArn"`
}

