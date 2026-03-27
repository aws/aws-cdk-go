package awsroute53globalresolver

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnFirewallDomainListPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnFirewallDomainListMixinProps := &CfnFirewallDomainListMixinProps{
//   	ClientToken: jsii.String("clientToken"),
//   	Description: jsii.String("description"),
//   	DomainFileUrl: jsii.String("domainFileUrl"),
//   	Domains: []*string{
//   		jsii.String("domains"),
//   	},
//   	GlobalResolverId: jsii.String("globalResolverId"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-firewalldomainlist.html
//
type CfnFirewallDomainListMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-firewalldomainlist.html#cfn-route53globalresolver-firewalldomainlist-clienttoken
	//
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-firewalldomainlist.html#cfn-route53globalresolver-firewalldomainlist-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// S3 URL to import domains from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-firewalldomainlist.html#cfn-route53globalresolver-firewalldomainlist-domainfileurl
	//
	DomainFileUrl *string `field:"optional" json:"domainFileUrl" yaml:"domainFileUrl"`
	// An inline list of domains to use for this domain list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-firewalldomainlist.html#cfn-route53globalresolver-firewalldomainlist-domains
	//
	Domains *[]*string `field:"optional" json:"domains" yaml:"domains"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-firewalldomainlist.html#cfn-route53globalresolver-firewalldomainlist-globalresolverid
	//
	GlobalResolverId *string `field:"optional" json:"globalResolverId" yaml:"globalResolverId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-firewalldomainlist.html#cfn-route53globalresolver-firewalldomainlist-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-firewalldomainlist.html#cfn-route53globalresolver-firewalldomainlist-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

