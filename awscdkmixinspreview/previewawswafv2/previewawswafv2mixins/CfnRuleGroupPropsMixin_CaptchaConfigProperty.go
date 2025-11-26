package previewawswafv2mixins


// Specifies how AWS WAF should handle `CAPTCHA` evaluations.
//
// This is available at the web ACL level and in each rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   captchaConfigProperty := &CaptchaConfigProperty{
//   	ImmunityTimeProperty: &ImmunityTimePropertyProperty{
//   		ImmunityTime: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-captchaconfig.html
//
type CfnRuleGroupPropsMixin_CaptchaConfigProperty struct {
	// Determines how long a `CAPTCHA` timestamp in the token remains valid after the client successfully solves a `CAPTCHA` puzzle.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-captchaconfig.html#cfn-wafv2-rulegroup-captchaconfig-immunitytimeproperty
	//
	ImmunityTimeProperty interface{} `field:"optional" json:"immunityTimeProperty" yaml:"immunityTimeProperty"`
}

