# Amazon CloudFront Construct Library

Amazon CloudFront is a web service that speeds up distribution of your static and dynamic web content, such as .html, .css, .js, and image files, to
your users. CloudFront delivers your content through a worldwide network of data centers called edge locations. When a user requests content that
you're serving with CloudFront, the user is routed to the edge location that provides the lowest latency, so that content is delivered with the best
possible performance.

## Distribution API

The `Distribution` API replaces the `CloudFrontWebDistribution` API which is now deprecated. The `Distribution` API is optimized for the
most common use cases of CloudFront distributions (e.g., single origin and behavior, few customizations) while still providing the ability for more
advanced use cases. The API focuses on simplicity for the common use cases, and convenience methods for creating the behaviors and origins necessary
for more complex use cases.

### Creating a distribution

CloudFront distributions deliver your content from one or more origins; an origin is the location where you store the original version of your
content. Origins can be created from S3 buckets or a custom origin (HTTP server). Constructs to define origins are in the `aws-cdk-lib/aws-cloudfront-origins` module.

Each distribution has a default behavior which applies to all requests to that distribution, and routes requests to a primary origin.
Additional behaviors may be specified for an origin with a given URL path pattern. Behaviors allow routing with multiple origins,
controlling which HTTP methods to support, whether to require users to use HTTPS, and what query strings or cookies to forward to your origin,
among other settings.

#### From an S3 Bucket

An S3 bucket can be added as an origin. An S3 bucket origin can either be configured as a standard bucket or as a website endpoint (see AWS docs for [Using an S3 Bucket](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/DownloadDistS3AndCustomOrigins.html#using-s3-as-origin)).

```go
// Creates a distribution from an S3 bucket with origin access control
myBucket := s3.NewBucket(this, jsii.String("myBucket"))
cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.S3BucketOrigin_WithOriginAccessControl(myBucket),
	},
})
```

See the README of the [`aws-cdk-lib/aws-cloudfront-origins`](https://github.com/aws/aws-cdk/blob/main/packages/aws-cdk-lib/aws-cloudfront-origins/README.md) module for more information on setting up S3 origins and origin access control (OAC).

#### ELBv2 Load Balancer

An Elastic Load Balancing (ELB) v2 load balancer may be used as an origin. In order for a load balancer to serve as an origin, it must be publicly
accessible (`internetFacing` is true). Both Application and Network load balancers are supported.

```go
// Creates a distribution from an ELBv2 load balancer
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

#### From an HTTP endpoint

Origins can also be created from any other HTTP endpoint, given the domain name, and optionally, other origin properties.

```go
// Creates a distribution from an HTTP endpoint
// Creates a distribution from an HTTP endpoint
cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewHttpOrigin(jsii.String("www.example.com")),
	},
})
```

### VPC origins

You can use CloudFront to deliver content from applications that are hosted in your virtual private cloud (VPC) private subnets.
You can use Application Load Balancers (ALBs), Network Load Balancers (NLBs), and EC2 instances in private subnets as VPC origins.

Learn more about [Restrict access with VPC origins](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-vpc-origins.html).

See the README of the `aws-cdk-lib/aws-cloudfront-origins` module for more information on setting up VPC origins.

### Domain Names and Certificates

When you create a distribution, CloudFront assigns a domain name for the distribution, for example: `d111111abcdef8.cloudfront.net`; this value can
be retrieved from `distribution.distributionDomainName`. CloudFront distributions use a default certificate (`*.cloudfront.net`) to support HTTPS by
default. If you want to use your own domain name, such as `www.example.com`, you must associate a certificate with your distribution that contains
your domain name, and provide one (or more) domain names from the certificate for the distribution.

The certificate must be present in the AWS Certificate Manager (ACM) service in the US East (N. Virginia) region; the certificate
may either be created by ACM, or created elsewhere and imported into ACM. When a certificate is used, the distribution will support HTTPS connections
from SNI only and a minimum protocol version of TLSv1.2_2021 if the `@aws-cdk/aws-cloudfront:defaultSecurityPolicyTLSv1.2_2021` feature flag is set, and TLSv1.2_2019 otherwise.

```go
// To use your own domain name in a Distribution, you must associate a certificate
import "github.com/aws/aws-cdk-go/awscdk"
import route53 "github.com/aws/aws-cdk-go/awscdk"

var hostedZone hostedZone

var myBucket bucket

myCertificate := acm.NewCertificate(this, jsii.String("mySiteCert"), &CertificateProps{
	DomainName: jsii.String("www.example.com"),
	Validation: acm.CertificateValidation_FromDns(hostedZone),
})
cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewS3Origin(myBucket),
	},
	DomainNames: []*string{
		jsii.String("www.example.com"),
	},
	Certificate: myCertificate,
})
```

However, you can customize the minimum protocol version for the certificate while creating the distribution using `minimumProtocolVersion` property.

```go
// Create a Distribution with a custom domain name and a minimum protocol version.
var myBucket bucket

cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewS3Origin(myBucket),
	},
	DomainNames: []*string{
		jsii.String("www.example.com"),
	},
	MinimumProtocolVersion: cloudfront.SecurityPolicyProtocol_TLS_V1_2016,
	SslSupportMethod: cloudfront.SSLMethod_SNI,
})
```

#### Moving an alternate domain name to a different distribution

When you try to add an alternate domain name to a distribution but the alternate domain name is already in use on a different distribution, you get a `CNAMEAlreadyExists` error (One or more of the CNAMEs you provided are already associated with a different resource).

In that case, you might want to move the existing alternate domain name from one distribution (the source distribution) to another (the target distribution). The following steps are an overview of the process. For more information, see [Moving an alternate domain name to a different distribution](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/alternate-domain-names-move.html).

1. Deploy the stack with the target distribution. The `certificate` property must be specified but the `domainNames` should be absent.
2. Move the alternate domain name by running CloudFront `associate-alias` command. For the example and preconditions, see the AWS documentation above.
3. Specify the `domainNames` property with the alternative domain name, then deploy the stack again to resolve the drift at the alternative domain name.

#### Cross Region Certificates

> **This feature is currently experimental**

You can enable the Stack property `crossRegionReferences`
in order to access resources in a different stack *and* region. With this feature flag
enabled it is possible to do something like creating a CloudFront distribution in `us-east-2` and
an ACM certificate in `us-east-1`.

```go
import "github.com/aws/aws-cdk-go/awscdk"
import route53 "github.com/aws/aws-cdk-go/awscdk"

var app app


