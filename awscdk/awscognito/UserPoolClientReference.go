package awscognito


// A reference to a UserPoolClient resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userPoolClientReference := &UserPoolClientReference{
//   	ClientId: jsii.String("clientId"),
//   	UserPoolId: jsii.String("userPoolId"),
//   }
//
type UserPoolClientReference struct {
	// The ClientId of the UserPoolClient resource.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The UserPoolId of the UserPoolClient resource.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
}

