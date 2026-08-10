package interfacesawswellarchitected


// A reference to a Lens resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lensReference := &LensReference{
//   	LensArn: jsii.String("lensArn"),
//   }
//
type LensReference struct {
	// The LensArn of the Lens resource.
	LensArn *string `field:"required" json:"lensArn" yaml:"lensArn"`
}