stack1 := awscdk.Newstack(app, jsii.String("Stack1"), &StackProps{
	Env: &Environment{
		Region: jsii.String("us-east-1"),
	},
	CrossRegionReferences: jsii.Boolean(true),
})
cert := acm.NewCertificate(stack1, jsii.String("Cert"), &CertificateProps{
	DomainName: jsii.String("*.example.com"),
	Validation: acm.CertificateValidation_FromDns(route53.PublicHostedZone_FromHostedZoneId(stack1, jsii.String("Zone"), jsii.String("Z0329774B51CGXTDQV3X"))),
})

stack2 := awscdk.Newstack(app, jsii.String("Stack2"), &StackProps{
	Env: &Environment{
		Region: jsii.String("us-east-2"),
	},
	CrossRegionReferences: jsii.Boolean(true),
})
cloudfront.NewDistribution(stack2, jsii.String("Distribution"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewHttpOrigin(jsii.String("example.com")),
	},
	DomainNames: []*string{
		jsii.String("dev.example.com"),
	},
	Certificate: cert,
})
```

### Multiple Behaviors & Origins

Each distribution has a default behavior which applies to all requests to that distribution; additional behaviors may be specified for a
given URL path pattern. Behaviors allow routing with multiple origins, controlling which HTTP methods to support, whether to require users to
use HTTPS, and what query strings or cookies to forward to your origin, among others.

The properties of the default behavior can be adjusted as part of the distribution creation. The following example shows configuring the HTTP
methods and viewer protocol policy of the cache.

```go
// Create a Distribution with configured HTTP methods and viewer protocol policy of the cache.
var myBucket bucket

myWebDistribution := cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewS3Origin(myBucket),
		AllowedMethods: cloudfront.AllowedMethods_ALLOW_ALL(),
		ViewerProtocolPolicy: cloudfront.ViewerProtocolPolicy_REDIRECT_TO_HTTPS,
	},
})
```

Additional behaviors can be specified at creation, or added after the initial creation. Each additional behavior is associated with an origin,
and enable customization for a specific set of resources based on a URL path pattern. For example, we can add a behavior to `myWebDistribution` to
override the default viewer protocol policy for all of the images.

```go
// Add a behavior to a Distribution after initial creation.
var myBucket bucket
var myWebDistribution distribution

myWebDistribution.AddBehavior(jsii.String("/images/*.jpg"), origins.NewS3Origin(myBucket), &AddBehaviorOptions{
	ViewerProtocolPolicy: cloudfront.ViewerProtocolPolicy_REDIRECT_TO_HTTPS,
})
```

These behaviors can also be specified at distribution creation time.

```go
// Create a Distribution with additional behaviors at creation time.
var myBucket bucket

bucketOrigin := origins.NewS3Origin(myBucket)
cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: bucketOrigin,
		AllowedMethods: cloudfront.AllowedMethods_ALLOW_ALL(),
		ViewerProtocolPolicy: cloudfront.ViewerProtocolPolicy_REDIRECT_TO_HTTPS,
	},
	AdditionalBehaviors: map[string]behaviorOptions{
		"/images/*.jpg": &behaviorOptions{
			"origin": bucketOrigin,
			"viewerProtocolPolicy": cloudfront.ViewerProtocolPolicy_REDIRECT_TO_HTTPS,
		},
	},
})
```

### Attaching WAF Web Acls

You can attach the AWS WAF web ACL to a CloudFront distribution.

To specify a web ACL created using the latest version of AWS WAF, use the ACL ARN, for example
`arn:aws:wafv2:us-east-1:123456789012:global/webacl/ExampleWebACL/473e64fd-f30b-4765-81a0-62ad96dd167a`.
The web ACL must be in the `us-east-1` region.

To specify a web ACL created using AWS WAF Classic, use the ACL ID, for example `473e64fd-f30b-4765-81a0-62ad96dd167a`.

```go
var bucketOrigin s3Origin
var webAcl cfnWebACL

distribution := cloudfront.NewDistribution(this, jsii.String("Distribution"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: bucketOrigin,
	},
	WebAclId: webAcl.AttrArn,
})
```

You can also attach a web ACL to a distribution after creation.

```go
var bucketOrigin s3Origin
var webAcl cfnWebACL

distribution := cloudfront.NewDistribution(this, jsii.String("Distribution"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: bucketOrigin,
	},
})

distribution.AttachWebAclId(webAcl.AttrArn)
```

### Customizing Cache Keys and TTLs with Cache Policies

You can use a cache policy to improve your cache hit ratio by controlling the values (URL query strings, HTTP headers, and cookies)
that are included in the cache key, and/or adjusting how long items remain in the cache via the time-to-live (TTL) settings.
CloudFront provides some predefined cache policies, known as managed policies, for common use cases. You can use these managed policies,
or you can create your own cache policy that’s specific to your needs.
See [https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html) for more details.

```go
// Using an existing cache policy for a Distribution
var bucketOrigin s3Origin

cloudfront.NewDistribution(this, jsii.String("myDistManagedPolicy"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: bucketOrigin,
		CachePolicy: cloudfront.CachePolicy_CACHING_OPTIMIZED(),
	},
})
```

```go
// Creating a custom cache policy for a Distribution -- all parameters optional
var bucketOrigin s3Origin

myCachePolicy := cloudfront.NewCachePolicy(this, jsii.String("myCachePolicy"), &CachePolicyProps{
	CachePolicyName: jsii.String("MyPolicy"),
	Comment: jsii.String("A default policy"),
	DefaultTtl: awscdk.Duration_Days(jsii.Number(2)),
	MinTtl: awscdk.Duration_Minutes(jsii.Number(1)),
	MaxTtl: awscdk.Duration_*Days(jsii.Number(10)),
	CookieBehavior: cloudfront.CacheCookieBehavior_All(),
	HeaderBehavior: cloudfront.CacheHeaderBehavior_AllowList(jsii.String("X-CustomHeader")),
	QueryStringBehavior: cloudfront.CacheQueryStringBehavior_DenyList(jsii.String("username")),
	EnableAcceptEncodingGzip: jsii.Boolean(true),
	EnableAcceptEncodingBrotli: jsii.Boolean(true),
})
cloudfront.NewDistribution(this, jsii.String("myDistCustomPolicy"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: bucketOrigin,
		CachePolicy: myCachePolicy,
	},
})
```

### Customizing Origin Requests with Origin Request Policies

When CloudFront makes a request to an origin, the URL path, request body (if present), and a few standard headers are included.
Other information from the viewer request, such as URL query strings, HTTP headers, and cookies, is not included in the origin request by default.
You can use an origin request policy to control the information that’s included in an origin request.
CloudFront provides some predefined origin request policies, known as managed policies, for common use cases. You can use these managed policies,
or you can create your own origin request policy that’s specific to your needs.
See [https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html) for more details.

```go
// Using an existing origin request policy for a Distribution
var bucketOrigin s3Origin

