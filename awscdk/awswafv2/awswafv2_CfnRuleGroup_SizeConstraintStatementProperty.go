package awswafv2


// A rule statement that compares a number of bytes against the size of a request component, using a comparison operator, such as greater than (>) or less than (<).
//
// For example, you can use a size constraint statement to look for query strings that are longer than 100 bytes.
//
// If you configure AWS WAF to inspect the request body, AWS WAF inspects only the first 8192 bytes (8 KB). If the request body for your web requests never exceeds 8192 bytes, you could use a size constraint statement to block requests that have a request body greater than 8192 bytes.
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
//   sizeConstraintStatementProperty := &sizeConstraintStatementProperty{
//   	comparisonOperator: jsii.String("comparisonOperator"),
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
//   	size: jsii.Number(123),
//   	textTransformations: []interface{}{
//   		&textTransformationProperty{
//   			priority: jsii.Number(123),
//   			type: jsii.String("type"),
//   		},
//   	},
//   }
//
type CfnRuleGroup_SizeConstraintStatementProperty struct {
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

