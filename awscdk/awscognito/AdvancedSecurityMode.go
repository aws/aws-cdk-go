package awscognito


// The different ways in which a user pool's Advanced Security Mode can be configured.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &UserPoolProps{
//   	// ...
//   	AdvancedSecurityMode: cognito.AdvancedSecurityMode_ENFORCED,
//   })
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-userpooladdons.html
//
// Deprecated: Advanced Security Mode is deprecated due to user pool feature plans. Use StandardThreatProtectionMode and CustomThreatProtectionMode to set Thread Protection level.
type AdvancedSecurityMode string

const (
	// Enable advanced security mode.
	// Deprecated: Advanced Security Mode is deprecated due to user pool feature plans. Use StandardThreatProtectionMode and CustomThreatProtectionMode to set Thread Protection level.
	AdvancedSecurityMode_ENFORCED AdvancedSecurityMode = "ENFORCED"
	// gather metrics on detected risks without taking action.
	//
	// Metrics are published to Amazon CloudWatch.
	// Deprecated: Advanced Security Mode is deprecated due to user pool feature plans. Use StandardThreatProtectionMode and CustomThreatProtectionMode to set Thread Protection level.
	AdvancedSecurityMode_AUDIT AdvancedSecurityMode = "AUDIT"
	// Advanced security mode is disabled.
	// Deprecated: Advanced Security Mode is deprecated due to user pool feature plans. Use StandardThreatProtectionMode and CustomThreatProtectionMode to set Thread Protection level.
	AdvancedSecurityMode_OFF AdvancedSecurityMode = "OFF"
)

