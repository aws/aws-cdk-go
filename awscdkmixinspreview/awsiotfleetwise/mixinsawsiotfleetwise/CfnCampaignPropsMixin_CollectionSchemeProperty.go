package mixinsawsiotfleetwise


// Specifies what data to collect and how often or when to collect it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   collectionSchemeProperty := &CollectionSchemeProperty{
//   	ConditionBasedCollectionScheme: &ConditionBasedCollectionSchemeProperty{
//   		ConditionLanguageVersion: jsii.Number(123),
//   		Expression: jsii.String("expression"),
//   		MinimumTriggerIntervalMs: jsii.Number(123),
//   		TriggerMode: jsii.String("triggerMode"),
//   	},
//   	TimeBasedCollectionScheme: &TimeBasedCollectionSchemeProperty{
//   		PeriodMs: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-collectionscheme.html
//
type CfnCampaignPropsMixin_CollectionSchemeProperty struct {
	// Information about a collection scheme that uses a simple logical expression to recognize what data to collect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-collectionscheme.html#cfn-iotfleetwise-campaign-collectionscheme-conditionbasedcollectionscheme
	//
	ConditionBasedCollectionScheme interface{} `field:"optional" json:"conditionBasedCollectionScheme" yaml:"conditionBasedCollectionScheme"`
	// Information about a collection scheme that uses a time period to decide how often to collect data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-collectionscheme.html#cfn-iotfleetwise-campaign-collectionscheme-timebasedcollectionscheme
	//
	TimeBasedCollectionScheme interface{} `field:"optional" json:"timeBasedCollectionScheme" yaml:"timeBasedCollectionScheme"`
}

