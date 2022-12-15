package awswafv2


// Additional information that's used by a managed rule group. Many managed rule groups don't require this.
//
// Use the `AWSManagedRulesBotControlRuleSet` configuration object to configure the protection level that you want the Bot Control rule group to use.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedRuleGroupConfigProperty := &managedRuleGroupConfigProperty{
//   	loginPath: jsii.String("loginPath"),
//   	passwordField: &fieldIdentifierProperty{
//   		identifier: jsii.String("identifier"),
//   	},
//   	payloadType: jsii.String("payloadType"),
//   	usernameField: &fieldIdentifierProperty{
//   		identifier: jsii.String("identifier"),
//   	},
//   }
//
type CfnWebACL_ManagedRuleGroupConfigProperty struct {
	// The path of the login endpoint for your application.
	//
	// For example, for the URL `https://example.com/web/login` , you would provide the path `/web/login` .
	LoginPath *string `field:"optional" json:"loginPath" yaml:"loginPath"`
	// Details about your login page password field.
	PasswordField interface{} `field:"optional" json:"passwordField" yaml:"passwordField"`
	// The payload type for your login endpoint, either JSON or form encoded.
	PayloadType *string `field:"optional" json:"payloadType" yaml:"payloadType"`
	// Details about your login page username field.
	UsernameField interface{} `field:"optional" json:"usernameField" yaml:"usernameField"`
}

