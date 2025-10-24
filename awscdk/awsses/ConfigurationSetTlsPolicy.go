package awsses


// TLS policy for a configuration set.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var myPool IDedicatedIpPool
//
//
//   ses.NewConfigurationSet(this, jsii.String("ConfigurationSet"), &ConfigurationSetProps{
//   	TlsPolicy: ses.ConfigurationSetTlsPolicy_REQUIRE,
//   	DedicatedIpPool: myPool,
//   	// Specify maximum delivery time
//   	// This configuration can be useful in such cases as time-sensitive emails (like those containing a one-time-password),
//   	// transactional emails, and email that you want to ensure isn't delivered during non-business hours.
//   	MaxDeliveryDuration: awscdk.Duration_Minutes(jsii.Number(10)),
//   })
//
type ConfigurationSetTlsPolicy string

const (
	// Messages are only delivered if a TLS connection can be established.
	ConfigurationSetTlsPolicy_REQUIRE ConfigurationSetTlsPolicy = "REQUIRE"
	// Messages can be delivered in plain text if a TLS connection can't be established.
	ConfigurationSetTlsPolicy_OPTIONAL ConfigurationSetTlsPolicy = "OPTIONAL"
)

