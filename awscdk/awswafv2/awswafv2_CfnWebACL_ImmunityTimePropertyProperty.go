package awswafv2


// Used for CAPTCHA and challenge token settings.
//
// Determines how long a `CAPTCHA` or challenge timestamp remains valid after AWS WAF updates it for a successful `CAPTCHA` or challenge response.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   immunityTimePropertyProperty := &immunityTimePropertyProperty{
//   	immunityTime: jsii.Number(123),
//   }
//
type CfnWebACL_ImmunityTimePropertyProperty struct {
	// The amount of time, in seconds, that a `CAPTCHA` or challenge timestamp is considered valid by AWS WAF .
	//
	// The default setting is 300.
	//
	// For the Challenge action, the minimum setting is 300.
	ImmunityTime *float64 `field:"required" json:"immunityTime" yaml:"immunityTime"`
}

