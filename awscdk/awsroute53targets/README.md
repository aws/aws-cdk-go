# Route53 Alias Record Targets for the CDK Route53 Library

This library contains Route53 Alias Record targets for:

* API Gateway custom domains

```go
import apigw "github.com/aws/aws-cdk-go/awscdk"

var zone HostedZone
var restApi LambdaRestApi


route53.NewARecord(this, jsii.String("AliasRecord"), &ARecordProps{
	Zone: Zone,
	Target: route53.RecordTarget_FromAlias(targets.NewApiGateway(restApi)),
})
```

* API Gateway V2 custom domains

```go
import apigwv2 "github.com/aws/aws-cdk-go/awscdk"

var zone HostedZone
var domainName DomainName


route53.NewARecord(this, jsii.String("AliasRecord"), &ARecordProps{
	Zone: Zone,
	Target: route53.RecordTarget_FromAlias(targets.NewApiGatewayv2DomainProperties(domainName.regionalDomainName, domainName.regionalHostedZoneId)),
})
```

* AppSync custom domains

```go
import appsync "github.com/aws/aws-cdk-go/awscdk"

var zone HostedZone
var graphqlApi GraphqlApi


route53.NewARecord(this, jsii.String("AliasRecord"), &ARecordProps{
	Zone: Zone,
	Target: route53.RecordTarget_FromAlias(targets.NewAppSyncTarget(graphqlApi)),
})
```

* CloudFront distributions

```go
import cloudfront "github.com/aws/aws-cdk-go/awscdk"

var zone HostedZone
var distribution CloudFrontWebDistribution


route53.NewARecord(this, jsii.String("AliasRecord"), &ARecordProps{
	Zone: Zone,
	Target: route53.RecordTarget_FromAlias(targets.NewCloudFrontTarget(distribution)),
})
```

* ELBv2 load balancers

By providing optional properties, you can specify whether to evaluate target health.

```go
import elbv2 "github.com/aws/aws-cdk-go/awscdk"

var zone HostedZone
var lb ApplicationLoadBalancer


route53.NewARecord(this, jsii.String("AliasRecord"), &ARecordProps{
	Zone: Zone,
	Target: route53.RecordTarget_FromAlias(
	targets.NewLoadBalancerTarget(lb, map[string]*bool{
		"evaluateTargetHealth": jsii.Boolean(true),
	})),
})
```

* Classic load balancers

By providing optional properties, you can specify whether to evaluate target health.

```go
import elb "github.com/aws/aws-cdk-go/awscdk"

var zone HostedZone
var lb LoadBalancer


route53.NewARecord(this, jsii.String("AliasRecord"), &ARecordProps{
	Zone: Zone,
	Target: route53.RecordTarget_FromAlias(
	targets.NewClassicLoadBalancerTarget(lb, map[string]*bool{
		"evaluateTargetHealth": jsii.Boolean(true),
	})),
})
```

