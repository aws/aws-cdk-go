package mixinsawsdlm


// *[Custom snapshot policies only]* Specifies a snapshot archiving rule for a schedule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   archiveRuleProperty := &ArchiveRuleProperty{
//   	RetainRule: &ArchiveRetainRuleProperty{
//   		RetentionArchiveTier: &RetentionArchiveTierProperty{
//   			Count: jsii.Number(123),
//   			Interval: jsii.Number(123),
//   			IntervalUnit: jsii.String("intervalUnit"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-archiverule.html
//
type CfnLifecyclePolicyPropsMixin_ArchiveRuleProperty struct {
	// Information about the retention period for the snapshot archiving rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-archiverule.html#cfn-dlm-lifecyclepolicy-archiverule-retainrule
	//
	RetainRule interface{} `field:"optional" json:"retainRule" yaml:"retainRule"`
}

