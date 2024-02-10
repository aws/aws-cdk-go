# Amazon Route53 Construct Library

To add a public hosted zone:

```go
route53.NewPublicHostedZone(this, jsii.String("HostedZone"), &PublicHostedZoneProps{
	ZoneName: jsii.String("fully.qualified.domain.com"),
})
```

To add a private hosted zone, use `PrivateHostedZone`. Note that
`enableDnsHostnames` and `enableDnsSupport` must have been enabled for the
VPC you're configuring for private hosted zones.

```go
var vpc vpc


zone := route53.NewPrivateHostedZone(this, jsii.String("HostedZone"), &PrivateHostedZoneProps{
	ZoneName: jsii.String("fully.qualified.domain.com"),
	Vpc: Vpc,
})
```

Additional VPCs can be added with `zone.addVpc()`.

## Adding Records

To add a TXT record to your zone:

```go
var myZone hostedZone


route53.NewTxtRecord(this, jsii.String("TXTRecord"), &TxtRecordProps{
	Zone: myZone,
	RecordName: jsii.String("_foo"),
	 // If the name ends with a ".", it will be used as-is;
	// if it ends with a "." followed by the zone name, a trailing "." will be added automatically;
	// otherwise, a ".", the zone name, and a trailing "." will be added automatically.
	// Defaults to zone root if not specified.
	Values: []*string{
		jsii.String("Bar!"),
		jsii.String("Baz?"),
	},
	Ttl: awscdk.Duration_Minutes(jsii.Number(90)),
})
```

To add a NS record to your zone:

```go
var myZone hostedZone


route53.NewNsRecord(this, jsii.String("NSRecord"), &NsRecordProps{
	Zone: myZone,
	RecordName: jsii.String("foo"),
	Values: []*string{
		jsii.String("ns-1.awsdns.co.uk."),
		jsii.String("ns-2.awsdns.com."),
	},
	Ttl: awscdk.Duration_Minutes(jsii.Number(90)),
})
```

To add a DS record to your zone:

```go
var myZone hostedZone


route53.NewDsRecord(this, jsii.String("DSRecord"), &DsRecordProps{
	Zone: myZone,
	RecordName: jsii.String("foo"),
	Values: []*string{
		jsii.String("12345 3 1 123456789abcdef67890123456789abcdef67890"),
	},
	Ttl: awscdk.Duration_Minutes(jsii.Number(90)),
})
```

To add an A record to your zone:

```go
var myZone hostedZone


route53.NewARecord(this, jsii.String("ARecord"), &ARecordProps{
	Zone: myZone,
	Target: route53.RecordTarget_FromIpAddresses(jsii.String("1.2.3.4"), jsii.String("5.6.7.8")),
})
```

To add an A record for an EC2 instance with an Elastic IP (EIP) to your zone:

```go
var instance instance

var myZone hostedZone


elasticIp := ec2.NewCfnEIP(this, jsii.String("EIP"), &CfnEIPProps{
	Domain: jsii.String("vpc"),
	InstanceId: instance.InstanceId,
})
route53.NewARecord(this, jsii.String("ARecord"), &ARecordProps{
	Zone: myZone,
	Target: route53.RecordTarget_FromIpAddresses(elasticIp.ref),
})
```

To add an AAAA record pointing to a CloudFront distribution:

```go
import cloudfront "github.com/aws/aws-cdk-go/awscdk"

var myZone hostedZone
var distribution cloudFrontWebDistribution

route53.NewAaaaRecord(this, jsii.String("Alias"), &AaaaRecordProps{
	Zone: myZone,
	Target: route53.RecordTarget_FromAlias(targets.NewCloudFrontTarget(distribution)),
})
```

[Geolocation routing](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy-geo.html) can be enabled for continent, country or subdivision:

