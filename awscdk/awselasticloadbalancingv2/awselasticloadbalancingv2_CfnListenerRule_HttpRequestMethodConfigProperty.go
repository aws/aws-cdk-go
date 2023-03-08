package awselasticloadbalancingv2


// Information about an HTTP method condition.
//
// HTTP defines a set of request methods, also referred to as HTTP verbs. For more information, see the [HTTP Method Registry](https://docs.aws.amazon.com/https://www.iana.org/assignments/http-methods/http-methods.xhtml) . You can also define custom HTTP methods.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpRequestMethodConfigProperty := &httpRequestMethodConfigProperty{
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnListenerRule_HttpRequestMethodConfigProperty struct {
	// The name of the request method.
	//
	// The maximum size is 40 characters. The allowed characters are A-Z, hyphen (-), and underscore (_). The comparison is case sensitive. Wildcards are not supported; therefore, the method name must be an exact match.
	//
	// If you specify multiple strings, the condition is satisfied if one of the strings matches the HTTP request method. We recommend that you route GET and HEAD requests in the same way, because the response to a HEAD request may be cached.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

