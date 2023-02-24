package awsses


// TLS policy for a configuration set.
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
type ConfigurationSetTlsPolicy string

const (
	// Messages are only delivered if a TLS connection can be established.
	ConfigurationSetTlsPolicy_REQUIRE ConfigurationSetTlsPolicy = "REQUIRE"
	// Messages can be delivered in plain text if a TLS connection can't be established.
	ConfigurationSetTlsPolicy_OPTIONAL ConfigurationSetTlsPolicy = "OPTIONAL"
)

