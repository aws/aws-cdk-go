package awselasticache


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recurringChargeProperty := &RecurringChargeProperty{
//   	RecurringChargeAmount: jsii.Number(123),
//   	RecurringChargeFrequency: jsii.String("recurringChargeFrequency"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-reservedcachenode-recurringcharge.html
//
type CfnReservedCacheNode_RecurringChargeProperty struct {
	// The monetary amount of the recurring charge.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-reservedcachenode-recurringcharge.html#cfn-elasticache-reservedcachenode-recurringcharge-recurringchargeamount
	//
	RecurringChargeAmount *float64 `field:"optional" json:"recurringChargeAmount" yaml:"recurringChargeAmount"`
	// The frequency of the recurring charge.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-reservedcachenode-recurringcharge.html#cfn-elasticache-reservedcachenode-recurringcharge-recurringchargefrequency
	//
	RecurringChargeFrequency *string `field:"optional" json:"recurringChargeFrequency" yaml:"recurringChargeFrequency"`
}

