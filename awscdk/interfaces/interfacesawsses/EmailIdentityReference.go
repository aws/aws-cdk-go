package interfacesawsses


// A reference to a EmailIdentity resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   emailIdentityReference := &EmailIdentityReference{
//   	EmailIdentity: jsii.String("emailIdentity"),
//   }
//
type EmailIdentityReference struct {
	// The EmailIdentity of the EmailIdentity resource.
	EmailIdentity *string `field:"required" json:"emailIdentity" yaml:"emailIdentity"`
}

