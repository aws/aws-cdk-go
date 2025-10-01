package awscognito


// A reference to a UserPoolResourceServer resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userPoolResourceServerReference := &UserPoolResourceServerReference{
//   	Identifier: jsii.String("identifier"),
//   	UserPoolId: jsii.String("userPoolId"),
//   }
//
type UserPoolResourceServerReference struct {
	// The Identifier of the UserPoolResourceServer resource.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// The UserPoolId of the UserPoolResourceServer resource.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
}

