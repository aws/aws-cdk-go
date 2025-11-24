package mixinsawsdlm


// *[Custom snapshot policies only]* Specifies information about the archive storage tier retention period.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnLifecyclePolicyPropsMixin_ArchiveRetainRuleProperty struct {
	// Information about retention period in the Amazon EBS Snapshots Archive.
	//
	// For more information, see [Archive Amazon EBS snapshots](https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/snapshot-archive.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-archiveretainrule.html#cfn-dlm-lifecyclepolicy-archiveretainrule-retentionarchivetier
	//
	RetentionArchiveTier interface{} `field:"optional" json:"retentionArchiveTier" yaml:"retentionArchiveTier"`
}