cloudfront.NewDistribution(this, jsii.String("myDistManagedPolicy"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: bucketOrigin,
		OriginRequestPolicy: cloudfront.OriginRequestPolicy_CORS_S3_ORIGIN(),
	},
})
```

```go
// Creating a custom origin request policy for a Distribution -- all parameters optional
var bucketOrigin s3Origin

myOriginRequestPolicy := cloudfront.NewOriginRequestPolicy(this, jsii.String("OriginRequestPolicy"), &OriginRequestPolicyProps{
	OriginRequestPolicyName: jsii.String("MyPolicy"),
	Comment: jsii.String("A default policy"),
	CookieBehavior: cloudfront.OriginRequestCookieBehavior_None(),
	HeaderBehavior: cloudfront.OriginRequestHeaderBehavior_All(jsii.String("CloudFront-Is-Android-Viewer")),
	QueryStringBehavior: cloudfront.OriginRequestQueryStringBehavior_AllowList(jsii.String("username")),
})

cloudfront.NewDistribution(this, jsii.String("myDistCustomPolicy"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: bucketOrigin,
		OriginRequestPolicy: myOriginRequestPolicy,
	},
})
```

### Customizing Response Headers with Response Headers Policies

You can configure CloudFront to add one or more HTTP headers to the responses that it sends to viewers (web browsers or other clients), without making any changes to the origin or writing any code.
To specify the headers that CloudFront adds to HTTP responses, you use a response headers policy. CloudFront adds the headers regardless of whether it serves the object from the cache or has to retrieve the object from the origin. If the origin response includes one or more of the headers that’s in a response headers policy, the policy can specify whether CloudFront uses the header it received from the origin or overwrites it with the one in the policy.
See [https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/adding-response-headers.html](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/adding-response-headers.html)

> [!NOTE]
> If xssProtection `reportUri` is specified, then `modeBlock` cannot be set to `true`.

```go
// Using an existing managed response headers policy
var bucketOrigin s3Origin

cloudfront.NewDistribution(this, jsii.String("myDistManagedPolicy"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: bucketOrigin,
		ResponseHeadersPolicy: cloudfront.ResponseHeadersPolicy_CORS_ALLOW_ALL_ORIGINS(),
	},
})

// Creating a custom response headers policy -- all parameters optional
myResponseHeadersPolicy := cloudfront.NewResponseHeadersPolicy(this, jsii.String("ResponseHeadersPolicy"), &ResponseHeadersPolicyProps{
	ResponseHeadersPolicyName: jsii.String("MyPolicy"),
	Comment: jsii.String("A default policy"),
	CorsBehavior: &ResponseHeadersCorsBehavior{
		AccessControlAllowCredentials: jsii.Boolean(false),
		AccessControlAllowHeaders: []*string{
			jsii.String("X-Custom-Header-1"),
			jsii.String("X-Custom-Header-2"),
		},
		AccessControlAllowMethods: []*string{
			jsii.String("GET"),
			jsii.String("POST"),
		},
		AccessControlAllowOrigins: []*string{
			jsii.String("*"),
		},
		AccessControlExposeHeaders: []*string{
			jsii.String("X-Custom-Header-1"),
			jsii.String("X-Custom-Header-2"),
		},
		AccessControlMaxAge: awscdk.Duration_Seconds(jsii.Number(600)),
		OriginOverride: jsii.Boolean(true),
	},
	CustomHeadersBehavior: &ResponseCustomHeadersBehavior{
		CustomHeaders: []responseCustomHeader{
			&responseCustomHeader{
				Header: jsii.String("X-Amz-Date"),
				Value: jsii.String("some-value"),
				Override: jsii.Boolean(true),
			},
			&responseCustomHeader{
				Header: jsii.String("X-Amz-Security-Token"),
				Value: jsii.String("some-value"),
				Override: jsii.Boolean(false),
			},
		},
	},
	SecurityHeadersBehavior: &ResponseSecurityHeadersBehavior{
		ContentSecurityPolicy: &ResponseHeadersContentSecurityPolicy{
			ContentSecurityPolicy: jsii.String("default-src https:;"),
			Override: jsii.Boolean(true),
		},
		ContentTypeOptions: &ResponseHeadersContentTypeOptions{
			Override: jsii.Boolean(true),
		},
		FrameOptions: &ResponseHeadersFrameOptions{
			FrameOption: cloudfront.HeadersFrameOption_DENY,
			Override: jsii.Boolean(true),
		},
		ReferrerPolicy: &ResponseHeadersReferrerPolicy{
			ReferrerPolicy: cloudfront.HeadersReferrerPolicy_NO_REFERRER,
			Override: jsii.Boolean(true),
		},
		StrictTransportSecurity: &ResponseHeadersStrictTransportSecurity{
			AccessControlMaxAge: awscdk.Duration_*Seconds(jsii.Number(600)),
			IncludeSubdomains: jsii.Boolean(true),
			Override: jsii.Boolean(true),
		},
		XssProtection: &ResponseHeadersXSSProtection{
			Protection: jsii.Boolean(true),
			ModeBlock: jsii.Boolean(false),
			ReportUri: jsii.String("https://example.com/csp-report"),
			Override: jsii.Boolean(true),
		},
	},
	RemoveHeaders: []*string{
		jsii.String("Server"),
	},
	ServerTimingSamplingRate: jsii.Number(50),
})
cloudfront.NewDistribution(this, jsii.String("myDistCustomPolicy"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: bucketOrigin,
		ResponseHeadersPolicy: myResponseHeadersPolicy,
	},
})
```

### Validating signed URLs or signed cookies with Trusted Key Groups

CloudFront Distribution supports validating signed URLs or signed cookies using key groups.
When a cache behavior contains trusted key groups, CloudFront requires signed URLs or signed
cookies for all requests that match the cache behavior.

```go
// Validating signed URLs or signed cookies with Trusted Key Groups

// public key in PEM format
var publicKey string

pubKey := cloudfront.NewPublicKey(this, jsii.String("MyPubKey"), &PublicKeyProps{
	EncodedKey: publicKey,
})

keyGroup := cloudfront.NewKeyGroup(this, jsii.String("MyKeyGroup"), &KeyGroupProps{
	Items: []iPublicKey{
		pubKey,
	},
})

