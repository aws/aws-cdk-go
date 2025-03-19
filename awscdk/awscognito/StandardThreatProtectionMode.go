package awscognito


// The Type of Threat Protection Enabled for Standard Authentication.
//
// This feature only functions if your FeaturePlan is set to FeaturePlan.PLUS
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-userpooladdons.html
//
type StandardThreatProtectionMode string

const (
	// Cognito automatically takes preventative actions in response to different levels of risk that you configure for your user pool.
	StandardThreatProtectionMode_FULL_FUNCTION StandardThreatProtectionMode = "FULL_FUNCTION"
	// Cognito gathers metrics on detected risks, but doesn't take automatic action.
	StandardThreatProtectionMode_AUDIT_ONLY StandardThreatProtectionMode = "AUDIT_ONLY"
	// Cognito doesn't gather metrics on detected risks or automatically take preventative actions.
	StandardThreatProtectionMode_NO_ENFORCEMENT StandardThreatProtectionMode = "NO_ENFORCEMENT"
)

