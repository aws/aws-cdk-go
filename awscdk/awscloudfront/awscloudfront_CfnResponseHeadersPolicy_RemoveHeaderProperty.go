package awscloudfront


// The name of an HTTP header that CloudFront removes from HTTP responses to requests that match the cache behavior that this response headers policy is attached to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   removeHeaderProperty := &removeHeaderProperty{
//   	header: jsii.String("header"),
//   }
//
type CfnResponseHeadersPolicy_RemoveHeaderProperty struct {
	// The HTTP header name.
	Header *string `field:"required" json:"header" yaml:"header"`
}

