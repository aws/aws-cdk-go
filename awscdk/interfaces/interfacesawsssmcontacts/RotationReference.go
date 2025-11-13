package interfacesawsssmcontacts


// A reference to a Rotation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rotationReference := &RotationReference{
//   	RotationArn: jsii.String("rotationArn"),
//   }
//
type RotationReference struct {
	// The Arn of the Rotation resource.
	RotationArn *string `field:"required" json:"rotationArn" yaml:"rotationArn"`
}

