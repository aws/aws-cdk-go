package mixinsawsses


// The verdict to evaluate in a verdict condition expression.
//
// > This data type is a UNION, so only one of the following members can be specified when used or returned.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ruleVerdictToEvaluateProperty := &RuleVerdictToEvaluateProperty{
//   	Analysis: &AnalysisProperty{
//   		Analyzer: jsii.String("analyzer"),
//   		ResultField: jsii.String("resultField"),
//   	},
//   	Attribute: jsii.String("attribute"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-ruleverdicttoevaluate.html
//
type CfnMailManagerRuleSetPropsMixin_RuleVerdictToEvaluateProperty struct {
	// The Add On ARN and its returned value to evaluate in a verdict condition expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-ruleverdicttoevaluate.html#cfn-ses-mailmanagerruleset-ruleverdicttoevaluate-analysis
	//
	Analysis interface{} `field:"optional" json:"analysis" yaml:"analysis"`
	// The email verdict attribute to evaluate in a string verdict expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-ruleverdicttoevaluate.html#cfn-ses-mailmanagerruleset-ruleverdicttoevaluate-attribute
	//
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
}

