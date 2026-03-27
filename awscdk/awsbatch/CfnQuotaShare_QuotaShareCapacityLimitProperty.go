package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   quotaShareCapacityLimitProperty := &QuotaShareCapacityLimitProperty{
//   	CapacityUnit: jsii.String("capacityUnit"),
//   	MaxCapacity: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-quotashare-quotasharecapacitylimit.html
//
type CfnQuotaShare_QuotaShareCapacityLimitProperty struct {
	// The unit of compute capacity for the capacityLimit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-quotashare-quotasharecapacitylimit.html#cfn-batch-quotashare-quotasharecapacitylimit-capacityunit
	//
	CapacityUnit *string `field:"required" json:"capacityUnit" yaml:"capacityUnit"`
	// The maximum capacity available for the quota share.
	//
	// This value represents the maximum amount of resources that can be allocated to jobs in the quota share without borrowing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-quotashare-quotasharecapacitylimit.html#cfn-batch-quotashare-quotasharecapacitylimit-maxcapacity
	//
	MaxCapacity *float64 `field:"required" json:"maxCapacity" yaml:"maxCapacity"`
}

