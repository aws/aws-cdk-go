package awswafv2


// The part of the web request that you want AWS WAF to inspect.
//
// Include the single `FieldToMatch` type that you want to inspect, with additional specifications as needed, according to the type. You specify a single request component in `FieldToMatch` for each rule statement that requires it. To inspect more than one component of the web request, create a separate rule statement for each component.
//
// Example JSON for a `QueryString` field to match:
//
// `"FieldToMatch": { "QueryString": {} }`
//
// Example JSON for a `Method` field to match specification:
//
// `"FieldToMatch": { "Method": { "Name": "DELETE" } }`.
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
//   fieldToMatchProperty := &fieldToMatchProperty{
//   	allQueryArguments: allQueryArguments,
//   	body: &bodyProperty{
//   		oversizeHandling: jsii.String("oversizeHandling"),
//   	},
//   	cookies: &cookiesProperty{
//   		matchPattern: &cookieMatchPatternProperty{
//   			all: all,
//   			excludedCookies: []*string{
//   				jsii.String("excludedCookies"),
//   			},
//   			includedCookies: []*string{
//   				jsii.String("includedCookies"),
//   			},
//   		},
//   		matchScope: jsii.String("matchScope"),
//   		oversizeHandling: jsii.String("oversizeHandling"),
//   	},
//   	headers: &headersProperty{
//   		matchPattern: &headerMatchPatternProperty{
//   			all: all,
//   			excludedHeaders: []*string{
//   				jsii.String("excludedHeaders"),
//   			},
//   			includedHeaders: []*string{
//   				jsii.String("includedHeaders"),
//   			},
//   		},
//   		matchScope: jsii.String("matchScope"),
//   		oversizeHandling: jsii.String("oversizeHandling"),
//   	},
//   	jsonBody: &jsonBodyProperty{
//   		matchPattern: &jsonMatchPatternProperty{
//   			all: all,
//   			includedPaths: []*string{
//   				jsii.String("includedPaths"),
//   			},
//   		},
//   		matchScope: jsii.String("matchScope"),
//
//   		// the properties below are optional
//   		invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   		oversizeHandling: jsii.String("oversizeHandling"),
//   	},
//   	method: method,
//   	queryString: queryString,
//   	singleHeader: singleHeader,
//   	singleQueryArgument: singleQueryArgument,
//   	uriPath: uriPath,
//   }
//
type CfnRuleGroup_FieldToMatchProperty struct {
	// Inspect all query arguments.
	AllQueryArguments interface{} `field:"optional" json:"allQueryArguments" yaml:"allQueryArguments"`
	// Inspect the request body as plain text.
	//
	// The request body immediately follows the request headers. This is the part of a request that contains any additional data that you want to send to your web server as the HTTP request body, such as data from a form.
	//
	// Only the first 8 KB (8192 bytes) of the request body are forwarded to AWS WAF for inspection by the underlying host service. For information about how to handle oversized request bodies, see the `Body` object configuration.
	Body interface{} `field:"optional" json:"body" yaml:"body"`
	// Inspect the request cookies.
	//
	// You must configure scope and pattern matching filters in the `Cookies` object, to define the set of cookies and the parts of the cookies that AWS WAF inspects.
	//
	// Only the first 8 KB (8192 bytes) of a request's cookies and only the first 200 cookies are forwarded to AWS WAF for inspection by the underlying host service. You must configure how to handle any oversize cookie content in the `Cookies` object. AWS WAF applies the pattern matching filters to the cookies that it receives from the underlying host service.
	Cookies interface{} `field:"optional" json:"cookies" yaml:"cookies"`
	// Inspect the request headers.
	//
	// You must configure scope and pattern matching filters in the `Headers` object, to define the set of headers to and the parts of the headers that AWS WAF inspects.
	//
	// Only the first 8 KB (8192 bytes) of a request's headers and only the first 200 headers are forwarded to AWS WAF for inspection by the underlying host service. You must configure how to handle any oversize header content in the `Headers` object. AWS WAF applies the pattern matching filters to the headers that it receives from the underlying host service.
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// Inspect the request body as JSON.
	//
	// The request body immediately follows the request headers. This is the part of a request that contains any additional data that you want to send to your web server as the HTTP request body, such as data from a form.
	//
	// Only the first 8 KB (8192 bytes) of the request body are forwarded to AWS WAF for inspection by the underlying host service. For information about how to handle oversized request bodies, see the `JsonBody` object configuration.
	JsonBody interface{} `field:"optional" json:"jsonBody" yaml:"jsonBody"`
	// Inspect the HTTP method.
	//
	// The method indicates the type of operation that the request is asking the origin to perform.
	Method interface{} `field:"optional" json:"method" yaml:"method"`
	// Inspect the query string.
	//
	// This is the part of a URL that appears after a `?` character, if any.
	QueryString interface{} `field:"optional" json:"queryString" yaml:"queryString"`
	// Inspect a single header.
	//
	// Provide the name of the header to inspect, for example, `User-Agent` or `Referer` . This setting isn't case sensitive.
	//
	// Example JSON: `"SingleHeader": { "Name": "haystack" }`
	//
	// Alternately, you can filter and inspect all headers with the `Headers` `FieldToMatch` setting.
	SingleHeader interface{} `field:"optional" json:"singleHeader" yaml:"singleHeader"`
	// Inspect a single query argument.
	//
	// Provide the name of the query argument to inspect, such as *UserName* or *SalesRegion* . The name can be up to 30 characters long and isn't case sensitive.
	//
	// Example JSON: `"SingleQueryArgument": { "Name": "myArgument" }`.
	SingleQueryArgument interface{} `field:"optional" json:"singleQueryArgument" yaml:"singleQueryArgument"`
	// Inspect the request URI path.
	//
	// This is the part of the web request that identifies a resource, for example, `/images/daily-ad.jpg` .
	UriPath interface{} `field:"optional" json:"uriPath" yaml:"uriPath"`
}

