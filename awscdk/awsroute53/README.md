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
var vpc Vpc


zone := route53.NewPrivateHostedZone(this, jsii.String("HostedZone"), &PrivateHostedZoneProps{
	ZoneName: jsii.String("fully.qualified.domain.com"),
	Vpc: Vpc,
})
```

Additional VPCs can be added with `zone.addVpc()`.

## Adding Records

To add a TXT record to your zone:

```go
var myZone HostedZone


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
var myZone HostedZone


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
var myZone HostedZone


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
var myZone HostedZone


route53.NewARecord(this, jsii.String("ARecord"), &ARecordProps{
	Zone: myZone,
	Target: route53.RecordTarget_FromIpAddresses(jsii.String("1.2.3.4"), jsii.String("5.6.7.8")),
})
```

To add an A record for an EC2 instance with an Elastic IP (EIP) to your zone:

```go
var instance Instance

var myZone HostedZone


elasticIp := ec2.NewCfnEIP(this, jsii.String("EIP"), &CfnEIPProps{
	Domain: jsii.String("vpc"),
	InstanceId: instance.InstanceId,
})
route53.NewARecord(this, jsii.String("ARecord"), &ARecordProps{
	Zone: myZone,
	Target: route53.RecordTarget_FromIpAddresses(elasticIp.ref),
})
```

To create an A record of type alias with target set to another record created outside CDK:

This function registers the given input i.e. DNS Name(string) of an existing record as an AliasTarget to the new ARecord. To register a target that is created as part of CDK use this instead.

Detailed information can be found in the [documentation](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_route53_targets-readme.html).

```go
var myZone HostedZone


targetRecord := "existing.record.cdk.local"
record := route53.ARecord_FromARecordAttributes(this, jsii.String("A"), &ARecordAttrs{
	Zone: myZone,
	RecordName: jsii.String("test"),
	TargetDNS: targetRecord,
})
```

To add an AAAA record pointing to a CloudFront distribution:

```go
import cloudfront "github.com/aws/aws-cdk-go/awscdk"

var myZone HostedZone
var distribution CloudFrontWebDistribution

route53.NewAaaaRecord(this, jsii.String("Alias"), &AaaaRecordProps{
	Zone: myZone,
	Target: route53.RecordTarget_FromAlias(targets.NewCloudFrontTarget(distribution)),
})
```

To add an HTTPS record:

```go
import cloudfront "github.com/aws/aws-cdk-go/awscdk"

var myZone HostedZone
var distribution CloudFrontWebDistribution

// Alias to CloudFront target
// Alias to CloudFront target
route53.NewHttpsRecord(this, jsii.String("HttpsRecord-CloudFrontAlias"), &HttpsRecordProps{
	Zone: myZone,
	Target: route53.RecordTarget_FromAlias(targets.NewCloudFrontTarget(distribution)),
})
// ServiceMode (priority >= 1)
// ServiceMode (priority >= 1)
route53.NewHttpsRecord(this, jsii.String("HttpsRecord-ServiceMode"), &HttpsRecordProps{
	Zone: myZone,
	Values: []HttpsRecordValue{
		route53.HttpsRecordValue_Service(&HttpsRecordServiceModeProps{
			Alpn: []Alpn{
				route53.Alpn_H3(),
				route53.Alpn_H2(),
			},
		}),
	},
})
// AliasMode (priority = 0)
// AliasMode (priority = 0)
route53.NewHttpsRecord(this, jsii.String("HttpsRecord-AliasMode"), &HttpsRecordProps{
	Zone: myZone,
	Values: []HttpsRecordValue{
		route53.HttpsRecordValue_Alias(jsii.String("service.example.com")),
	},
})
```

[Geolocation routing](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy-geo.html) can be enabled for continent, country or subdivision:

```go
var myZone HostedZone


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
var myZone HostedZone


route53.NewARecord(this, jsii.String("ARecordWeighted1"), &ARecordProps{
	Zone: myZone,
	Target: route53.RecordTarget_FromIpAddresses(jsii.String("1.2.3.4")),
	Weight: jsii.Number(10),
})
```

To enable [latency based routing](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy-latency.html), use the `region` parameter:

```go
var myZone HostedZone


route53.NewARecord(this, jsii.String("ARecordLatency1"), &ARecordProps{
	Zone: myZone,
	Target: route53.RecordTarget_FromIpAddresses(jsii.String("1.2.3.4")),
	Region: jsii.String("us-east-1"),
})
```

To enable [multivalue answer routing](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy-multivalue.html), use the `multivalueAnswer` parameter:

```go
var myZone HostedZone


