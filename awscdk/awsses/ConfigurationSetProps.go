package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for a configuration set.
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
type ConfigurationSetProps struct {
	// A name for the configuration set.
	// Default: - a CloudFormation generated name.
	//
	ConfigurationSetName *string `field:"optional" json:"configurationSetName" yaml:"configurationSetName"`
	// The custom subdomain that is used to redirect email recipients to the Amazon SES event tracking domain.
	// Default: - use the default awstrack.me domain
	//
	CustomTrackingRedirectDomain *string `field:"optional" json:"customTrackingRedirectDomain" yaml:"customTrackingRedirectDomain"`
	// The dedicated IP pool to associate with the configuration set.
	// Default: - do not use a dedicated IP pool.
	//
	DedicatedIpPool IDedicatedIpPool `field:"optional" json:"dedicatedIpPool" yaml:"dedicatedIpPool"`
	// The maximum amount of time that Amazon SES API v2 will attempt delivery of email.
	//
	// This value must be greater than or equal to 5 minutes and less than or equal to 14 hours.
	// Default: undefined - SES defaults to 14 hours.
	//
	MaxDeliveryDuration awscdk.Duration `field:"optional" json:"maxDeliveryDuration" yaml:"maxDeliveryDuration"`
	// Whether to publish reputation metrics for the configuration set, such as bounce and complaint rates, to Amazon CloudWatch.
	// Default: true.
	//
	ReputationMetrics *bool `field:"optional" json:"reputationMetrics" yaml:"reputationMetrics"`
	// Whether email sending is enabled.
	// Default: true.
	//
	SendingEnabled *bool `field:"optional" json:"sendingEnabled" yaml:"sendingEnabled"`
	// The reasons for which recipient email addresses should be automatically added to your account's suppression list.
	// Default: - use account level settings.
	//
	SuppressionReasons SuppressionReasons `field:"optional" json:"suppressionReasons" yaml:"suppressionReasons"`
	// Specifies whether messages that use the configuration set are required to use Transport Layer Security (TLS).
	// Default: ConfigurationSetTlsPolicy.OPTIONAL
	//
	TlsPolicy ConfigurationSetTlsPolicy `field:"optional" json:"tlsPolicy" yaml:"tlsPolicy"`
	// The Virtual Deliverability Manager (VDM) options that apply to the configuration set.
	// Default: - VDM options not configured at the configuration set level. In this case, use account level settings. (To set the account level settings using CDK, use the `VdmAttributes` Construct.)
	//
	VdmOptions *VdmOptions `field:"optional" json:"vdmOptions" yaml:"vdmOptions"`
}

