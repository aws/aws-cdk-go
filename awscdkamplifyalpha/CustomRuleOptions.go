package awscdkamplifyalpha


// Options for a custom rewrite/redirect rule for an Amplify App.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import amplify_alpha "github.com/aws/aws-cdk-go/awscdkamplifyalpha"
//
//   customRuleOptions := &CustomRuleOptions{
//   	Source: jsii.String("source"),
//   	Target: jsii.String("target"),
//
//   	// the properties below are optional
//   	Condition: jsii.String("condition"),
//   	Status: amplify_alpha.RedirectStatus_REWRITE,
//   }
//
// Experimental.
type CustomRuleOptions struct {
	// The source pattern for a URL rewrite or redirect rule.
	// See: https://docs.aws.amazon.com/amplify/latest/userguide/redirects.html
	//
	// Experimental.
	Source *string `field:"required" json:"source" yaml:"source"`
	// The target pattern for a URL rewrite or redirect rule.
	// See: https://docs.aws.amazon.com/amplify/latest/userguide/redirects.html
	//
	// Experimental.
	Target *string `field:"required" json:"target" yaml:"target"`
	// The condition for a URL rewrite or redirect rule, e.g. country code.
	// See: https://docs.aws.amazon.com/amplify/latest/userguide/redirects.html
	//
	// Experimental.
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// The status code for a URL rewrite or redirect rule.
	// See: https://docs.aws.amazon.com/amplify/latest/userguide/redirects.html
	//
	// Experimental.
	Status RedirectStatus `field:"optional" json:"status" yaml:"status"`
}

