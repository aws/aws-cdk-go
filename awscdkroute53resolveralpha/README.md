# Amazon Route53 Resolver Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

## DNS Firewall

With Route 53 Resolver DNS Firewall, you can filter and regulate outbound DNS traffic for your
virtual private connections (VPCs). To do this, you create reusable collections of filtering rules
in DNS Firewall rule groups and associate the rule groups to your VPC.

DNS Firewall provides protection for outbound DNS requests from your VPCs. These requests route
through Resolver for domain name resolution. A primary use of DNS Firewall protections is to help
prevent DNS exfiltration of your data. DNS exfiltration can happen when a bad actor compromises
an application instance in your VPC and then uses DNS lookup to send data out of the VPC to a domain
that they control. With DNS Firewall, you can monitor and control the domains that your applications
can query. You can deny access to the domains that you know to be bad and allow all other queries
to pass through. Alternately, you can deny access to all domains except for the ones that you
explicitly trust.

### Domain lists

Domain lists can be created using a list of strings, a text file stored in Amazon S3 or a local
text file:

```go
blockList := route53resolver.NewFirewallDomainList(this, jsii.String("BlockList"), &FirewallDomainListProps{
	Domains: route53resolver.FirewallDomains_FromList([]*string{
		jsii.String("bad-domain.com"),
		jsii.String("bot-domain.net"),
	}),
})

s3List := route53resolver.NewFirewallDomainList(this, jsii.String("S3List"), &FirewallDomainListProps{
	Domains: route53resolver.FirewallDomains_FromS3Url(jsii.String("s3://bucket/prefix/object")),
})

assetList := route53resolver.NewFirewallDomainList(this, jsii.String("AssetList"), &FirewallDomainListProps{
	Domains: route53resolver.FirewallDomains_FromAsset(jsii.String("/path/to/domains.txt")),
})
```

The file must be a text file and must contain a single domain per line.

Use `FirewallDomainList.fromFirewallDomainListId()` to import an existing or [AWS managed domain list](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/resolver-dns-firewall-managed-domain-lists.html):

```go
// AWSManagedDomainsMalwareDomainList in us-east-1
malwareList := route53resolver.FirewallDomainList_FromFirewallDomainListId(this, jsii.String("Malware"), jsii.String("rslvr-fdl-2c46f2ecbfec4dcc"))
```

### Rule group

Create a rule group:

```go
var myBlockList FirewallDomainList

route53resolver.NewFirewallRuleGroup(this, jsii.String("RuleGroup"), &FirewallRuleGroupProps{
	Rules: []FirewallRule{
		&FirewallRule{
			Priority: jsii.Number(10),
			FirewallDomainList: myBlockList,
			// block and reply with NODATA
			Action: route53resolver.FirewallRuleAction_Block(),
		},
	},
})
```

Rules can be added at construction time or using `addRule()`:

```go
var myBlockList FirewallDomainList
var ruleGroup FirewallRuleGroup


ruleGroup.AddRule(&FirewallRule{
	Priority: jsii.Number(10),
	FirewallDomainList: myBlockList,
	// block and reply with NXDOMAIN
	Action: route53resolver.FirewallRuleAction_Block(route53resolver.DnsBlockResponse_NxDomain()),
})

ruleGroup.AddRule(&FirewallRule{
	Priority: jsii.Number(20),
	FirewallDomainList: myBlockList,
	// block and override DNS response with a custom domain
	Action: route53resolver.FirewallRuleAction_*Block(route53resolver.DnsBlockResponse_Override(jsii.String("amazon.com"))),
})
```

Use `associate()` to associate a rule group with a VPC:

```go
import ec2 "github.com/aws/aws-cdk-go/awscdk"

var ruleGroup FirewallRuleGroup
var myVpc Vpc


ruleGroup.Associate(jsii.String("Association"), &FirewallRuleGroupAssociationOptions{
	Priority: jsii.Number(101),
	Vpc: myVpc,
})
```
