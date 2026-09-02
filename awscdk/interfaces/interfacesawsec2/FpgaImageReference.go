package interfacesawsec2


// A reference to a FpgaImage resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fpgaImageReference := &FpgaImageReference{
//   	FpgaImageArn: jsii.String("fpgaImageArn"),
//   }
//
type FpgaImageReference struct {
	// The Arn of the FpgaImage resource.
	FpgaImageArn *string `field:"required" json:"fpgaImageArn" yaml:"fpgaImageArn"`
}

