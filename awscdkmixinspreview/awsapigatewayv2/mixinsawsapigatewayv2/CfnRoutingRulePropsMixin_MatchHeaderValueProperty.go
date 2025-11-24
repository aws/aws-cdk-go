package mixinsawsapigatewayv2


// Represents a `MatchHeaderValue` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   matchHeaderValueProperty := &MatchHeaderValueProperty{
//   	Header: jsii.String("header"),
//   	ValueGlob: jsii.String("valueGlob"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-routingrule-matchheadervalue.html
//
type CfnRoutingRulePropsMixin_MatchHeaderValueProperty struct {
	// The case insensitive header name to be matched.
	//
	// The header name must be less than 40 characters and the only allowed characters are `a-z` , `A-Z` , `0-9` , and the following special characters: `*?-!#$%&'.^_`|~.` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-routingrule-matchheadervalue.html#cfn-apigatewayv2-routingrule-matchheadervalue-header
	//
	Header *string `field:"optional" json:"header" yaml:"header"`
	// The case sensitive header glob value to be matched against entire header value.
	//
	// The header glob value must be less than 128 characters and the only allowed characters are `a-z` , `A-Z` , `0-9` , and the following special characters: `*?-!#$%&'.^_`|~` . Wildcard matching is supported for header glob values but must be for `*prefix-match` , `suffix-match*` , or `*infix*-match` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-routingrule-matchheadervalue.html#cfn-apigatewayv2-routingrule-matchheadervalue-valueglob
	//
	ValueGlob *string `field:"optional" json:"valueGlob" yaml:"valueGlob"`
}

