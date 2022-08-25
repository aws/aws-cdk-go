package awswafv2


// A rule statement that defines a string match search for AWS WAF to apply to web requests.
//
// The byte match statement provides the bytes to search for, the location in requests that you want AWS WAF to search, and other settings. The bytes to search for are typically a string that corresponds with ASCII characters. In the AWS WAF console and the developer guide, this is called a string match statement.
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
//   byteMatchStatementProperty := &byteMatchStatementProperty{
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
//   	positionalConstraint: jsii.String("positionalConstraint"),
//   	textTransformations: []interface{}{
//   		&textTransformationProperty{
//   			priority: jsii.Number(123),
//   			type: jsii.String("type"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	searchString: jsii.String("searchString"),
//   	searchStringBase64: jsii.String("searchStringBase64"),
//   }
//
type CfnRuleGroup_ByteMatchStatementProperty struct {
	// The part of the web request that you want AWS WAF to inspect.
	FieldToMatch interface{} `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
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
	PositionalConstraint *string `field:"required" json:"positionalConstraint" yaml:"positionalConstraint"`
	// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection.
	//
	// If you specify one or more transformations in a rule statement, AWS WAF performs all transformations on the content of the request component identified by `FieldToMatch` , starting from the lowest priority setting, before inspecting the content for a match.
	TextTransformations interface{} `field:"required" json:"textTransformations" yaml:"textTransformations"`
	// A string value that you want AWS WAF to search for.
	//
	// AWS WAF searches only in the part of web requests that you designate for inspection in `FieldToMatch` . The maximum length of the value is 50 bytes. For alphabetic characters A-Z and a-z, the value is case sensitive.
	//
	// Don't encode this string. Provide the value that you want AWS WAF to search for. AWS CloudFormation automatically base64 encodes the value for you.
	//
	// For example, suppose the value of `Type` is `HEADER` and the value of `Data` is `User-Agent` . If you want to search the `User-Agent` header for the value `BadBot` , you provide the string `BadBot` in the value of `SearchString` .
	//
	// You must specify either `SearchString` or `SearchStringBase64` in a `ByteMatchStatement` .
	SearchString *string `field:"optional" json:"searchString" yaml:"searchString"`
	// String to search for in a web request component, base64-encoded.
	//
	// If you don't want to encode the string, specify the unencoded value in `SearchString` instead.
	//
	// You must specify either `SearchString` or `SearchStringBase64` in a `ByteMatchStatement` .
	SearchStringBase64 *string `field:"optional" json:"searchStringBase64" yaml:"searchStringBase64"`
}

