package awsses


// Reasons for which recipient email addresses should be automatically added to your account's suppression list.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var myPool iDedicatedIpPool
//
//
//   ses.NewConfigurationSet(this, jsii.String("ConfigurationSet"), &ConfigurationSetProps{
//   	CustomTrackingRedirectDomain: jsii.String("track.cdk.dev"),
//   	SuppressionReasons: ses.SuppressionReasons_COMPLAINTS_ONLY,
//   	TlsPolicy: ses.ConfigurationSetTlsPolicy_REQUIRE,
//   	DedicatedIpPool: myPool,
//   	// Specify maximum delivery time
//   	// This configuration can be useful in such cases as time-sensitive emails (like those containing a one-time-password),
//   	// transactional emails, and email that you want to ensure isn't delivered during non-business hours.
//   	MaxDeliveryDuration: awscdk.Duration_Minutes(jsii.Number(10)),
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

