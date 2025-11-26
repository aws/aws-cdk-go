package previewawsamplifymixins


// The CustomRule property type allows you to specify redirects, rewrites, and reverse proxies.
//
// Redirects enable a web app to reroute navigation from one URL to another.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customRuleProperty := &CustomRuleProperty{
//   	Condition: jsii.String("condition"),
//   	Source: jsii.String("source"),
//   	Status: jsii.String("status"),
//   	Target: jsii.String("target"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-app-customrule.html
//
type CfnAppPropsMixin_CustomRuleProperty struct {
	// The condition for a URL rewrite or redirect rule, such as a country code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-app-customrule.html#cfn-amplify-app-customrule-condition
	//
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// The source pattern for a URL rewrite or redirect rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-app-customrule.html#cfn-amplify-app-customrule-source
	//
	Source *string `field:"optional" json:"source" yaml:"source"`
	// The status code for a URL rewrite or redirect rule.
	//
	// - **200** - Represents a 200 rewrite rule.
	// - **301** - Represents a 301 (moved pemanently) redirect rule. This and all future requests should be directed to the target URL.
	// - **302** - Represents a 302 temporary redirect rule.
	// - **404** - Represents a 404 redirect rule.
	// - **404-200** - Represents a 404 rewrite rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-app-customrule.html#cfn-amplify-app-customrule-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The target pattern for a URL rewrite or redirect rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-app-customrule.html#cfn-amplify-app-customrule-target
	//
	Target *string `field:"optional" json:"target" yaml:"target"`
}

