package awsses


// Properties for a configuration set.
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
type ConfigurationSetProps struct {
	// A name for the configuration set.
	ConfigurationSetName *string `field:"optional" json:"configurationSetName" yaml:"configurationSetName"`
	// The custom subdomain that is used to redirect email recipients to the Amazon SES event tracking domain.
	CustomTrackingRedirectDomain *string `field:"optional" json:"customTrackingRedirectDomain" yaml:"customTrackingRedirectDomain"`
	// The dedicated IP pool to associate with the configuration set.
	DedicatedIpPool IDedicatedIpPool `field:"optional" json:"dedicatedIpPool" yaml:"dedicatedIpPool"`
	// Whether to publish reputation metrics for the configuration set, such as bounce and complaint rates, to Amazon CloudWatch.
	ReputationMetrics *bool `field:"optional" json:"reputationMetrics" yaml:"reputationMetrics"`
	// Whether email sending is enabled.
	SendingEnabled *bool `field:"optional" json:"sendingEnabled" yaml:"sendingEnabled"`
	// The reasons for which recipient email addresses should be automatically added to your account's suppression list.
	SuppressionReasons SuppressionReasons `field:"optional" json:"suppressionReasons" yaml:"suppressionReasons"`
	// Specifies whether messages that use the configuration set are required to use Transport Layer Security (TLS).
	TlsPolicy ConfigurationSetTlsPolicy `field:"optional" json:"tlsPolicy" yaml:"tlsPolicy"`
}

