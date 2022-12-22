package awsiotfleetwise


// Specifies what data to collect and how often or when to collect it.
//
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
	// Information about a collection scheme that uses a simple logical expression to recognize what data to collect.
	ConditionBasedCollectionScheme interface{} `field:"optional" json:"conditionBasedCollectionScheme" yaml:"conditionBasedCollectionScheme"`
	// Information about a collection scheme that uses a time period to decide how often to collect data.
	TimeBasedCollectionScheme interface{} `field:"optional" json:"timeBasedCollectionScheme" yaml:"timeBasedCollectionScheme"`
}

