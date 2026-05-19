package awsbedrockagentcore


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
type BrowserSigning string

const (
	// Browser signing is enabled.
	BrowserSigning_ENABLED BrowserSigning = "ENABLED"
	// Browser signing is disabled.
	BrowserSigning_DISABLED BrowserSigning = "DISABLED"
)

