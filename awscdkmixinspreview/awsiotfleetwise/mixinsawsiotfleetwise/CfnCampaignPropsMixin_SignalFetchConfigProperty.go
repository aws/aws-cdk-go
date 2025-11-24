package mixinsawsiotfleetwise


// The configuration of the signal fetch operation.
//
// > Access to certain AWS IoT FleetWise features is currently gated. For more information, see [AWS Region and feature availability](https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/fleetwise-regions.html) in the *AWS IoT FleetWise Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   signalFetchConfigProperty := &SignalFetchConfigProperty{
//   	ConditionBased: &ConditionBasedSignalFetchConfigProperty{
//   		ConditionExpression: jsii.String("conditionExpression"),
//   		TriggerMode: jsii.String("triggerMode"),
//   	},
//   	TimeBased: &TimeBasedSignalFetchConfigProperty{
//   		ExecutionFrequencyMs: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-signalfetchconfig.html
//
type CfnCampaignPropsMixin_SignalFetchConfigProperty struct {
	// The configuration of a condition-based signal fetch operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-signalfetchconfig.html#cfn-iotfleetwise-campaign-signalfetchconfig-conditionbased
	//
	ConditionBased interface{} `field:"optional" json:"conditionBased" yaml:"conditionBased"`
	// The configuration of a time-based signal fetch operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-signalfetchconfig.html#cfn-iotfleetwise-campaign-signalfetchconfig-timebased
	//
	TimeBased interface{} `field:"optional" json:"timeBased" yaml:"timeBased"`
}

