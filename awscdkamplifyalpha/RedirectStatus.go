package awscdkamplifyalpha


// The status code for a URL rewrite or redirect rule.
//
// Example:
//   var amplifyApp app
//
//   amplifyApp.AddCustomRule(map[string]interface{}{
//   	"source": jsii.String("/docs/specific-filename.html"),
//   	"target": jsii.String("/documents/different-filename.html"),
//   	"status": amplify.RedirectStatus_TEMPORARY_REDIRECT,
//   })
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

