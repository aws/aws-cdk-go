package awsamplify


// Use the BasicAuthConfig property type to set password protection for a specific branch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   basicAuthConfigProperty := &BasicAuthConfigProperty{
//   	Password: jsii.String("password"),
//   	Username: jsii.String("username"),
//
//   	// the properties below are optional
//   	EnableBasicAuth: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-branch-basicauthconfig.html
//
type CfnBranch_BasicAuthConfigProperty struct {
	// The password for basic authorization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-branch-basicauthconfig.html#cfn-amplify-branch-basicauthconfig-password
	//
	Password *string `field:"required" json:"password" yaml:"password"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-branch-basicauthconfig.html#cfn-amplify-branch-basicauthconfig-username
	//
	Username *string `field:"required" json:"username" yaml:"username"`
	// Enables basic authorization for the branch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-branch-basicauthconfig.html#cfn-amplify-branch-basicauthconfig-enablebasicauth
	//
	EnableBasicAuth interface{} `field:"optional" json:"enableBasicAuth" yaml:"enableBasicAuth"`
}

