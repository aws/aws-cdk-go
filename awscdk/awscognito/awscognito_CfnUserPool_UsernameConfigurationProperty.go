package awscognito


// The `UsernameConfiguration` property type specifies case sensitivity on the username input for the selected sign-in option.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   usernameConfigurationProperty := &usernameConfigurationProperty{
//   	caseSensitive: jsii.Boolean(false),
//   }
//
type CfnUserPool_UsernameConfigurationProperty struct {
	// Specifies whether user name case sensitivity will be applied for all users in the user pool through Amazon Cognito APIs.
	//
	// Valid values include:
	//
	// - **True** - Enables case sensitivity for all username input. When this option is set to `True` , users must sign in using the exact capitalization of their given username, such as “UserName”. This is the default value.
	// - **False** - Enables case insensitivity for all username input. For example, when this option is set to `False` , users can sign in using either "username" or "Username". This option also enables both `preferred_username` and `email` alias to be case insensitive, in addition to the `username` attribute.
	CaseSensitive interface{} `field:"optional" json:"caseSensitive" yaml:"caseSensitive"`
}

