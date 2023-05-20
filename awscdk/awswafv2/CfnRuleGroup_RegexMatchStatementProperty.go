package awswafv2


// A rule statement used to search web request components for a match against a single regular expression.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//   		JsonBody: &JsonBodyProperty{
//   			MatchPattern: &JsonMatchPatternProperty{
//   				All: all,
//   				IncludedPaths: []*string{
//   					jsii.String("includedPaths"),
//   				},
//   			},
//   			MatchScope: jsii.String("matchScope"),
//
//   			// the properties below are optional
//   			InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   			OversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		Method: method,
//   		QueryString: queryString,
//   		SingleHeader: singleHeader,
//   		SingleQueryArgument: singleQueryArgument,
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
type CfnRuleGroup_RegexMatchStatementProperty struct {
	// The part of the web request that you want AWS WAF to inspect.
	FieldToMatch interface{} `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// The string representing the regular expression.
	RegexString *string `field:"required" json:"regexString" yaml:"regexString"`
	// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection.
	//
	// Text transformations are used in rule match statements, to transform the `FieldToMatch` request component before inspecting it, and they're used in rate-based rule statements, to transform request components before using them as custom aggregation keys. If you specify one or more transformations to apply, AWS WAF performs all transformations on the specified content, starting from the lowest priority setting, and then uses the component contents.
	TextTransformations interface{} `field:"required" json:"textTransformations" yaml:"textTransformations"`
}

