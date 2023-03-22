package awswafv2


// The identifier of the username or password field, used in the `ManagedRuleGroupConfig` settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldIdentifierProperty := &fieldIdentifierProperty{
//   	identifier: jsii.String("identifier"),
//   }
//
type CfnWebACL_FieldIdentifierProperty struct {
	// The name of the username or password field, used in the `ManagedRuleGroupConfig` settings.
	//
	// When the `PayloadType` is `JSON` , the identifier must be in JSON pointer syntax. For example `/form/username` . For information about the JSON Pointer syntax, see the Internet Engineering Task Force (IETF) documentation [JavaScript Object Notation (JSON) Pointer](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc6901) .
	//
	// When the `PayloadType` is `FORM_ENCODED` , use the HTML form names. For example, `username` .
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
}

