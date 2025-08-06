package awscognito


// The Type of Threat Protection Enabled for Custom Authentication.
//
// This feature only functions if your FeaturePlan is set to FeaturePlan.PLUS
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-userpooladdons.html
//
type CustomThreatProtectionMode string

const (
	// Cognito automatically takes preventative actions in response to different levels of risk that you configure for your user pool.
	CustomThreatProtectionMode_FULL_FUNCTION CustomThreatProtectionMode = "FULL_FUNCTION"
	// Cognito gathers metrics on detected risks, but doesn't take automatic action.
	CustomThreatProtectionMode_AUDIT_ONLY CustomThreatProtectionMode = "AUDIT_ONLY"
)

