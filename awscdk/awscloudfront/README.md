# Amazon CloudFront Construct Library

Amazon CloudFront is a web service that speeds up distribution of your static and dynamic web content, such as .html, .css, .js, and image files, to
your users. CloudFront delivers your content through a worldwide network of data centers called edge locations. When a user requests content that
you're serving with CloudFront, the user is routed to the edge location that provides the lowest latency, so that content is delivered with the best
possible performance.

## Distribution API

The `Distribution` API is currently being built to replace the existing `CloudFrontWebDistribution` API. The `Distribution` API is optimized for the
most common use cases of CloudFront distributions (e.g., single origin and behavior, few customizations) while still providing the ability for more
advanced use cases. The API focuses on simplicity for the common use cases, and convenience methods for creating the behaviors and origins necessary
for more complex use cases.

### Creating a distribution

CloudFront distributions deliver your content from one or more origins; an origin is the location where you store the original version of your
content. Origins can be created from S3 buckets or a custom origin (HTTP server). Constructs to define origins are in the `@aws-cdk/aws-cloudfront-origins` module.

Each distribution has a default behavior which applies to all requests to that distribution, and routes requests to a primary origin.
Additional behaviors may be specified for an origin with a given URL path pattern. Behaviors allow routing with multiple origins,
controlling which HTTP methods to support, whether to require users to use HTTPS, and what query strings or cookies to forward to your origin,
among other settings.

#### From an S3 Bucket

An S3 bucket can be added as an origin. If the bucket is configured as a website endpoint, the distribution can use S3 redirects and S3 custom error
documents.

```go
// Creates a distribution from an S3 bucket.
myBucket := s3.NewBucket(this, jsii.String("myBucket"))
cloudfront.NewDistribution(this, jsii.String("myDist"), &distributionProps{
	defaultBehavior: &behaviorOptions{
		origin: origins.NewS3Origin(myBucket),
	},
})
```

The above will treat the bucket differently based on if `IBucket.isWebsite` is set or not. If the bucket is configured as a website, the bucket is
treated as an HTTP origin, and the built-in S3 redirects and error pages can be used. Otherwise, the bucket is handled as a bucket origin and
CloudFront's redirect and error handling will be used. In the latter case, the Origin will create an origin access identity and grant it access to the
underlying bucket. This can be used in conjunction with a bucket that is not public to require that your users access your content using CloudFront
URLs and not S3 URLs directly.

#### ELBv2 Load Balancer

An Elastic Load Balancing (ELB) v2 load balancer may be used as an origin. In order for a load balancer to serve as an origin, it must be publicly
accessible (`internetFacing` is true). Both Application and Network load balancers are supported.

```go
// Creates a distribution from an ELBv2 load balancer
var vpc vpc

// Create an application load balancer in a VPC. 'internetFacing' must be 'true'
// for CloudFront to access the load balancer and use it as an origin.
lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &applicationLoadBalancerProps{
	vpc: vpc,
	internetFacing: jsii.Boolean(true),
})
cloudfront.NewDistribution(this, jsii.String("myDist"), &distributionProps{
	defaultBehavior: &behaviorOptions{
		origin: origins.NewLoadBalancerV2Origin(lb),
	},
})
```

#### From an HTTP endpoint

Origins can also be created from any other HTTP endpoint, given the domain name, and optionally, other origin properties.

```go
// Creates a distribution from an HTTP endpoint
// Creates a distribution from an HTTP endpoint
cloudfront.NewDistribution(this, jsii.String("myDist"), &distributionProps{
	defaultBehavior: &behaviorOptions{
		origin: origins.NewHttpOrigin(jsii.String("www.example.com")),
	},
})
```

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
import acm "github.com/aws/aws-cdk-go/awscdk"
import route53 "github.com/aws/aws-cdk-go/awscdk"

var hostedZone hostedZone

var myBucket bucket

myCertificate := acm.NewDnsValidatedCertificate(this, jsii.String("mySiteCert"), &dnsValidatedCertificateProps{
	domainName: jsii.String("www.example.com"),
	hostedZone: hostedZone,
})
cloudfront.NewDistribution(this, jsii.String("myDist"), &distributionProps{
	defaultBehavior: &behaviorOptions{
		origin: origins.NewS3Origin(myBucket),
	},
	domainNames: []*string{
		jsii.String("www.example.com"),
	},
	certificate: myCertificate,
})
```

However, you can customize the minimum protocol version for the certificate while creating the distribution using `minimumProtocolVersion` property.

```go
// Create a Distribution with a custom domain name and a minimum protocol version.
var myBucket bucket