cloudfront.NewDistribution(this, jsii.String("Dist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewHttpOrigin(jsii.String("www.example.com")),
		TrustedKeyGroups: []iKeyGroup{
			keyGroup,
		},
	},
})
```

### Lambda@Edge

Lambda@Edge is an extension of AWS Lambda, a compute service that lets you execute
functions that customize the content that CloudFront delivers. You can author Node.js
or Python functions in the US East (N. Virginia) region, and then execute them in AWS
locations globally that are closer to the viewer, without provisioning or managing servers.
Lambda@Edge functions are associated with a specific behavior and event type. Lambda@Edge
can be used to rewrite URLs, alter responses based on headers or cookies, or authorize
requests based on headers or authorization tokens.

The following shows a Lambda@Edge function added to the default behavior and triggered
on every request:

```go
var myBucket bucket
// A Lambda@Edge function added to default behavior of a Distribution
// and triggered on every request
myFunc := experimental.NewEdgeFunction(this, jsii.String("MyFunction"), &EdgeFunctionProps{
	Runtime: lambda.Runtime_NODEJS_LATEST(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
})
cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewS3Origin(myBucket),
		EdgeLambdas: []edgeLambda{
			&edgeLambda{
				FunctionVersion: myFunc.currentVersion,
				EventType: cloudfront.LambdaEdgeEventType_VIEWER_REQUEST,
			},
		},
	},
})
```

> **Note:** Lambda@Edge functions must be created in the `us-east-1` region, regardless of the region of the CloudFront distribution and stack.
> To make it easier to request functions for Lambda@Edge, the `EdgeFunction` construct can be used.
> The `EdgeFunction` construct will automatically request a function in `us-east-1`, regardless of the region of the current stack.
> `EdgeFunction` has the same interface as `Function` and can be created and used interchangeably.
> Please note that using `EdgeFunction` requires that the `us-east-1` region has been bootstrapped.
> See [https://docs.aws.amazon.com/cdk/latest/guide/bootstrapping.html](https://docs.aws.amazon.com/cdk/latest/guide/bootstrapping.html) for more about bootstrapping regions.

If the stack is in `us-east-1`, a "normal" `lambda.Function` can be used instead of an `EdgeFunction`.

```go
// Using a lambda Function instead of an EdgeFunction for stacks in `us-east-`.
myFunc := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_LATEST(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
})
```

If the stack is not in `us-east-1`, and you need references from different applications on the same account,
you can also set a specific stack ID for each Lambda@Edge.

```go
// Setting stackIds for EdgeFunctions that can be referenced from different applications
// on the same account.
myFunc1 := experimental.NewEdgeFunction(this, jsii.String("MyFunction1"), &EdgeFunctionProps{
	Runtime: lambda.Runtime_NODEJS_LATEST(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler1"))),
	StackId: jsii.String("edge-lambda-stack-id-1"),
})

myFunc2 := experimental.NewEdgeFunction(this, jsii.String("MyFunction2"), &EdgeFunctionProps{
	Runtime: lambda.Runtime_NODEJS_LATEST(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_*FromAsset(path.join(__dirname, jsii.String("lambda-handler2"))),
	StackId: jsii.String("edge-lambda-stack-id-2"),
})
```

Lambda@Edge functions can also be associated with additional behaviors,
either at or after Distribution creation time.

```go
// Associating a Lambda@Edge function with additional behaviors.

var myFunc edgeFunction
// assigning at Distribution creation
var myBucket bucket

myOrigin := origins.NewS3Origin(myBucket)
myDistribution := cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: myOrigin,
	},
	AdditionalBehaviors: map[string]behaviorOptions{
		"images/*": &behaviorOptions{
			"origin": myOrigin,
			"edgeLambdas": []EdgeLambda{
				&EdgeLambda{
					"functionVersion": myFunc.currentVersion,
					"eventType": cloudfront.LambdaEdgeEventType_ORIGIN_REQUEST,
					"includeBody": jsii.Boolean(true),
				},
			},
		},
	},
})

// assigning after creation
myDistribution.AddBehavior(jsii.String("images/*"), myOrigin, &AddBehaviorOptions{
	EdgeLambdas: []edgeLambda{
		&edgeLambda{
			FunctionVersion: myFunc.currentVersion,
			EventType: cloudfront.LambdaEdgeEventType_VIEWER_RESPONSE,
		},
	},
})
```

Adding an existing Lambda@Edge function created in a different stack to a CloudFront distribution.

```go
// Adding an existing Lambda@Edge function created in a different stack
// to a CloudFront distribution.
var s3Bucket bucket

functionVersion := lambda.Version_FromVersionArn(this, jsii.String("Version"), jsii.String("arn:aws:lambda:us-east-1:123456789012:function:functionName:1"))

cloudfront.NewDistribution(this, jsii.String("distro"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewS3Origin(s3Bucket),
		EdgeLambdas: []edgeLambda{
			&edgeLambda{
				FunctionVersion: *FunctionVersion,
				EventType: cloudfront.LambdaEdgeEventType_VIEWER_REQUEST,
			},
		},
	},
})
```

### CloudFront Function

You can also deploy CloudFront functions and add them to a CloudFront distribution.

```go
var s3Bucket bucket
// Add a cloudfront Function to a Distribution
cfFunction := cloudfront.NewFunction(this, jsii.String("Function"), &FunctionProps{
	Code: cloudfront.FunctionCode_FromInline(jsii.String("function handler(event) { return event.request }")),
	Runtime: cloudfront.FunctionRuntime_JS_2_0(),
})
cloudfront.NewDistribution(this, jsii.String("distro"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewS3Origin(s3Bucket),
		FunctionAssociations: []functionAssociation{
			&functionAssociation{
				Function: cfFunction,
				EventType: cloudfront.FunctionEventType_VIEWER_REQUEST,
			},
		},
	},
})
```

It will auto-generate the name of the function and deploy it to the `live` stage.

Additionally, you can load the function's code from a file using the `FunctionCode.fromFile()` method.

If you set `autoPublish` to false, the function will not be automatically published to the LIVE stage when it’s created.

```go
cloudfront.NewFunction(this, jsii.String("Function"), &FunctionProps{
	Code: cloudfront.FunctionCode_FromInline(jsii.String("function handler(event) { return event.request }")),
	Runtime: cloudfront.FunctionRuntime_JS_2_0(),
	AutoPublish: jsii.Boolean(false),
})
```

### Key Value Store

A CloudFront Key Value Store can be created and optionally have data imported from a JSON file
by default.

To create an empty Key Value Store:

```go
store := cloudfront.NewKeyValueStore(this, jsii.String("KeyValueStore"))
```

To also include an initial set of values, the `source` property can be specified, either from a
local file or an inline string. For the structure of this file, see [Creating a file of key value pairs](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/kvs-with-functions-create-s3-kvp.html).

```go
storeAsset := cloudfront.NewKeyValueStore(this, jsii.String("KeyValueStoreAsset"), &KeyValueStoreProps{
	KeyValueStoreName: jsii.String("KeyValueStoreAsset"),
	Source: cloudfront.ImportSource_FromAsset(jsii.String("path-to-data.json")),
})

