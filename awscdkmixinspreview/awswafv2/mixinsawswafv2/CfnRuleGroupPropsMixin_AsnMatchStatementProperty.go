package mixinsawswafv2


// A rule statement that inspects web traffic based on the Autonomous System Number (ASN) associated with the request's IP address.
//
// For additional details, see [ASN match rule statement](https://docs.aws.amazon.com/waf/latest/developerguide/waf-rule-statement-type-asn-match.html) in the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   asnMatchStatementProperty := &AsnMatchStatementProperty{
//   	AsnList: []interface{}{
//   		jsii.Number(123),
//   	},
//   	ForwardedIpConfig: &ForwardedIPConfigurationProperty{
//   		FallbackBehavior: jsii.String("fallbackBehavior"),
//   		HeaderName: jsii.String("headerName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-asnmatchstatement.html
//
type CfnRuleGroupPropsMixin_AsnMatchStatementProperty struct {
	// Contains one or more Autonomous System Numbers (ASNs).
	//
	// ASNs are unique identifiers assigned to large internet networks managed by organizations such as internet service providers, enterprises, universities, or government agencies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-asnmatchstatement.html#cfn-wafv2-rulegroup-asnmatchstatement-asnlist
	//
	AsnList interface{} `field:"optional" json:"asnList" yaml:"asnList"`
	// The configuration for inspecting IP addresses to match against an ASN in an HTTP header that you specify, instead of using the IP address that's reported by the web request origin.
	//
	// Commonly, this is the X-Forwarded-For (XFF) header, but you can specify any header name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-asnmatchstatement.html#cfn-wafv2-rulegroup-asnmatchstatement-forwardedipconfig
	//
	ForwardedIpConfig interface{} `field:"optional" json:"forwardedIpConfig" yaml:"forwardedIpConfig"`
}

