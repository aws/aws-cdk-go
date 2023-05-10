package awswafv2


// A rule statement that compares a number of bytes against the size of a request component, using a comparison operator, such as greater than (>) or less than (<).
//
// For example, you can use a size constraint statement to look for query strings that are longer than 100 bytes.
//
// If you configure AWS WAF to inspect the request body, AWS WAF inspects only the number of bytes of the body up to the limit for the web ACL. By default, for regional web ACLs, this limit is 8 KB (8,192 kilobytes) and for CloudFront web ACLs, this limit is 16 KB (16,384 kilobytes). For CloudFront web ACLs, you can increase the limit in the web ACL `AssociationConfig` , for additional fees. If you know that the request body for your web requests should never exceed the inspection limit, you could use a size constraint statement to block requests that have a larger request body size.
//
// If you choose URI for the value of Part of the request to filter on, the slash (/) in the URI counts as one character. For example, the URI `/logo.jpg` is nine characters long.
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
//   sizeConstraintStatementProperty := &SizeConstraintStatementProperty{
//   	ComparisonOperator: jsii.String("comparisonOperator"),
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
//   	Size: jsii.Number(123),
//   	TextTransformations: []interface{}{
//   		&TextTransformationProperty{
//   			Priority: jsii.Number(123),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
type CfnWebACL_SizeConstraintStatementProperty struct {
	// The operator to use to compare the request part to the size setting.
	ComparisonOperator *string `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The part of the web request that you want AWS WAF to inspect.
	FieldToMatch interface{} `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// The size, in byte, to compare to the request part, after any transformations.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection.
	//
	// If you specify one or more transformations in a rule statement, AWS WAF performs all transformations on the content of the request component identified by `FieldToMatch` , starting from the lowest priority setting, before inspecting the content for a match.
	TextTransformations interface{} `field:"required" json:"textTransformations" yaml:"textTransformations"`
}

