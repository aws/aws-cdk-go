package mixinsawsses


// Properties for CfnReceiptRuleSetPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnReceiptRuleSetMixinProps := &CfnReceiptRuleSetMixinProps{
//   	RuleSetName: jsii.String("ruleSetName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-receiptruleset.html
//
type CfnReceiptRuleSetMixinProps struct {
	// The name of the receipt rule set to make active.
	//
	// Setting this value to null disables all email receiving.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-receiptruleset.html#cfn-ses-receiptruleset-rulesetname
	//
	RuleSetName *string `field:"optional" json:"ruleSetName" yaml:"ruleSetName"`
}

