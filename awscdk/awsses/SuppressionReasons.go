package awsses


// Reasons for which recipient email addresses should be automatically added to your account's suppression list.
//
// Example:
//   // Only bounces will be suppressed.
//   // Only bounces will be suppressed.
//   ses.NewConfigurationSet(this, jsii.String("ConfigurationSet"), &ConfigurationSetProps{
//   	SuppressionReasons: ses.SuppressionReasons_BOUNCES_ONLY,
//   })
//
//   // Only complaints will be suppressed.
//   // Only complaints will be suppressed.
//   ses.NewConfigurationSet(this, jsii.String("ConfigurationSet"), &ConfigurationSetProps{
//   	SuppressionReasons: ses.SuppressionReasons_COMPLAINTS_ONLY,
//   })
//
//   // Both bounces and complaints will be suppressed.
//   // Both bounces and complaints will be suppressed.
//   ses.NewConfigurationSet(this, jsii.String("ConfigurationSet"), &ConfigurationSetProps{
//   	SuppressionReasons: ses.SuppressionReasons_BOUNCES_AND_COMPLAINTS,
//   })
//
type SuppressionReasons string

const (
	// Bounces and complaints.
	SuppressionReasons_BOUNCES_AND_COMPLAINTS SuppressionReasons = "BOUNCES_AND_COMPLAINTS"
	// Bounces only.
	SuppressionReasons_BOUNCES_ONLY SuppressionReasons = "BOUNCES_ONLY"
	// Complaints only.
	SuppressionReasons_COMPLAINTS_ONLY SuppressionReasons = "COMPLAINTS_ONLY"
)