storeInline := cloudfront.NewKeyValueStore(this, jsii.String("KeyValueStoreInline"), &KeyValueStoreProps{
	KeyValueStoreName: jsii.String("KeyValueStoreInline"),
	Source: cloudfront.ImportSource_FromInline(jSON.stringify(map[string][]map[string]*string{
		"data": []map[string]*string{
			map[string]*string{
				"key": jsii.String("key1"),
				"value": jsii.String("value1"),
			},
			map[string]*string{
				"key": jsii.String("key2"),
				"value": jsii.String("value2"),
			},
		},
	})),
})
```

The Key Value Store can then be associated to a function using the `cloudfront-js-2.0` runtime
or newer:

```go
store := cloudfront.NewKeyValueStore(this, jsii.String("KeyValueStore"))
cloudfront.NewFunction(this, jsii.String("Function"), &FunctionProps{
	Code: cloudfront.FunctionCode_FromInline(jsii.String("function handler(event) { return event.request }")),
	// Note that JS_2_0 must be used for Key Value Store support
	Runtime: cloudfront.FunctionRuntime_JS_2_0(),
	KeyValueStore: store,
})
```

### Logging

You can configure CloudFront to create log files that contain detailed information about every user request that CloudFront receives.
The logs can go to either an existing bucket, or a bucket will be created for you.

```go
// Configure logging for Distributions

// Simplest form - creates a new bucket and logs to it.
// Configure logging for Distributions
// Simplest form - creates a new bucket and logs to it.
cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewHttpOrigin(jsii.String("www.example.com")),
	},
	EnableLogging: jsii.Boolean(true),
})

// You can optionally log to a specific bucket, configure whether cookies are logged, and give the log files a prefix.
// You can optionally log to a specific bucket, configure whether cookies are logged, and give the log files a prefix.
cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewHttpOrigin(jsii.String("www.example.com")),
	},
	EnableLogging: jsii.Boolean(true),
	 // Optional, this is implied if logBucket is specified
	LogBucket: s3.NewBucket(this, jsii.String("LogBucket"), &BucketProps{
		ObjectOwnership: s3.ObjectOwnership_OBJECT_WRITER,
	}),
	LogFilePrefix: jsii.String("distribution-access-logs/"),
	LogIncludesCookies: jsii.Boolean(true),
})
```

### CloudFront Distribution Metrics

You can view operational metrics about your CloudFront distributions.

#### Default CloudFront Distribution Metrics

The [following metrics are available by default](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/viewing-cloudfront-metrics.html#monitoring-console.distributions) for all CloudFront distributions:

* Total requests: The total number of viewer requests received by CloudFront for all HTTP methods and for both HTTP and HTTPS requests.
* Total bytes uploaded: The total number of bytes that viewers uploaded to your origin with CloudFront, using POST and PUT requests.
* Total bytes downloaded: The total number of bytes downloaded by viewers for GET, HEAD, and OPTIONS requests.
* Total error rate: The percentage of all viewer requests for which the response's HTTP status code was 4xx or 5xx.
* 4xx error rate: The percentage of all viewer requests for which the response's HTTP status code was 4xx.
* 5xx error rate: The percentage of all viewer requests for which the response's HTTP status code was 5xx.

```go
dist := cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewHttpOrigin(jsii.String("www.example.com")),
	},
})

// Retrieving default distribution metrics
requestsMetric := dist.MetricRequests()
bytesUploadedMetric := dist.MetricBytesUploaded()
bytesDownloadedMetric := dist.MetricBytesDownloaded()
totalErrorRateMetric := dist.MetricTotalErrorRate()
http4xxErrorRateMetric := dist.Metric4xxErrorRate()
http5xxErrorRateMetric := dist.Metric5xxErrorRate()
```

#### Additional CloudFront Distribution Metrics

You can enable [additional CloudFront distribution metrics](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/viewing-cloudfront-metrics.html#monitoring-console.distributions-additional), which include the following metrics:

* 4xx and 5xx error rates: View 4xx and 5xx error rates by the specific HTTP status code, as a percentage of total requests.
* Origin latency: See the total time spent from when CloudFront receives a request to when it provides a response to the network (not the viewer), for responses that are served from the origin, not the CloudFront cache.
* Cache hit rate: View cache hits as a percentage of total cacheable requests, excluding errors.

```go
dist := cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewHttpOrigin(jsii.String("www.example.com")),
	},
	PublishAdditionalMetrics: jsii.Boolean(true),
})

// Retrieving additional distribution metrics
latencyMetric := dist.MetricOriginLatency()
cacheHitRateMetric := dist.MetricCacheHitRate()
http401ErrorRateMetric := dist.Metric401ErrorRate()
http403ErrorRateMetric := dist.Metric403ErrorRate()
http404ErrorRateMetric := dist.Metric404ErrorRate()
http502ErrorRateMetric := dist.Metric502ErrorRate()
http503ErrorRateMetric := dist.Metric503ErrorRate()
http504ErrorRateMetric := dist.Metric504ErrorRate()
```

### HTTP Versions

You can configure CloudFront to use a particular version of the HTTP protocol. By default,
newly created distributions use HTTP/2 but can be configured to use both HTTP/2 and HTTP/3 or
just HTTP/3. For all supported HTTP versions, see the `HttpVerson` enum.

```go
// Configure a distribution to use HTTP/2 and HTTP/3
// Configure a distribution to use HTTP/2 and HTTP/3
cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewHttpOrigin(jsii.String("www.example.com")),
	},
	HttpVersion: cloudfront.HttpVersion_HTTP2_AND_3,
})
```

### Importing Distributions

Existing distributions can be imported as well; note that like most imported constructs, an imported distribution cannot be modified.
However, it can be used as a reference for other higher-level constructs.

```go
// Using a reference to an imported Distribution
distribution := cloudfront.Distribution_FromDistributionAttributes(this, jsii.String("ImportedDist"), &DistributionAttributes{
	DomainName: jsii.String("d111111abcdef8.cloudfront.net"),
	DistributionId: jsii.String("012345ABCDEF"),
})
```

### Permissions

Use the `grant()` method to allow actions on the distribution.
`grantCreateInvalidation()` is a shorthand to allow `CreateInvalidation`.

```go
var distribution distribution
var lambdaFn function

distribution.Grant(lambdaFn, jsii.String("cloudfront:ListInvalidations"), jsii.String("cloudfront:GetInvalidation"))
distribution.GrantCreateInvalidation(lambdaFn)
```

### Realtime Log Config

CloudFront supports realtime log delivery from your distribution to a Kinesis stream.

See [Real-time logs](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html) in the CloudFront User Guide.

Example:

```go
// Adding realtime logs config to a Cloudfront Distribution on default behavior.
import kinesis "github.com/aws/aws-cdk-go/awscdk"

var stream stream


