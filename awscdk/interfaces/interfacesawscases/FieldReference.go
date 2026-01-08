package interfacesawscases


// A reference to a Field resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldReference := &FieldReference{
//   	FieldArn: jsii.String("fieldArn"),
//   }
//
type FieldReference struct {
	// The FieldArn of the Field resource.
	FieldArn *string `field:"required" json:"fieldArn" yaml:"fieldArn"`
}

