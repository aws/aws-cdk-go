package awsdlm


// Specifies a rule for cross-Region snapshot copies.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   crossRegionCopyRuleProperty := &crossRegionCopyRuleProperty{
//   	encrypted: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	cmkArn: jsii.String("cmkArn"),
//   	copyTags: jsii.Boolean(false),
//   	deprecateRule: &crossRegionCopyDeprecateRuleProperty{
//   		interval: jsii.Number(123),
//   		intervalUnit: jsii.String("intervalUnit"),
//   	},
//   	retainRule: &crossRegionCopyRetainRuleProperty{
//   		interval: jsii.Number(123),
//   		intervalUnit: jsii.String("intervalUnit"),
//   	},
//   	target: jsii.String("target"),
//   	targetRegion: jsii.String("targetRegion"),
//   }
//
type CfnLifecyclePolicy_CrossRegionCopyRuleProperty struct {
	// To encrypt a copy of an unencrypted snapshot if encryption by default is not enabled, enable encryption using this parameter.
	//
	// Copies of encrypted snapshots are encrypted, even if this parameter is false or if encryption by default is not enabled.
	Encrypted interface{} `field:"required" json:"encrypted" yaml:"encrypted"`
	// The Amazon Resource Name (ARN) of the AWS KMS key to use for EBS encryption.
	//
	// If this parameter is not specified, the default KMS key for the account is used.
	CmkArn *string `field:"optional" json:"cmkArn" yaml:"cmkArn"`
	// Indicates whether to copy all user-defined tags from the source snapshot to the cross-Region snapshot copy.
	CopyTags interface{} `field:"optional" json:"copyTags" yaml:"copyTags"`
	// The AMI deprecation rule for cross-Region AMI copies created by the rule.
	DeprecateRule interface{} `field:"optional" json:"deprecateRule" yaml:"deprecateRule"`
	// The retention rule that indicates how long snapshot copies are to be retained in the destination Region.
	RetainRule interface{} `field:"optional" json:"retainRule" yaml:"retainRule"`
	// The target Region or the Amazon Resource Name (ARN) of the target Outpost for the snapshot copies.
	//
	// Use this parameter instead of *TargetRegion* . Do not specify both.
	Target *string `field:"optional" json:"target" yaml:"target"`
	// Avoid using this parameter when creating new policies.
	//
	// Instead, use *Target* to specify a target Region or a target Outpost for snapshot copies.
	//
	// For policies created before the *Target* parameter was introduced, this parameter indicates the target Region for snapshot copies.
	TargetRegion *string `field:"optional" json:"targetRegion" yaml:"targetRegion"`
}

