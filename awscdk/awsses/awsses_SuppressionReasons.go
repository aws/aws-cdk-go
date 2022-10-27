package awsses


// Reasons for which recipient email addresses should be automatically added to your account's suppression list.
//
// Example:
//   var myPool iDedicatedIpPool
//
//
//   ses.NewConfigurationSet(this, jsii.String("ConfigurationSet"), &configurationSetProps{
//   	customTrackingRedirectDomain: jsii.String("track.cdk.dev"),
//   	suppressionReasons: ses.suppressionReasons_COMPLAINTS_ONLY,
//   	tlsPolicy: ses.configurationSetTlsPolicy_REQUIRE,
//   	dedicatedIpPool: myPool,
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

