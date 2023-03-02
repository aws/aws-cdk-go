package awsamplify


// The CustomRule property type allows you to specify redirects, rewrites, and reverse proxies.
//
// Redirects enable a web app to reroute navigation from one URL to another.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customRuleProperty := &customRuleProperty{
//   	source: jsii.String("source"),
//   	target: jsii.String("target"),
//
//   	// the properties below are optional
//   	condition: jsii.String("condition"),
//   	status: jsii.String("status"),
//   }
//
type CfnApp_CustomRuleProperty struct {
	// The source pattern for a URL rewrite or redirect rule.
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 2048.
	//
	// *Pattern:* (?s).+
	Source *string `field:"required" json:"source" yaml:"source"`
	// The target pattern for a URL rewrite or redirect rule.
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 2048.
	//
	// *Pattern:* (?s).+
	Target *string `field:"required" json:"target" yaml:"target"`
	// The condition for a URL rewrite or redirect rule, such as a country code.
	//
	// *Length Constraints:* Minimum length of 0. Maximum length of 2048.
	//
	// *Pattern:* (?s).*
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// The status code for a URL rewrite or redirect rule.
	//
	// - **200** - Represents a 200 rewrite rule.
	// - **301** - Represents a 301 (moved pemanently) redirect rule. This and all future requests should be directed to the target URL.
	// - **302** - Represents a 302 temporary redirect rule.
	// - **404** - Represents a 404 redirect rule.
	// - **404-200** - Represents a 404 rewrite rule.
	//
	// *Length Constraints:* Minimum length of 3. Maximum length of 7.
	//
	// *Pattern:* .{3,7}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

