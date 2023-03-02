package awscloudfront


// A complex type that contains `HeaderName` and `HeaderValue` elements, if any, for this distribution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   originCustomHeaderProperty := &originCustomHeaderProperty{
//   	headerName: jsii.String("headerName"),
//   	headerValue: jsii.String("headerValue"),
//   }
//
type CfnDistribution_OriginCustomHeaderProperty struct {
	// The name of a header that you want CloudFront to send to your origin.
	//
	// For more information, see [Adding Custom Headers to Origin Requests](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/forward-custom-headers.html) in the *Amazon CloudFront Developer Guide* .
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
	// The value for the header that you specified in the `HeaderName` field.
	HeaderValue *string `field:"required" json:"headerValue" yaml:"headerValue"`
}

