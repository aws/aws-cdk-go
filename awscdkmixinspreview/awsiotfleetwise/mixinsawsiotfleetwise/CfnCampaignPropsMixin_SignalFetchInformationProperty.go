package mixinsawsiotfleetwise


// Information about the signal to be fetched.
//
// > Access to certain AWS IoT FleetWise features is currently gated. For more information, see [AWS Region and feature availability](https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/fleetwise-regions.html) in the *AWS IoT FleetWise Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   signalFetchInformationProperty := &SignalFetchInformationProperty{
//   	Actions: []*string{
//   		jsii.String("actions"),
//   	},
//   	ConditionLanguageVersion: jsii.Number(123),
//   	FullyQualifiedName: jsii.String("fullyQualifiedName"),
//   	SignalFetchConfig: &SignalFetchConfigProperty{
//   		ConditionBased: &ConditionBasedSignalFetchConfigProperty{
//   			ConditionExpression: jsii.String("conditionExpression"),
//   			TriggerMode: jsii.String("triggerMode"),
//   		},
//   		TimeBased: &TimeBasedSignalFetchConfigProperty{
//   			ExecutionFrequencyMs: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-signalfetchinformation.html
//
type CfnCampaignPropsMixin_SignalFetchInformationProperty struct {
	// The actions to be performed by the signal fetch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-signalfetchinformation.html#cfn-iotfleetwise-campaign-signalfetchinformation-actions
	//
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// The version of the condition language used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-signalfetchinformation.html#cfn-iotfleetwise-campaign-signalfetchinformation-conditionlanguageversion
	//
	ConditionLanguageVersion *float64 `field:"optional" json:"conditionLanguageVersion" yaml:"conditionLanguageVersion"`
	// The fully qualified name of the signal to be fetched.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-signalfetchinformation.html#cfn-iotfleetwise-campaign-signalfetchinformation-fullyqualifiedname
	//
	FullyQualifiedName *string `field:"optional" json:"fullyQualifiedName" yaml:"fullyQualifiedName"`
	// The configuration of the signal fetch operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-signalfetchinformation.html#cfn-iotfleetwise-campaign-signalfetchinformation-signalfetchconfig
	//
	SignalFetchConfig interface{} `field:"optional" json:"signalFetchConfig" yaml:"signalFetchConfig"`
}

