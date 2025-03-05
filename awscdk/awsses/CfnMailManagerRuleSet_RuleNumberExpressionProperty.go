package awsses


// A number expression to match numeric conditions with integers from the incoming email.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleNumberExpressionProperty := &RuleNumberExpressionProperty{
//   	Evaluate: &RuleNumberToEvaluateProperty{
//   		Attribute: jsii.String("attribute"),
//   	},
//   	Operator: jsii.String("operator"),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-rulenumberexpression.html
//
type CfnMailManagerRuleSet_RuleNumberExpressionProperty struct {
	// The number to evaluate in a numeric condition expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-rulenumberexpression.html#cfn-ses-mailmanagerruleset-rulenumberexpression-evaluate
	//
	Evaluate interface{} `field:"required" json:"evaluate" yaml:"evaluate"`
	// The operator for a numeric condition expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-rulenumberexpression.html#cfn-ses-mailmanagerruleset-rulenumberexpression-operator
	//
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// The value to evaluate in a numeric condition expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-rulenumberexpression.html#cfn-ses-mailmanagerruleset-rulenumberexpression-value
	//
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

