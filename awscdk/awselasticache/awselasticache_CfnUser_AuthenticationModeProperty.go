package awselasticache


// Specifies the authentication mode to use.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authenticationModeProperty := &authenticationModeProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	passwords: []*string{
//   		jsii.String("passwords"),
//   	},
//   }
//
type CfnUser_AuthenticationModeProperty struct {
	// Specifies the authentication type.
	//
	// Possible options are IAM authentication, password and no password.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Specifies the passwords to use for authentication if `Type` is set to `password` .
	Passwords *[]*string `field:"optional" json:"passwords" yaml:"passwords"`
}

