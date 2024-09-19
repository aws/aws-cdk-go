package awsiotfleetwise


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   signalFetchInformationProperty := &SignalFetchInformationProperty{
//   	Actions: []*string{
//   		jsii.String("actions"),
//   	},
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
//
//   	// the properties below are optional
//   	ConditionLanguageVersion: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-signalfetchinformation.html
//
type CfnCampaign_SignalFetchInformationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-signalfetchinformation.html#cfn-iotfleetwise-campaign-signalfetchinformation-actions
	//
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-signalfetchinformation.html#cfn-iotfleetwise-campaign-signalfetchinformation-fullyqualifiedname
	//
	FullyQualifiedName *string `field:"required" json:"fullyQualifiedName" yaml:"fullyQualifiedName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-signalfetchinformation.html#cfn-iotfleetwise-campaign-signalfetchinformation-signalfetchconfig
	//
	SignalFetchConfig interface{} `field:"required" json:"signalFetchConfig" yaml:"signalFetchConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-signalfetchinformation.html#cfn-iotfleetwise-campaign-signalfetchinformation-conditionlanguageversion
	//
	ConditionLanguageVersion *float64 `field:"optional" json:"conditionLanguageVersion" yaml:"conditionLanguageVersion"`
}

