package awselasticloadbalancingv2


// Information about a query string condition.
//
// The query string component of a URI starts after the first '?' character and is terminated by either a '#' character or the end of the URI. A typical query string contains key/value pairs separated by '&' characters. The allowed characters are specified by RFC 3986. Any character can be percentage encoded.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queryStringConfigProperty := &queryStringConfigProperty{
//   	values: []interface{}{
//   		&queryStringKeyValueProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnListenerRule_QueryStringConfigProperty struct {
	// One or more key/value pairs or values to find in the query string.
	//
	// The maximum size of each string is 128 characters. The comparison is case insensitive. The following wildcard characters are supported: * (matches 0 or more characters) and ? (matches exactly 1 character). To search for a literal '*' or '?' character in a query string, you must escape these characters in `Values` using a '\' character.
	//
	// If you specify multiple key/value pairs or values, the condition is satisfied if one of them is found in the query string.
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

