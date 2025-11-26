package previewawsdlmmixins


// *[Custom snapshot policies only]* Specifies a rule for sharing snapshots across AWS accounts .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   shareRuleProperty := &ShareRuleProperty{
//   	TargetAccounts: []*string{
//   		jsii.String("targetAccounts"),
//   	},
//   	UnshareInterval: jsii.Number(123),
//   	UnshareIntervalUnit: jsii.String("unshareIntervalUnit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-sharerule.html
//
type CfnLifecyclePolicyPropsMixin_ShareRuleProperty struct {
	// The IDs of the AWS accounts with which to share the snapshots.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-sharerule.html#cfn-dlm-lifecyclepolicy-sharerule-targetaccounts
	//
	TargetAccounts *[]*string `field:"optional" json:"targetAccounts" yaml:"targetAccounts"`
	// The period after which snapshots that are shared with other AWS accounts are automatically unshared.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-sharerule.html#cfn-dlm-lifecyclepolicy-sharerule-unshareinterval
	//
	UnshareInterval *float64 `field:"optional" json:"unshareInterval" yaml:"unshareInterval"`
	// The unit of time for the automatic unsharing interval.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-sharerule.html#cfn-dlm-lifecyclepolicy-sharerule-unshareintervalunit
	//
	UnshareIntervalUnit *string `field:"optional" json:"unshareIntervalUnit" yaml:"unshareIntervalUnit"`
}

