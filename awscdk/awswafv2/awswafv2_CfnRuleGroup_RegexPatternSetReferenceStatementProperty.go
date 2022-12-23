package awswafv2


// A rule statement used to search web request components for matches with regular expressions.
//
// To use this, create a `RegexPatternSet` that specifies the expressions that you want to detect, then use the ARN of that set in this statement. A web request matches the pattern set rule statement if the request component matches any of the patterns in the set.
//
// Each regex pattern set rule statement references a regex pattern set. You create and maintain the set independent of your rules. This allows you to use the single set in multiple rules. When you update the referenced set, AWS WAF automatically updates all rules that reference it.
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
//   regexPatternSetReferenceStatementProperty := &regexPatternSetReferenceStatementProperty{
//   	arn: jsii.String("arn"),
//   	fieldToMatch: &fieldToMatchProperty{
//   		allQueryArguments: allQueryArguments,
//   		body: &bodyProperty{
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		cookies: &cookiesProperty{
//   			matchPattern: &cookieMatchPatternProperty{
//   				all: all,
//   				excludedCookies: []*string{
//   					jsii.String("excludedCookies"),
//   				},
//   				includedCookies: []*string{
//   					jsii.String("includedCookies"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		headers: &headersProperty{
//   			matchPattern: &headerMatchPatternProperty{
//   				all: all,
//   				excludedHeaders: []*string{
//   					jsii.String("excludedHeaders"),
//   				},
//   				includedHeaders: []*string{
//   					jsii.String("includedHeaders"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		jsonBody: &jsonBodyProperty{
//   			matchPattern: &jsonMatchPatternProperty{
//   				all: all,
//   				includedPaths: []*string{
//   					jsii.String("includedPaths"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//
//   			// the properties below are optional
//   			invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		method: method,
//   		queryString: queryString,
//   		singleHeader: singleHeader,
//   		singleQueryArgument: singleQueryArgument,
//   		uriPath: uriPath,
//   	},
//   	textTransformations: []interface{}{
//   		&textTransformationProperty{
//   			priority: jsii.Number(123),
//   			type: jsii.String("type"),
//   		},
//   	},
//   }
//
type CfnRuleGroup_RegexPatternSetReferenceStatementProperty struct {
	// The Amazon Resource Name (ARN) of the `RegexPatternSet` that this statement references.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The part of the web request that you want AWS WAF to inspect.
	FieldToMatch interface{} `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection.
	//
	// If you specify one or more transformations in a rule statement, AWS WAF performs all transformations on the content of the request component identified by `FieldToMatch` , starting from the lowest priority setting, before inspecting the content for a match.
	TextTransformations interface{} `field:"required" json:"textTransformations" yaml:"textTransformations"`
}

