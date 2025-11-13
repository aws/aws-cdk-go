package interfacesawsssmcontacts


// A reference to a Contact resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contactReference := &ContactReference{
//   	ContactArn: jsii.String("contactArn"),
//   }
//
type ContactReference struct {
	// The Arn of the Contact resource.
	ContactArn *string `field:"required" json:"contactArn" yaml:"contactArn"`
}

