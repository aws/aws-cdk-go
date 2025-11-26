package previewawssesmixins


// A verdict expression is evaluated against verdicts of the email.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ruleVerdictExpressionProperty := &RuleVerdictExpressionProperty{
//   	Evaluate: &RuleVerdictToEvaluateProperty{
//   		Analysis: &AnalysisProperty{
//   			Analyzer: jsii.String("analyzer"),
//   			ResultField: jsii.String("resultField"),
//   		},
//   		Attribute: jsii.String("attribute"),
//   	},
//   	Operator: jsii.String("operator"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-ruleverdictexpression.html
//
type CfnMailManagerRuleSetPropsMixin_RuleVerdictExpressionProperty struct {
	// The verdict to evaluate in a verdict condition expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-ruleverdictexpression.html#cfn-ses-mailmanagerruleset-ruleverdictexpression-evaluate
	//
	Evaluate interface{} `field:"optional" json:"evaluate" yaml:"evaluate"`
	// The matching operator for a verdict condition expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-ruleverdictexpression.html#cfn-ses-mailmanagerruleset-ruleverdictexpression-operator
	//
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// The values to match with the email's verdict using the given operator.
	//
	// For the EQUALS operator, if multiple values are given, the condition is deemed to match if any of the given verdicts match that of the email. For the NOT_EQUALS operator, if multiple values are given, the condition is deemed to match of none of the given verdicts match the verdict of the email.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-ruleverdictexpression.html#cfn-ses-mailmanagerruleset-ruleverdictexpression-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

