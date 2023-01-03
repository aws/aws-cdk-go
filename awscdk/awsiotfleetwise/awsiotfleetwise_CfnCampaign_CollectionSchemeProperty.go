package awsiotfleetwise


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   collectionSchemeProperty := &collectionSchemeProperty{
//   	conditionBasedCollectionScheme: &conditionBasedCollectionSchemeProperty{
//   		expression: jsii.String("expression"),
//
//   		// the properties below are optional
//   		conditionLanguageVersion: jsii.Number(123),
//   		minimumTriggerIntervalMs: jsii.Number(123),
//   		triggerMode: jsii.String("triggerMode"),
//   	},
//   	timeBasedCollectionScheme: &timeBasedCollectionSchemeProperty{
//   		periodMs: jsii.Number(123),
//   	},
//   }
//
type CfnCampaign_CollectionSchemeProperty struct {
	// `CfnCampaign.CollectionSchemeProperty.ConditionBasedCollectionScheme`.
	ConditionBasedCollectionScheme interface{} `field:"optional" json:"conditionBasedCollectionScheme" yaml:"conditionBasedCollectionScheme"`
	// `CfnCampaign.CollectionSchemeProperty.TimeBasedCollectionScheme`.
	TimeBasedCollectionScheme interface{} `field:"optional" json:"timeBasedCollectionScheme" yaml:"timeBasedCollectionScheme"`
}

