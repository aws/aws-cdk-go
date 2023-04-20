package awsdlm


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   archiveRetainRuleProperty := &ArchiveRetainRuleProperty{
//   	RetentionArchiveTier: &RetentionArchiveTierProperty{
//   		Count: jsii.Number(123),
//   		Interval: jsii.Number(123),
//   		IntervalUnit: jsii.String("intervalUnit"),
//   	},
//   }
//
type CfnLifecyclePolicy_ArchiveRetainRuleProperty struct {
	// `CfnLifecyclePolicy.ArchiveRetainRuleProperty.RetentionArchiveTier`.
	RetentionArchiveTier interface{} `field:"required" json:"retentionArchiveTier" yaml:"retentionArchiveTier"`
}

