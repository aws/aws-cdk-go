package awswafv2


// The identifier of a field in the web request payload that contains customer data.
//
// This data type is used to specify fields in the `RequestInspection` configurations, for the managed rule group configuration `AWSManagedRulesATPRuleSet` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldIdentifierProperty := &FieldIdentifierProperty{
//   	Identifier: jsii.String("identifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-fieldidentifier.html
//
type CfnWebACL_FieldIdentifierProperty struct {
	// The name of the field.
	//
	// When the `PayloadType` in the request inspection is `JSON` , this identifier must be in JSON pointer syntax. For example `/form/username` . For information about the JSON Pointer syntax, see the Internet Engineering Task Force (IETF) documentation [JavaScript Object Notation (JSON) Pointer](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc6901) .
	//
	// When the `PayloadType` is `FORM_ENCODED` , use the HTML form names. For example, `username` .
	//
	// For more information, see the descriptions for each field type in the request inspection properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-fieldidentifier.html#cfn-wafv2-webacl-fieldidentifier-identifier
	//
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
}

