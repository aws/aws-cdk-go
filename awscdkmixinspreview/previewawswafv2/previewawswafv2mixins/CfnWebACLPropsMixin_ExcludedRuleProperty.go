package previewawswafv2mixins


// Specifies a single rule in a rule group whose action you want to override to `Count` .
//
// > Instead of this option, use `RuleActionOverrides` . It accepts any valid action setting, including `Count` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   excludedRuleProperty := &ExcludedRuleProperty{
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-excludedrule.html
//
type CfnWebACLPropsMixin_ExcludedRuleProperty struct {
	// The name of the rule whose action you want to override to `Count` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-excludedrule.html#cfn-wafv2-webacl-excludedrule-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

