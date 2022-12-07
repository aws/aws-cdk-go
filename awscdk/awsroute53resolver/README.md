# Amazon Route53 Resolver Construct Library

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
blockList := route53resolver.NewFirewallDomainList(this, jsii.String("BlockList"), &firewallDomainListProps{
	domains: route53resolver.firewallDomains.fromList([]*string{
		jsii.String("bad-domain.com"),
		jsii.String("bot-domain.net"),
	}),
})

s3List := route53resolver.NewFirewallDomainList(this, jsii.String("S3List"), &firewallDomainListProps{
	domains: route53resolver.*firewallDomains.fromS3Url(jsii.String("s3://bucket/prefix/object")),
})

assetList := route53resolver.NewFirewallDomainList(this, jsii.String("AssetList"), &firewallDomainListProps{
	domains: route53resolver.*firewallDomains.fromAsset(jsii.String("/path/to/domains.txt")),
})
```

The file must be a text file and must contain a single domain per line.

Use `FirewallDomainList.fromFirewallDomainListId()` to import an existing or [AWS managed domain list](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/resolver-dns-firewall-managed-domain-lists.html):

```go
// AWSManagedDomainsMalwareDomainList in us-east-1
malwareList := route53resolver.firewallDomainList.fromFirewallDomainListId(this, jsii.String("Malware"), jsii.String("rslvr-fdl-2c46f2ecbfec4dcc"))
```

### Rule group

Create a rule group:

```go
var myBlockList firewallDomainList

route53resolver.NewFirewallRuleGroup(this, jsii.String("RuleGroup"), &firewallRuleGroupProps{
	rules: []firewallRule{
		&firewallRule{
			priority: jsii.Number(10),
			firewallDomainList: myBlockList,
			// block and reply with NODATA
			action: route53resolver.firewallRuleAction.block(),
		},
	},
})
```

Rules can be added at construction time or using `addRule()`:

```go
var myBlockList firewallDomainList
var ruleGroup firewallRuleGroup


ruleGroup.addRule(&firewallRule{
	priority: jsii.Number(10),
	firewallDomainList: myBlockList,
	// block and reply with NXDOMAIN
	action: route53resolver.firewallRuleAction.block(route53resolver.dnsBlockResponse.nxDomain()),
})

ruleGroup.addRule(&firewallRule{
	priority: jsii.Number(20),
	firewallDomainList: myBlockList,
	// block and override DNS response with a custom domain
	action: route53resolver.*firewallRuleAction.block(route53resolver.*dnsBlockResponse.override(jsii.String("amazon.com"))),
})
```

Use `associate()` to associate a rule group with a VPC:

```go
import ec2 "github.com/aws/aws-cdk-go/awscdk"

var ruleGroup firewallRuleGroup
var myVpc vpc


ruleGroup.associate(jsii.String("Association"), &firewallRuleGroupAssociationOptions{
	priority: jsii.Number(101),
	vpc: myVpc,
})
```
