package awsbedrockagentcorealpha


// Browser signing.
//
// Specifies whether browser signing is enabled.
// When enabled, the browser will cryptographically sign HTTP requests to identify
// itself as an AI agent to bot control vendors.
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type BrowserSigning string

const (
	// Browser signing is enabled.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	BrowserSigning_ENABLED BrowserSigning = "ENABLED"
	// Browser signing is disabled.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	BrowserSigning_DISABLED BrowserSigning = "DISABLED"
)

