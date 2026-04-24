package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   quotaShareResourceSharingConfigurationProperty := &QuotaShareResourceSharingConfigurationProperty{
//   	Strategy: jsii.String("strategy"),
//
//   	// the properties below are optional
//   	BorrowLimit: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-quotashare-quotashareresourcesharingconfiguration.html
//
type CfnQuotaShare_QuotaShareResourceSharingConfigurationProperty struct {
	// The resource sharing strategy for the quota share.
	//
	// The `RESERVE` strategy allows a quota share to reserve idle capacity for itself. `LEND` configures the share to lend its idle capacity to another share in need of capacity. The `LEND_AND_BORROW` strategy configures the share to borrow idle capacity from an underutilized share, as well as lend to another share.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-quotashare-quotashareresourcesharingconfiguration.html#cfn-batch-quotashare-quotashareresourcesharingconfiguration-strategy
	//
	Strategy *string `field:"required" json:"strategy" yaml:"strategy"`
	// The maximum percentage of additional capacity that the quota share can borrow from other shares.
	//
	// `BorrowLimit` can only be applied to quota shares with a strategy of `LEND_AND_BORROW`. This value is expressed as a percentage of the quota share's configured CapacityLimits. The `BorrowLimit` is applied uniformly across all capacity units. For example, if the `BorrowLimit` is 200, the quota share can borrow up to 200% of its configured `maxCapacity` for each capacity unit. The default `BorrowLimit` is -1, which indicates unlimited borrowing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-quotashare-quotashareresourcesharingconfiguration.html#cfn-batch-quotashare-quotashareresourcesharingconfiguration-borrowlimit
	//
	BorrowLimit *float64 `field:"optional" json:"borrowLimit" yaml:"borrowLimit"`
}

