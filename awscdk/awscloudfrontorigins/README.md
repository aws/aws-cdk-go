# CloudFront Origins for the CDK CloudFront Library

This library contains convenience methods for defining origins for a CloudFront distribution. You can use this library to create origins from
S3 buckets, Elastic Load Balancing v2 load balancers, or any other domain name.

## S3 Bucket

An S3 bucket can be added as an origin. If the bucket is configured as a website endpoint, the distribution can use S3 redirects and S3 custom error
documents.

```go
myBucket := s3.NewBucket(this, jsii.String("myBucket"))
cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewS3Origin(myBucket),
	},
})
```

The above will treat the bucket differently based on if `IBucket.isWebsite` is set or not. If the bucket is configured as a website, the bucket is
treated as an HTTP origin, and the built-in S3 redirects and error pages can be used. Otherwise, the bucket is handled as a bucket origin and
CloudFront's redirect and error handling will be used. In the latter case, the Origin will create an origin access identity and grant it access to the
underlying bucket. This can be used in conjunction with a bucket that is not public to require that your users access your content using CloudFront
URLs and not S3 URLs directly. Alternatively, a custom origin access identity can be passed to the S3 origin in the properties.

### Adding Custom Headers

You can configure CloudFront to add custom headers to the requests that it sends to your origin. These custom headers enable you to send and gather information from your origin that you don’t get with typical viewer requests. These headers can even be customized for each origin. CloudFront supports custom headers for both for custom and Amazon S3 origins.

```go
myBucket := s3.NewBucket(this, jsii.String("myBucket"))
cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewS3Origin(myBucket, &S3OriginProps{
			CustomHeaders: map[string]*string{
				"Foo": jsii.String("bar"),
			},
		}),
	},
})
```

## ELBv2 Load Balancer

An Elastic Load Balancing (ELB) v2 load balancer may be used as an origin. In order for a load balancer to serve as an origin, it must be publicly
accessible (`internetFacing` is true). Both Application and Network load balancers are supported.

```go
import ec2 "github.com/aws/aws-cdk-go/awscdk"
import elbv2 "github.com/aws/aws-cdk-go/awscdk"

var vpc vpc

// Create an application load balancer in a VPC. 'internetFacing' must be 'true'
// for CloudFront to access the load balancer and use it as an origin.
lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &ApplicationLoadBalancerProps{
	Vpc: Vpc,
	InternetFacing: jsii.Boolean(true),
})
cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewLoadBalancerV2Origin(lb),
	},
})
```

The origin can also be customized to respond on different ports, have different connection properties, etc.

```go
import elbv2 "github.com/aws/aws-cdk-go/awscdk"

var loadBalancer applicationLoadBalancer

origin := origins.NewLoadBalancerV2Origin(loadBalancer, &LoadBalancerV2OriginProps{
	ConnectionAttempts: jsii.Number(3),
	ConnectionTimeout: awscdk.Duration_Seconds(jsii.Number(5)),
	ReadTimeout: awscdk.Duration_*Seconds(jsii.Number(45)),
	KeepaliveTimeout: awscdk.Duration_*Seconds(jsii.Number(45)),
	ProtocolPolicy: cloudfront.OriginProtocolPolicy_MATCH_VIEWER,
})
```

Note that the `readTimeout` and `keepaliveTimeout` properties can extend their values over 60 seconds only if a limit increase request for CloudFront origin response timeout
quota has been approved in the target account; otherwise, values over 60 seconds will produce an error at deploy time. Consider that this value is
still limited to a maximum value of 180 seconds, which is a hard limit for that quota.

## From an HTTP endpoint

Origins can also be created from any other HTTP endpoint, given the domain name, and optionally, other origin properties.

```go
cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewHttpOrigin(jsii.String("www.example.com")),
	},
})
```

See the documentation of `aws-cdk-lib/aws-cloudfront` for more information.

## Failover Origins (Origin Groups)

You can set up CloudFront with origin failover for scenarios that require high availability.
To get started, you create an origin group with two origins: a primary and a secondary.
If the primary origin is unavailable, or returns specific HTTP response status codes that indicate a failure,
CloudFront automatically switches to the secondary origin.
You achieve that behavior in the CDK using the `OriginGroup` class:

```go
myBucket := s3.NewBucket(this, jsii.String("myBucket"))
cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewOriginGroup(&OriginGroupProps{
			PrimaryOrigin: origins.NewS3Origin(myBucket),
			FallbackOrigin: origins.NewHttpOrigin(jsii.String("www.example.com")),
			// optional, defaults to: 500, 502, 503 and 504
			FallbackStatusCodes: []*f64{
				jsii.Number(404),
			},
		}),
	},
})
```

## From an API Gateway REST API

Origins can be created from an API Gateway REST API. It is recommended to use a
[regional API](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-api-endpoint-types.html) in this case. The origin path will automatically be set as the stage name.

```go
var api restApi

cloudfront.NewDistribution(this, jsii.String("Distribution"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewRestApiOrigin(api),
	},
})
```

If you want to use a different origin path, you can specify it in the `originPath` property.

```go
var api restApi

cloudfront.NewDistribution(this, jsii.String("Distribution"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewRestApiOrigin(api, &RestApiOriginProps{
			OriginPath: jsii.String("/custom-origin-path"),
		}),
	},
})
```

## From a Lambda Function URL

Lambda Function URLs enable direct invocation of Lambda functions via HTTP(S), without intermediaries. They can be set as CloudFront origins for streamlined function execution behind a CDN, leveraging caching and custom domains.

```go
import lambda "github.com/aws/aws-cdk-go/awscdk"

var fn function

fnUrl := fn.AddFunctionUrl(&FunctionUrlOptions{
	AuthType: lambda.FunctionUrlAuthType_NONE,
})

cloudfront.NewDistribution(this, jsii.String("Distribution"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewFunctionUrlOrigin(fnUrl),
	},
})
```
