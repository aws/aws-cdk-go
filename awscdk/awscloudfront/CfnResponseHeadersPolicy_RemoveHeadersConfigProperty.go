package awscloudfront


// A list of HTTP header names that CloudFront removes from HTTP responses to requests that match the cache behavior that this response headers policy is attached to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   removeHeadersConfigProperty := &RemoveHeadersConfigProperty{
//   	Items: []interface{}{
//   		&RemoveHeaderProperty{
//   			Header: jsii.String("header"),
//   		},
//   	},
//   }
//
type CfnResponseHeadersPolicy_RemoveHeadersConfigProperty struct {
	// The list of HTTP header names.
	Items interface{} `field:"required" json:"items" yaml:"items"`
}

