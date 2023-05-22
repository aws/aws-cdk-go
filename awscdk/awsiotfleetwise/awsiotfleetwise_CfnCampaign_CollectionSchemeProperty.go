package awsiotfleetwise


// Specifies what data to collect and how often or when to collect it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   collectionSchemeProperty := &CollectionSchemeProperty{
//   	ConditionBasedCollectionScheme: &ConditionBasedCollectionSchemeProperty{
//   		Expression: jsii.String("expression"),
//
//   		// the properties below are optional
//   		ConditionLanguageVersion: jsii.Number(123),
//   		MinimumTriggerIntervalMs: jsii.Number(123),
//   		TriggerMode: jsii.String("triggerMode"),
//   	},
//   	TimeBasedCollectionScheme: &TimeBasedCollectionSchemeProperty{
//   		PeriodMs: jsii.Number(123),
//   	},
//   }
//
type CfnCampaign_CollectionSchemeProperty struct {
	// `CfnCampaign.CollectionSchemeProperty.ConditionBasedCollectionScheme`.
	ConditionBasedCollectionScheme interface{} `field:"optional" json:"conditionBasedCollectionScheme" yaml:"conditionBasedCollectionScheme"`
	// `CfnCampaign.CollectionSchemeProperty.TimeBasedCollectionScheme`.
	TimeBasedCollectionScheme interface{} `field:"optional" json:"timeBasedCollectionScheme" yaml:"timeBasedCollectionScheme"`
}

