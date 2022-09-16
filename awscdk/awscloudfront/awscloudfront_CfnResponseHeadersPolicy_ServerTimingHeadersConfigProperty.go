package awscloudfront


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverTimingHeadersConfigProperty := &serverTimingHeadersConfigProperty{
//   	enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	samplingRate: jsii.Number(123),
//   }
//
type CfnResponseHeadersPolicy_ServerTimingHeadersConfigProperty struct {
	// `CfnResponseHeadersPolicy.ServerTimingHeadersConfigProperty.Enabled`.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// `CfnResponseHeadersPolicy.ServerTimingHeadersConfigProperty.SamplingRate`.
	SamplingRate *float64 `field:"optional" json:"samplingRate" yaml:"samplingRate"`
}

