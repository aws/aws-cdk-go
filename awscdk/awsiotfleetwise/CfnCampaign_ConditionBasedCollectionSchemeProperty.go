package awsiotfleetwise


// Information about a collection scheme that uses a simple logical expression to recognize what data to collect.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionBasedCollectionSchemeProperty := &ConditionBasedCollectionSchemeProperty{
//   	Expression: jsii.String("expression"),
//
//   	// the properties below are optional
//   	ConditionLanguageVersion: jsii.Number(123),
//   	MinimumTriggerIntervalMs: jsii.Number(123),
//   	TriggerMode: jsii.String("triggerMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-conditionbasedcollectionscheme.html
//
type CfnCampaign_ConditionBasedCollectionSchemeProperty struct {
	// The logical expression used to recognize what data to collect.
	//
	// For example, `$variable.Vehicle.OutsideAirTemperature >= 105.0` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-conditionbasedcollectionscheme.html#cfn-iotfleetwise-campaign-conditionbasedcollectionscheme-expression
	//
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Specifies the version of the conditional expression language.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-conditionbasedcollectionscheme.html#cfn-iotfleetwise-campaign-conditionbasedcollectionscheme-conditionlanguageversion
	//
	ConditionLanguageVersion *float64 `field:"optional" json:"conditionLanguageVersion" yaml:"conditionLanguageVersion"`
	// The minimum duration of time between two triggering events to collect data, in milliseconds.
	//
	// > If a signal changes often, you might want to collect data at a slower rate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-conditionbasedcollectionscheme.html#cfn-iotfleetwise-campaign-conditionbasedcollectionscheme-minimumtriggerintervalms
	//
	MinimumTriggerIntervalMs *float64 `field:"optional" json:"minimumTriggerIntervalMs" yaml:"minimumTriggerIntervalMs"`
	// Whether to collect data for all triggering events ( `ALWAYS` ).
	//
	// Specify ( `RISING_EDGE` ), or specify only when the condition first evaluates to false. For example, triggering on "AirbagDeployed"; Users aren't interested on triggering when the airbag is already exploded; they only care about the change from not deployed => deployed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-conditionbasedcollectionscheme.html#cfn-iotfleetwise-campaign-conditionbasedcollectionscheme-triggermode
	//
	TriggerMode *string `field:"optional" json:"triggerMode" yaml:"triggerMode"`
}

