package awscloudfront


// A list of HTTP response header names and their values.
//
// CloudFront includes these headers in HTTP responses that it sends for requests that match a cache behavior that’s associated with this response headers policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customHeadersConfigProperty := &customHeadersConfigProperty{
//   	items: []interface{}{
//   		&customHeaderProperty{
//   			header: jsii.String("header"),
//   			override: jsii.Boolean(false),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnResponseHeadersPolicy_CustomHeadersConfigProperty struct {
	// The list of HTTP response headers and their values.
	Items interface{} `field:"required" json:"items" yaml:"items"`
}

