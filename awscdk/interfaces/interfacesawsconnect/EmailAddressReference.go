package interfacesawsconnect


// A reference to a EmailAddress resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   emailAddressReference := &EmailAddressReference{
//   	EmailAddressArn: jsii.String("emailAddressArn"),
//   }
//
type EmailAddressReference struct {
	// The EmailAddressArn of the EmailAddress resource.
	EmailAddressArn *string `field:"required" json:"emailAddressArn" yaml:"emailAddressArn"`
}

