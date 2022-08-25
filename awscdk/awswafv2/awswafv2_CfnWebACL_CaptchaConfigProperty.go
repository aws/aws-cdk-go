package awswafv2


// Specifies how AWS WAF should handle `CAPTCHA` evaluations for rules that don't have their own `CaptchaConfig` settings.
//
// If you don't specify this, AWS WAF uses its default settings for `CaptchaConfig` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   captchaConfigProperty := &captchaConfigProperty{
//   	immunityTimeProperty: &immunityTimePropertyProperty{
//   		immunityTime: jsii.Number(123),
//   	},
//   }
//
type CfnWebACL_CaptchaConfigProperty struct {
	// Determines how long a `CAPTCHA` token remains valid after the client successfully solves a `CAPTCHA` puzzle.
	ImmunityTimeProperty interface{} `field:"optional" json:"immunityTimeProperty" yaml:"immunityTimeProperty"`
}

