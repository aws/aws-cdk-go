package interfacesawslightsail


// A reference to a ContactMethod resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contactMethodReference := &ContactMethodReference{
//   	ContactMethodArn: jsii.String("contactMethodArn"),
//   }
//
type ContactMethodReference struct {
	// The Arn of the ContactMethod resource.
	ContactMethodArn *string `field:"required" json:"contactMethodArn" yaml:"contactMethodArn"`
}

