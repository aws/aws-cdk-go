package awsiotfleetwise


// Information about a collection scheme that uses a simple logical expression to recognize what data to collect.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionBasedCollectionSchemeProperty := &conditionBasedCollectionSchemeProperty{
//   	expression: jsii.String("expression"),
//
//   	// the properties below are optional
//   	conditionLanguageVersion: jsii.Number(123),
//   	minimumTriggerIntervalMs: jsii.Number(123),
//   	triggerMode: jsii.String("triggerMode"),
//   }
//
type CfnCampaign_ConditionBasedCollectionSchemeProperty struct {
	// The logical expression used to recognize what data to collect.
	//
	// For example, `$variable.Vehicle.OutsideAirTemperature >= 105.0` .
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Specifies the version of the conditional expression language.
	ConditionLanguageVersion *float64 `field:"optional" json:"conditionLanguageVersion" yaml:"conditionLanguageVersion"`
	// The minimum duration of time between two triggering events to collect data, in milliseconds.
	//
	// > If a signal changes often, you might want to collect data at a slower rate.
	MinimumTriggerIntervalMs *float64 `field:"optional" json:"minimumTriggerIntervalMs" yaml:"minimumTriggerIntervalMs"`
	// Whether to collect data for all triggering events ( `ALWAYS` ).
	//
	// Specify ( `RISING_EDGE` ), or specify only when the condition first evaluates to false. For example, triggering on "AirbagDeployed"; Users aren't interested on triggering when the airbag is already exploded; they only care about the change from not deployed => deployed.
	TriggerMode *string `field:"optional" json:"triggerMode" yaml:"triggerMode"`
}

