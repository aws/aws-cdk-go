package awsnetworksecuritymanager


// Associates either a template or a rule with a policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   associatedTemplateAndRuleProperty := &AssociatedTemplateAndRuleProperty{
//   	RuleArn: jsii.String("ruleArn"),
//   	TemplateArn: jsii.String("templateArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networksecuritymanager-policy-associatedtemplateandrule.html
//
type CfnPolicy_AssociatedTemplateAndRuleProperty struct {
	// ARN of the associated rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networksecuritymanager-policy-associatedtemplateandrule.html#cfn-networksecuritymanager-policy-associatedtemplateandrule-rulearn
	//
	RuleArn *string `field:"optional" json:"ruleArn" yaml:"ruleArn"`
	// ARN of the associated template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networksecuritymanager-policy-associatedtemplateandrule.html#cfn-networksecuritymanager-policy-associatedtemplateandrule-templatearn
	//
	TemplateArn *string `field:"optional" json:"templateArn" yaml:"templateArn"`
}

