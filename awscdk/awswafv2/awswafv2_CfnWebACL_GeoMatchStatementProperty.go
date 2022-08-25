package awswafv2


// A rule statement used to identify web requests based on country of origin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geoMatchStatementProperty := &geoMatchStatementProperty{
//   	countryCodes: []*string{
//   		jsii.String("countryCodes"),
//   	},
//   	forwardedIpConfig: &forwardedIPConfigurationProperty{
//   		fallbackBehavior: jsii.String("fallbackBehavior"),
//   		headerName: jsii.String("headerName"),
//   	},
//   }
//
type CfnWebACL_GeoMatchStatementProperty struct {
	// An array of two-character country codes, for example, `[ "US", "CN" ]` , from the alpha-2 country ISO codes of the ISO 3166 international standard.
	CountryCodes *[]*string `field:"optional" json:"countryCodes" yaml:"countryCodes"`
	// The configuration for inspecting IP addresses in an HTTP header that you specify, instead of using the IP address that's reported by the web request origin.
	//
	// Commonly, this is the X-Forwarded-For (XFF) header, but you can specify any header name.
	//
	// > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.
	ForwardedIpConfig interface{} `field:"optional" json:"forwardedIpConfig" yaml:"forwardedIpConfig"`
}

