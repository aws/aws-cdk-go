package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for a configuration set.
//
// Example:
//   ses.NewConfigurationSet(this, jsii.String("ConfigurationSet"), &ConfigurationSetProps{
//   	CustomTrackingRedirectDomain: jsii.String("track.cdk.dev"),
//   	CustomTrackingHttpsPolicy: ses.HttpsPolicy_REQUIRE,
//   })
//
type ConfigurationSetProps struct {
	// A name for the configuration set.
	// Default: - a CloudFormation generated name.
	//
	ConfigurationSetName *string `field:"optional" json:"configurationSetName" yaml:"configurationSetName"`
	// The https policy to use for tracking open and click events.
	// Default: - HttpsPolicy.OPTIONAL if customTrackingRedirectDomain is set, otherwise undefined
	//
	CustomTrackingHttpsPolicy HttpsPolicy `field:"optional" json:"customTrackingHttpsPolicy" yaml:"customTrackingHttpsPolicy"`
	// The custom subdomain that is used to redirect email recipients to the Amazon SES event tracking domain.
	// Default: - use the default awstrack.me domain
	//
	CustomTrackingRedirectDomain *string `field:"optional" json:"customTrackingRedirectDomain" yaml:"customTrackingRedirectDomain"`
	// The dedicated IP pool to associate with the configuration set.
	// Default: - do not use a dedicated IP pool.
	//
	DedicatedIpPool IDedicatedIpPool `field:"optional" json:"dedicatedIpPool" yaml:"dedicatedIpPool"`
	// If true, account-level suppression list is disabled;
	//
	// email sent with this configuration set
	// will not use any suppression settings at all.
	// Default: false.
	//
	DisableSuppressionList *bool `field:"optional" json:"disableSuppressionList" yaml:"disableSuppressionList"`
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

