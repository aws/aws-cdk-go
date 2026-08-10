package interfacesawsomics


// A reference to a Reference resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   referenceReference := &ReferenceReference{
//   	ReferenceArn: jsii.String("referenceArn"),
//   }
//
type ReferenceReference struct {
	// The Arn of the Reference resource.
	ReferenceArn *string `field:"required" json:"referenceArn" yaml:"referenceArn"`
}

