package awswafv2


// The parts of the request that you want to keep out of the logs.
//
// This is used in the logging configuration `RedactedFields` specification.
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
//   fieldToMatchProperty := &FieldToMatchProperty{
//   	JsonBody: jsonBody,
//   	Method: method,
//   	QueryString: queryString,
//   	SingleHeader: singleHeader,
//   	UriPath: uriPath,
//   }
//
type CfnLoggingConfiguration_FieldToMatchProperty struct {
	// Redact the request body JSON.
	JsonBody interface{} `field:"optional" json:"jsonBody" yaml:"jsonBody"`
	// Redact the indicated HTTP method.
	//
	// The method indicates the type of operation that the request is asking the origin to perform.
	Method interface{} `field:"optional" json:"method" yaml:"method"`
	// Redact the query string.
	//
	// This is the part of a URL that appears after a `?` character, if any.
	QueryString interface{} `field:"optional" json:"queryString" yaml:"queryString"`
	// Redact a single header.
	//
	// Provide the name of the header to inspect, for example, `User-Agent` or `Referer` . This setting isn't case sensitive.
	//
	// Example JSON: `"SingleHeader": { "Name": "haystack" }`.
	SingleHeader interface{} `field:"optional" json:"singleHeader" yaml:"singleHeader"`
	// Redact the request URI path.
	//
	// This is the part of the web request that identifies a resource, for example, `/images/daily-ad.jpg` .
	UriPath interface{} `field:"optional" json:"uriPath" yaml:"uriPath"`
}

