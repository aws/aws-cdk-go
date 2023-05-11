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
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

app := cdk.NewApp()
stack := cdk.NewStack(app, jsii.String("aws-servicediscovery-integ"))

namespace := servicediscovery.NewHttpNamespace(stack, jsii.String("MyNamespace"), &HttpNamespaceProps{
	Name: jsii.String("MyHTTPNamespace"),
})

service1 := namespace.CreateService(jsii.String("NonIpService"), &BaseServiceProps{
	Description: jsii.String("service registering non-ip instances"),
})

service1.RegisterNonIpInstance(jsii.String("NonIpInstance"), &NonIpInstanceBaseProps{
	CustomAttributes: map[string]*string{
		"arn": jsii.String("arn:aws:s3:::mybucket"),
	},
})

service2 := namespace.CreateService(jsii.String("IpService"), &BaseServiceProps{
	Description: jsii.String("service registering ip instances"),
	HealthCheck: &HealthCheckConfig{
		Type: servicediscovery.HealthCheckType_HTTP,
		ResourcePath: jsii.String("/check"),
	},
})

service2.RegisterIpInstance(jsii.String("IpInstance"), &IpInstanceBaseProps{
	Ipv4: jsii.String("54.239.25.192"),
})

app.Synth()
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
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

app := cdk.NewApp()
stack := cdk.NewStack(app, jsii.String("aws-servicediscovery-integ"))

vpc := ec2.NewVpc(stack, jsii.String("Vpc"), &VpcProps{
	MaxAzs: jsii.Number(2),
})

namespace := servicediscovery.NewPrivateDnsNamespace(stack, jsii.String("Namespace"), &PrivateDnsNamespaceProps{
	Name: jsii.String("boobar.com"),
	Vpc: Vpc,
})

service := namespace.CreateService(jsii.String("Service"), &DnsServiceProps{
	DnsRecordType: servicediscovery.DnsRecordType_A_AAAA,
	DnsTtl: cdk.Duration_Seconds(jsii.Number(30)),
	LoadBalancer: jsii.Boolean(true),
})

loadbalancer := elbv2.NewApplicationLoadBalancer(stack, jsii.String("LB"), &ApplicationLoadBalancerProps{
	Vpc: Vpc,
	InternetFacing: jsii.Boolean(true),
})

service.RegisterLoadBalancer(jsii.String("Loadbalancer"), loadbalancer)

arnService := namespace.CreateService(jsii.String("ArnService"), &DnsServiceProps{
	DiscoveryType: servicediscovery.DiscoveryType_API,
})

arnService.RegisterNonIpInstance(jsii.String("NonIpInstance"), &NonIpInstanceBaseProps{
	CustomAttributes: map[string]*string{
		"arn": jsii.String("arn://"),
	},
})

app.Synth()
```

## Public DNS Namespace Example

The following example creates an AWS Cloud Map namespace that
supports both API calls and public DNS queries, creates a service in
that namespace, and registers an IP instance:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

app := cdk.NewApp()
stack := cdk.NewStack(app, jsii.String("aws-servicediscovery-integ"))

namespace := servicediscovery.NewPublicDnsNamespace(stack, jsii.String("Namespace"), &PublicDnsNamespaceProps{
	Name: jsii.String("foobar.com"),
})

service := namespace.CreateService(jsii.String("Service"), &DnsServiceProps{
	Name: jsii.String("foo"),
	DnsRecordType: servicediscovery.DnsRecordType_A,
	DnsTtl: cdk.Duration_Seconds(jsii.Number(30)),
	HealthCheck: &HealthCheckConfig{
		Type: servicediscovery.HealthCheckType_HTTPS,
		ResourcePath: jsii.String("/healthcheck"),
		FailureThreshold: jsii.Number(2),
	},
})

service.RegisterIpInstance(jsii.String("IpInstance"), &IpInstanceBaseProps{
	Ipv4: jsii.String("54.239.25.192"),
	Port: jsii.Number(443),
})

app.Synth()
```

For DNS namespaces, you can also register instances to services with CNAME records:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

app := cdk.NewApp()
stack := cdk.NewStack(app, jsii.String("aws-servicediscovery-integ"))

namespace := servicediscovery.NewPublicDnsNamespace(stack, jsii.String("Namespace"), &PublicDnsNamespaceProps{
	Name: jsii.String("foobar.com"),
})

service := namespace.CreateService(jsii.String("Service"), &DnsServiceProps{
	Name: jsii.String("foo"),
	DnsRecordType: servicediscovery.DnsRecordType_CNAME,
	DnsTtl: cdk.Duration_Seconds(jsii.Number(30)),
})

service.RegisterCnameInstance(jsii.String("CnameInstance"), &CnameInstanceBaseProps{
	InstanceCname: jsii.String("service.pizza"),
})

app.Synth()
```
