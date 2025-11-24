package awsbedrockagentcorealpha


// Browser signing.
//
// Specifies whether browser signing is enabled.
// When enabled, the browser will cryptographically sign HTTP requests to identify
// itself as an AI agent to bot control vendors.
//
// Example:
//   browser := agentcore.NewBrowserCustom(this, jsii.String("test-browser"), &BrowserCustomProps{
//   	BrowserCustomName: jsii.String("test_browser"),
//   	BrowserSigning: agentcore.BrowserSigning_ENABLED,
//   })
//
// Experimental.
type BrowserSigning string

const (
	// Browser signing is enabled.
	// Experimental.
	BrowserSigning_ENABLED BrowserSigning = "ENABLED"
	// Browser signing is disabled.
	// Experimental.
	BrowserSigning_DISABLED BrowserSigning = "DISABLED"
)