realTimeConfig := cloudfront.NewRealtimeLogConfig(this, jsii.String("realtimeLog"), &RealtimeLogConfigProps{
	EndPoints: []endpoint{
		cloudfront.*endpoint_FromKinesisStream(stream),
	},
	Fields: []*string{
		jsii.String("timestamp"),
		jsii.String("c-ip"),
		jsii.String("time-to-first-byte"),
		jsii.String("sc-status"),
	},
	RealtimeLogConfigName: jsii.String("my-delivery-stream"),
	SamplingRate: jsii.Number(100),
})

cloudfront.NewDistribution(this, jsii.String("myCdn"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewHttpOrigin(jsii.String("www.example.com")),
		RealtimeLogConfig: realTimeConfig,
	},
})
```

### gRPC

CloudFront supports gRPC, an open-source remote procedure call (RPC) framework built on HTTP/2. gRPC offers bi-directional streaming and
binary protocol that buffers payloads, making it suitable for applications that require low latency communications.

To enable your distribution to handle gRPC requests, you must include HTTP/2 as one of the supported HTTP versions and allow HTTP methods,
including POST.

See [Using gRPC with CloudFront distributions](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-using-grpc.html)
in the CloudFront User Guide.

Example:

```go
cloudfront.NewDistribution(this, jsii.String("myCdn"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewHttpOrigin(jsii.String("www.example.com")),
		AllowedMethods: cloudfront.AllowedMethods_ALLOW_ALL(),
		 // `AllowedMethods.ALLOW_ALL` is required if `enableGrpc` is true
		EnableGrpc: jsii.Boolean(true),
	},
})
```

## Migrating from the original CloudFrontWebDistribution to the newer Distribution construct

It's possible to migrate a distribution from the original to the modern API.
The changes necessary are the following:

### The Distribution

Replace `new CloudFrontWebDistribution` with `new Distribution`. Some
configuration properties have been changed:

| Old API                | New API                                                                                        |
| ---------------------- | ---------------------------------------------------------------------------------------------- |
| `originConfigs`        | `defaultBehavior`; use `additionalBehaviors` if necessary                                      |
| `viewerCertificate`    | `certificate`; use `domainNames` for aliases                                                   |
| `errorConfigurations`  | `errorResponses`                                                                               |
| `loggingConfig`        | `enableLogging`; configure with `logBucket` `logFilePrefix` and `logIncludesCookies`           |
| `viewerProtocolPolicy` | removed; set on each behavior instead. default changed from `REDIRECT_TO_HTTPS` to `ALLOW_ALL` |

After switching constructs, you need to maintain the same logical ID for the underlying [CfnDistribution](https://docs.aws.amazon.com/cdk/api/v1/docs/@aws-cdk_aws-cloudfront.CfnDistribution.html) if you wish to avoid the deletion and recreation of your distribution.
To do this, use [escape hatches](https://docs.aws.amazon.com/cdk/v2/guide/cfn_layer.html) to override the logical ID created by the new Distribution construct with the logical ID created by the old construct.

Example:

```go
var sourceBucket bucket


myDistribution := cloudfront.NewDistribution(this, jsii.String("MyCfWebDistribution"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewS3Origin(sourceBucket),
	},
})
cfnDistribution := myDistribution.Node.defaultChild.(cfnDistribution)
cfnDistribution.OverrideLogicalId(jsii.String("MyDistributionCFDistribution3H55TI9Q"))
```

### Behaviors

The modern API makes use of the [CloudFront Origins](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_cloudfront_origins-readme.html) module to easily configure your origin. Replace your origin configuration with the relevant CloudFront Origins class. For example, here's a behavior with an S3 origin:

```go
var sourceBucket bucket
var oai originAccessIdentity


cloudfront.NewCloudFrontWebDistribution(this, jsii.String("MyCfWebDistribution"), &CloudFrontWebDistributionProps{
	OriginConfigs: []sourceConfiguration{
		&sourceConfiguration{
			S3OriginSource: &S3OriginConfig{
				S3BucketSource: sourceBucket,
				OriginAccessIdentity: oai,
			},
			Behaviors: []behavior{
				&behavior{
					IsDefaultBehavior: jsii.Boolean(true),
				},
			},
		},
	},
})
```

Becomes:

```go
var sourceBucket bucket


distribution := cloudfront.NewDistribution(this, jsii.String("MyCfWebDistribution"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewS3Origin(sourceBucket),
	},
})
```

In the original API all behaviors are defined in the `originConfigs` property. The new API is optimized for a single origin and behavior, so the default behavior and additional behaviors will be defined separately.

```go
var sourceBucket bucket
var oai originAccessIdentity


cloudfront.NewCloudFrontWebDistribution(this, jsii.String("MyCfWebDistribution"), &CloudFrontWebDistributionProps{
	OriginConfigs: []sourceConfiguration{
		&sourceConfiguration{
			S3OriginSource: &S3OriginConfig{
				S3BucketSource: sourceBucket,
				OriginAccessIdentity: oai,
			},
			Behaviors: []behavior{
				&behavior{
					IsDefaultBehavior: jsii.Boolean(true),
				},
			},
		},
		&sourceConfiguration{
			CustomOriginSource: &CustomOriginConfig{
				DomainName: jsii.String("MYALIAS"),
			},
			Behaviors: []*behavior{
				&behavior{
					PathPattern: jsii.String("/somewhere"),
				},
			},
		},
	},
})
```

Becomes:

```go
var sourceBucket bucket


distribution := cloudfront.NewDistribution(this, jsii.String("MyCfWebDistribution"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewS3Origin(sourceBucket),
	},
	AdditionalBehaviors: map[string]behaviorOptions{
		"/somewhere": &behaviorOptions{
			"origin": origins.NewHttpOrigin(jsii.String("MYALIAS")),
		},
	},
})
```

### Certificates

If you are using an ACM certificate, you can pass the certificate directly to the `certificate` prop.
Any aliases used before in the `ViewerCertificate` class should be passed in to the `domainNames` prop in the modern API.

```go
import acm "github.com/aws/aws-cdk-go/awscdk"
var certificate certificate
var sourceBucket bucket


viewerCertificate := cloudfront.ViewerCertificate_FromAcmCertificate(certificate, &ViewerCertificateOptions{
	Aliases: []*string{
		jsii.String("MYALIAS"),
	},
})

cloudfront.NewCloudFrontWebDistribution(this, jsii.String("MyCfWebDistribution"), &CloudFrontWebDistributionProps{
	OriginConfigs: []sourceConfiguration{
		&sourceConfiguration{
			S3OriginSource: &S3OriginConfig{
				S3BucketSource: sourceBucket,
			},
			Behaviors: []behavior{
				&behavior{
					IsDefaultBehavior: jsii.Boolean(true),
				},
			},
		},
	},
	ViewerCertificate: viewerCertificate,
})
```

Becomes:

```go
import acm "github.com/aws/aws-cdk-go/awscdk"
var certificate certificate
var sourceBucket bucket


