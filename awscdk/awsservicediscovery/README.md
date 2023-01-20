# Amazon ECS Service Discovery Construct Library

This module is part of the [AWS Cloud Development Kit](https://github.com/aws/aws-cdk) project.

This package contains constructs for working with **AWS Cloud Map**

AWS Cloud Map is a fully managed service that you can use to create and
maintain a map of the backend services and resources that your applications
depend on.

For further information on AWS Cloud Map,
see the [AWS Cloud Map documentation](https://docs.aws.amazon.com/cloud-map)

## HTTP Namespace Example

The following example creates an AWS Cloud Map namespace that
supports API calls, creates a service in that namespace, and
registers an instance to it:

```go
import cdk "github.com/aws/aws-cdk-go/awscdk"
import servicediscovery "github.com/aws/aws-cdk-go/awscdk"

app := cdk.NewApp()
stack := cdk.NewStack(app, jsii.String("aws-servicediscovery-integ"))

namespace := servicediscovery.NewHttpNamespace(stack, jsii.String("MyNamespace"), &httpNamespaceProps{
	name: jsii.String("MyHTTPNamespace"),
})

service1 := namespace.createService(jsii.String("NonIpService"), &baseServiceProps{
	description: jsii.String("service registering non-ip instances"),
})

service1.registerNonIpInstance(jsii.String("NonIpInstance"), &nonIpInstanceBaseProps{
	customAttributes: map[string]*string{
		"arn": jsii.String("arn:aws:s3:::mybucket"),
	},
})

service2 := namespace.createService(jsii.String("IpService"), &baseServiceProps{
	description: jsii.String("service registering ip instances"),
	healthCheck: &healthCheckConfig{
		type: servicediscovery.healthCheckType_HTTP,
		resourcePath: jsii.String("/check"),
	},
})

service2.registerIpInstance(jsii.String("IpInstance"), &ipInstanceBaseProps{
	ipv4: jsii.String("54.239.25.192"),
})

app.synth()
```

## Private DNS Namespace Example

The following example creates an AWS Cloud Map namespace that
supports both API calls and DNS queries within a vpc, creates a
service in that namespace, and registers a loadbalancer as an
instance.

A secondary service is also configured which only supports API based discovery, a
non ip based resource is registered to this service:

```go
import ec2 "github.com/aws/aws-cdk-go/awscdk"
import elbv2 "github.com/aws/aws-cdk-go/awscdk"
import cdk "github.com/aws/aws-cdk-go/awscdk"
import servicediscovery "github.com/aws/aws-cdk-go/awscdk"

app := cdk.NewApp()
stack := cdk.NewStack(app, jsii.String("aws-servicediscovery-integ"))

vpc := ec2.NewVpc(stack, jsii.String("Vpc"), &vpcProps{
	maxAzs: jsii.Number(2),
})

namespace := servicediscovery.NewPrivateDnsNamespace(stack, jsii.String("Namespace"), &privateDnsNamespaceProps{
	name: jsii.String("boobar.com"),
	vpc: vpc,
})

service := namespace.createService(jsii.String("Service"), &dnsServiceProps{
	dnsRecordType: servicediscovery.dnsRecordType_A_AAAA,
	dnsTtl: cdk.duration.seconds(jsii.Number(30)),
	loadBalancer: jsii.Boolean(true),
})

loadbalancer := elbv2.NewApplicationLoadBalancer(stack, jsii.String("LB"), &applicationLoadBalancerProps{
	vpc: vpc,
	internetFacing: jsii.Boolean(true),
})

service.registerLoadBalancer(jsii.String("Loadbalancer"), loadbalancer)

arnService := namespace.createService(jsii.String("ArnService"), &dnsServiceProps{
	discoveryType: servicediscovery.discoveryType_API,
})

arnService.registerNonIpInstance(jsii.String("NonIpInstance"), &nonIpInstanceBaseProps{
	customAttributes: map[string]*string{
		"arn": jsii.String("arn://"),
	},
})

app.synth()
```

## Public DNS Namespace Example

The following example creates an AWS Cloud Map namespace that
supports both API calls and public DNS queries, creates a service in
that namespace, and registers an IP instance:

```go
import cdk "github.com/aws/aws-cdk-go/awscdk"
import servicediscovery "github.com/aws/aws-cdk-go/awscdk"

app := cdk.NewApp()
stack := cdk.NewStack(app, jsii.String("aws-servicediscovery-integ"))

namespace := servicediscovery.NewPublicDnsNamespace(stack, jsii.String("Namespace"), &publicDnsNamespaceProps{
	name: jsii.String("foobar.com"),
})

service := namespace.createService(jsii.String("Service"), &dnsServiceProps{
	name: jsii.String("foo"),
	dnsRecordType: servicediscovery.dnsRecordType_A,
	dnsTtl: cdk.duration.seconds(jsii.Number(30)),
	healthCheck: &healthCheckConfig{
		type: servicediscovery.healthCheckType_HTTPS,
		resourcePath: jsii.String("/healthcheck"),
		failureThreshold: jsii.Number(2),
	},
})

service.registerIpInstance(jsii.String("IpInstance"), &ipInstanceBaseProps{
	ipv4: jsii.String("54.239.25.192"),
	port: jsii.Number(443),
})

app.synth()
```

For DNS namespaces, you can also register instances to services with CNAME records:

```go
import cdk "github.com/aws/aws-cdk-go/awscdk"
import servicediscovery "github.com/aws/aws-cdk-go/awscdk"

app := cdk.NewApp()
stack := cdk.NewStack(app, jsii.String("aws-servicediscovery-integ"))

namespace := servicediscovery.NewPublicDnsNamespace(stack, jsii.String("Namespace"), &publicDnsNamespaceProps{
	name: jsii.String("foobar.com"),
})

service := namespace.createService(jsii.String("Service"), &dnsServiceProps{
	name: jsii.String("foo"),
	dnsRecordType: servicediscovery.dnsRecordType_CNAME,
	dnsTtl: cdk.duration.seconds(jsii.Number(30)),
})

service.registerCnameInstance(jsii.String("CnameInstance"), &cnameInstanceBaseProps{
	instanceCname: jsii.String("service.pizza"),
})

app.synth()
```
