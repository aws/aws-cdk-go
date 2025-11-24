package mixinsawsbedrock


// A rule within the policy definition that defines logical constraints.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   policyDefinitionRuleProperty := &PolicyDefinitionRuleProperty{
//   	AlternateExpression: jsii.String("alternateExpression"),
//   	Expression: jsii.String("expression"),
//   	Id: jsii.String("id"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-automatedreasoningpolicy-policydefinitionrule.html
//
type CfnAutomatedReasoningPolicyPropsMixin_PolicyDefinitionRuleProperty struct {
	// An alternative expression for the policy rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-automatedreasoningpolicy-policydefinitionrule.html#cfn-bedrock-automatedreasoningpolicy-policydefinitionrule-alternateexpression
	//
	AlternateExpression *string `field:"optional" json:"alternateExpression" yaml:"alternateExpression"`
	// The logical expression that defines the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-automatedreasoningpolicy-policydefinitionrule.html#cfn-bedrock-automatedreasoningpolicy-policydefinitionrule-expression
	//
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// The unique identifier for the policy definition rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-automatedreasoningpolicy-policydefinitionrule.html#cfn-bedrock-automatedreasoningpolicy-policydefinitionrule-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
}

