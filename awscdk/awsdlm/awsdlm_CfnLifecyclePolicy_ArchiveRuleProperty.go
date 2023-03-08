package awsdlm


// *[Snapshot policies only]* Specifies a snapshot archiving rule for a schedule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   archiveRuleProperty := &archiveRuleProperty{
//   	retainRule: &archiveRetainRuleProperty{
//   		retentionArchiveTier: &retentionArchiveTierProperty{
//   			count: jsii.Number(123),
//   			interval: jsii.Number(123),
//   			intervalUnit: jsii.String("intervalUnit"),
//   		},
//   	},
//   }
//
type CfnLifecyclePolicy_ArchiveRuleProperty struct {
	// Information about the retention period for the snapshot archiving rule.
	RetainRule interface{} `field:"required" json:"retainRule" yaml:"retainRule"`
}

