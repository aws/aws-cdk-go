package previewawsdlmmixins


// *[Custom snapshot policies only]* Describes the retention rule for archived snapshots.
//
// Once the archive retention threshold is met, the snapshots are permanently deleted from the archive tier.
//
// > The archive retention rule must retain snapshots in the archive tier for a minimum of 90 days.
//
// For *count-based schedules* , you must specify *Count* . For *age-based schedules* , you must specify *Interval* and *IntervalUnit* .
//
// For more information about using snapshot archiving, see [Considerations for snapshot lifecycle policies](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/snapshot-ami-policy.html#dlm-archive) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   retentionArchiveTierProperty := &RetentionArchiveTierProperty{
//   	Count: jsii.Number(123),
//   	Interval: jsii.Number(123),
//   	IntervalUnit: jsii.String("intervalUnit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-retentionarchivetier.html
//
type CfnLifecyclePolicyPropsMixin_RetentionArchiveTierProperty struct {
	// The maximum number of snapshots to retain in the archive storage tier for each volume.
	//
	// The count must ensure that each snapshot remains in the archive tier for at least 90 days. For example, if the schedule creates snapshots every 30 days, you must specify a count of 3 or more to ensure that each snapshot is archived for at least 90 days.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-retentionarchivetier.html#cfn-dlm-lifecyclepolicy-retentionarchivetier-count
	//
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Specifies the period of time to retain snapshots in the archive tier.
	//
	// After this period expires, the snapshot is permanently deleted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-retentionarchivetier.html#cfn-dlm-lifecyclepolicy-retentionarchivetier-interval
	//
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The unit of time in which to measure the *Interval* .
	//
	// For example, to retain a snapshots in the archive tier for 6 months, specify `Interval=6` and `IntervalUnit=MONTHS` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-retentionarchivetier.html#cfn-dlm-lifecyclepolicy-retentionarchivetier-intervalunit
	//
	IntervalUnit *string `field:"optional" json:"intervalUnit" yaml:"intervalUnit"`
}

