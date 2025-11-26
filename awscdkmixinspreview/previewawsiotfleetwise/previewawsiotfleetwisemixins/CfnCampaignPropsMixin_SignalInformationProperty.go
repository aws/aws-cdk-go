package previewawsiotfleetwisemixins


// Information about a signal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   signalInformationProperty := &SignalInformationProperty{
//   	DataPartitionId: jsii.String("dataPartitionId"),
//   	MaxSampleCount: jsii.Number(123),
//   	MinimumSamplingIntervalMs: jsii.Number(123),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-signalinformation.html
//
type CfnCampaignPropsMixin_SignalInformationProperty struct {
	// The ID of the data partition this signal is associated with.
	//
	// The ID must match one of the IDs provided in `dataPartitions` . This is accomplished either by specifying a particular data partition ID or by using `default` for an established default partition. You can establish a default partition in the `DataPartition` data type.
	//
	// > If you upload a signal as a condition for a campaign's data partition, the same signal must be included in `signalsToCollect` . > Access to certain AWS IoT FleetWise features is currently gated. For more information, see [AWS Region and feature availability](https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/fleetwise-regions.html) in the *AWS IoT FleetWise Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-signalinformation.html#cfn-iotfleetwise-campaign-signalinformation-datapartitionid
	//
	DataPartitionId *string `field:"optional" json:"dataPartitionId" yaml:"dataPartitionId"`
	// The maximum number of samples to collect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-signalinformation.html#cfn-iotfleetwise-campaign-signalinformation-maxsamplecount
	//
	MaxSampleCount *float64 `field:"optional" json:"maxSampleCount" yaml:"maxSampleCount"`
	// The minimum duration of time (in milliseconds) between two triggering events to collect data.
	//
	// > If a signal changes often, you might want to collect data at a slower rate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-signalinformation.html#cfn-iotfleetwise-campaign-signalinformation-minimumsamplingintervalms
	//
	MinimumSamplingIntervalMs *float64 `field:"optional" json:"minimumSamplingIntervalMs" yaml:"minimumSamplingIntervalMs"`
	// The name of the signal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-signalinformation.html#cfn-iotfleetwise-campaign-signalinformation-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