cloudfront.NewDistribution(this, jsii.String("myDist"), &distributionProps{
	defaultBehavior: &behaviorOptions{
		origin: origins.NewS3Origin(myBucket),
	},
	domainNames: []*string{
		jsii.String("www.example.com"),
	},
	minimumProtocolVersion: cloudfront.securityPolicyProtocol_TLS_V1_2016,
	sslSupportMethod: cloudfront.sSLMethod_SNI,
})
```

#### Cross Region Certificates

> **This feature is currently experimental**

You can enable the Stack property `crossRegionReferences`
in order to access resources in a different stack *and* region. With this feature flag
enabled it is possible to do something like creating a CloudFront distribution in `us-east-2` and
an ACM certificate in `us-east-1`.

```go
// Example automatically generated from non-compiling source. May contain errors.
stack1 := awscdk.Newstack(app, jsii.String("Stack1"), &stackProps{
	env: &environment{
		region: jsii.String("us-east-1"),
	},
	crossRegionReferences: jsii.Boolean(true),
})
cert := acm.NewCertificate(stack1, jsii.String("Cert"), map[string]interface{}{
	"domainName": jsii.String("*.example.com"),
	"validation": acm.CertificateValidation_fromDns(route53.PublicHostedZone_fromHostedZoneId(stack1, jsii.String("Zone"), jsii.String("Z0329774B51CGXTDQV3X"))),
})

