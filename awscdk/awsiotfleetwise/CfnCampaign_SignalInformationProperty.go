package awsiotfleetwise


// Information about a signal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   signalInformationProperty := &SignalInformationProperty{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	MaxSampleCount: jsii.Number(123),
//   	MinimumSamplingIntervalMs: jsii.Number(123),
//   }
//
type CfnCampaign_SignalInformationProperty struct {
	// The name of the signal.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The maximum number of samples to collect.
	MaxSampleCount *float64 `field:"optional" json:"maxSampleCount" yaml:"maxSampleCount"`
	// The minimum duration of time (in milliseconds) between two triggering events to collect data.
	//
	// > If a signal changes often, you might want to collect data at a slower rate.
	MinimumSamplingIntervalMs *float64 `field:"optional" json:"minimumSamplingIntervalMs" yaml:"minimumSamplingIntervalMs"`
}

