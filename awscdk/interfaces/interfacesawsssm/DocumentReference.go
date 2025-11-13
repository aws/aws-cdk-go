package interfacesawsssm


// A reference to a Document resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentReference := &DocumentReference{
//   	DocumentName: jsii.String("documentName"),
//   }
//
type DocumentReference struct {
	// The Name of the Document resource.
	DocumentName *string `field:"required" json:"documentName" yaml:"documentName"`
}