stack2 := awscdk.Newstack(app, jsii.String("Stack2"), &stackProps{
	env: &environment{
		region: jsii.String("us-east-2"),
	},
	crossRegionReferences: jsii.Boolean(true),
})
cloudfront.NewDistribution(stack2, jsii.String("Distribution"), &distributionProps{
	defaultBehavior: &behaviorOptions{
		origin: origins.NewHttpOrigin(jsii.String("example.com")),
	},
	domainNames: []*string{
		jsii.String("dev.example.com"),
	},
	certificate: cert,
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

myWebDistribution := cloudfront.NewDistribution(this, jsii.String("myDist"), &distributionProps{
	defaultBehavior: &behaviorOptions{
		origin: origins.NewS3Origin(myBucket),
		allowedMethods: cloudfront.allowedMethods_ALLOW_ALL(),
		viewerProtocolPolicy: cloudfront.viewerProtocolPolicy_REDIRECT_TO_HTTPS,
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

myWebDistribution.addBehavior(jsii.String("/images/*.jpg"), origins.NewS3Origin(myBucket), &addBehaviorOptions{
	viewerProtocolPolicy: cloudfront.viewerProtocolPolicy_REDIRECT_TO_HTTPS,
})
```

These behaviors can also be specified at distribution creation time.

```go
// Create a Distribution with additional behaviors at creation time.
var myBucket bucket

bucketOrigin := origins.NewS3Origin(myBucket)
cloudfront.NewDistribution(this, jsii.String("myDist"), &distributionProps{
	defaultBehavior: &behaviorOptions{
		origin: bucketOrigin,
		allowedMethods: cloudfront.allowedMethods_ALLOW_ALL(),
		viewerProtocolPolicy: cloudfront.viewerProtocolPolicy_REDIRECT_TO_HTTPS,
	},
	additionalBehaviors: map[string]*behaviorOptions{
		"/images/*.jpg": &behaviorOptions{
			"origin": bucketOrigin,
			"viewerProtocolPolicy": cloudfront.*viewerProtocolPolicy_REDIRECT_TO_HTTPS,
		},
	},
})
```

### Customizing Cache Keys and TTLs with Cache Policies

You can use a cache policy to improve your cache hit ratio by controlling the values (URL query strings, HTTP headers, and cookies)
that are included in the cache key, and/or adjusting how long items remain in the cache via the time-to-live (TTL) settings.
CloudFront provides some predefined cache policies, known as managed policies, for common use cases. You can use these managed policies,
or you can create your own cache policy that’s specific to your needs.
See https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html for more details.

```go
// Using an existing cache policy for a Distribution
var bucketOrigin s3Origin

cloudfront.NewDistribution(this, jsii.String("myDistManagedPolicy"), &distributionProps{
	defaultBehavior: &behaviorOptions{
		origin: bucketOrigin,
		cachePolicy: cloudfront.cachePolicy_CACHING_OPTIMIZED(),
	},
})
```

```go
// Creating a custom cache policy for a Distribution -- all parameters optional
var bucketOrigin s3Origin

myCachePolicy := cloudfront.NewCachePolicy(this, jsii.String("myCachePolicy"), &cachePolicyProps{
	cachePolicyName: jsii.String("MyPolicy"),
	comment: jsii.String("A default policy"),
	defaultTtl: awscdk.Duration.days(jsii.Number(2)),
	minTtl: awscdk.Duration.minutes(jsii.Number(1)),
	maxTtl: awscdk.Duration.days(jsii.Number(10)),
	cookieBehavior: cloudfront.cacheCookieBehavior.all(),
	headerBehavior: cloudfront.cacheHeaderBehavior.allowList(jsii.String("X-CustomHeader")),
	queryStringBehavior: cloudfront.cacheQueryStringBehavior.denyList(jsii.String("username")),
	enableAcceptEncodingGzip: jsii.Boolean(true),
	enableAcceptEncodingBrotli: jsii.Boolean(true),
})
cloudfront.NewDistribution(this, jsii.String("myDistCustomPolicy"), &distributionProps{
	defaultBehavior: &behaviorOptions{
		origin: bucketOrigin,
		cachePolicy: myCachePolicy,
	},
})
```

### Customizing Origin Requests with Origin Request Policies

When CloudFront makes a request to an origin, the URL path, request body (if present), and a few standard headers are included.
Other information from the viewer request, such as URL query strings, HTTP headers, and cookies, is not included in the origin request by default.
You can use an origin request policy to control the information that’s included in an origin request.
CloudFront provides some predefined origin request policies, known as managed policies, for common use cases. You can use these managed policies,
or you can create your own origin request policy that’s specific to your needs.
See https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html for more details.

```go
// Using an existing origin request policy for a Distribution
var bucketOrigin s3Origin

cloudfront.NewDistribution(this, jsii.String("myDistManagedPolicy"), &distributionProps{
	defaultBehavior: &behaviorOptions{
		origin: bucketOrigin,
		originRequestPolicy: cloudfront.originRequestPolicy_CORS_S3_ORIGIN(),
	},
})
```

```go
// Creating a custom origin request policy for a Distribution -- all parameters optional
var bucketOrigin s3Origin

myOriginRequestPolicy := cloudfront.NewOriginRequestPolicy(this, jsii.String("OriginRequestPolicy"), &originRequestPolicyProps{
	originRequestPolicyName: jsii.String("MyPolicy"),
	comment: jsii.String("A default policy"),
	cookieBehavior: cloudfront.originRequestCookieBehavior.none(),
	headerBehavior: cloudfront.originRequestHeaderBehavior.all(jsii.String("CloudFront-Is-Android-Viewer")),
	queryStringBehavior: cloudfront.originRequestQueryStringBehavior.allowList(jsii.String("username")),
})

cloudfront.NewDistribution(this, jsii.String("myDistCustomPolicy"), &distributionProps{
	defaultBehavior: &behaviorOptions{
		origin: bucketOrigin,
		originRequestPolicy: myOriginRequestPolicy,
	},
})
```

### Customizing Response Headers with Response Headers Policies

You can configure CloudFront to add one or more HTTP headers to the responses that it sends to viewers (web browsers or other clients), without making any changes to the origin or writing any code.
To specify the headers that CloudFront adds to HTTP responses, you use a response headers policy. CloudFront adds the headers regardless of whether it serves the object from the cache or has to retrieve the object from the origin. If the origin response includes one or more of the headers that’s in a response headers policy, the policy can specify whether CloudFront uses the header it received from the origin or overwrites it with the one in the policy.
See https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/adding-response-headers.html

```go
// Using an existing managed response headers policy
var bucketOrigin s3Origin

cloudfront.NewDistribution(this, jsii.String("myDistManagedPolicy"), &distributionProps{
	defaultBehavior: &behaviorOptions{
		origin: bucketOrigin,
		responseHeadersPolicy: cloudfront.responseHeadersPolicy_CORS_ALLOW_ALL_ORIGINS(),
	},
})

// Creating a custom response headers policy -- all parameters optional
myResponseHeadersPolicy := cloudfront.NewResponseHeadersPolicy(this, jsii.String("ResponseHeadersPolicy"), &responseHeadersPolicyProps{
	responseHeadersPolicyName: jsii.String("MyPolicy"),
	comment: jsii.String("A default policy"),
	corsBehavior: &responseHeadersCorsBehavior{
		accessControlAllowCredentials: jsii.Boolean(false),
		accessControlAllowHeaders: []*string{
			jsii.String("X-Custom-Header-1"),
			jsii.String("X-Custom-Header-2"),
		},
		accessControlAllowMethods: []*string{
			jsii.String("GET"),
			jsii.String("POST"),
		},
		accessControlAllowOrigins: []*string{
			jsii.String("*"),
		},
		accessControlExposeHeaders: []*string{
			jsii.String("X-Custom-Header-1"),
			jsii.String("X-Custom-Header-2"),
		},
		accessControlMaxAge: awscdk.Duration.seconds(jsii.Number(600)),
		originOverride: jsii.Boolean(true),
	},
	customHeadersBehavior: &responseCustomHeadersBehavior{
		customHeaders: []responseCustomHeader{
			&responseCustomHeader{
				header: jsii.String("X-Amz-Date"),
				value: jsii.String("some-value"),
				override: jsii.Boolean(true),
			},
			&responseCustomHeader{
				header: jsii.String("X-Amz-Security-Token"),
				value: jsii.String("some-value"),
				override: jsii.Boolean(false),
			},
		},
	},
	securityHeadersBehavior: &responseSecurityHeadersBehavior{
		contentSecurityPolicy: &responseHeadersContentSecurityPolicy{
			contentSecurityPolicy: jsii.String("default-src https:;"),
			override: jsii.Boolean(true),
		},
		contentTypeOptions: &responseHeadersContentTypeOptions{
			override: jsii.Boolean(true),
		},
		frameOptions: &responseHeadersFrameOptions{
			frameOption: cloudfront.headersFrameOption_DENY,
			override: jsii.Boolean(true),
		},
		referrerPolicy: &responseHeadersReferrerPolicy{
			referrerPolicy: cloudfront.headersReferrerPolicy_NO_REFERRER,
			override: jsii.Boolean(true),
		},
		strictTransportSecurity: &responseHeadersStrictTransportSecurity{
			accessControlMaxAge: awscdk.Duration.seconds(jsii.Number(600)),
			includeSubdomains: jsii.Boolean(true),
			override: jsii.Boolean(true),
		},
		xssProtection: &responseHeadersXSSProtection{
			protection: jsii.Boolean(true),
			modeBlock: jsii.Boolean(true),
			reportUri: jsii.String("https://example.com/csp-report"),
			override: jsii.Boolean(true),
		},
	},
})
cloudfront.NewDistribution(this, jsii.String("myDistCustomPolicy"), &distributionProps{
	defaultBehavior: &behaviorOptions{
		origin: bucketOrigin,
		responseHeadersPolicy: myResponseHeadersPolicy,
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

pubKey := cloudfront.NewPublicKey(this, jsii.String("MyPubKey"), &publicKeyProps{
	encodedKey: publicKey,
})

keyGroup := cloudfront.NewKeyGroup(this, jsii.String("MyKeyGroup"), &keyGroupProps{
	items: []iPublicKey{
		pubKey,
	},
})

cloudfront.NewDistribution(this, jsii.String("Dist"), &distributionProps{
	defaultBehavior: &behaviorOptions{
		origin: origins.NewHttpOrigin(jsii.String("www.example.com")),
		trustedKeyGroups: []iKeyGroup{
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
myFunc := #error#.NewEdgeFunction(this, jsii.String("MyFunction"), &edgeFunctionProps{
	runtime: lambda.runtime_NODEJS_14_X(),
	handler: jsii.String("index.handler"),
	code: lambda.code.fromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
})
cloudfront.NewDistribution(this, jsii.String("myDist"), &distributionProps{
	defaultBehavior: &behaviorOptions{
		origin: origins.NewS3Origin(myBucket),
		edgeLambdas: []edgeLambda{
			&edgeLambda{
				functionVersion: myFunc.currentVersion,
				eventType: cloudfront.lambdaEdgeEventType_VIEWER_REQUEST,
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
> See https://docs.aws.amazon.com/cdk/latest/guide/bootstrapping.html for more about bootstrapping regions.

If the stack is in `us-east-1`, a "normal" `lambda.Function` can be used instead of an `EdgeFunction`.

```go
// Using a lambda Function instead of an EdgeFunction for stacks in `us-east-`.
myFunc := lambda.NewFunction(this, jsii.String("MyFunction"), &functionProps{
	runtime: lambda.runtime_NODEJS_14_X(),
	handler: jsii.String("index.handler"),
	code: lambda.code.fromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
})
```

If the stack is not in `us-east-1`, and you need references from different applications on the same account,
you can also set a specific stack ID for each Lambda@Edge.

```go
// Setting stackIds for EdgeFunctions that can be referenced from different applications
// on the same account.
myFunc1 := #error#.NewEdgeFunction(this, jsii.String("MyFunction1"), &edgeFunctionProps{
	runtime: lambda.runtime_NODEJS_14_X(),
	handler: jsii.String("index.handler"),
	code: lambda.code.fromAsset(path.join(__dirname, jsii.String("lambda-handler1"))),
	stackId: jsii.String("edge-lambda-stack-id-1"),
})

myFunc2 := #error#.NewEdgeFunction(this, jsii.String("MyFunction2"), &edgeFunctionProps{
	runtime: lambda.*runtime_NODEJS_14_X(),
	handler: jsii.String("index.handler"),
	code: lambda.*code.fromAsset(path.join(__dirname, jsii.String("lambda-handler2"))),
	stackId: jsii.String("edge-lambda-stack-id-2"),
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
myDistribution := cloudfront.NewDistribution(this, jsii.String("myDist"), &distributionProps{
	defaultBehavior: &behaviorOptions{
		origin: myOrigin,
	},
	additionalBehaviors: map[string]*behaviorOptions{
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
myDistribution.addBehavior(jsii.String("images/*"), myOrigin, &addBehaviorOptions{
	edgeLambdas: []edgeLambda{
		&edgeLambda{
			functionVersion: myFunc.currentVersion,
			eventType: cloudfront.lambdaEdgeEventType_VIEWER_RESPONSE,
		},
	},
})
```

Adding an existing Lambda@Edge function created in a different stack to a CloudFront distribution.

```go
// Adding an existing Lambda@Edge function created in a different stack
// to a CloudFront distribution.
var s3Bucket bucket

functionVersion := lambda.version.fromVersionArn(this, jsii.String("Version"), jsii.String("arn:aws:lambda:us-east-1:123456789012:function:functionName:1"))

cloudfront.NewDistribution(this, jsii.String("distro"), &distributionProps{
	defaultBehavior: &behaviorOptions{
		origin: origins.NewS3Origin(s3Bucket),
		edgeLambdas: []edgeLambda{
			&edgeLambda{
				functionVersion: functionVersion,
				eventType: cloudfront.lambdaEdgeEventType_VIEWER_REQUEST,
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
cfFunction := cloudfront.NewFunction(this, jsii.String("Function"), &functionProps{
	code: cloudfront.functionCode.fromInline(jsii.String("function handler(event) { return event.request }")),
})
cloudfront.NewDistribution(this, jsii.String("distro"), &distributionProps{
	defaultBehavior: &behaviorOptions{
		origin: origins.NewS3Origin(s3Bucket),
		functionAssociations: []functionAssociation{
			&functionAssociation{
				function: cfFunction,
				eventType: cloudfront.functionEventType_VIEWER_REQUEST,
			},
		},
	},
})
```

It will auto-generate the name of the function and deploy it to the `live` stage.

Additionally, you can load the function's code from a file using the `FunctionCode.fromFile()` method.

### Logging

You can configure CloudFront to create log files that contain detailed information about every user request that CloudFront receives.
The logs can go to either an existing bucket, or a bucket will be created for you.

```go
// Configure logging for Distributions

// Simplest form - creates a new bucket and logs to it.
// Configure logging for Distributions
// Simplest form - creates a new bucket and logs to it.
cloudfront.NewDistribution(this, jsii.String("myDist"), &distributionProps{
	defaultBehavior: &behaviorOptions{
		origin: origins.NewHttpOrigin(jsii.String("www.example.com")),
	},
	enableLogging: jsii.Boolean(true),
})

// You can optionally log to a specific bucket, configure whether cookies are logged, and give the log files a prefix.
// You can optionally log to a specific bucket, configure whether cookies are logged, and give the log files a prefix.
cloudfront.NewDistribution(this, jsii.String("myDist"), &distributionProps{
	defaultBehavior: &behaviorOptions{
		origin: origins.NewHttpOrigin(jsii.String("www.example.com")),
	},
	enableLogging: jsii.Boolean(true),
	 // Optional, this is implied if logBucket is specified
	logBucket: s3.NewBucket(this, jsii.String("LogBucket")),
	logFilePrefix: jsii.String("distribution-access-logs/"),
	logIncludesCookies: jsii.Boolean(true),
})
```

### HTTP Versions

You can configure CloudFront to use a particular version of the HTTP protocol. By default,
newly created distributions use HTTP/2 but can be configured to use both HTTP/2 and HTTP/3 or
just HTTP/3. For all supported HTTP versions, see the `HttpVerson` enum.

```go
// Example automatically generated from non-compiling source. May contain errors.
// Configure a distribution to use HTTP/2 and HTTP/3
// Configure a distribution to use HTTP/2 and HTTP/3
cloudfront.NewDistribution(this, jsii.String("myDist"), &distributionProps{
	defaultBehavior: &behaviorOptions{
		origin: origins.NewHttpOrigin(jsii.String("www.example.com")),
	},
	httpVersion: cloudfront.httpVersion_HTTP2_AND_3,
})
```

### Importing Distributions

Existing distributions can be imported as well; note that like most imported constructs, an imported distribution cannot be modified.
However, it can be used as a reference for other higher-level constructs.

```go
// Using a reference to an imported Distribution
distribution := cloudfront.distribution.fromDistributionAttributes(this, jsii.String("ImportedDist"), &distributionAttributes{
	domainName: jsii.String("d111111abcdef8.cloudfront.net"),
	distributionId: jsii.String("012345ABCDEF"),
})
```

### Permissions

Use the `grant()` method to allow actions on the distribution.
`grantCreateInvalidation()` is a shorthand to allow `CreateInvalidation`.

```go
var distribution distribution
var lambdaFn function

distribution.grant(lambdaFn, jsii.String("cloudfront:ListInvalidations"), jsii.String("cloudfront:GetInvalidation"))
distribution.grantCreateInvalidation(lambdaFn)
```

## Migrating from the original CloudFrontWebDistribution to the newer Distribution construct

It's possible to migrate a distribution from the original to the modern API.
The changes necessary are the following:

### The Distribution

Replace `new CloudFrontWebDistribution` with `new Distribution`. Some
configuration properties have been changed:

| Old API                        | New API                                                                                        |
|--------------------------------|------------------------------------------------------------------------------------------------|
| `originConfigs`                | `defaultBehavior`; use `additionalBehaviors` if necessary                                      |
| `viewerCertificate`            | `certificate`; use `domainNames` for aliases                                                   |
| `errorConfigurations`          | `errorResponses`                                                                               |
| `loggingConfig`                | `enableLogging`; configure with `logBucket` `logFilePrefix` and `logIncludesCookies`           |
| `viewerProtocolPolicy`         | removed; set on each behavior instead. default changed from `REDIRECT_TO_HTTPS` to `ALLOW_ALL` |

After switching constructs, you need to maintain the same logical ID for the underlying [CfnDistribution](https://docs.aws.amazon.com/cdk/api/v1/docs/@aws-cdk_aws-cloudfront.CfnDistribution.html) if you wish to avoid the deletion and recreation of your distribution.
To do this, use [escape hatches](https://docs.aws.amazon.com/cdk/v2/guide/cfn_layer.html) to override the logical ID created by the new Distribution construct with the logical ID created by the old construct.

Example:

```go
var sourceBucket bucket


myDistribution := cloudfront.NewDistribution(this, jsii.String("MyCfWebDistribution"), &distributionProps{
	defaultBehavior: &behaviorOptions{
		origin: origins.NewS3Origin(sourceBucket),
	},
})
cfnDistribution := myDistribution.node.defaultChild.(cfnDistribution)
cfnDistribution.overrideLogicalId(jsii.String("MyDistributionCFDistribution3H55TI9Q"))
```

### Behaviors

The modern API makes use of the [CloudFront Origins](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_cloudfront_origins-readme.html) module to easily configure your origin. Replace your origin configuration with the relevant CloudFront Origins class. For example, here's a behavior with an S3 origin:

```go
var sourceBucket bucket
var oai originAccessIdentity


cloudfront.NewCloudFrontWebDistribution(this, jsii.String("MyCfWebDistribution"), &cloudFrontWebDistributionProps{
	originConfigs: []sourceConfiguration{
		&sourceConfiguration{
			s3OriginSource: &s3OriginConfig{
				s3BucketSource: sourceBucket,
				originAccessIdentity: oai,
			},
			behaviors: []behavior{
				&behavior{
					isDefaultBehavior: jsii.Boolean(true),
				},
			},
		},
	},
})
```

Becomes:

```go
var sourceBucket bucket


distribution := cloudfront.NewDistribution(this, jsii.String("MyCfWebDistribution"), &distributionProps{
	defaultBehavior: &behaviorOptions{
		origin: origins.NewS3Origin(sourceBucket),
	},
})
```

In the original API all behaviors are defined in the `originConfigs` property. The new API is optimized for a single origin and behavior, so the default behavior and additional behaviors will be defined separately.

```go
var sourceBucket bucket
var oai originAccessIdentity


cloudfront.NewCloudFrontWebDistribution(this, jsii.String("MyCfWebDistribution"), &cloudFrontWebDistributionProps{
	originConfigs: []sourceConfiguration{
		&sourceConfiguration{
			s3OriginSource: &s3OriginConfig{
				s3BucketSource: sourceBucket,
				originAccessIdentity: oai,
			},
			behaviors: []behavior{
				&behavior{
					isDefaultBehavior: jsii.Boolean(true),
				},
			},
		},
		&sourceConfiguration{
			customOriginSource: &customOriginConfig{
				domainName: jsii.String("MYALIAS"),
			},
			behaviors: []*behavior{
				&behavior{
					pathPattern: jsii.String("/somewhere"),
				},
			},
		},
	},
})
```

Becomes:

```go
var sourceBucket bucket


distribution := cloudfront.NewDistribution(this, jsii.String("MyCfWebDistribution"), &distributionProps{
	defaultBehavior: &behaviorOptions{
		origin: origins.NewS3Origin(sourceBucket),
	},
	additionalBehaviors: map[string]*behaviorOptions{
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


viewerCertificate := cloudfront.viewerCertificate.fromAcmCertificate(certificate, &viewerCertificateOptions{
	aliases: []*string{
		jsii.String("MYALIAS"),
	},
})

cloudfront.NewCloudFrontWebDistribution(this, jsii.String("MyCfWebDistribution"), &cloudFrontWebDistributionProps{
	originConfigs: []sourceConfiguration{
		&sourceConfiguration{
			s3OriginSource: &s3OriginConfig{
				s3BucketSource: sourceBucket,
			},
			behaviors: []behavior{
				&behavior{
					isDefaultBehavior: jsii.Boolean(true),
				},
			},
		},
	},
	viewerCertificate: viewerCertificate,
})
```

Becomes:

```go
import acm "github.com/aws/aws-cdk-go/awscdk"
var certificate certificate
var sourceBucket bucket


distribution := cloudfront.NewDistribution(this, jsii.String("MyCfWebDistribution"), &distributionProps{
	defaultBehavior: &behaviorOptions{
		origin: origins.NewS3Origin(sourceBucket),
	},
	domainNames: []*string{
		jsii.String("MYALIAS"),
	},
	certificate: certificate,
})
```

IAM certificates aren't directly supported by the new API, but can be easily configured through [escape hatches](https://docs.aws.amazon.com/cdk/v2/guide/cfn_layer.html)

```go
var sourceBucket bucket

viewerCertificate := cloudfront.viewerCertificate.fromIamCertificate(jsii.String("MYIAMROLEIDENTIFIER"), &viewerCertificateOptions{
	aliases: []*string{
		jsii.String("MYALIAS"),
	},
})

cloudfront.NewCloudFrontWebDistribution(this, jsii.String("MyCfWebDistribution"), &cloudFrontWebDistributionProps{
	originConfigs: []sourceConfiguration{
		&sourceConfiguration{
			s3OriginSource: &s3OriginConfig{
				s3BucketSource: sourceBucket,
			},
			behaviors: []behavior{
				&behavior{
					isDefaultBehavior: jsii.Boolean(true),
				},
			},
		},
	},
	viewerCertificate: viewerCertificate,
})
```

Becomes:

```go
var sourceBucket bucket

distribution := cloudfront.NewDistribution(this, jsii.String("MyCfWebDistribution"), &distributionProps{
	defaultBehavior: &behaviorOptions{
		origin: origins.NewS3Origin(sourceBucket),
	},
	domainNames: []*string{
		jsii.String("MYALIAS"),
	},
})

cfnDistribution := distribution.node.defaultChild.(cfnDistribution)

cfnDistribution.addPropertyOverride(jsii.String("ViewerCertificate.IamCertificateId"), jsii.String("MYIAMROLEIDENTIFIER"))
cfnDistribution.addPropertyOverride(jsii.String("ViewerCertificate.SslSupportMethod"), jsii.String("sni-only"))
```

### Other changes

A number of default settings have changed on the new API when creating a new distribution, behavior, and origin.
After making the major changes needed for the migration, run `cdk diff` to see what settings have changed.
If no changes are desired during migration, you will at the least be able to use [escape hatches](https://docs.aws.amazon.com/cdk/v2/guide/cfn_layer.html) to override what the CDK synthesizes, if you can't change the properties directly.

## CloudFrontWebDistribution API

> The `CloudFrontWebDistribution` construct is the original construct written for working with CloudFront distributions.
> Users are encouraged to use the newer `Distribution` instead, as it has a simpler interface and receives new features faster.

Example usage:

```go
// Using a CloudFrontWebDistribution construct.

var sourceBucket bucket

distribution := cloudfront.NewCloudFrontWebDistribution(this, jsii.String("MyDistribution"), &cloudFrontWebDistributionProps{
	originConfigs: []sourceConfiguration{
		&sourceConfiguration{
			s3OriginSource: &s3OriginConfig{
				s3BucketSource: sourceBucket,
			},
			behaviors: []behavior{
				&behavior{
					isDefaultBehavior: jsii.Boolean(true),
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

distribution := cloudfront.NewCloudFrontWebDistribution(this, jsii.String("AnAmazingWebsiteProbably"), &cloudFrontWebDistributionProps{
	originConfigs: []sourceConfiguration{
		&sourceConfiguration{
			s3OriginSource: &s3OriginConfig{
				s3BucketSource: s3BucketSource,
			},
			behaviors: []behavior{
				&behavior{
					isDefaultBehavior: jsii.Boolean(true),
				},
			},
		},
	},
	viewerCertificate: cloudfront.viewerCertificate.fromCloudFrontDefaultCertificate(jsii.String("www.example.com")),
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

certificate := certificatemanager.NewCertificate(this, jsii.String("Certificate"), &certificateProps{
	domainName: jsii.String("example.com"),
	subjectAlternativeNames: []*string{
		jsii.String("*.example.com"),
	},
})

distribution := cloudfront.NewCloudFrontWebDistribution(this, jsii.String("AnAmazingWebsiteProbably"), &cloudFrontWebDistributionProps{
	originConfigs: []sourceConfiguration{
		&sourceConfiguration{
			s3OriginSource: &s3OriginConfig{
				s3BucketSource: s3BucketSource,
			},
			behaviors: []behavior{
				&behavior{
					isDefaultBehavior: jsii.Boolean(true),
				},
			},
		},
	},
	viewerCertificate: cloudfront.viewerCertificate.fromAcmCertificate(certificate, &viewerCertificateOptions{
		aliases: []*string{
			jsii.String("example.com"),
			jsii.String("www.example.com"),
		},
		securityPolicy: cloudfront.securityPolicyProtocol_TLS_V1,
		 // default
		sslMethod: cloudfront.sSLMethod_SNI,
	}),
})
```

#### IAM certificate

You can also import a certificate into the IAM certificate store.

See [Importing an SSL/TLS Certificate](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cnames-and-https-procedures.html#cnames-and-https-uploading-certificates) in the CloudFront User Guide.

Example:

```go
s3BucketSource := s3.NewBucket(this, jsii.String("Bucket"))

distribution := cloudfront.NewCloudFrontWebDistribution(this, jsii.String("AnAmazingWebsiteProbably"), &cloudFrontWebDistributionProps{
	originConfigs: []sourceConfiguration{
		&sourceConfiguration{
			s3OriginSource: &s3OriginConfig{
				s3BucketSource: s3BucketSource,
			},
			behaviors: []behavior{
				&behavior{
					isDefaultBehavior: jsii.Boolean(true),
				},
			},
		},
	},
	viewerCertificate: cloudfront.viewerCertificate.fromIamCertificate(jsii.String("certificateId"), &viewerCertificateOptions{
		aliases: []*string{
			jsii.String("example.com"),
		},
		securityPolicy: cloudfront.securityPolicyProtocol_SSL_V3,
		 // default
		sslMethod: cloudfront.sSLMethod_SNI,
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

pubKey := cloudfront.NewPublicKey(this, jsii.String("MyPubKey"), &publicKeyProps{
	encodedKey: publicKey,
})

keyGroup := cloudfront.NewKeyGroup(this, jsii.String("MyKeyGroup"), &keyGroupProps{
	items: []iPublicKey{
		pubKey,
	},
})

cloudfront.NewCloudFrontWebDistribution(this, jsii.String("AnAmazingWebsiteProbably"), &cloudFrontWebDistributionProps{
	originConfigs: []sourceConfiguration{
		&sourceConfiguration{
			s3OriginSource: &s3OriginConfig{
				s3BucketSource: sourceBucket,
			},
			behaviors: []behavior{
				&behavior{
					isDefaultBehavior: jsii.Boolean(true),
					trustedKeyGroups: []iKeyGroup{
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

cloudfront.NewCloudFrontWebDistribution(this, jsii.String("MyDistribution"), &cloudFrontWebDistributionProps{
	originConfigs: []sourceConfiguration{
		&sourceConfiguration{
			s3OriginSource: &s3OriginConfig{
				s3BucketSource: sourceBucket,
			},
			behaviors: []behavior{
				&behavior{
					isDefaultBehavior: jsii.Boolean(true),
				},
			},
		},
	},
	geoRestriction: cloudfront.geoRestriction.allowlist(jsii.String("US"), jsii.String("GB")),
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
distribution := cloudfront.NewCloudFrontWebDistribution(this, jsii.String("MyDistribution"), &cloudFrontWebDistributionProps{
	originConfigs: []sourceConfiguration{
		&sourceConfiguration{
			connectionAttempts: jsii.Number(3),
			connectionTimeout: awscdk.Duration.seconds(jsii.Number(10)),
			behaviors: []behavior{
				&behavior{
					isDefaultBehavior: jsii.Boolean(true),
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
cloudfront.NewCloudFrontWebDistribution(this, jsii.String("ADistribution"), &cloudFrontWebDistributionProps{
	originConfigs: []sourceConfiguration{
		&sourceConfiguration{
			s3OriginSource: &s3OriginConfig{
				s3BucketSource: s3.bucket.fromBucketName(this, jsii.String("aBucket"), jsii.String("myoriginbucket")),
				originPath: jsii.String("/"),
				originHeaders: map[string]*string{
					"myHeader": jsii.String("42"),
				},
				originShieldRegion: jsii.String("us-west-2"),
			},
			failoverS3OriginSource: &s3OriginConfig{
				s3BucketSource: s3.*bucket.fromBucketName(this, jsii.String("aBucketFallback"), jsii.String("myoriginbucketfallback")),
				originPath: jsii.String("/somewhere"),
				originHeaders: map[string]*string{
					"myHeader2": jsii.String("21"),
				},
				originShieldRegion: jsii.String("us-east-1"),
			},
			failoverCriteriaStatusCodes: []failoverStatusCode{
				cloudfront.*failoverStatusCode_INTERNAL_SERVER_ERROR,
			},
			behaviors: []behavior{
				&behavior{
					isDefaultBehavior: jsii.Boolean(true),
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
cloudfront.NewKeyGroup(this, jsii.String("MyKeyGroup"), &keyGroupProps{
	items: []iPublicKey{
		cloudfront.NewPublicKey(this, jsii.String("MyPublicKey"), &publicKeyProps{
			encodedKey: jsii.String("..."),
		}),
	},
})
```

See:

* https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html
* https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-trusted-signers.html
