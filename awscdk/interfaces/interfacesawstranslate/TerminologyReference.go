package interfacesawstranslate


// A reference to a Terminology resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   terminologyReference := &TerminologyReference{
//   	TerminologyArn: jsii.String("terminologyArn"),
//   }
//
type TerminologyReference struct {
	// The Arn of the Terminology resource.
	TerminologyArn *string `field:"required" json:"terminologyArn" yaml:"terminologyArn"`
}

