package awsses


// HTTPS policy option for the protocol of the open and click tracking links for your custom redirect domain.
//
// Example:
//   ses.NewConfigurationSet(this, jsii.String("ConfigurationSet"), &ConfigurationSetProps{
//   	CustomTrackingRedirectDomain: jsii.String("track.cdk.dev"),
//   	CustomTrackingHttpsPolicy: ses.HttpsPolicy_REQUIRE,
//   })
//
type HttpsPolicy string

const (
	// Open and Click tracking links will both be wrapped using HTTPS.
	HttpsPolicy_REQUIRE HttpsPolicy = "REQUIRE"
	// Open tracking links will be wrapped using HTTPS.
	//
	// Click tracking links will be wrapped using the original protocol of the link.
	HttpsPolicy_REQUIRE_OPEN_ONLY HttpsPolicy = "REQUIRE_OPEN_ONLY"
	// Open tracking links will be wrapped using HTTP.
	//
	// Click tracking links will be wrapped using the original protocol of the link.
	HttpsPolicy_OPTIONAL HttpsPolicy = "OPTIONAL"
)

