package awsrds


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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-reserveddbinstance-recurringcharge.html
//
type CfnReservedDBInstance_RecurringChargeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-reserveddbinstance-recurringcharge.html#cfn-rds-reserveddbinstance-recurringcharge-recurringchargeamount
	//
	RecurringChargeAmount *float64 `field:"optional" json:"recurringChargeAmount" yaml:"recurringChargeAmount"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-reserveddbinstance-recurringcharge.html#cfn-rds-reserveddbinstance-recurringcharge-recurringchargefrequency
	//
	RecurringChargeFrequency *string `field:"optional" json:"recurringChargeFrequency" yaml:"recurringChargeFrequency"`
}

