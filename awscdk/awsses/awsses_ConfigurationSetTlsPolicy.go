package awsses


// TLS policy for a configuration set.
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
type ConfigurationSetTlsPolicy string

const (
	// Messages are only delivered if a TLS connection can be established.
	ConfigurationSetTlsPolicy_REQUIRE ConfigurationSetTlsPolicy = "REQUIRE"
	// Messages can be delivered in plain text if a TLS connection can't be established.
	ConfigurationSetTlsPolicy_OPTIONAL ConfigurationSetTlsPolicy = "OPTIONAL"
)

