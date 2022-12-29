# Amazon Route53 Construct Library

To add a public hosted zone:

```go
route53.NewPublicHostedZone(this, jsii.String("HostedZone"), &publicHostedZoneProps{
	zoneName: jsii.String("fully.qualified.domain.com"),
})
```

To add a private hosted zone, use `PrivateHostedZone`. Note that
`enableDnsHostnames` and `enableDnsSupport` must have been enabled for the
VPC you're configuring for private hosted zones.

```go
var vpc vpc


zone := route53.NewPrivateHostedZone(this, jsii.String("HostedZone"), &privateHostedZoneProps{
	zoneName: jsii.String("fully.qualified.domain.com"),
	vpc: vpc,
})
```

Additional VPCs can be added with `zone.addVpc()`.

## Adding Records

To add a TXT record to your zone:

```go
var myZone hostedZone


route53.NewTxtRecord(this, jsii.String("TXTRecord"), &txtRecordProps{
	zone: myZone,
	recordName: jsii.String("_foo"),
	 // If the name ends with a ".", it will be used as-is;
	// if it ends with a "." followed by the zone name, a trailing "." will be added automatically;
	// otherwise, a ".", the zone name, and a trailing "." will be added automatically.
	// Defaults to zone root if not specified.
	values: []*string{
		jsii.String("Bar!"),
		jsii.String("Baz?"),
	},
	ttl: awscdk.Duration.minutes(jsii.Number(90)),
})
```

To add a NS record to your zone:

```go
var myZone hostedZone


route53.NewNsRecord(this, jsii.String("NSRecord"), &nsRecordProps{
	zone: myZone,
	recordName: jsii.String("foo"),
	values: []*string{
		jsii.String("ns-1.awsdns.co.uk."),
		jsii.String("ns-2.awsdns.com."),
	},
	ttl: awscdk.Duration.minutes(jsii.Number(90)),
})
```

To add a DS record to your zone:

```go
var myZone hostedZone


route53.NewDsRecord(this, jsii.String("DSRecord"), &dsRecordProps{
	zone: myZone,
	recordName: jsii.String("foo"),
	values: []*string{
		jsii.String("12345 3 1 123456789abcdef67890123456789abcdef67890"),
	},
	ttl: awscdk.Duration.minutes(jsii.Number(90)),
})
```

To add an A record to your zone:

```go
var myZone hostedZone


route53.NewARecord(this, jsii.String("ARecord"), &aRecordProps{
	zone: myZone,
	target: route53.recordTarget.fromIpAddresses(jsii.String("1.2.3.4"), jsii.String("5.6.7.8")),
})
```

To add an A record for an EC2 instance with an Elastic IP (EIP) to your zone:

```go
var instance instance

var myZone hostedZone


elasticIp := ec2.NewCfnEIP(this, jsii.String("EIP"), &cfnEIPProps{
	domain: jsii.String("vpc"),
	instanceId: instance.instanceId,
})
route53.NewARecord(this, jsii.String("ARecord"), &aRecordProps{
	zone: myZone,
	target: route53.recordTarget.fromIpAddresses(elasticIp.ref),
})
```

To add an AAAA record pointing to a CloudFront distribution:

```go
import cloudfront "github.com/aws/aws-cdk-go/awscdk"

var myZone hostedZone
var distribution cloudFrontWebDistribution

route53.NewAaaaRecord(this, jsii.String("Alias"), &aaaaRecordProps{
	zone: myZone,
	target: route53.recordTarget.fromAlias(targets.NewCloudFrontTarget(distribution)),
})
```

Constructs are available for A, AAAA, CAA, CNAME, MX, NS, SRV and TXT records.

Use the `CaaAmazonRecord` construct to easily restrict certificate authorities
allowed to issue certificates for a domain to Amazon only.

### Working with existing record sets

Use the `deleteExisting` prop to delete an existing record set before deploying the new one.
This is useful if you want to minimize downtime and avoid "manual" actions while deploying a
stack with a record set that already exists. This is typically the case for record sets that
are not already "owned" by CloudFormation or "owned" by another stack or construct that is
going to be deleted (migration).

```go
var myZone hostedZone


route53.NewARecord(this, jsii.String("ARecord"), &aRecordProps{
	zone: myZone,
	target: route53.recordTarget.fromIpAddresses(jsii.String("1.2.3.4"), jsii.String("5.6.7.8")),
	deleteExisting: jsii.Boolean(true),
})
```

