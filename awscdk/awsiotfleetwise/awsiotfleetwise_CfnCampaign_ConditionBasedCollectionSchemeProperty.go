package awsiotfleetwise


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
	// `CfnCampaign.ConditionBasedCollectionSchemeProperty.Expression`.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// `CfnCampaign.ConditionBasedCollectionSchemeProperty.ConditionLanguageVersion`.
	ConditionLanguageVersion *float64 `field:"optional" json:"conditionLanguageVersion" yaml:"conditionLanguageVersion"`
	// `CfnCampaign.ConditionBasedCollectionSchemeProperty.MinimumTriggerIntervalMs`.
	MinimumTriggerIntervalMs *float64 `field:"optional" json:"minimumTriggerIntervalMs" yaml:"minimumTriggerIntervalMs"`
	// `CfnCampaign.ConditionBasedCollectionSchemeProperty.TriggerMode`.
	TriggerMode *string `field:"optional" json:"triggerMode" yaml:"triggerMode"`
}