```go
var myZone hostedZone


// continent
// continent
route53.NewARecord(this, jsii.String("ARecordGeoLocationContinent"), &ARecordProps{
	Zone: myZone,
	Target: route53.RecordTarget_FromIpAddresses(jsii.String("1.2.3.0"), jsii.String("5.6.7.0")),
	GeoLocation: route53.GeoLocation_Continent(route53.Continent_EUROPE),
})

// country
// country
route53.NewARecord(this, jsii.String("ARecordGeoLocationCountry"), &ARecordProps{
	Zone: myZone,
	Target: route53.RecordTarget_*FromIpAddresses(jsii.String("1.2.3.1"), jsii.String("5.6.7.1")),
	GeoLocation: route53.GeoLocation_Country(jsii.String("DE")),
})

// subdivision
// subdivision
route53.NewARecord(this, jsii.String("ARecordGeoLocationSubDividion"), &ARecordProps{
	Zone: myZone,
	Target: route53.RecordTarget_*FromIpAddresses(jsii.String("1.2.3.2"), jsii.String("5.6.7.2")),
	GeoLocation: route53.GeoLocation_Subdivision(jsii.String("WA")),
})

// default (wildcard record if no specific record is found)
// default (wildcard record if no specific record is found)
route53.NewARecord(this, jsii.String("ARecordGeoLocationDefault"), &ARecordProps{
	Zone: myZone,
	Target: route53.RecordTarget_*FromIpAddresses(jsii.String("1.2.3.3"), jsii.String("5.6.7.3")),
	GeoLocation: route53.GeoLocation_Default(),
})
```

To enable [weighted routing](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy-weighted.html), use the `weight` parameter:

```go
var myZone hostedZone


route53.NewARecord(this, jsii.String("ARecordWeighted1"), &ARecordProps{
	Zone: myZone,
	Target: route53.RecordTarget_FromIpAddresses(jsii.String("1.2.3.4")),
	Weight: jsii.Number(10),
})
```

To enable [latency based routing](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy-latency.html), use the `region` parameter:

```go
var myZone hostedZone


route53.NewARecord(this, jsii.String("ARecordLatency1"), &ARecordProps{
	Zone: myZone,
	Target: route53.RecordTarget_FromIpAddresses(jsii.String("1.2.3.4")),
	Region: jsii.String("us-east-1"),
})
```

To enable [multivalue answer routing](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy-multivalue.html), use the `multivalueAnswer` parameter:

```go
var myZone hostedZone


route53.NewARecord(this, jsii.String("ARecordMultiValue1"), &ARecordProps{
	Zone: myZone,
	Target: route53.RecordTarget_FromIpAddresses(jsii.String("1.2.3.4")),
	MultiValueAnswer: jsii.Boolean(true),
})
```

To specify a unique identifier to differentiate among multiple resource record sets that have the same combination of name and type, use the `setIdentifier` parameter:

```go
var myZone hostedZone


route53.NewARecord(this, jsii.String("ARecordWeighted1"), &ARecordProps{
	Zone: myZone,
	Target: route53.RecordTarget_FromIpAddresses(jsii.String("1.2.3.4")),
	Weight: jsii.Number(10),
	SetIdentifier: jsii.String("weighted-record-id"),
})
```

**Warning** It is not possible to specify `setIdentifier` for a simple routing policy.

Constructs are available for A, AAAA, CAA, CNAME, MX, NS, SRV and TXT records.

Use the `CaaAmazonRecord` construct to easily restrict certificate authorities
allowed to issue certificates for a domain to Amazon only.

### Replacing existing record sets (dangerous!)

Use the `deleteExisting` prop to delete an existing record set before deploying the new one.
This is useful if you want to minimize downtime and avoid "manual" actions while deploying a
stack with a record set that already exists. This is typically the case for record sets that
are not already "owned" by CloudFormation or "owned" by another stack or construct that is
going to be deleted (migration).

> **N.B.:** this feature is dangerous, use with caution! It can only be used safely when
> `deleteExisting` is set to `true` as soon as the resource is added to the stack. Changing
> an existing Record Set's `deleteExisting` property from `false -> true` after deployment
> will delete the record!

```go
var myZone hostedZone


route53.NewARecord(this, jsii.String("ARecord"), &ARecordProps{
	Zone: myZone,
	Target: route53.RecordTarget_FromIpAddresses(jsii.String("1.2.3.4"), jsii.String("5.6.7.8")),
	DeleteExisting: jsii.Boolean(true),
})
```

### Cross Account Zone Delegation

If you want to have your root domain hosted zone in one account and your subdomain hosted
zone in a different one, you can use `CrossAccountZoneDelegationRecord` to set up delegation
between them.

In the account containing the parent hosted zone:

