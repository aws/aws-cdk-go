package mixinsawsroute53recoverycontrol


// The rule configuration for an assertion rule.
//
// That is, the criteria that you set for specific assertion controls (routing controls) that specify how many controls must be enabled after a transaction completes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ruleConfigProperty := &RuleConfigProperty{
//   	Inverted: jsii.Boolean(false),
//   	Threshold: jsii.Number(123),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoverycontrol-safetyrule-ruleconfig.html
//
type CfnSafetyRulePropsMixin_RuleConfigProperty struct {
	// Logical negation of the rule.
	//
	// If the rule would usually evaluate true, it's evaluated as false, and vice versa.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoverycontrol-safetyrule-ruleconfig.html#cfn-route53recoverycontrol-safetyrule-ruleconfig-inverted
	//
	Inverted interface{} `field:"optional" json:"inverted" yaml:"inverted"`
	// The value of N, when you specify an `ATLEAST` rule type.
	//
	// That is, `Threshold` is the number of controls that must be set when you specify an `ATLEAST` type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoverycontrol-safetyrule-ruleconfig.html#cfn-route53recoverycontrol-safetyrule-ruleconfig-threshold
	//
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
	// A rule can be one of the following: `ATLEAST` , `AND` , or `OR` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoverycontrol-safetyrule-ruleconfig.html#cfn-route53recoverycontrol-safetyrule-ruleconfig-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

