package previewawsdlmmixins


// *[Custom snapshot and AMI policies only]* Specifies a cross-Region copy rule for a snapshot and AMI policies.
//
// > To specify a cross-Region copy action for event-based polices, use [CrossRegionCopyAction](https://docs.aws.amazon.com/dlm/latest/APIReference/API_CrossRegionCopyAction.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   crossRegionCopyRuleProperty := &CrossRegionCopyRuleProperty{
//   	CmkArn: jsii.String("cmkArn"),
//   	CopyTags: jsii.Boolean(false),
//   	DeprecateRule: &CrossRegionCopyDeprecateRuleProperty{
//   		Interval: jsii.Number(123),
//   		IntervalUnit: jsii.String("intervalUnit"),
//   	},
//   	Encrypted: jsii.Boolean(false),
//   	RetainRule: &CrossRegionCopyRetainRuleProperty{
//   		Interval: jsii.Number(123),
//   		IntervalUnit: jsii.String("intervalUnit"),
//   	},
//   	Target: jsii.String("target"),
//   	TargetRegion: jsii.String("targetRegion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-crossregioncopyrule.html
//
type CfnLifecyclePolicyPropsMixin_CrossRegionCopyRuleProperty struct {
	// The Amazon Resource Name (ARN) of the AWS KMS key to use for EBS encryption.
	//
	// If this parameter is not specified, the default KMS key for the account is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-crossregioncopyrule.html#cfn-dlm-lifecyclepolicy-crossregioncopyrule-cmkarn
	//
	CmkArn *string `field:"optional" json:"cmkArn" yaml:"cmkArn"`
	// Indicates whether to copy all user-defined tags from the source snapshot or AMI to the cross-Region copy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-crossregioncopyrule.html#cfn-dlm-lifecyclepolicy-crossregioncopyrule-copytags
	//
	CopyTags interface{} `field:"optional" json:"copyTags" yaml:"copyTags"`
	// *[Custom AMI policies only]* The AMI deprecation rule for cross-Region AMI copies created by the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-crossregioncopyrule.html#cfn-dlm-lifecyclepolicy-crossregioncopyrule-deprecaterule
	//
	DeprecateRule interface{} `field:"optional" json:"deprecateRule" yaml:"deprecateRule"`
	// To encrypt a copy of an unencrypted snapshot if encryption by default is not enabled, enable encryption using this parameter.
	//
	// Copies of encrypted snapshots are encrypted, even if this parameter is false or if encryption by default is not enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-crossregioncopyrule.html#cfn-dlm-lifecyclepolicy-crossregioncopyrule-encrypted
	//
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// The retention rule that indicates how long the cross-Region snapshot or AMI copies are to be retained in the destination Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-crossregioncopyrule.html#cfn-dlm-lifecyclepolicy-crossregioncopyrule-retainrule
	//
	RetainRule interface{} `field:"optional" json:"retainRule" yaml:"retainRule"`
	// > Use this parameter for snapshot policies only. For AMI policies, use *TargetRegion* instead.
	//
	// *[Custom snapshot policies only]* The target Region or the Amazon Resource Name (ARN) of the target Outpost for the snapshot copies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-crossregioncopyrule.html#cfn-dlm-lifecyclepolicy-crossregioncopyrule-target
	//
	Target *string `field:"optional" json:"target" yaml:"target"`
	// > Use this parameter for AMI policies only.
	//
	// For snapshot policies, use *Target* instead. For snapshot policies created before the *Target* parameter was introduced, this parameter indicates the target Region for snapshot copies.
	//
	// *[Custom AMI policies only]* The target Region or the Amazon Resource Name (ARN) of the target Outpost for the snapshot copies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-crossregioncopyrule.html#cfn-dlm-lifecyclepolicy-crossregioncopyrule-targetregion
	//
	TargetRegion *string `field:"optional" json:"targetRegion" yaml:"targetRegion"`
}

