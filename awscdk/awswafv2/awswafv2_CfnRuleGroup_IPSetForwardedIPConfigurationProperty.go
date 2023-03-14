package awswafv2


// The configuration for inspecting IP addresses in an HTTP header that you specify, instead of using the IP address that's reported by the web request origin.
//
// Commonly, this is the X-Forwarded-For (XFF) header, but you can specify any header name.
//
// > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.
//
// This configuration is used only for `IPSetReferenceStatement` . For `GeoMatchStatement` and `RateBasedStatement` , use `ForwardedIPConfig` instead.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iPSetForwardedIPConfigurationProperty := map[string]*string{
//   	"fallbackBehavior": jsii.String("fallbackBehavior"),
//   	"headerName": jsii.String("headerName"),
//   	"position": jsii.String("position"),
//   }
//
type CfnRuleGroup_IPSetForwardedIPConfigurationProperty struct {
	// The match status to assign to the web request if the request doesn't have a valid IP address in the specified position.
	//
	// > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.
	//
	// You can specify the following fallback behaviors:
	//
	// - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
	// - `NO_MATCH` - Treat the web request as not matching the rule statement.
	FallbackBehavior *string `field:"required" json:"fallbackBehavior" yaml:"fallbackBehavior"`
	// The name of the HTTP header to use for the IP address.
	//
	// For example, to use the X-Forwarded-For (XFF) header, set this to `X-Forwarded-For` .
	//
	// > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
	// The position in the header to search for the IP address.
	//
	// The header can contain IP addresses of the original client and also of proxies. For example, the header value could be `10.1.1.1, 127.0.0.0, 10.10.10.10` where the first IP address identifies the original client and the rest identify proxies that the request went through.
	//
	// The options for this setting are the following:
	//
	// - FIRST - Inspect the first IP address in the list of IP addresses in the header. This is usually the client's original IP.
	// - LAST - Inspect the last IP address in the list of IP addresses in the header.
	// - ANY - Inspect all IP addresses in the header for a match. If the header contains more than 10 IP addresses, AWS WAF inspects the last 10.
	Position *string `field:"required" json:"position" yaml:"position"`
}

