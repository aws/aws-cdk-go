package interfacesawsmpa


// A reference to a IdentitySource resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   identitySourceReference := &IdentitySourceReference{
//   	IdentitySourceArn: jsii.String("identitySourceArn"),
//   }
//
type IdentitySourceReference struct {
	// The IdentitySourceArn of the IdentitySource resource.
	IdentitySourceArn *string `field:"required" json:"identitySourceArn" yaml:"identitySourceArn"`
}

