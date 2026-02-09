package previewawswafv2mixins


// A rule statement used to search web request components for a match against a single regular expression.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var uriPath interface{}
//
//   regexMatchStatementProperty := &RegexMatchStatementProperty{
//   	FieldToMatch: &FieldToMatchProperty{
//   		AllQueryArguments: allQueryArguments,
//   		Body: &BodyProperty{
//   			OversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		Cookies: &CookiesProperty{
//   			MatchPattern: &CookieMatchPatternProperty{
//   				All: all,
//   				ExcludedCookies: []*string{
//   					jsii.String("excludedCookies"),
//   				},
//   				IncludedCookies: []*string{
//   					jsii.String("includedCookies"),
//   				},
//   			},
//   			MatchScope: jsii.String("matchScope"),
//   			OversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		HeaderOrder: &HeaderOrderProperty{
//   			OversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		Headers: &HeadersProperty{
//   			MatchPattern: &HeaderMatchPatternProperty{
//   				All: all,
//   				ExcludedHeaders: []*string{
//   					jsii.String("excludedHeaders"),
//   				},
//   				IncludedHeaders: []*string{
//   					jsii.String("includedHeaders"),
//   				},
//   			},
//   			MatchScope: jsii.String("matchScope"),
//   			OversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		Ja3Fingerprint: &JA3FingerprintProperty{
//   			FallbackBehavior: jsii.String("fallbackBehavior"),
//   		},
//   		Ja4Fingerprint: &JA4FingerprintProperty{
//   			FallbackBehavior: jsii.String("fallbackBehavior"),
//   		},
//   		JsonBody: &JsonBodyProperty{
//   			InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   			MatchPattern: &JsonMatchPatternProperty{
//   				All: all,
//   				IncludedPaths: []*string{
//   					jsii.String("includedPaths"),
//   				},
//   			},
//   			MatchScope: jsii.String("matchScope"),
//   			OversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		Method: method,
//   		QueryString: queryString,
//   		SingleHeader: singleHeader,
//   		SingleQueryArgument: singleQueryArgument,
//   		UriFragment: &UriFragmentProperty{
//   			FallbackBehavior: jsii.String("fallbackBehavior"),
//   		},
//   		UriPath: uriPath,
//   	},
//   	RegexString: jsii.String("regexString"),
//   	TextTransformations: []interface{}{
//   		&TextTransformationProperty{
//   			Priority: jsii.Number(123),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-regexmatchstatement.html
//
type CfnWebACLPropsMixin_RegexMatchStatementProperty struct {
	// The part of the web request that you want AWS WAF to inspect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-regexmatchstatement.html#cfn-wafv2-webacl-regexmatchstatement-fieldtomatch
	//
	FieldToMatch interface{} `field:"optional" json:"fieldToMatch" yaml:"fieldToMatch"`
	// The string representing the regular expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-regexmatchstatement.html#cfn-wafv2-webacl-regexmatchstatement-regexstring
	//
	RegexString *string `field:"optional" json:"regexString" yaml:"regexString"`
	// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection.
	//
	// If you specify one or more transformations in a rule statement, AWS WAF performs all transformations on the content of the request component identified by `FieldToMatch` , starting from the lowest priority setting, before inspecting the content for a match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-regexmatchstatement.html#cfn-wafv2-webacl-regexmatchstatement-texttransformations
	//
	TextTransformations interface{} `field:"optional" json:"textTransformations" yaml:"textTransformations"`
}