distribution := cloudfront.NewDistribution(this, jsii.String("MyCfWebDistribution"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewS3Origin(sourceBucket),
	},
	DomainNames: []*string{
		jsii.String("MYALIAS"),
	},
	Certificate: certificate,
})
```

IAM certificates aren't directly supported by the new API, but can be easily configured through [escape hatches](https://docs.aws.amazon.com/cdk/v2/guide/cfn_layer.html)

```go
var sourceBucket bucket

viewerCertificate := cloudfront.ViewerCertificate_FromIamCertificate(jsii.String("MYIAMROLEIDENTIFIER"), &ViewerCertificateOptions{
	Aliases: []*string{
		jsii.String("MYALIAS"),
	},
})

cloudfront.NewCloudFrontWebDistribution(this, jsii.String("MyCfWebDistribution"), &CloudFrontWebDistributionProps{
	OriginConfigs: []sourceConfiguration{
		&sourceConfiguration{
			S3OriginSource: &S3OriginConfig{
				S3BucketSource: sourceBucket,
			},
			Behaviors: []behavior{
				&behavior{
					IsDefaultBehavior: jsii.Boolean(true),
				},
			},
		},
	},
	ViewerCertificate: viewerCertificate,
})
```

Becomes:

```go
var sourceBucket bucket

distribution := cloudfront.NewDistribution(this, jsii.String("MyCfWebDistribution"), &DistributionProps{
	DefaultBehavior: &BehaviorOptions{
		Origin: origins.NewS3Origin(sourceBucket),
	},
	DomainNames: []*string{
		jsii.String("MYALIAS"),
	},
})

cfnDistribution := distribution.Node.defaultChild.(cfnDistribution)

cfnDistribution.AddPropertyOverride(jsii.String("ViewerCertificate.IamCertificateId"), jsii.String("MYIAMROLEIDENTIFIER"))
cfnDistribution.AddPropertyOverride(jsii.String("ViewerCertificate.SslSupportMethod"), jsii.String("sni-only"))
```

### Other changes

A number of default settings have changed on the new API when creating a new distribution, behavior, and origin.
After making the major changes needed for the migration, run `cdk diff` to see what settings have changed.
If no changes are desired during migration, you will at the least be able to use [escape hatches](https://docs.aws.amazon.com/cdk/v2/guide/cfn_layer.html) to override what the CDK synthesizes, if you can't change the properties directly.

## CloudFrontWebDistribution API

> The `CloudFrontWebDistribution` construct is the original construct written for working with CloudFront distributions and has been marked as deprecated.
> Users are encouraged to use the newer `Distribution` instead, as it has a simpler interface and receives new features faster.

Example usage:

```go
// Using a CloudFrontWebDistribution construct.

var sourceBucket bucket

distribution := cloudfront.NewCloudFrontWebDistribution(this, jsii.String("MyDistribution"), &CloudFrontWebDistributionProps{
	OriginConfigs: []sourceConfiguration{
		&sourceConfiguration{
			S3OriginSource: &S3OriginConfig{
				S3BucketSource: sourceBucket,
			},
			Behaviors: []behavior{
				&behavior{
					IsDefaultBehavior: jsii.Boolean(true),
				},
			},
		},
	},
})
```

### Viewer certificate

By default, CloudFront Web Distributions will answer HTTPS requests with CloudFront's default certificate,
only containing the distribution `domainName` (e.g. d111111abcdef8.cloudfront.net).
You can customize the viewer certificate property to provide a custom certificate and/or list of domain name aliases to fit your needs.

See [Using Alternate Domain Names and HTTPS](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-https-alternate-domain-names.html) in the CloudFront User Guide.

#### Default certificate

You can customize the default certificate aliases. This is intended to be used in combination with CNAME records in your DNS zone.

Example:

```go
s3BucketSource := s3.NewBucket(this, jsii.String("Bucket"))

distribution := cloudfront.NewCloudFrontWebDistribution(this, jsii.String("AnAmazingWebsiteProbably"), &CloudFrontWebDistributionProps{
	OriginConfigs: []sourceConfiguration{
		&sourceConfiguration{
			S3OriginSource: &S3OriginConfig{
				S3BucketSource: *S3BucketSource,
			},
			Behaviors: []behavior{
				&behavior{
					IsDefaultBehavior: jsii.Boolean(true),
				},
			},
		},
	},
	ViewerCertificate: cloudfront.ViewerCertificate_FromCloudFrontDefaultCertificate(jsii.String("www.example.com")),
})
```

#### ACM certificate

You can change the default certificate by one stored AWS Certificate Manager, or ACM.
Those certificate can either be generated by AWS, or purchased by another CA imported into ACM.

For more information, see
[the aws-certificatemanager module documentation](https://docs.aws.amazon.com/cdk/api/latest/docs/aws-certificatemanager-readme.html)
or [Importing Certificates into AWS Certificate Manager](https://docs.aws.amazon.com/acm/latest/userguide/import-certificate.html)
in the AWS Certificate Manager User Guide.

Example:

```go
s3BucketSource := s3.NewBucket(this, jsii.String("Bucket"))

certificate := certificatemanager.NewCertificate(this, jsii.String("Certificate"), &CertificateProps{
	DomainName: jsii.String("example.com"),
	SubjectAlternativeNames: []*string{
		jsii.String("*.example.com"),
	},
})

distribution := cloudfront.NewCloudFrontWebDistribution(this, jsii.String("AnAmazingWebsiteProbably"), &CloudFrontWebDistributionProps{
	OriginConfigs: []sourceConfiguration{
		&sourceConfiguration{
			S3OriginSource: &S3OriginConfig{
				S3BucketSource: *S3BucketSource,
			},
			Behaviors: []behavior{
				&behavior{
					IsDefaultBehavior: jsii.Boolean(true),
				},
			},
		},
	},
	ViewerCertificate: cloudfront.ViewerCertificate_FromAcmCertificate(certificate, &ViewerCertificateOptions{
		Aliases: []*string{
			jsii.String("example.com"),
			jsii.String("www.example.com"),
		},
		SecurityPolicy: cloudfront.SecurityPolicyProtocol_TLS_V1,
		 // default
		SslMethod: cloudfront.SSLMethod_SNI,
	}),
})
```

#### IAM certificate

You can also import a certificate into the IAM certificate store.

See [Importing an SSL/TLS Certificate](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cnames-and-https-procedures.html#cnames-and-https-uploading-certificates) in the CloudFront User Guide.

Example:

```go
s3BucketSource := s3.NewBucket(this, jsii.String("Bucket"))

