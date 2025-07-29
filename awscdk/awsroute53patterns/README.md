# CDK Construct library for higher-level Route 53 Constructs

This library provides higher-level Amazon Route 53 constructs which follow common
architectural patterns.

## HTTPS Redirect

If you want to speed up delivery of your web content, you can use Amazon CloudFront,
the AWS content delivery network (CDN). CloudFront can deliver your entire website
—including dynamic, static, streaming, and interactive content—by using a global
network of edge locations. Requests for your content are automatically routed to the
edge location that gives your users the lowest latency.

This construct allows creating a redirect from domainA to domainB using Amazon
CloudFront and Amazon S3. You can specify multiple domains to be redirected.
[Learn more](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-to-cloudfront-distribution.html) about routing traffic to a CloudFront web distribution.

The `HttpsRedirect` constructs creates:

* Amazon CloudFront distribution - makes website available from data centres
  around the world
* Amazon S3 bucket - empty bucket used for website hosting redirect (`websiteRedirect`) capabilities.
* Amazon Route 53 A/AAAA Alias records - routes traffic to the CloudFront distribution
* AWS Certificate Manager certificate - SSL/TLS certificate used by
  CloudFront for your domain

⚠️ The stack/construct can be used in any region for configuring an HTTPS redirect.
The certificate created in Amazon Certificate Manager (ACM) will be in US East (N. Virginia)
region. If you use an existing certificate, the AWS region of the certificate
must be in US East (N. Virginia).

The following example creates an HTTPS redirect from `foo.example.com` to `bar.example.com`
As an existing certificate is not provided, one will be created in `us-east-1` by the CDK.

```go
patterns.NewHttpsRedirect(this, jsii.String("Redirect"), &HttpsRedirectProps{
	RecordNames: []*string{
		jsii.String("foo.example.com"),
	},
	TargetDomain: jsii.String("bar.example.com"),
	Zone: route53.HostedZone_FromHostedZoneAttributes(this, jsii.String("HostedZone"), &HostedZoneAttributes{
		HostedZoneId: jsii.String("ID"),
		ZoneName: jsii.String("example.com"),
	}),
})
```

To have `HttpsRedirect` use the `Certificate` construct as the default
created certificate instead of the deprecated `DnsValidatedCertificate`
construct, enable the `@aws-cdk/aws-route53-patters:useCertificate`
feature flag. If you are creating the stack in a region other than `us-east-1`
you must also enable `crossRegionReferences` on the stack.

```go
var app app

stack := awscdk.Newstack(app, jsii.String("Stack"), &StackProps{
	CrossRegionReferences: jsii.Boolean(true),
	Env: &Environment{
		Region: jsii.String("us-east-2"),
	},
})

patterns.NewHttpsRedirect(this, jsii.String("Redirect"), &HttpsRedirectProps{
	RecordNames: []*string{
		jsii.String("foo.example.com"),
	},
	TargetDomain: jsii.String("bar.example.com"),
	Zone: route53.HostedZone_FromHostedZoneAttributes(this, jsii.String("HostedZone"), &HostedZoneAttributes{
		HostedZoneId: jsii.String("ID"),
		ZoneName: jsii.String("example.com"),
	}),
})
```

It is safe to upgrade to `@aws-cdk/aws-route53-patterns:useCertificate` since
the new certificate will be created and updated on the CloudFront distribution
before the old certificate is deleted.
