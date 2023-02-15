package awscognito


// The different ways in which a user pool's Advanced Security Mode can be configured.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &userPoolProps{
//   	// ...
//   	advancedSecurityMode: cognito.advancedSecurityMode_ENFORCED,
//   })
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-userpooladdons.html#cfn-cognito-userpool-userpooladdons-advancedsecuritymode
//
type AdvancedSecurityMode string

const (
	// Enable advanced security mode.
	AdvancedSecurityMode_ENFORCED AdvancedSecurityMode = "ENFORCED"
	// gather metrics on detected risks without taking action.
	//
	// Metrics are published to Amazon CloudWatch.
	AdvancedSecurityMode_AUDIT AdvancedSecurityMode = "AUDIT"
	// Advanced security mode is disabled.
	AdvancedSecurityMode_OFF AdvancedSecurityMode = "OFF"
)

