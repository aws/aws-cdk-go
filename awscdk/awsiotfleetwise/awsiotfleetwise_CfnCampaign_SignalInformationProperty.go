package awsiotfleetwise


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   signalInformationProperty := &signalInformationProperty{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	maxSampleCount: jsii.Number(123),
//   	minimumSamplingIntervalMs: jsii.Number(123),
//   }
//
type CfnCampaign_SignalInformationProperty struct {
	// `CfnCampaign.SignalInformationProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnCampaign.SignalInformationProperty.MaxSampleCount`.
	MaxSampleCount *float64 `field:"optional" json:"maxSampleCount" yaml:"maxSampleCount"`
	// `CfnCampaign.SignalInformationProperty.MinimumSamplingIntervalMs`.
	MinimumSamplingIntervalMs *float64 `field:"optional" json:"minimumSamplingIntervalMs" yaml:"minimumSamplingIntervalMs"`
}

