package awsroute53globalresolver

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDnsView`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDnsViewProps := &CfnDnsViewProps{
//   	GlobalResolverId: jsii.String("globalResolverId"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	ClientToken: jsii.String("clientToken"),
//   	Description: jsii.String("description"),
//   	DnssecValidation: jsii.String("dnssecValidation"),
//   	EdnsClientSubnet: jsii.String("ednsClientSubnet"),
//   	FirewallRulesFailOpen: jsii.String("firewallRulesFailOpen"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-dnsview.html
//
type CfnDnsViewProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-dnsview.html#cfn-route53globalresolver-dnsview-globalresolverid
	//
	GlobalResolverId *string `field:"required" json:"globalResolverId" yaml:"globalResolverId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-dnsview.html#cfn-route53globalresolver-dnsview-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-dnsview.html#cfn-route53globalresolver-dnsview-clienttoken
	//
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-dnsview.html#cfn-route53globalresolver-dnsview-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-dnsview.html#cfn-route53globalresolver-dnsview-dnssecvalidation
	//
	DnssecValidation *string `field:"optional" json:"dnssecValidation" yaml:"dnssecValidation"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-dnsview.html#cfn-route53globalresolver-dnsview-ednsclientsubnet
	//
	EdnsClientSubnet *string `field:"optional" json:"ednsClientSubnet" yaml:"ednsClientSubnet"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-dnsview.html#cfn-route53globalresolver-dnsview-firewallrulesfailopen
	//
	FirewallRulesFailOpen *string `field:"optional" json:"firewallRulesFailOpen" yaml:"firewallRulesFailOpen"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-dnsview.html#cfn-route53globalresolver-dnsview-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