**Important:** Based on [AWS documentation](https://aws.amazon.com/de/premiumsupport/knowledge-center/alias-resource-record-set-route53-cli/), all alias record in Route 53 that points to a Elastic Load Balancer will always include *dualstack* for the DNSName to resolve IPv4/IPv6 addresses (without *dualstack* IPv6 will not resolve).

For example, if the Amazon-provided DNS for the load balancer is `ALB-xxxxxxx.us-west-2.elb.amazonaws.com`, CDK will create alias target in Route 53 will be `dualstack.ALB-xxxxxxx.us-west-2.elb.amazonaws.com`.

* GlobalAccelerator

By providing optional properties, you can specify whether to evaluate target health.

```go
import globalaccelerator "github.com/aws/aws-cdk-go/awscdk"

var zone HostedZone
var accelerator Accelerator


route53.NewARecord(this, jsii.String("AliasRecord"), &ARecordProps{
	Zone: Zone,
	Target: route53.RecordTarget_FromAlias(
	targets.NewGlobalAcceleratorTarget(accelerator, map[string]*bool{
		"evaluateTargetHealth": jsii.Boolean(true),
	})),
})
```

**Important:** If you use GlobalAcceleratorDomainTarget, passing a string rather than an instance of IAccelerator, ensure that the string is a valid domain name of an existing Global Accelerator instance.
See [the documentation on DNS addressing](https://docs.aws.amazon.com/global-accelerator/latest/dg/dns-addressing-custom-domains.dns-addressing.html) with Global Accelerator for more info.

* InterfaceVpcEndpoints

**Important:** Based on the CFN docs for VPCEndpoints - [see here](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#aws-resource-ec2-vpcendpoint-return-values) - the attributes returned for DnsEntries in CloudFormation is a combination of the hosted zone ID and the DNS name. The entries are ordered as follows: regional public DNS, zonal public DNS, private DNS, and wildcard DNS. This order is not enforced for AWS Marketplace services, and therefore this CDK construct is ONLY guaranteed to work with non-marketplace services.

```go
import ec2 "github.com/aws/aws-cdk-go/awscdk"

var zone HostedZone
var interfaceVpcEndpoint InterfaceVpcEndpoint


route53.NewARecord(this, jsii.String("AliasRecord"), &ARecordProps{
	Zone: Zone,
	Target: route53.RecordTarget_FromAlias(targets.NewInterfaceVpcEndpointTarget(interfaceVpcEndpoint)),
})
```

* S3 Bucket Website:

**Important:** The Bucket name must strictly match the full DNS name.
See [the Developer Guide](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/getting-started.html) for more info.

By providing optional properties, you can specify whether to evaluate target health.

```go
import s3 "github.com/aws/aws-cdk-go/awscdk"


recordName := "www"
domainName := "example.com"

bucketWebsite := s3.NewBucket(this, jsii.String("BucketWebsite"), &BucketProps{
	BucketName: []*string{
		recordName,
		domainName,
	}.join(jsii.String(".")),
	 // www.example.com
	PublicReadAccess: jsii.Boolean(true),
	WebsiteIndexDocument: jsii.String("index.html"),
})

zone := route53.HostedZone_FromLookup(this, jsii.String("Zone"), &HostedZoneProviderProps{
	DomainName: jsii.String(DomainName),
}) // example.com

 // example.com
route53.NewARecord(this, jsii.String("AliasRecord"), &ARecordProps{
	Zone: Zone,
	RecordName: jsii.String(RecordName),
	 // www
	Target: route53.RecordTarget_FromAlias(
	targets.NewBucketWebsiteTarget(bucketWebsite, map[string]*bool{
		"evaluateTargetHealth": jsii.Boolean(true),
	})),
})
```

* User pool domain

```go
import cognito "github.com/aws/aws-cdk-go/awscdk"

var zone HostedZone
var domain UserPoolDomain

route53.NewARecord(this, jsii.String("AliasRecord"), &ARecordProps{
	Zone: Zone,
	Target: route53.RecordTarget_FromAlias(targets.NewUserPoolDomainTarget(domain)),
})
```

* Route 53 record

```go
var zone HostedZone
var record ARecord

route53.NewARecord(this, jsii.String("AliasRecord"), &ARecordProps{
	Zone: Zone,
	Target: route53.RecordTarget_FromAlias(targets.NewRoute53RecordTarget(record)),
})
```

* Elastic Beanstalk environment:

**Important:** Only supports Elastic Beanstalk environments created after 2016 that have a regional endpoint.

By providing optional properties, you can specify whether to evaluate target health.

```go
var zone HostedZone
var ebsEnvironmentUrl string


route53.NewARecord(this, jsii.String("AliasRecord"), &ARecordProps{
	Zone: Zone,
	Target: route53.RecordTarget_FromAlias(
	targets.NewElasticBeanstalkEnvironmentEndpointTarget(ebsEnvironmentUrl, map[string]*bool{
		"evaluateTargetHealth": jsii.Boolean(true),
	})),
})
```

If Elastic Beanstalk environment URL is not avaiable at synth time, you can specify Hosted Zone ID of the target

```go
import "github.com/aws/aws-cdk-go/awscdk"

var zone HostedZone
var ebsEnvironmentUrl string


route53.NewARecord(this, jsii.String("AliasRecord"), &ARecordProps{
	Zone: Zone,
	Target: route53.RecordTarget_FromAlias(
	targets.NewElasticBeanstalkEnvironmentEndpointTarget(ebsEnvironmentUrl, map[string]*string{
		"hostedZoneId": awscdk.RegionInfo_get(jsii.String("us-east-1")).ebsEnvEndpointHostedZoneId,
	})),
})
```

Or you can specify Stack region for CDK to generate the correct Hosted Zone ID.

```go
import "github.com/aws/aws-cdk-go/awscdk"

var app App
var zone HostedZone
var ebsEnvironmentUrl string


stack := awscdk.Newstack(app, jsii.String("my-stack"), &StackProps{
	Env: &Environment{
		Region: jsii.String("us-east-1"),
	},
})

route53.NewARecord(stack, jsii.String("AliasRecord"), &ARecordProps{
	Zone: Zone,
	Target: route53.RecordTarget_FromAlias(
	targets.NewElasticBeanstalkEnvironmentEndpointTarget(ebsEnvironmentUrl)),
})
```

See the documentation of `aws-cdk-lib/aws-route53` for more information.