```go
parentZone := route53.NewPublicHostedZone(this, jsii.String("HostedZone"), &PublicHostedZoneProps{
	ZoneName: jsii.String("someexample.com"),
})
crossAccountRole := iam.NewRole(this, jsii.String("CrossAccountRole"), &RoleProps{
	// The role name must be predictable
	RoleName: jsii.String("MyDelegationRole"),
	// The other account
	AssumedBy: iam.NewAccountPrincipal(jsii.String("12345678901")),
	// You can scope down this role policy to be least privileged.
	// If you want the other account to be able to manage specific records,
	// you can scope down by resource and/or normalized record names
	InlinePolicies: map[string]policyDocument{
		"crossAccountPolicy": iam.NewPolicyDocument(&PolicyDocumentProps{
			"statements": []PolicyStatement{
				iam.NewPolicyStatement(&PolicyStatementProps{
					"sid": jsii.String("ListHostedZonesByName"),
					"effect": iam.Effect_ALLOW,
					"actions": []*string{
						jsii.String("route53:ListHostedZonesByName"),
					},
					"resources": []*string{
						jsii.String("*"),
					},
				}),
				iam.NewPolicyStatement(&PolicyStatementProps{
					"sid": jsii.String("GetHostedZoneAndChangeResourceRecordSet"),
					"effect": iam.Effect_ALLOW,
					"actions": []*string{
						jsii.String("route53:GetHostedZone"),
						jsii.String("route53:ChangeResourceRecordSet"),
					},
					// This example assumes the RecordSet subdomain.somexample.com
					// is contained in the HostedZone
					"resources": []*string{
						jsii.String("arn:aws:route53:::hostedzone/HZID00000000000000000"),
					},
					"conditions": map[string]interface{}{
						"ForAllValues:StringLike": map[string][]*string{
							"route53:ChangeResourceRecordSetsNormalizedRecordNames": []*string{
								jsii.String("subdomain.someexample.com"),
							},
						},
					},
				}),
			},
		}),
	},
})
parentZone.GrantDelegation(crossAccountRole)
```

In the account containing the child zone to be delegated:

```go
subZone := route53.NewPublicHostedZone(this, jsii.String("SubZone"), &PublicHostedZoneProps{
	ZoneName: jsii.String("sub.someexample.com"),
})

// import the delegation role by constructing the roleArn
delegationRoleArn := awscdk.stack_Of(this).FormatArn(&ArnComponents{
	Region: jsii.String(""),
	 // IAM is global in each partition
	Service: jsii.String("iam"),
	Account: jsii.String("parent-account-id"),
	Resource: jsii.String("role"),
	ResourceName: jsii.String("MyDelegationRole"),
})
delegationRole := iam.Role_FromRoleArn(this, jsii.String("DelegationRole"), delegationRoleArn)

// create the record
// create the record
route53.NewCrossAccountZoneDelegationRecord(this, jsii.String("delegate"), &CrossAccountZoneDelegationRecordProps{
	DelegatedZone: subZone,
	ParentHostedZoneName: jsii.String("someexample.com"),
	 // or you can use parentHostedZoneId
	DelegationRole: DelegationRole,
})
```

### Add Trailing Dot to Domain Names

In order to continue managing existing domain names with trailing dots using CDK, you can set `addTrailingDot: false` to prevent the Construct from adding a dot at the end of the domain name.

```go
route53.NewPublicHostedZone(this, jsii.String("HostedZone"), &PublicHostedZoneProps{
	ZoneName: jsii.String("fully.qualified.domain.com."),
	AddTrailingDot: jsii.Boolean(false),
})
```

## Imports

If you don't know the ID of the Hosted Zone to import, you can use the
`HostedZone.fromLookup`:

```go
route53.HostedZone_FromLookup(this, jsii.String("MyZone"), &HostedZoneProviderProps{
	DomainName: jsii.String("example.com"),
})
```

