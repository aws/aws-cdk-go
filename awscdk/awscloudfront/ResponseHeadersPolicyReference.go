package awscloudfront


// A reference to a ResponseHeadersPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   responseHeadersPolicyReference := &ResponseHeadersPolicyReference{
//   	ResponseHeadersPolicyId: jsii.String("responseHeadersPolicyId"),
//   }
//
type ResponseHeadersPolicyReference struct {
	// The Id of the ResponseHeadersPolicy resource.
	ResponseHeadersPolicyId *string `field:"required" json:"responseHeadersPolicyId" yaml:"responseHeadersPolicyId"`
}

