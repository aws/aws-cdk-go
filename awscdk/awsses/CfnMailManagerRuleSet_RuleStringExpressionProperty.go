package awsses


// A string expression is evaluated against strings or substrings of the email.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleStringExpressionProperty := &RuleStringExpressionProperty{
//   	Evaluate: &RuleStringToEvaluateProperty{
//   		Attribute: jsii.String("attribute"),
//   		MimeHeaderAttribute: jsii.String("mimeHeaderAttribute"),
//   	},
//   	Operator: jsii.String("operator"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-rulestringexpression.html
//
type CfnMailManagerRuleSet_RuleStringExpressionProperty struct {
	// The string to evaluate in a string condition expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-rulestringexpression.html#cfn-ses-mailmanagerruleset-rulestringexpression-evaluate
	//
	Evaluate interface{} `field:"required" json:"evaluate" yaml:"evaluate"`
	// The matching operator for a string condition expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-rulestringexpression.html#cfn-ses-mailmanagerruleset-rulestringexpression-operator
	//
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// The string(s) to be evaluated in a string condition expression.
	//
	// For all operators, except for NOT_EQUALS, if multiple values are given, the values are processed as an OR. That is, if any of the values match the email's string using the given operator, the condition is deemed to match. However, for NOT_EQUALS, the condition is only deemed to match if none of the given strings match the email's string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-rulestringexpression.html#cfn-ses-mailmanagerruleset-rulestringexpression-values
	//
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

