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
	// The resource sharing strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-quotashare-quotashareresourcesharingconfiguration.html#cfn-batch-quotashare-quotashareresourcesharingconfiguration-strategy
	//
	Strategy *string `field:"required" json:"strategy" yaml:"strategy"`
	// The maximum amount of compute capacity that can be borrowed.
	//
	// Use -1 for unlimited borrowing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-quotashare-quotashareresourcesharingconfiguration.html#cfn-batch-quotashare-quotashareresourcesharingconfiguration-borrowlimit
	//
	BorrowLimit *float64 `field:"optional" json:"borrowLimit" yaml:"borrowLimit"`
}

