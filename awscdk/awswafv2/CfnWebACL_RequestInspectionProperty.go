package awswafv2


// The criteria for inspecting login requests, used by the ATP rule group to validate credentials usage.
//
// This is part of the `AWSManagedRulesATPRuleSet` configuration in `ManagedRuleGroupConfig` .
//
// In these settings, you specify how your application accepts login attempts by providing the request payload type and the names of the fields within the request body where the username and password are provided.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   requestInspectionProperty := &RequestInspectionProperty{
//   	PasswordField: &FieldIdentifierProperty{
//   		Identifier: jsii.String("identifier"),
//   	},
//   	PayloadType: jsii.String("payloadType"),
//   	UsernameField: &FieldIdentifierProperty{
//   		Identifier: jsii.String("identifier"),
//   	},
//   }
//
type CfnWebACL_RequestInspectionProperty struct {
	// The name of the field in the request payload that contains your customer's password.
	//
	// How you specify this depends on the request inspection payload type.
	//
	// - For JSON payloads, specify the field name in JSON pointer syntax. For information about the JSON Pointer syntax, see the Internet Engineering Task Force (IETF) documentation [JavaScript Object Notation (JSON) Pointer](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc6901) .
	//
	// For example, for the JSON payload `{ "form": { "password": "THE_PASSWORD" } }` , the password field specification is `/form/password` .
	// - For form encoded payload types, use the HTML form names.
	//
	// For example, for an HTML form with the input element named `password1` , the password field specification is `password1` .
	PasswordField interface{} `field:"required" json:"passwordField" yaml:"passwordField"`
	// The payload type for your login endpoint, either JSON or form encoded.
	PayloadType *string `field:"required" json:"payloadType" yaml:"payloadType"`
	// The name of the field in the request payload that contains your customer's username.
	//
	// How you specify this depends on the request inspection payload type.
	//
	// - For JSON payloads, specify the field name in JSON pointer syntax. For information about the JSON Pointer syntax, see the Internet Engineering Task Force (IETF) documentation [JavaScript Object Notation (JSON) Pointer](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc6901) .
	//
	// For example, for the JSON payload `{ "form": { "username": "THE_USERNAME" } }` , the username field specification is `/form/username` .
	// - For form encoded payload types, use the HTML form names.
	//
	// For example, for an HTML form with the input element named `username1` , the username field specification is `username1`.
	UsernameField interface{} `field:"required" json:"usernameField" yaml:"usernameField"`
}

