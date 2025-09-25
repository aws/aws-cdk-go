package awscognito


// A reference to a UserPoolDomain resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userPoolDomainReference := &UserPoolDomainReference{
//   	Domain: jsii.String("domain"),
//   	UserPoolId: jsii.String("userPoolId"),
//   }
//
type UserPoolDomainReference struct {
	// The Domain of the UserPoolDomain resource.
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// The UserPoolId of the UserPoolDomain resource.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
}

