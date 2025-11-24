package mixinsawsses


// The conditional expression used to evaluate an email for determining if a rule action should be taken.
//
// > This data type is a UNION, so only one of the following members can be specified when used or returned.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ruleConditionProperty := &RuleConditionProperty{
//   	BooleanExpression: &RuleBooleanExpressionProperty{
//   		Evaluate: &RuleBooleanToEvaluateProperty{
//   			Analysis: &AnalysisProperty{
//   				Analyzer: jsii.String("analyzer"),
//   				ResultField: jsii.String("resultField"),
//   			},
//   			Attribute: jsii.String("attribute"),
//   			IsInAddressList: &RuleIsInAddressListProperty{
//   				AddressLists: []*string{
//   					jsii.String("addressLists"),
//   				},
//   				Attribute: jsii.String("attribute"),
//   			},
//   		},
//   		Operator: jsii.String("operator"),
//   	},
//   	DmarcExpression: &RuleDmarcExpressionProperty{
//   		Operator: jsii.String("operator"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	IpExpression: &RuleIpExpressionProperty{
//   		Evaluate: &RuleIpToEvaluateProperty{
//   			Attribute: jsii.String("attribute"),
//   		},
//   		Operator: jsii.String("operator"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	NumberExpression: &RuleNumberExpressionProperty{
//   		Evaluate: &RuleNumberToEvaluateProperty{
//   			Attribute: jsii.String("attribute"),
//   		},
//   		Operator: jsii.String("operator"),
//   		Value: jsii.Number(123),
//   	},
//   	StringExpression: &RuleStringExpressionProperty{
//   		Evaluate: &RuleStringToEvaluateProperty{
//   			Analysis: &AnalysisProperty{
//   				Analyzer: jsii.String("analyzer"),
//   				ResultField: jsii.String("resultField"),
//   			},
//   			Attribute: jsii.String("attribute"),
//   			MimeHeaderAttribute: jsii.String("mimeHeaderAttribute"),
//   		},
//   		Operator: jsii.String("operator"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	VerdictExpression: &RuleVerdictExpressionProperty{
//   		Evaluate: &RuleVerdictToEvaluateProperty{
//   			Analysis: &AnalysisProperty{
//   				Analyzer: jsii.String("analyzer"),
//   				ResultField: jsii.String("resultField"),
//   			},
//   			Attribute: jsii.String("attribute"),
//   		},
//   		Operator: jsii.String("operator"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-rulecondition.html
//
type CfnMailManagerRuleSetPropsMixin_RuleConditionProperty struct {
	// The condition applies to a boolean expression passed in this field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-rulecondition.html#cfn-ses-mailmanagerruleset-rulecondition-booleanexpression
	//
	BooleanExpression interface{} `field:"optional" json:"booleanExpression" yaml:"booleanExpression"`
	// The condition applies to a DMARC policy expression passed in this field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-rulecondition.html#cfn-ses-mailmanagerruleset-rulecondition-dmarcexpression
	//
	DmarcExpression interface{} `field:"optional" json:"dmarcExpression" yaml:"dmarcExpression"`
	// The condition applies to an IP address expression passed in this field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-rulecondition.html#cfn-ses-mailmanagerruleset-rulecondition-ipexpression
	//
	IpExpression interface{} `field:"optional" json:"ipExpression" yaml:"ipExpression"`
	// The condition applies to a number expression passed in this field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-rulecondition.html#cfn-ses-mailmanagerruleset-rulecondition-numberexpression
	//
	NumberExpression interface{} `field:"optional" json:"numberExpression" yaml:"numberExpression"`
	// The condition applies to a string expression passed in this field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-rulecondition.html#cfn-ses-mailmanagerruleset-rulecondition-stringexpression
	//
	StringExpression interface{} `field:"optional" json:"stringExpression" yaml:"stringExpression"`
	// The condition applies to a verdict expression passed in this field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-rulecondition.html#cfn-ses-mailmanagerruleset-rulecondition-verdictexpression
	//
	VerdictExpression interface{} `field:"optional" json:"verdictExpression" yaml:"verdictExpression"`
}