route53.NewARecord(this, jsii.String("ARecordMultiValue1"), &ARecordProps{
	Zone: myZone,
	Target: route53.RecordTarget_FromIpAddresses(jsii.String("1.2.3.4")),
	MultiValueAnswer: jsii.Boolean(true),
})
```

To enable [IP-based routing](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy-ipbased.html), use the `cidrRoutingConfig` parameter:

```go
var myZone HostedZone


cidrCollection := route53.NewCfnCidrCollection(this, jsii.String("CidrCollection"), &CfnCidrCollectionProps{
	Name: jsii.String("test-collection"),
	Locations: []interface{}{
		&LocationProperty{
			CidrList: []*string{
				jsii.String("192.168.1.0/24"),
			},
			LocationName: jsii.String("my_location"),
		},
	},
})

route53.NewARecord(this, jsii.String("CidrRoutingConfig"), &ARecordProps{
	Zone: myZone,
	Target: route53.RecordTarget_FromIpAddresses(jsii.String("1.2.3.4")),
	SetIdentifier: jsii.String("test"),
	CidrRoutingConfig: route53.CidrRoutingConfig_Create(&CidrRoutingConfigProps{
		CollectionId: cidrCollection.AttrId,
		LocationName: jsii.String("test_location"),
	}),
})
```

To use the default CIDR record, call the `route53.CidrRoutingConfig.default`. This sets the `locationName` to `*`. The `collectionId` is still required.

```go
var myZone HostedZone


cidrCollection := route53.NewCfnCidrCollection(this, jsii.String("CidrCollection"), &CfnCidrCollectionProps{
	Name: jsii.String("test-collection"),
	Locations: []interface{}{
		&LocationProperty{
			CidrList: []*string{
				jsii.String("192.168.1.0/24"),
			},
			LocationName: jsii.String("my_location"),
		},
	},
})

route53.NewARecord(this, jsii.String("DefaultCidrRoutingConfig"), &ARecordProps{
	Zone: myZone,
	Target: route53.RecordTarget_FromIpAddresses(jsii.String("5.6.7.8")),
	SetIdentifier: jsii.String("default"),
	CidrRoutingConfig: route53.CidrRoutingConfig_WithDefaultLocationName(cidrCollection.AttrId),
})
```

To specify a unique identifier to differentiate among multiple resource record sets that have the same combination of name and type, use the `setIdentifier` parameter:

```go
var myZone HostedZone


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

### Health Checks

See the [Route 53 Health Checks documentation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthcheckconfig.html#cfn-route53-healthcheck-healthcheckconfig-type) for possible types of health checks.

Route 53 has the ability to monitor the health of your application and only return records for healthy endpoints.
This is done using a `HealthCheck` construct.

In the following example, the `ARecord` will be returned by Route 53 in response to DNS queries only if the HTTP requests to the `example.com/health` endpoint return a 2XX or 3XX status code.

In case, when the endpoint is not healthy, the `ARecord2` will be returned by Route 53 in response to DNS queries.

```go
var myZone HostedZone


healthCheck := route53.NewHealthCheck(this, jsii.String("HealthCheck"), &HealthCheckProps{
	Type: route53.HealthCheckType_HTTP,
	Fqdn: jsii.String("example.com"),
	Port: jsii.Number(80),
	ResourcePath: jsii.String("/health"),
	FailureThreshold: jsii.Number(3),
	RequestInterval: awscdk.Duration_Seconds(jsii.Number(30)),
})

route53.NewARecord(this, jsii.String("ARecord"), &ARecordProps{
	Zone: myZone,
	Target: route53.RecordTarget_FromIpAddresses(jsii.String("1.2.3.4")),
	HealthCheck: HealthCheck,
	Weight: jsii.Number(100),
})
route53.NewARecord(this, jsii.String("ARecord2"), &ARecordProps{
	Zone: myZone,
	Target: route53.RecordTarget_*FromIpAddresses(jsii.String("5.6.7.8")),
	Weight: jsii.Number(0),
})
```

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
var myZone HostedZone


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
})
parentZone.GrantDelegation(crossAccountRole)
```

To restrict the records that can be created with the delegation IAM role, use the optional `delegatedZoneNames` property in the delegation options,
which enforces the `route53:ChangeResourceRecordSetsNormalizedRecordNames` condition key for record names that match those hosted zone names.
The `delegatedZoneNames` list may only consist of hosted zones names that are subzones of the parent hosted zone.

If the delegated zone name contains an unresolved token,
it must resolve to a zone name that satisfies the requirements according to the documentation:
https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/specifying-conditions-route53.html#route53_rrset_conditionkeys_normalization

> All letters must be lowercase.
> The DNS name must be without the trailing dot.
> Characters other than a–z, 0–9, - (hyphen), _ (underscore), and . (period, as a delimiter between labels) must use escape codes in the format \three-digit octal code. For example, \052 is the octal code for character *.

This feature allows you to better follow the minimum permissions privilege principle:

```go
var betaCrossAccountRole Role

