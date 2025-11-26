package previewawsamplifymixins


// Use the BasicAuthConfig property type to set password protection for a specific branch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   basicAuthConfigProperty := &BasicAuthConfigProperty{
//   	EnableBasicAuth: jsii.Boolean(false),
//   	Password: jsii.String("password"),
//   	Username: jsii.String("username"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-branch-basicauthconfig.html
//
type CfnBranchPropsMixin_BasicAuthConfigProperty struct {
	// Enables basic authorization for the branch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-branch-basicauthconfig.html#cfn-amplify-branch-basicauthconfig-enablebasicauth
	//
	EnableBasicAuth interface{} `field:"optional" json:"enableBasicAuth" yaml:"enableBasicAuth"`
	// The password for basic authorization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-branch-basicauthconfig.html#cfn-amplify-branch-basicauthconfig-password
	//
	Password *string `field:"optional" json:"password" yaml:"password"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-branch-basicauthconfig.html#cfn-amplify-branch-basicauthconfig-username
	//
	Username *string `field:"optional" json:"username" yaml:"username"`
}