### Cross Account Zone Delegation

If you want to have your root domain hosted zone in one account and your subdomain hosted
zone in a diferent one, you can use `CrossAccountZoneDelegationRecord` to set up delegation
between them.

In the account containing the parent hosted zone:

```go
parentZone := route53.NewPublicHostedZone(this, jsii.String("HostedZone"), &publicHostedZoneProps{
	zoneName: jsii.String("someexample.com"),
})
crossAccountRole := iam.NewRole(this, jsii.String("CrossAccountRole"), &roleProps{
	// The role name must be predictable
	roleName: jsii.String("MyDelegationRole"),
	// The other account
	assumedBy: iam.NewAccountPrincipal(jsii.String("12345678901")),
})
parentZone.grantDelegation(crossAccountRole)
```

In the account containing the child zone to be delegated:

```go
subZone := route53.NewPublicHostedZone(this, jsii.String("SubZone"), &publicHostedZoneProps{
	zoneName: jsii.String("sub.someexample.com"),
})

// import the delegation role by constructing the roleArn
delegationRoleArn := awscdk.stack.of(this).formatArn(&arnComponents{
	region: jsii.String(""),
	 // IAM is global in each partition
	service: jsii.String("iam"),
	account: jsii.String("parent-account-id"),
	resource: jsii.String("role"),
	resourceName: jsii.String("MyDelegationRole"),
})
delegationRole := iam.role.fromRoleArn(this, jsii.String("DelegationRole"), delegationRoleArn)

// create the record
// create the record
route53.NewCrossAccountZoneDelegationRecord(this, jsii.String("delegate"), &crossAccountZoneDelegationRecordProps{
	delegatedZone: subZone,
	parentHostedZoneName: jsii.String("someexample.com"),
	 // or you can use parentHostedZoneId
	delegationRole: delegationRole,
})
```

## Imports

If you don't know the ID of the Hosted Zone to import, you can use the
`HostedZone.fromLookup`:

```go
route53.hostedZone.fromLookup(this, jsii.String("MyZone"), &hostedZoneProviderProps{
	domainName: jsii.String("example.com"),
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
zone := route53.hostedZone.fromHostedZoneAttributes(this, jsii.String("MyZone"), &hostedZoneAttributes{
	zoneName: jsii.String("example.com"),
	hostedZoneId: jsii.String("ZOJJZC49E0EPZ"),
})
```

Alternatively, use the `HostedZone.fromHostedZoneId` to import hosted zones if
you know the ID and the retrieval for the `zoneName` is undesirable.

```go
zone := route53.hostedZone.fromHostedZoneId(this, jsii.String("MyZone"), jsii.String("ZOJJZC49E0EPZ"))
```

You can import a Public Hosted Zone as well with the similar `PublicHostedZone.fromPublicHostedZoneId` and `PublicHostedZone.fromPublicHostedZoneAttributes` methods:

```go
zoneFromAttributes := route53.publicHostedZone.fromPublicHostedZoneAttributes(this, jsii.String("MyZone"), &publicHostedZoneAttributes{
	zoneName: jsii.String("example.com"),
	hostedZoneId: jsii.String("ZOJJZC49E0EPZ"),
})

// Does not know zoneName
zoneFromId := route53.publicHostedZone.fromPublicHostedZoneId(this, jsii.String("MyZone"), jsii.String("ZOJJZC49E0EPZ"))
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
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

stack := awscdk.NewStack()
vpc := awscdk.NewVpc(stack, jsii.String("VPC"))
nlb := awscdk.NewNetworkLoadBalancer(stack, jsii.String("NLB"), &networkLoadBalancerProps{
	vpc: vpc,
})
vpces := awscdk.NewVpcEndpointService(stack, jsii.String("VPCES"), &vpcEndpointServiceProps{
	vpcEndpointServiceLoadBalancers: []iVpcEndpointServiceLoadBalancer{
		nlb,
	},
})
// You must use a public hosted zone so domain ownership can be verified
zone := awscdk.NewPublicHostedZone(stack, jsii.String("PHZ"), &publicHostedZoneProps{
	zoneName: jsii.String("aws-cdk.dev"),
})
awscdk.NewVpcEndpointServiceDomainName(stack, jsii.String("EndpointDomain"), &vpcEndpointServiceDomainNameProps{
	endpointService: vpces,
	domainName: jsii.String("my-stuff.aws-cdk.dev"),
	publicHostedZone: zone,
})
```