distribution := cloudfront.NewCloudFrontWebDistribution(this, jsii.String("AnAmazingWebsiteProbably"), &CloudFrontWebDistributionProps{
	OriginConfigs: []sourceConfiguration{
		&sourceConfiguration{
			S3OriginSource: &S3OriginConfig{
				S3BucketSource: *S3BucketSource,
			},
			Behaviors: []behavior{
				&behavior{
					IsDefaultBehavior: jsii.Boolean(true),
				},
			},
		},
	},
	ViewerCertificate: cloudfront.ViewerCertificate_FromIamCertificate(jsii.String("certificateId"), &ViewerCertificateOptions{
		Aliases: []*string{
			jsii.String("example.com"),
		},
		SecurityPolicy: cloudfront.SecurityPolicyProtocol_SSL_V3,
		 // default
		SslMethod: cloudfront.SSLMethod_SNI,
	}),
})
```

### Trusted Key Groups

CloudFront Web Distributions supports validating signed URLs or signed cookies using key groups.
When a cache behavior contains trusted key groups, CloudFront requires signed URLs or signed cookies for all requests that match the cache behavior.

Example:

```go
// Using trusted key groups for Cloudfront Web Distributions.
var sourceBucket bucket
var publicKey string

pubKey := cloudfront.NewPublicKey(this, jsii.String("MyPubKey"), &PublicKeyProps{
	EncodedKey: publicKey,
})

keyGroup := cloudfront.NewKeyGroup(this, jsii.String("MyKeyGroup"), &KeyGroupProps{
	Items: []iPublicKey{
		pubKey,
	},
})

cloudfront.NewCloudFrontWebDistribution(this, jsii.String("AnAmazingWebsiteProbably"), &CloudFrontWebDistributionProps{
	OriginConfigs: []sourceConfiguration{
		&sourceConfiguration{
			S3OriginSource: &S3OriginConfig{
				S3BucketSource: sourceBucket,
			},
			Behaviors: []behavior{
				&behavior{
					IsDefaultBehavior: jsii.Boolean(true),
					TrustedKeyGroups: []iKeyGroup{
						keyGroup,
					},
				},
			},
		},
	},
})
```

### Restrictions

CloudFront supports adding restrictions to your distribution.

See [Restricting the Geographic Distribution of Your Content](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/georestrictions.html) in the CloudFront User Guide.

Example:

```go
// Adding restrictions to a Cloudfront Web Distribution.
var sourceBucket bucket

cloudfront.NewCloudFrontWebDistribution(this, jsii.String("MyDistribution"), &CloudFrontWebDistributionProps{
	OriginConfigs: []sourceConfiguration{
		&sourceConfiguration{
			S3OriginSource: &S3OriginConfig{
				S3BucketSource: sourceBucket,
			},
			Behaviors: []behavior{
				&behavior{
					IsDefaultBehavior: jsii.Boolean(true),
				},
			},
		},
	},
	GeoRestriction: cloudfront.GeoRestriction_Allowlist(jsii.String("US"), jsii.String("GB")),
})
```

### Connection behaviors between CloudFront and your origin

CloudFront provides you even more control over the connection behaviors between CloudFront and your origin.
You can now configure the number of connection attempts CloudFront will make to your origin and the origin connection timeout for each attempt.

See [Origin Connection Attempts](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#origin-connection-attempts)

See [Origin Connection Timeout](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#origin-connection-timeout)

Example usage:

```go
// Configuring connection behaviors between Cloudfront and your origin
distribution := cloudfront.NewCloudFrontWebDistribution(this, jsii.String("MyDistribution"), &CloudFrontWebDistributionProps{
	OriginConfigs: []sourceConfiguration{
		&sourceConfiguration{
			ConnectionAttempts: jsii.Number(3),
			ConnectionTimeout: awscdk.Duration_Seconds(jsii.Number(10)),
			Behaviors: []behavior{
				&behavior{
					IsDefaultBehavior: jsii.Boolean(true),
				},
			},
		},
	},
})
```

#### Origin Fallback

In case the origin source is not available and answers with one of the
specified status codes the failover origin source will be used.

```go
// Configuring origin fallback options for the CloudFrontWebDistribution
// Configuring origin fallback options for the CloudFrontWebDistribution
cloudfront.NewCloudFrontWebDistribution(this, jsii.String("ADistribution"), &CloudFrontWebDistributionProps{
	OriginConfigs: []sourceConfiguration{
		&sourceConfiguration{
			S3OriginSource: &S3OriginConfig{
				S3BucketSource: s3.Bucket_FromBucketName(this, jsii.String("aBucket"), jsii.String("amzn-s3-demo-bucket")),
				OriginPath: jsii.String("/"),
				OriginHeaders: map[string]*string{
					"myHeader": jsii.String("42"),
				},
				OriginShieldRegion: jsii.String("us-west-2"),
			},
			FailoverS3OriginSource: &S3OriginConfig{
				S3BucketSource: s3.Bucket_*FromBucketName(this, jsii.String("aBucketFallback"), jsii.String("amzn-s3-demo-bucket1")),
				OriginPath: jsii.String("/somewhere"),
				OriginHeaders: map[string]*string{
					"myHeader2": jsii.String("21"),
				},
				OriginShieldRegion: jsii.String("us-east-1"),
			},
			FailoverCriteriaStatusCodes: []failoverStatusCode{
				cloudfront.*failoverStatusCode_INTERNAL_SERVER_ERROR,
			},
			Behaviors: []behavior{
				&behavior{
					IsDefaultBehavior: jsii.Boolean(true),
				},
			},
		},
	},
})
```

## KeyGroup & PublicKey API

You can create a key group to use with CloudFront signed URLs and signed cookies
You can add public keys to use with CloudFront features such as signed URLs, signed cookies, and field-level encryption.

The following example command uses OpenSSL to generate an RSA key pair with a length of 2048 bits and save to the file named `private_key.pem`.

```bash
openssl genrsa -out private_key.pem 2048
```

The resulting file contains both the public and the private key. The following example command extracts the public key from the file named `private_key.pem` and stores it in `public_key.pem`.

```bash
openssl rsa -pubout -in private_key.pem -out public_key.pem
```

Note: Don't forget to copy/paste the contents of `public_key.pem` file including `-----BEGIN PUBLIC KEY-----` and `-----END PUBLIC KEY-----` lines into `encodedKey` parameter when creating a `PublicKey`.

Example:

```go
// Create a key group to use with CloudFront signed URLs and signed cookies.
// Create a key group to use with CloudFront signed URLs and signed cookies.
cloudfront.NewKeyGroup(this, jsii.String("MyKeyGroup"), &KeyGroupProps{
	Items: []iPublicKey{
		cloudfront.NewPublicKey(this, jsii.String("MyPublicKey"), &PublicKeyProps{
			EncodedKey: jsii.String("..."),
		}),
	},
})
```

See:

* [https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html)
* [https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-trusted-signers.html](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-trusted-signers.html)
