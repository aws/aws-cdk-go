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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-archiveretainrule.html
//
type CfnLifecyclePolicy_ArchiveRetainRuleProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-archiveretainrule.html#cfn-dlm-lifecyclepolicy-archiveretainrule-retentionarchivetier
	//
	RetentionArchiveTier interface{} `field:"required" json:"retentionArchiveTier" yaml:"retentionArchiveTier"`
}

