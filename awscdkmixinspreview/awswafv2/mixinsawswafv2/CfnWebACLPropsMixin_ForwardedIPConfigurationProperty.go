package mixinsawswafv2


// The configuration for inspecting IP addresses in an HTTP header that you specify, instead of using the IP address that's reported by the web request origin.
//
// Commonly, this is the X-Forwarded-For (XFF) header, but you can specify any header name.
//
// > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.
//
// This configuration is used for `GeoMatchStatement` , `AsnMatchStatement` , and `RateBasedStatement` . For `IPSetReferenceStatement` , use `IPSetForwardedIPConfig` instead.
//
// AWS WAF only evaluates the first IP address found in the specified HTTP header.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   forwardedIPConfigurationProperty := &ForwardedIPConfigurationProperty{
//   	FallbackBehavior: jsii.String("fallbackBehavior"),
//   	HeaderName: jsii.String("headerName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-forwardedipconfiguration.html
//
type CfnWebACLPropsMixin_ForwardedIPConfigurationProperty struct {
	// The match status to assign to the web request if the request doesn't have a valid IP address in the specified position.
	//
	// > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.
	//
	// You can specify the following fallback behaviors:
	//
	// - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
	// - `NO_MATCH` - Treat the web request as not matching the rule statement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-forwardedipconfiguration.html#cfn-wafv2-webacl-forwardedipconfiguration-fallbackbehavior
	//
	FallbackBehavior *string `field:"optional" json:"fallbackBehavior" yaml:"fallbackBehavior"`
	// The name of the HTTP header to use for the IP address.
	//
	// For example, to use the X-Forwarded-For (XFF) header, set this to `X-Forwarded-For` .
	//
	// > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-forwardedipconfiguration.html#cfn-wafv2-webacl-forwardedipconfiguration-headername
	//
	HeaderName *string `field:"optional" json:"headerName" yaml:"headerName"`
}

