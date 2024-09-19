package awsiotfleetwise


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnCampaign_SignalFetchConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-signalfetchconfig.html#cfn-iotfleetwise-campaign-signalfetchconfig-conditionbased
	//
	ConditionBased interface{} `field:"optional" json:"conditionBased" yaml:"conditionBased"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-signalfetchconfig.html#cfn-iotfleetwise-campaign-signalfetchconfig-timebased
	//
	TimeBased interface{} `field:"optional" json:"timeBased" yaml:"timeBased"`
}

