package awswafv2


// The patterns to look for in the JSON body.
//
// AWS WAF inspects the results of these pattern matches against the rule inspection criteria. This is used with the `FieldToMatch` option `JsonBody` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//
//   jsonMatchPatternProperty := &jsonMatchPatternProperty{
//   	all: all,
//   	includedPaths: []*string{
//   		jsii.String("includedPaths"),
//   	},
//   }
//
type CfnWebACL_JsonMatchPatternProperty struct {
	// Match all of the elements. See also `MatchScope` in the `JsonBody` `FieldToMatch` specification.
	//
	// You must specify either this setting or the `IncludedPaths` setting, but not both.
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// Match only the specified include paths. See also `MatchScope` in the `JsonBody` `FieldToMatch` specification.
	//
	// Provide the include paths using JSON Pointer syntax. For example, `"IncludedPaths": ["/dogs/0/name", "/dogs/1/name"]` . For information about this syntax, see the Internet Engineering Task Force (IETF) documentation [JavaScript Object Notation (JSON) Pointer](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc6901) .
	//
	// You must specify either this setting or the `All` setting, but not both.
	//
	// > Don't use this option to include all paths. Instead, use the `All` setting.
	IncludedPaths *[]*string `field:"optional" json:"includedPaths" yaml:"includedPaths"`
}

