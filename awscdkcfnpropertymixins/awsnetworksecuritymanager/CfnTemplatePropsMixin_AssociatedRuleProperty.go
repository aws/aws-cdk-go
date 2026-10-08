package awsnetworksecuritymanager


// An association between a template and a rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   associatedRuleProperty := &AssociatedRuleProperty{
//   	RuleArn: jsii.String("ruleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networksecuritymanager-template-associatedrule.html
//
type CfnTemplatePropsMixin_AssociatedRuleProperty struct {
	// ARN of the associated rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networksecuritymanager-template-associatedrule.html#cfn-networksecuritymanager-template-associatedrule-rulearn
	//
	RuleArn *string `field:"optional" json:"ruleArn" yaml:"ruleArn"`
}

