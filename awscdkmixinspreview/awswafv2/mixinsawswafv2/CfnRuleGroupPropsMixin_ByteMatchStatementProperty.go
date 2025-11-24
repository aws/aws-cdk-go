package mixinsawswafv2


// A rule statement that defines a string match search for AWS WAF to apply to web requests.
//
// The byte match statement provides the bytes to search for, the location in requests that you want AWS WAF to search, and other settings. The bytes to search for are typically a string that corresponds with ASCII characters. In the AWS WAF console and the developer guide, this is called a string match statement.
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
//   byteMatchStatementProperty := &ByteMatchStatementProperty{
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
//   	PositionalConstraint: jsii.String("positionalConstraint"),
//   	SearchString: jsii.String("searchString"),
//   	SearchStringBase64: jsii.String("searchStringBase64"),
//   	TextTransformations: []interface{}{
//   		&TextTransformationProperty{
//   			Priority: jsii.Number(123),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-bytematchstatement.html
//
type CfnRuleGroupPropsMixin_ByteMatchStatementProperty struct {
	// The part of the web request that you want AWS WAF to inspect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-bytematchstatement.html#cfn-wafv2-rulegroup-bytematchstatement-fieldtomatch
	//
	FieldToMatch interface{} `field:"optional" json:"fieldToMatch" yaml:"fieldToMatch"`
	// The area within the portion of the web request that you want AWS WAF to search for `SearchString` .
	//
	// Valid values include the following:
	//
	// *CONTAINS*
	//
	// The specified part of the web request must include the value of `SearchString` , but the location doesn't matter.
	//
	// *CONTAINS_WORD*
	//
	// The specified part of the web request must include the value of `SearchString` , and `SearchString` must contain only alphanumeric characters or underscore (A-Z, a-z, 0-9, or _). In addition, `SearchString` must be a word, which means that both of the following are true:
	//
	// - `SearchString` is at the beginning of the specified part of the web request or is preceded by a character other than an alphanumeric character or underscore (_). Examples include the value of a header and `;BadBot` .
	// - `SearchString` is at the end of the specified part of the web request or is followed by a character other than an alphanumeric character or underscore (_), for example, `BadBot;` and `-BadBot;` .
	//
	// *EXACTLY*
	//
	// The value of the specified part of the web request must exactly match the value of `SearchString` .
	//
	// *STARTS_WITH*
	//
	// The value of `SearchString` must appear at the beginning of the specified part of the web request.
	//
	// *ENDS_WITH*
	//
	// The value of `SearchString` must appear at the end of the specified part of the web request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-bytematchstatement.html#cfn-wafv2-rulegroup-bytematchstatement-positionalconstraint
	//
	PositionalConstraint *string `field:"optional" json:"positionalConstraint" yaml:"positionalConstraint"`
	// A string value that you want AWS WAF to search for.
	//
	// AWS WAF searches only in the part of web requests that you designate for inspection in `FieldToMatch` . The maximum length of the value is 200 bytes. For alphabetic characters A-Z and a-z, the value is case sensitive.
	//
	// Don't encode this string. Provide the value that you want AWS WAF to search for. AWS CloudFormation automatically base64 encodes the value for you.
	//
	// For example, suppose the value of `Type` is `HEADER` and the value of `Data` is `User-Agent` . If you want to search the `User-Agent` header for the value `BadBot` , you provide the string `BadBot` in the value of `SearchString` .
	//
	// You must specify either `SearchString` or `SearchStringBase64` in a `ByteMatchStatement` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-bytematchstatement.html#cfn-wafv2-rulegroup-bytematchstatement-searchstring
	//
	SearchString *string `field:"optional" json:"searchString" yaml:"searchString"`
	// String to search for in a web request component, base64-encoded.
	//
	// If you don't want to encode the string, specify the unencoded value in `SearchString` instead.
	//
	// You must specify either `SearchString` or `SearchStringBase64` in a `ByteMatchStatement` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-bytematchstatement.html#cfn-wafv2-rulegroup-bytematchstatement-searchstringbase64
	//
	SearchStringBase64 *string `field:"optional" json:"searchStringBase64" yaml:"searchStringBase64"`
	// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection.
	//
	// If you specify one or more transformations in a rule statement, AWS WAF performs all transformations on the content of the request component identified by `FieldToMatch` , starting from the lowest priority setting, before inspecting the content for a match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-bytematchstatement.html#cfn-wafv2-rulegroup-bytematchstatement-texttransformations
	//
	TextTransformations interface{} `field:"optional" json:"textTransformations" yaml:"textTransformations"`
}