`HostedZone.fromLookup` requires an environment to be configured. Check
out the [documentation](https://docs.aws.amazon.com/cdk/latest/guide/environments.html) for more documentation and examples. CDK
automatically looks into your `~/.aws/config` file for the `[default]` profile.
If you want to specify a different account run `cdk deploy --profile [profile]`.

```text
new MyDevStack(app, 'dev', {
  env: {
    account: process.env.CDK_DEFAULT_ACCOUNT,
    region: process.env.CDK_DEFAULT_REGION,
  },
});
```

If you know the ID and Name of a Hosted Zone, you can import it directly:

```go
zone := route53.HostedZone_FromHostedZoneAttributes(this, jsii.String("MyZone"), &HostedZoneAttributes{
	ZoneName: jsii.String("example.com"),
	HostedZoneId: jsii.String("ZOJJZC49E0EPZ"),
})
```

Alternatively, use the `HostedZone.fromHostedZoneId` to import hosted zones if
you know the ID and the retrieval for the `zoneName` is undesirable.

```go
zone := route53.HostedZone_FromHostedZoneId(this, jsii.String("MyZone"), jsii.String("ZOJJZC49E0EPZ"))
```

You can import a Public Hosted Zone as well with the similar `PublicHostedZone.fromPublicHostedZoneId` and `PublicHostedZone.fromPublicHostedZoneAttributes` methods:

```go
zoneFromAttributes := route53.PublicHostedZone_FromPublicHostedZoneAttributes(this, jsii.String("MyZone"), &PublicHostedZoneAttributes{
	ZoneName: jsii.String("example.com"),
	HostedZoneId: jsii.String("ZOJJZC49E0EPZ"),
})

// Does not know zoneName
zoneFromId := route53.PublicHostedZone_FromPublicHostedZoneId(this, jsii.String("MyZone"), jsii.String("ZOJJZC49E0EPZ"))
```

You can use `CrossAccountZoneDelegationRecord` on imported Hosted Zones with the `grantDelegation` method:

```go
crossAccountRole := iam.NewRole(this, jsii.String("CrossAccountRole"), &RoleProps{
	// The role name must be predictable
	RoleName: jsii.String("MyDelegationRole"),
	// The other account
	AssumedBy: iam.NewAccountPrincipal(jsii.String("12345678901")),
})

zoneFromId := route53.HostedZone_FromHostedZoneId(this, jsii.String("MyZone"), jsii.String("zone-id"))
zoneFromId.GrantDelegation(crossAccountRole)

publicZoneFromId := route53.PublicHostedZone_FromPublicHostedZoneId(this, jsii.String("MyPublicZone"), jsii.String("public-zone-id"))
publicZoneFromId.GrantDelegation(crossAccountRole)

privateZoneFromId := route53.PrivateHostedZone_FromPrivateHostedZoneId(this, jsii.String("MyPrivateZone"), jsii.String("private-zone-id"))
privateZoneFromId.GrantDelegation(crossAccountRole)
```

## VPC Endpoint Service Private DNS

When you create a VPC endpoint service, AWS generates endpoint-specific DNS hostnames that consumers use to communicate with the service.
For example, vpce-1234-abcdev-us-east-1.vpce-svc-123345.us-east-1.vpce.amazonaws.com.
By default, your consumers access the service with that DNS name.
This can cause problems with HTTPS traffic because the DNS will not match the backend certificate:

```console
curl: (60) SSL: no alternative certificate subject name matches target host name 'vpce-abcdefghijklmnopq-rstuvwx.vpce-svc-abcdefghijklmnopq.us-east-1.vpce.amazonaws.com'
```

Effectively, the endpoint appears untrustworthy. To mitigate this, clients have to create an alias for this DNS name in Route53.

Private DNS for an endpoint service lets you configure a private DNS name so consumers can
access the service using an existing DNS name without creating this Route53 DNS alias
This DNS name can also be guaranteed to match up with the backend certificate.

Before consumers can use the private DNS name, you must verify that you have control of the domain/subdomain.

Assuming your account has ownership of the particular domain/subdomain,
this construct sets up the private DNS configuration on the endpoint service,
creates all the necessary Route53 entries, and verifies domain ownership.

```go
import "github.com/aws/aws-cdk-go/awscdk"


vpc := ec2.NewVpc(this, jsii.String("VPC"))
nlb := awscdk.NewNetworkLoadBalancer(this, jsii.String("NLB"), &NetworkLoadBalancerProps{
	Vpc: Vpc,
})
vpces := ec2.NewVpcEndpointService(this, jsii.String("VPCES"), &VpcEndpointServiceProps{
	VpcEndpointServiceLoadBalancers: []iVpcEndpointServiceLoadBalancer{
		nlb,
	},
})
// You must use a public hosted zone so domain ownership can be verified
zone := route53.NewPublicHostedZone(this, jsii.String("PHZ"), &PublicHostedZoneProps{
	ZoneName: jsii.String("aws-cdk.dev"),
})
route53.NewVpcEndpointServiceDomainName(this, jsii.String("EndpointDomain"), &VpcEndpointServiceDomainNameProps{
	EndpointService: vpces,
	DomainName: jsii.String("my-stuff.aws-cdk.dev"),
	PublicHostedZone: zone,
})
```
