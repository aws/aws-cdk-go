package awscloudfront


// A list of HTTP headers that CloudFront includes as values for the `Access-Control-Expose-Headers` HTTP response header.
//
// For more information about the `Access-Control-Expose-Headers` HTTP response header, see [Access-Control-Expose-Headers](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Expose-Headers) in the MDN Web Docs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessControlExposeHeadersProperty := &accessControlExposeHeadersProperty{
//   	items: []*string{
//   		jsii.String("items"),
//   	},
//   }
//
type CfnResponseHeadersPolicy_AccessControlExposeHeadersProperty struct {
	// The list of HTTP headers.
	//
	// You can specify `*` to expose all headers.
	Items *[]*string `field:"required" json:"items" yaml:"items"`
}

