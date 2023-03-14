package awsamplify


// Use the BasicAuthConfig property type to set password protection for a specific branch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   basicAuthConfigProperty := &basicAuthConfigProperty{
//   	password: jsii.String("password"),
//   	username: jsii.String("username"),
//
//   	// the properties below are optional
//   	enableBasicAuth: jsii.Boolean(false),
//   }
//
type CfnBranch_BasicAuthConfigProperty struct {
	// The password for basic authorization.
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 255.
	Password *string `field:"required" json:"password" yaml:"password"`
	// The user name for basic authorization.
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 255.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Enables basic authorization for the branch.
	EnableBasicAuth interface{} `field:"optional" json:"enableBasicAuth" yaml:"enableBasicAuth"`
}

