package awsamplify


// Use the BasicAuthConfig property type to set password protection at an app level to all your branches.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   basicAuthConfigProperty := &basicAuthConfigProperty{
//   	enableBasicAuth: jsii.Boolean(false),
//   	password: jsii.String("password"),
//   	username: jsii.String("username"),
//   }
//
type CfnApp_BasicAuthConfigProperty struct {
	// Enables basic authorization for the Amplify app's branches.
	EnableBasicAuth interface{} `field:"optional" json:"enableBasicAuth" yaml:"enableBasicAuth"`
	// The password for basic authorization.
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 255.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The user name for basic authorization.
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 255.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

