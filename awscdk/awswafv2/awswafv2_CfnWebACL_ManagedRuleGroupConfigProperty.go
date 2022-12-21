package awswafv2


// Additional information that's used by a managed rule group. Most managed rule groups don't require this.
//
// Use this for the account takeover prevention managed rule group `AWSManagedRulesATPRuleSet` , to provide information about the sign-in page of your application.
//
// You can provide multiple individual `ManagedRuleGroupConfig` objects for any rule group configuration, for example `UsernameField` and `PasswordField` . The configuration that you provide depends on the needs of the managed rule group. For the ATP managed rule group, you provide the following individual configuration objects: `LoginPath` , `PasswordField` , `PayloadType` and `UsernameField` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedRuleGroupConfigProperty := &managedRuleGroupConfigProperty{
//   	awsManagedRulesBotControlRuleSet: &aWSManagedRulesBotControlRuleSetProperty{
//   		inspectionLevel: jsii.String("inspectionLevel"),
//   	},
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
	// `CfnWebACL.ManagedRuleGroupConfigProperty.AWSManagedRulesBotControlRuleSet`.
	AwsManagedRulesBotControlRuleSet interface{} `field:"optional" json:"awsManagedRulesBotControlRuleSet" yaml:"awsManagedRulesBotControlRuleSet"`
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

