package interfacesawscognito


// A reference to a Terms resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   termsReference := &TermsReference{
//   	TermsId: jsii.String("termsId"),
//   	UserPoolId: jsii.String("userPoolId"),
//   }
//
type TermsReference struct {
	// The TermsId of the Terms resource.
	TermsId *string `field:"required" json:"termsId" yaml:"termsId"`
	// The UserPoolId of the Terms resource.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
}

