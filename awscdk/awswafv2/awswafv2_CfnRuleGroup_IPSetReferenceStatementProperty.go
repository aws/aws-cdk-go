package awswafv2


// A rule statement used to detect web requests coming from particular IP addresses or address ranges.
//
// To use this, create an `IPSet` that specifies the addresses you want to detect, then use the ARN of that set in this statement.
//
// Each IP set rule statement references an IP set. You create and maintain the set independent of your rules. This allows you to use the single set in multiple rules. When you update the referenced set, AWS WAF automatically updates all rules that reference it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iPSetReferenceStatementProperty := map[string]interface{}{
//   	"arn": jsii.String("arn"),
//
//   	// the properties below are optional
//   	"ipSetForwardedIpConfig": map[string]*string{
//   		"fallbackBehavior": jsii.String("fallbackBehavior"),
//   		"headerName": jsii.String("headerName"),
//   		"position": jsii.String("position"),
//   	},
//   }
//
type CfnRuleGroup_IPSetReferenceStatementProperty struct {
	// The Amazon Resource Name (ARN) of the `IPSet` that this statement references.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The configuration for inspecting IP addresses in an HTTP header that you specify, instead of using the IP address that's reported by the web request origin.
	//
	// Commonly, this is the X-Forwarded-For (XFF) header, but you can specify any header name.
	//
	// > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.
	IpSetForwardedIpConfig interface{} `field:"optional" json:"ipSetForwardedIpConfig" yaml:"ipSetForwardedIpConfig"`
}

