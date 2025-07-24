package awsses


// The IP address to evaluate for this condition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleIpToEvaluateProperty := &RuleIpToEvaluateProperty{
//   	Attribute: jsii.String("attribute"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-ruleiptoevaluate.html
//
type CfnMailManagerRuleSet_RuleIpToEvaluateProperty struct {
	// The attribute of the email to evaluate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-ruleiptoevaluate.html#cfn-ses-mailmanagerruleset-ruleiptoevaluate-attribute
	//
	Attribute *string `field:"required" json:"attribute" yaml:"attribute"`
}

