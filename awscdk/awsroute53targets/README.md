# Route53 Alias Record Targets for the CDK Route53 Library

This library contains Route53 Alias Record targets for:

* API Gateway custom domains

  ```go
  import apigw "github.com/aws/aws-cdk-go/awscdk"

  var zone hostedZone
  var restApi lambdaRestApi


  route53.NewARecord(this, jsii.String("AliasRecord"), &aRecordProps{
  	zone: zone,
  	target: route53.recordTarget.fromAlias(targets.NewApiGateway(restApi)),
  })
  ```
* API Gateway V2 custom domains

  ```go
  // Example automatically generated from non-compiling source. May contain errors.
  import apigwv2 "github.com/aws/aws-cdk-go/awscdk"

  var zone hostedZone
  var domainName apigwv2.DomainName


  route53.NewARecord(this, jsii.String("AliasRecord"), &aRecordProps{
  	zone: zone,
  	target: route53.recordTarget.fromAlias(targets.NewApiGatewayv2DomainProperties(domainName.regionalDomainName, domainName.regionalHostedZoneId)),
  })
  ```
* CloudFront distributions

  ```go
  import cloudfront "github.com/aws/aws-cdk-go/awscdk"

  var zone hostedZone
  var distribution cloudFrontWebDistribution


  route53.NewARecord(this, jsii.String("AliasRecord"), &aRecordProps{
  	zone: zone,
  	target: route53.recordTarget.fromAlias(targets.NewCloudFrontTarget(distribution)),
  })
  ```
* ELBv2 load balancers

  ```go
  import elbv2 "github.com/aws/aws-cdk-go/awscdk"

  var zone hostedZone
  var lb applicationLoadBalancer


  route53.NewARecord(this, jsii.String("AliasRecord"), &aRecordProps{
  	zone: zone,
  	target: route53.recordTarget.fromAlias(targets.NewLoadBalancerTarget(lb)),
  })
  ```
* Classic load balancers

  ```go
  import elb "github.com/aws/aws-cdk-go/awscdk"

  var zone hostedZone
  var lb loadBalancer


  route53.NewARecord(this, jsii.String("AliasRecord"), &aRecordProps{
  	zone: zone,
  	target: route53.recordTarget.fromAlias(targets.NewClassicLoadBalancerTarget(lb)),
  })
  ```

**Important:** Based on [AWS documentation](https://aws.amazon.com/de/premiumsupport/knowledge-center/alias-resource-record-set-route53-cli/), all alias record in Route 53 that points to a Elastic Load Balancer will always include *dualstack* for the DNSName to resolve IPv4/IPv6 addresses (without *dualstack* IPv6 will not resolve).

For example, if the Amazon-provided DNS for the load balancer is `ALB-xxxxxxx.us-west-2.elb.amazonaws.com`, CDK will create alias target in Route 53 will be `dualstack.ALB-xxxxxxx.us-west-2.elb.amazonaws.com`.

* GlobalAccelerator

  ```go
  import globalaccelerator "github.com/aws/aws-cdk-go/awscdk"

  var zone hostedZone
  var accelerator accelerator


  route53.NewARecord(this, jsii.String("AliasRecord"), &aRecordProps{
  	zone: zone,
  	target: route53.recordTarget.fromAlias(targets.NewGlobalAcceleratorTarget(accelerator)),
  })
  ```

**Important:** If you use GlobalAcceleratorDomainTarget, passing a string rather than an instance of IAccelerator, ensure that the string is a valid domain name of an existing Global Accelerator instance.
See [the documentation on DNS addressing](https://docs.aws.amazon.com/global-accelerator/latest/dg/dns-addressing-custom-domains.dns-addressing.html) with Global Accelerator for more info.

* InterfaceVpcEndpoints

**Important:** Based on the CFN docs for VPCEndpoints - [see here](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#aws-resource-ec2-vpcendpoint-return-values) - the attributes returned for DnsEntries in CloudFormation is a combination of the hosted zone ID and the DNS name. The entries are ordered as follows: regional public DNS, zonal public DNS, private DNS, and wildcard DNS. This order is not enforced for AWS Marketplace services, and therefore this CDK construct is ONLY guaranteed to work with non-marketplace services.

```go
import ec2 "github.com/aws/aws-cdk-go/awscdk"

var zone hostedZone
var interfaceVpcEndpoint interfaceVpcEndpoint


route53.NewARecord(this, jsii.String("AliasRecord"), &aRecordProps{
	zone: zone,
	target: route53.recordTarget.fromAlias(targets.NewInterfaceVpcEndpointTarget(interfaceVpcEndpoint)),
})
```

* S3 Bucket Website:

**Important:** The Bucket name must strictly match the full DNS name.
See [the Developer Guide](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/getting-started.html) for more info.

```go
import s3 "github.com/aws/aws-cdk-go/awscdk"


recordName := "www"
domainName := "example.com"

bucketWebsite := s3.NewBucket(this, jsii.String("BucketWebsite"), &bucketProps{
	bucketName: []*string{
		recordName,
		domainName,
	}.join(jsii.String(".")),
	 // www.example.com
	publicReadAccess: jsii.Boolean(true),
	websiteIndexDocument: jsii.String("index.html"),
})

zone := route53.hostedZone.fromLookup(this, jsii.String("Zone"), &hostedZoneProviderProps{
	domainName: jsii.String(domainName),
}) // example.com

 // example.com
route53.NewARecord(this, jsii.String("AliasRecord"), &aRecordProps{
	zone: zone,
	recordName: jsii.String(recordName),
	 // www
	target: route53.recordTarget.fromAlias(targets.NewBucketWebsiteTarget(bucketWebsite)),
})
```

* User pool domain

  ```go
  import cognito "github.com/aws/aws-cdk-go/awscdk"

  var zone hostedZone
  var domain userPoolDomain

  route53.NewARecord(this, jsii.String("AliasRecord"), &aRecordProps{
  	zone: zone,
  	target: route53.recordTarget.fromAlias(targets.NewUserPoolDomainTarget(domain)),
  })
  ```
* Route 53 record

  ```go
  var zone hostedZone
  var record aRecord

  route53.NewARecord(this, jsii.String("AliasRecord"), &aRecordProps{
  	zone: zone,
  	target: route53.recordTarget.fromAlias(targets.NewRoute53RecordTarget(record)),
  })
  ```
* Elastic Beanstalk environment:

**Important:** Only supports Elastic Beanstalk environments created after 2016 that have a regional endpoint.

```go
var zone hostedZone
var ebsEnvironmentUrl string


route53.NewARecord(this, jsii.String("AliasRecord"), &aRecordProps{
	zone: zone,
	target: route53.recordTarget.fromAlias(targets.NewElasticBeanstalkEnvironmentEndpointTarget(ebsEnvironmentUrl)),
})
```

See the documentation of `@aws-cdk/aws-route53` for more information.
