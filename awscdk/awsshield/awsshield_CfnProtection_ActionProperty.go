package awsshield


// Specifies the action setting that Shield Advanced should use in the AWS WAF rules that it creates on behalf of the protected resource in response to DDoS attacks.
//
// You specify this as part of the configuration for the automatic application layer DDoS mitigation feature, when you enable or update automatic mitigation. Shield Advanced creates the AWS WAF rules in a Shield Advanced-managed rule group, inside the web ACL that you have associated with the resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var block interface{}
//   var count interface{}
//
//   actionProperty := &ActionProperty{
//   	Block: block,
//   	Count: count,
//   }
//
type CfnProtection_ActionProperty struct {
	// Specifies that Shield Advanced should configure its AWS WAF rules with the AWS WAF `Block` action.
	//
	// You must specify exactly one action, either `Block` or `Count` .
	//
	// Example JSON: `{ "Block": {} }`
	//
	// Example YAML: `Block: {}`.
	Block interface{} `field:"optional" json:"block" yaml:"block"`
	// Specifies that Shield Advanced should configure its AWS WAF rules with the AWS WAF `Count` action.
	//
	// You must specify exactly one action, either `Block` or `Count` .
	//
	// Example JSON: `{ "Count": {} }`
	//
	// Example YAML: `Count: {}`.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
}

