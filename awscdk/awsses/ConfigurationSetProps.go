package awsses


// Properties for a configuration set.
//
// Example:
//   var myPool iDedicatedIpPool
//
//
//   ses.NewConfigurationSet(this, jsii.String("ConfigurationSet"), &ConfigurationSetProps{
//   	CustomTrackingRedirectDomain: jsii.String("track.cdk.dev"),
//   	SuppressionReasons: ses.SuppressionReasons_COMPLAINTS_ONLY,
//   	TlsPolicy: ses.ConfigurationSetTlsPolicy_REQUIRE,
//   	DedicatedIpPool: myPool,
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
	// Whether to publish reputation metrics for the configuration set, such as bounce and complaint rates, to Amazon CloudWatch.
	// Default: false.
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
}

