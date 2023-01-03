package awscloudfront


// An HTTP response header name and its value.
//
// CloudFront includes this header in HTTP responses that it sends for requests that match a cache behavior that’s associated with this response headers policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   responseCustomHeader := &responseCustomHeader{
//   	header: jsii.String("header"),
//   	override: jsii.Boolean(false),
//   	value: jsii.String("value"),
//   }
//
type ResponseCustomHeader struct {
	// The HTTP response header name.
	Header *string `field:"required" json:"header" yaml:"header"`
	// A Boolean that determines whether CloudFront overrides a response header with the same name received from the origin with the header specified here.
	Override *bool `field:"required" json:"override" yaml:"override"`
	// The value for the HTTP response header.
	Value *string `field:"required" json:"value" yaml:"value"`
}

