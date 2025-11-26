package previewawscognitomixins


// Case sensitivity of the username input for the selected sign-in option.
//
// When case sensitivity is set to `False` (case insensitive), users can sign in with any combination of capital and lowercase letters. For example, `username` , `USERNAME` , or `UserName` , or for email, `email@example.com` or `EMaiL@eXamplE.Com` . For most use cases, set case sensitivity to `False` (case insensitive) as a best practice. When usernames and email addresses are case insensitive, Amazon Cognito treats any variation in case as the same user, and prevents a case variation from being assigned to the same attribute for a different user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   usernameConfigurationProperty := &UsernameConfigurationProperty{
//   	CaseSensitive: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-usernameconfiguration.html
//
type CfnUserPoolPropsMixin_UsernameConfigurationProperty struct {
	// Specifies whether user name case sensitivity will be applied for all users in the user pool through Amazon Cognito APIs.
	//
	// For most use cases, set case sensitivity to `False` (case insensitive) as a best practice. When usernames and email addresses are case insensitive, users can sign in as the same user when they enter a different capitalization of their user name.
	//
	// Valid values include:
	//
	// - **true** - Enables case sensitivity for all username input. When this option is set to `true` , users must sign in using the exact capitalization of their given username, such as “UserName”. This is the default value.
	// - **false** - Enables case insensitivity for all username input. For example, when this option is set to `false` , users can sign in using `username` , `USERNAME` , or `UserName` . This option also enables both `preferred_username` and `email` alias to be case insensitive, in addition to the `username` attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-usernameconfiguration.html#cfn-cognito-userpool-usernameconfiguration-casesensitive
	//
	CaseSensitive interface{} `field:"optional" json:"caseSensitive" yaml:"caseSensitive"`
}

