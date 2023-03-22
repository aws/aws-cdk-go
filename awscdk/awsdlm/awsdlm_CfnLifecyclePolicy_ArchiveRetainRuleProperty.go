package awsdlm


// *[Snapshot policies only]* Specifies information about the archive storage tier retention period.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   archiveRetainRuleProperty := &archiveRetainRuleProperty{
//   	retentionArchiveTier: &retentionArchiveTierProperty{
//   		count: jsii.Number(123),
//   		interval: jsii.Number(123),
//   		intervalUnit: jsii.String("intervalUnit"),
//   	},
//   }
//
type CfnLifecyclePolicy_ArchiveRetainRuleProperty struct {
	// Information about retention period in the Amazon EBS Snapshots Archive.
	//
	// For more information, see [Archive Amazon EBS snapshots](https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/snapshot-archive.html) .
	RetentionArchiveTier interface{} `field:"required" json:"retentionArchiveTier" yaml:"retentionArchiveTier"`
}

