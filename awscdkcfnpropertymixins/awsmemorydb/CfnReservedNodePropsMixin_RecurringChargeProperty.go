package awsmemorydb


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   recurringChargeProperty := &RecurringChargeProperty{
//   	RecurringChargeAmount: jsii.Number(123),
//   	RecurringChargeFrequency: jsii.String("recurringChargeFrequency"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-memorydb-reservednode-recurringcharge.html
//
type CfnReservedNodePropsMixin_RecurringChargeProperty struct {
	// The amount of the recurring charge.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-memorydb-reservednode-recurringcharge.html#cfn-memorydb-reservednode-recurringcharge-recurringchargeamount
	//
	RecurringChargeAmount *float64 `field:"optional" json:"recurringChargeAmount" yaml:"recurringChargeAmount"`
	// The frequency of the recurring charge.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-memorydb-reservednode-recurringcharge.html#cfn-memorydb-reservednode-recurringcharge-recurringchargefrequency
	//
	RecurringChargeFrequency *string `field:"optional" json:"recurringChargeFrequency" yaml:"recurringChargeFrequency"`
}

