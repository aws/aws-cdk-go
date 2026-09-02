package awsmediatailor


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   recurringConsumptionProperty := &RecurringConsumptionProperty{
//   	AvailMatchingCriteria: []interface{}{
//   		&AvailMatchingCriteriaProperty{
//   			DynamicVariable: jsii.String("dynamicVariable"),
//   			Operator: jsii.String("operator"),
//   		},
//   	},
//   	RetrievedAdExpirationSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-prefetchschedule-recurringconsumption.html
//
type CfnPrefetchSchedulePropsMixin_RecurringConsumptionProperty struct {
	// The configuration for the dynamic variables that determine which ad breaks that MediaTailor inserts prefetched ads in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-prefetchschedule-recurringconsumption.html#cfn-mediatailor-prefetchschedule-recurringconsumption-availmatchingcriteria
	//
	AvailMatchingCriteria interface{} `field:"optional" json:"availMatchingCriteria" yaml:"availMatchingCriteria"`
	// The number of seconds that an ad is available for insertion after it was prefetched.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-prefetchschedule-recurringconsumption.html#cfn-mediatailor-prefetchschedule-recurringconsumption-retrievedadexpirationseconds
	//
	RetrievedAdExpirationSeconds *float64 `field:"optional" json:"retrievedAdExpirationSeconds" yaml:"retrievedAdExpirationSeconds"`
}

