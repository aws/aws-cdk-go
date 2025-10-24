package awscdkamplifyalpha


// Options for a custom rewrite/redirect rule for an Amplify App.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkamplifyalpha"
//
//   var amplifyApp App
//
//   amplifyApp.AddCustomRule(awscdkamplifyalpha.NewCustomRule(&CustomRuleOptions{
//   	Source: jsii.String("/docs/specific-filename.html"),
//   	Target: jsii.String("/documents/different-filename.html"),
//   	Status: amplify.RedirectStatus_TEMPORARY_REDIRECT,
//   }))
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
	// Default: - no condition.
	//
	// Experimental.
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// The status code for a URL rewrite or redirect rule.
	// See: https://docs.aws.amazon.com/amplify/latest/userguide/redirects.html
	//
	// Default: PERMANENT_REDIRECT.
	//
	// Experimental.
	Status RedirectStatus `field:"optional" json:"status" yaml:"status"`
}

