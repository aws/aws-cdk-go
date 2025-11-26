package previewawsamplifymixins


// Use the BasicAuthConfig property type to set password protection at an app level to all your branches.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-app-basicauthconfig.html
//
type CfnAppPropsMixin_BasicAuthConfigProperty struct {
	// Enables basic authorization for the Amplify app's branches.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-app-basicauthconfig.html#cfn-amplify-app-basicauthconfig-enablebasicauth
	//
	EnableBasicAuth interface{} `field:"optional" json:"enableBasicAuth" yaml:"enableBasicAuth"`
	// The password for basic authorization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-app-basicauthconfig.html#cfn-amplify-app-basicauthconfig-password
	//
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The user name for basic authorization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-app-basicauthconfig.html#cfn-amplify-app-basicauthconfig-username
	//
	Username *string `field:"optional" json:"username" yaml:"username"`
}

