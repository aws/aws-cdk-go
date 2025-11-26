package previewawssesmixins


// The number to evaluate in a numeric condition expression.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ruleNumberToEvaluateProperty := &RuleNumberToEvaluateProperty{
//   	Attribute: jsii.String("attribute"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-rulenumbertoevaluate.html
//
type CfnMailManagerRuleSetPropsMixin_RuleNumberToEvaluateProperty struct {
	// An email attribute that is used as the number to evaluate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-rulenumbertoevaluate.html#cfn-ses-mailmanagerruleset-rulenumbertoevaluate-attribute
	//
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
}

