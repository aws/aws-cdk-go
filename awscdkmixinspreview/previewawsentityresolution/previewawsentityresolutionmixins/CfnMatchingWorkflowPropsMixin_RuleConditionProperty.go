package previewawsentityresolutionmixins


// An object that defines the `ruleCondition` and the `ruleName` to use in a matching workflow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ruleConditionProperty := &RuleConditionProperty{
//   	Condition: jsii.String("condition"),
//   	RuleName: jsii.String("ruleName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-rulecondition.html
//
type CfnMatchingWorkflowPropsMixin_RuleConditionProperty struct {
	// A statement that specifies the conditions for a matching rule.
	//
	// If your data is accurate, use an Exact matching function: `Exact` or `ExactManyToMany` .
	//
	// If your data has variations in spelling or pronunciation, use a Fuzzy matching function: `Cosine` , `Levenshtein` , or `Soundex` .
	//
	// Use operators if you want to combine ( `AND` ), separate ( `OR` ), or group matching functions `(...)` .
	//
	// For example: `(Cosine(a, 10) AND Exact(b, true)) OR ExactManyToMany(c, d)`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-rulecondition.html#cfn-entityresolution-matchingworkflow-rulecondition-condition
	//
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// A name for the matching rule.
	//
	// For example: `Rule1`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-matchingworkflow-rulecondition.html#cfn-entityresolution-matchingworkflow-rulecondition-rulename
	//
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
}

