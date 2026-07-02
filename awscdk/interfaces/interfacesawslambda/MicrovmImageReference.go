package interfacesawslambda


// A reference to a MicrovmImage resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   microvmImageReference := &MicrovmImageReference{
//   	ImageArn: jsii.String("imageArn"),
//   }
//
type MicrovmImageReference struct {
	// The ImageArn of the MicrovmImage resource.
	ImageArn *string `field:"required" json:"imageArn" yaml:"imageArn"`
}

