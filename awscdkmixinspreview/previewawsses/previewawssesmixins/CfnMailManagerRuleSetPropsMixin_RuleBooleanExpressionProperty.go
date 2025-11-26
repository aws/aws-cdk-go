package previewawssesmixins


// A boolean expression to be used in a rule condition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ruleBooleanExpressionProperty := &RuleBooleanExpressionProperty{
//   	Evaluate: &RuleBooleanToEvaluateProperty{
//   		Analysis: &AnalysisProperty{
//   			Analyzer: jsii.String("analyzer"),
//   			ResultField: jsii.String("resultField"),
//   		},
//   		Attribute: jsii.String("attribute"),
//   		IsInAddressList: &RuleIsInAddressListProperty{
//   			AddressLists: []*string{
//   				jsii.String("addressLists"),
//   			},
//   			Attribute: jsii.String("attribute"),
//   		},
//   	},
//   	Operator: jsii.String("operator"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-rulebooleanexpression.html
//
type CfnMailManagerRuleSetPropsMixin_RuleBooleanExpressionProperty struct {
	// The operand on which to perform a boolean condition operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-rulebooleanexpression.html#cfn-ses-mailmanagerruleset-rulebooleanexpression-evaluate
	//
	Evaluate interface{} `field:"optional" json:"evaluate" yaml:"evaluate"`
	// The matching operator for a boolean condition expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-rulebooleanexpression.html#cfn-ses-mailmanagerruleset-rulebooleanexpression-operator
	//
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
}

