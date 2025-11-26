package previewawsapigatewayv2mixins


// Represents a `MatchHeaders` condition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   matchHeadersProperty := &MatchHeadersProperty{
//   	AnyOf: []interface{}{
//   		&MatchHeaderValueProperty{
//   			Header: jsii.String("header"),
//   			ValueGlob: jsii.String("valueGlob"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-routingrule-matchheaders.html
//
type CfnRoutingRulePropsMixin_MatchHeadersProperty struct {
	// The header name and header value glob to be matched.
	//
	// The matchHeaders condition is matched if any of the header name and header value globs are matched.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-routingrule-matchheaders.html#cfn-apigatewayv2-routingrule-matchheaders-anyof
	//
	AnyOf interface{} `field:"optional" json:"anyOf" yaml:"anyOf"`
}

