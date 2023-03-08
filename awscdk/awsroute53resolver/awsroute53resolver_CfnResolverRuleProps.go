package awsroute53resolver

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnResolverRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResolverRuleProps := &cfnResolverRuleProps{
//   	domainName: jsii.String("domainName"),
//   	ruleType: jsii.String("ruleType"),
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	resolverEndpointId: jsii.String("resolverEndpointId"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	targetIps: []interface{}{
//   		&targetAddressProperty{
//   			ip: jsii.String("ip"),
//
//   			// the properties below are optional
//   			port: jsii.String("port"),
//   		},
//   	},
//   }
//
type CfnResolverRuleProps struct {
	// DNS queries for this domain name are forwarded to the IP addresses that are specified in `TargetIps` .
	//
	// If a query matches multiple Resolver rules (example.com and www.example.com), the query is routed using the Resolver rule that contains the most specific domain name (www.example.com).
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// When you want to forward DNS queries for specified domain name to resolvers on your network, specify `FORWARD` .
	//
	// When you have a forwarding rule to forward DNS queries for a domain to your network and you want Resolver to process queries for a subdomain of that domain, specify `SYSTEM` .
	//
	// For example, to forward DNS queries for example.com to resolvers on your network, you create a rule and specify `FORWARD` for `RuleType` . To then have Resolver process queries for apex.example.com, you create a rule and specify `SYSTEM` for `RuleType` .
	//
	// Currently, only Resolver can create rules that have a value of `RECURSIVE` for `RuleType` .
	RuleType *string `field:"required" json:"ruleType" yaml:"ruleType"`
	// The name for the Resolver rule, which you specified when you created the Resolver rule.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ID of the endpoint that the rule is associated with.
	ResolverEndpointId *string `field:"optional" json:"resolverEndpointId" yaml:"resolverEndpointId"`
	// Route 53 Resolver doesn't support updating tags through CloudFormation.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// An array that contains the IP addresses and ports that an outbound endpoint forwards DNS queries to.
	//
	// Typically, these are the IP addresses of DNS resolvers on your network. Specify IPv4 addresses. IPv6 is not supported.
	TargetIps interface{} `field:"optional" json:"targetIps" yaml:"targetIps"`
}

