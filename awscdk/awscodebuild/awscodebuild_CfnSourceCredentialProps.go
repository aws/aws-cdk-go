package awscodebuild


// Properties for defining a `CfnSourceCredential`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSourceCredentialProps := &cfnSourceCredentialProps{
//   	authType: jsii.String("authType"),
//   	serverType: jsii.String("serverType"),
//   	token: jsii.String("token"),
//
//   	// the properties below are optional
//   	username: jsii.String("username"),
//   }
//
type CfnSourceCredentialProps struct {
	// The type of authentication used by the credentials.
	//
	// Valid options are OAUTH, BASIC_AUTH, or PERSONAL_ACCESS_TOKEN.
	AuthType *string `field:"required" json:"authType" yaml:"authType"`
	// The type of source provider.
	//
	// The valid options are GITHUB, GITHUB_ENTERPRISE, or BITBUCKET.
	ServerType *string `field:"required" json:"serverType" yaml:"serverType"`
	// For GitHub or GitHub Enterprise, this is the personal access token.
	//
	// For Bitbucket, this is the app password.
	Token *string `field:"required" json:"token" yaml:"token"`
	// The Bitbucket username when the `authType` is BASIC_AUTH.
	//
	// This parameter is not valid for other types of source providers or connections.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

