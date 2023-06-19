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

Constructs are available for A, AAAA, CAA, CNAME, MX, NS, SRV and TXT records.

Use the `CaaAmazonRecord` construct to easily restrict certificate authorities
allowed to issue certificates for a domain to Amazon only.

To add a NS record to a HostedZone in different account you can do the following:

In the account containing the parent hosted zone:

```go
parentZone := route53.NewPublicHostedZone(this, jsii.String("HostedZone"), &PublicHostedZoneProps{
	ZoneName: jsii.String("someexample.com"),
	CrossAccountZoneDelegationPrincipal: iam.NewAccountPrincipal(jsii.String("12345678901")),
	CrossAccountZoneDelegationRoleName: jsii.String("MyDelegationRole"),
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

You can import a Public Hosted Zone as well with the similar `PubicHostedZone.fromPublicHostedZoneId` and `PubicHostedZone.fromPublicHostedZoneAttributes` methods:

```go
zoneFromAttributes := route53.PublicHostedZone_FromPublicHostedZoneAttributes(this, jsii.String("MyZone"), &PublicHostedZoneAttributes{
	ZoneName: jsii.String("example.com"),
	HostedZoneId: jsii.String("ZOJJZC49E0EPZ"),
})

// Does not know zoneName
zoneFromId := route53.PublicHostedZone_FromPublicHostedZoneId(this, jsii.String("MyZone"), jsii.String("ZOJJZC49E0EPZ"))
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
nlb := awscdk.NewNetworkLoadBalancer(stack, jsii.String("NLB"), &NetworkLoadBalancerProps{
	Vpc: Vpc,
})
vpces := awscdk.NewVpcEndpointService(stack, jsii.String("VPCES"), &VpcEndpointServiceProps{
	VpcEndpointServiceLoadBalancers: []iVpcEndpointServiceLoadBalancer{
		nlb,
	},
})
// You must use a public hosted zone so domain ownership can be verified
zone := awscdk.NewPublicHostedZone(stack, jsii.String("PHZ"), &PublicHostedZoneProps{
	ZoneName: jsii.String("aws-cdk.dev"),
})
awscdk.NewVpcEndpointServiceDomainName(stack, jsii.String("EndpointDomain"), &VpcEndpointServiceDomainNameProps{
	EndpointService: vpces,
	DomainName: jsii.String("my-stuff.aws-cdk.dev"),
	PublicHostedZone: zone,
})
```
