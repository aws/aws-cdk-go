package interfacesawssmsvoice


// A reference to a VerifiedDestinationNumber resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   verifiedDestinationNumberReference := &VerifiedDestinationNumberReference{
//   	VerifiedDestinationNumberArn: jsii.String("verifiedDestinationNumberArn"),
//   }
//
type VerifiedDestinationNumberReference struct {
	// The VerifiedDestinationNumberArn of the VerifiedDestinationNumber resource.
	VerifiedDestinationNumberArn *string `field:"required" json:"verifiedDestinationNumberArn" yaml:"verifiedDestinationNumberArn"`
}

