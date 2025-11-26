package previewawswafv2mixins


// Used for CAPTCHA and challenge token settings.
//
// Determines how long a `CAPTCHA` or challenge timestamp remains valid after AWS WAF updates it for a successful `CAPTCHA` or challenge response.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   immunityTimePropertyProperty := &ImmunityTimePropertyProperty{
//   	ImmunityTime: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-immunitytimeproperty.html
//
type CfnRuleGroupPropsMixin_ImmunityTimePropertyProperty struct {
	// The amount of time, in seconds, that a `CAPTCHA` or challenge timestamp is considered valid by AWS WAF .
	//
	// The default setting is 300.
	//
	// For the Challenge action, the minimum setting is 300.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-immunitytimeproperty.html#cfn-wafv2-rulegroup-immunitytimeproperty-immunitytime
	//
	ImmunityTime *float64 `field:"optional" json:"immunityTime" yaml:"immunityTime"`
}

