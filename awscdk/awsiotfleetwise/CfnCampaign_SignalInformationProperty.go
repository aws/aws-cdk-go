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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-signalinformation.html
//
type CfnCampaign_SignalInformationProperty struct {
	// The name of the signal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-signalinformation.html#cfn-iotfleetwise-campaign-signalinformation-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// (Optional) The maximum number of samples to collect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-signalinformation.html#cfn-iotfleetwise-campaign-signalinformation-maxsamplecount
	//
	MaxSampleCount *float64 `field:"optional" json:"maxSampleCount" yaml:"maxSampleCount"`
	// (Optional) The minimum duration of time (in milliseconds) between two triggering events to collect data.
	//
	// > If a signal changes often, you might want to collect data at a slower rate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-signalinformation.html#cfn-iotfleetwise-campaign-signalinformation-minimumsamplingintervalms
	//
	MinimumSamplingIntervalMs *float64 `field:"optional" json:"minimumSamplingIntervalMs" yaml:"minimumSamplingIntervalMs"`
}

