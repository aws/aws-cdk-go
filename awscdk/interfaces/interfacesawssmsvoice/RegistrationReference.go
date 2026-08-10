package interfacesawssmsvoice


// A reference to a Registration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   registrationReference := &RegistrationReference{
//   	RegistrationArn: jsii.String("registrationArn"),
//   }
//
type RegistrationReference struct {
	// The RegistrationArn of the Registration resource.
	RegistrationArn *string `field:"required" json:"registrationArn" yaml:"registrationArn"`
}

