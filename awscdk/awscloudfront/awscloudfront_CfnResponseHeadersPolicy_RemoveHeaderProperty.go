package awscloudfront


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
	// `CfnResponseHeadersPolicy.RemoveHeaderProperty.Header`.
	Header *string `field:"required" json:"header" yaml:"header"`
}

