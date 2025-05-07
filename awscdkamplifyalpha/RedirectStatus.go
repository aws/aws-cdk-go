package awscdkamplifyalpha


// The status code for a URL rewrite or redirect rule.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkamplifyalpha"
//
//   var amplifyApp app
//
//   amplifyApp.AddCustomRule(awscdkamplifyalpha.NewCustomRule(&CustomRuleOptions{
//   	Source: jsii.String("/docs/specific-filename.html"),
//   	Target: jsii.String("/documents/different-filename.html"),
//   	Status: amplify.RedirectStatus_TEMPORARY_REDIRECT,
//   }))
//
// Experimental.
type RedirectStatus string

const (
	// Rewrite (200).
	// Experimental.
	RedirectStatus_REWRITE RedirectStatus = "REWRITE"
	// Permanent redirect (301).
	// Experimental.
	RedirectStatus_PERMANENT_REDIRECT RedirectStatus = "PERMANENT_REDIRECT"
	// Temporary redirect (302).
	// Experimental.
	RedirectStatus_TEMPORARY_REDIRECT RedirectStatus = "TEMPORARY_REDIRECT"
	// Not found (404).
	// Experimental.
	RedirectStatus_NOT_FOUND RedirectStatus = "NOT_FOUND"
	// Not found rewrite (404).
	// Experimental.
	RedirectStatus_NOT_FOUND_REWRITE RedirectStatus = "NOT_FOUND_REWRITE"
)

