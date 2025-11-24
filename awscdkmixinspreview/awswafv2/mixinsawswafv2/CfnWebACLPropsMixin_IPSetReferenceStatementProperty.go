package mixinsawswafv2


// A rule statement used to detect web requests coming from particular IP addresses or address ranges.
//
// To use this, create an `IPSet` that specifies the addresses you want to detect, then use the ARN of that set in this statement.
//
// Each IP set rule statement references an IP set. You create and maintain the set independent of your rules. This allows you to use the single set in multiple rules. When you update the referenced set, AWS WAF automatically updates all rules that reference it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   iPSetReferenceStatementProperty := map[string]*string{
//   	"arn": jsii.String("arn"),
//   	"ipSetForwardedIpConfig": map[string]*string{
//   		"fallbackBehavior": jsii.String("fallbackBehavior"),
//   		"headerName": jsii.String("headerName"),
//   		"position": jsii.String("position"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ipsetreferencestatement.html
//
type CfnWebACLPropsMixin_IPSetReferenceStatementProperty struct {
	// The Amazon Resource Name (ARN) of the `IPSet` that this statement references.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ipsetreferencestatement.html#cfn-wafv2-webacl-ipsetreferencestatement-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The configuration for inspecting IP addresses in an HTTP header that you specify, instead of using the IP address that's reported by the web request origin.
	//
	// Commonly, this is the X-Forwarded-For (XFF) header, but you can specify any header name.
	//
	// > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ipsetreferencestatement.html#cfn-wafv2-webacl-ipsetreferencestatement-ipsetforwardedipconfig
	//
	IpSetForwardedIpConfig interface{} `field:"optional" json:"ipSetForwardedIpConfig" yaml:"ipSetForwardedIpConfig"`
}

