package awscloudfront


// A list of origins (domain names) that CloudFront can use as the value for the `Access-Control-Allow-Origin` HTTP response header.
//
// For more information about the `Access-Control-Allow-Origin` HTTP response header, see [Access-Control-Allow-Origin](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Origin) in the MDN Web Docs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessControlAllowOriginsProperty := &accessControlAllowOriginsProperty{
//   	items: []*string{
//   		jsii.String("items"),
//   	},
//   }
//
type CfnResponseHeadersPolicy_AccessControlAllowOriginsProperty struct {
	// The list of origins (domain names).
	//
	// You can specify `*` to allow all origins.
	Items *[]*string `field:"required" json:"items" yaml:"items"`
}

