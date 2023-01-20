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
//   var jsonBody interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var uriPath interface{}
//
//   fieldToMatchProperty := &fieldToMatchProperty{
//   	jsonBody: jsonBody,
//   	method: method,
//   	queryString: queryString,
//   	singleHeader: singleHeader,
//   	uriPath: uriPath,
//   }
//
type CfnLoggingConfiguration_FieldToMatchProperty struct {
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
	// Inspect the request URI path.
	//
	// This is the part of the web request that identifies a resource, for example, `/images/daily-ad.jpg` .
	UriPath interface{} `field:"optional" json:"uriPath" yaml:"uriPath"`
}