var prodCrossAccountRole Role
parentZone := route53.NewPublicHostedZone(this, jsii.String("HostedZone"), &PublicHostedZoneProps{
	ZoneName: jsii.String("someexample.com"),
})
parentZone.GrantDelegation(betaCrossAccountRole, &GrantDelegationOptions{
	DelegatedZoneNames: []*string{
		jsii.String("beta.someexample.com"),
	},
})
parentZone.GrantDelegation(prodCrossAccountRole, &GrantDelegationOptions{
	DelegatedZoneNames: []*string{
		jsii.String("prod.someexample.com"),
	},
})
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

Delegating the hosted zone requires assuming a role in the parent hosted zone's account.
In order for the assumed credentials to be valid, the resource must assume the role using
an STS endpoint in a region where both the subdomain's account and the parent's account
are opted-in. By default, this region is determined automatically, but if you need to
change the region used for the AssumeRole call, specify `assumeRoleRegion`:

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

route53.NewCrossAccountZoneDelegationRecord(this, jsii.String("delegate"), &CrossAccountZoneDelegationRecordProps{
	DelegatedZone: subZone,
	ParentHostedZoneName: jsii.String("someexample.com"),
	 // or you can use parentHostedZoneId
	DelegationRole: DelegationRole,
	AssumeRoleRegion: jsii.String("us-east-1"),
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

## Enabling DNSSEC

DNSSEC can be enabled for Hosted Zones. For detailed information, see
[Configuring DNSSEC signing in Amazon Route 53](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-configuring-dnssec.html).

Enabling DNSSEC requires an asymmetric KMS Customer-Managed Key using the `ECC_NIST_P256` key spec.
Additionally, that KMS key must be in `us-east-1`.

```go
kmsKey := kms.NewKey(this, jsii.String("KmsCMK"), &KeyProps{
	KeySpec: kms.KeySpec_ECC_NIST_P256,
	KeyUsage: kms.KeyUsage_SIGN_VERIFY,
})
hostedZone := route53.NewHostedZone(this, jsii.String("HostedZone"), &HostedZoneProps{
	ZoneName: jsii.String("example.com"),
})
// Enable DNSSEC signing for the zone
hostedZone.EnableDnssec(&ZoneSigningOptions{
	KmsKey: KmsKey,
})
```

The necessary permissions for Route 53 to use the key will automatically be added when using
this configuration. If it is necessary to create a key signing key manually, that can be done
using the `KeySigningKey` construct:

```go
var hostedZone HostedZone
var kmsKey Key

route53.NewKeySigningKey(this, jsii.String("KeySigningKey"), &KeySigningKeyProps{
	HostedZone: HostedZone,
	KmsKey: KmsKey,
	KeySigningKeyName: jsii.String("ksk"),
	Status: route53.KeySigningKeyStatus_ACTIVE,
})
```

When directly constructing the `KeySigningKey` resource, enabling DNSSEC signing for the hosted
zone will be need to be done explicitly (either using the `CfnDNSSEC` construct or via another
means).

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
Note that any records created with a hosted zone obtained this way must have their name be fully qualified

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

You can import a Private Hosted Zone with `PrivateHostedZone.fromPrivateHostedZoneId` and `PrivateHostedZone.fromPrivateHostedZoneAttributes` methods:

```go
privateZoneFromAttributes := route53.PrivateHostedZone_FromPrivateHostedZoneAttributes(this, jsii.String("MyPrivateZone"), &PrivateHostedZoneAttributes{
	ZoneName: jsii.String("example.local"),
	HostedZoneId: jsii.String("ZOJJZC49E0EPZ"),
})

// Does not know zoneName
privateZoneFromId := route53.PrivateHostedZone_FromPrivateHostedZoneId(this, jsii.String("MyPrivateZone"), jsii.String("ZOJJZC49E0EPZ"))
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
	VpcEndpointServiceLoadBalancers: []IVpcEndpointServiceLoadBalancer{
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
