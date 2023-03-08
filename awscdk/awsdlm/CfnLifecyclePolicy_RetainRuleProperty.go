package awsdlm


// *[Snapshot and AMI policies only]* Specifies a retention rule for snapshots created by snapshot policies, or for AMIs created by AMI policies.
//
// > For snapshot policies that have an `ArchiveRule` , this retention rule applies to standard tier retention. When the retention threshold is met, snapshots are moved from the standard to the archive tier.
// >
// > For snapshot policies that do not have an *ArchiveRule* , snapshots are permanently deleted when this retention threshold is met.
//
// You can retain snapshots based on either a count or a time interval.
//
// - *Count-based retention*
//
// You must specify *Count* . If you specify an `ArchiveRule` for the schedule, then you can specify a retention count of `0` to archive snapshots immediately after creation. If you specify a `FastRestoreRule` , `ShareRule` , or a `CrossRegionCopyRule` , then you must specify a retention count of `1` or more.
// - *Age-based retention*
//
// You must specify *Interval* and *IntervalUnit* . If you specify an `ArchiveRule` for the schedule, then you can specify a retention interval of `0` days to archive snapshots immediately after creation. If you specify a `FastRestoreRule` , `ShareRule` , or a `CrossRegionCopyRule` , then you must specify a retention interval of `1` day or more.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   retainRuleProperty := &RetainRuleProperty{
//   	Count: jsii.Number(123),
//   	Interval: jsii.Number(123),
//   	IntervalUnit: jsii.String("intervalUnit"),
//   }
//
type CfnLifecyclePolicy_RetainRuleProperty struct {
	// The number of snapshots to retain for each volume, up to a maximum of 1000.
	//
	// For example if you want to retain a maximum of three snapshots, specify `3` . When the fourth snapshot is created, the oldest retained snapshot is deleted, or it is moved to the archive tier if you have specified an `ArchiveRule` .
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// The amount of time to retain each snapshot.
	//
	// The maximum is 100 years. This is equivalent to 1200 months, 5200 weeks, or 36500 days.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The unit of time for time-based retention.
	//
	// For example, to retain snapshots for 3 months, specify `Interval=3` and `IntervalUnit=MONTHS` . Once the snapshot has been retained for 3 months, it is deleted, or it is moved to the archive tier if you have specified an `ArchiveRule` .
	IntervalUnit *string `field:"optional" json:"intervalUnit" yaml:"intervalUnit"`
}

