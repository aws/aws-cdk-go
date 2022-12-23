# AWS Certificate Manager Construct Library

AWS Certificate Manager (ACM) handles the complexity of creating, storing, and renewing public and private SSL/TLS X.509 certificates and keys that
protect your AWS websites and applications. ACM certificates can secure singular domain names, multiple specific domain names, wildcard domains, or
combinations of these. ACM wildcard certificates can protect an unlimited number of subdomains.

This package provides Constructs for provisioning and referencing ACM certificates which can be used with CloudFront and ELB.

After requesting a certificate, you will need to prove that you own the
domain in question before the certificate will be granted. The CloudFormation
deployment will wait until this verification process has been completed.

Because of this wait time, when using manual validation methods, it's better
to provision your certificates either in a separate stack from your main
service, or provision them manually and import them into your CDK application.

**Note:** There is a limit on total number of ACM certificates that can be requested on an account and region within a year.
The default limit is 2000, but this limit may be (much) lower on new AWS accounts.
See https://docs.aws.amazon.com/acm/latest/userguide/acm-limits.html for more information.

## DNS validation

DNS validation is the preferred method to validate domain ownership, as it has a number of advantages over email validation.
See also [Validate with DNS](https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-validate-dns.html)
in the AWS Certificate Manager User Guide.

If Amazon Route 53 is your DNS provider for the requested domain, the DNS record can be
created automatically:

```go
myHostedZone := route53.NewHostedZone(this, jsii.String("HostedZone"), &hostedZoneProps{
	zoneName: jsii.String("example.com"),
})
acm.NewCertificate(this, jsii.String("Certificate"), &certificateProps{
	domainName: jsii.String("hello.example.com"),
	certificateName: jsii.String("Hello World Service"),
	 // Optionally provide an certificate name
	validation: acm.certificateValidation.fromDns(myHostedZone),
})
```

If Route 53 is not your DNS provider, the DNS records must be added manually and the stack will not complete
creating until the records are added.

```go
acm.NewCertificate(this, jsii.String("Certificate"), &certificateProps{
	domainName: jsii.String("hello.example.com"),
	validation: acm.certificateValidation.fromDns(),
})
```

When working with multiple domains, use the `CertificateValidation.fromDnsMultiZone()`:

```go
exampleCom := route53.NewHostedZone(this, jsii.String("ExampleCom"), &hostedZoneProps{
	zoneName: jsii.String("example.com"),
})
exampleNet := route53.NewHostedZone(this, jsii.String("ExampleNet"), &hostedZoneProps{
	zoneName: jsii.String("example.net"),
})

cert := acm.NewCertificate(this, jsii.String("Certificate"), &certificateProps{
	domainName: jsii.String("test.example.com"),
	subjectAlternativeNames: []*string{
		jsii.String("cool.example.com"),
		jsii.String("test.example.net"),
	},
	validation: acm.certificateValidation.fromDnsMultiZone(map[string]iHostedZone{
		"test.example.com": exampleCom,
		"cool.example.com": exampleCom,
		"test.example.net": exampleNet,
	}),
})
```

## Email validation

Email-validated certificates (the default) are validated by receiving an
email on one of a number of predefined domains and following the instructions
in the email.

See [Validate with Email](https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-validate-email.html)
in the AWS Certificate Manager User Guide.

```go
acm.NewCertificate(this, jsii.String("Certificate"), &certificateProps{
	domainName: jsii.String("hello.example.com"),
	validation: acm.certificateValidation.fromEmail(),
})
```

## Cross-region Certificates

ACM certificates that are used with CloudFront -- or higher-level constructs which rely on CloudFront -- must be in the `us-east-1` region.
The `DnsValidatedCertificate` construct exists to facilitate creating these certificates cross-region. This resource can only be used with
Route53-based DNS validation.

```go
var myHostedZone hostedZone

acm.NewDnsValidatedCertificate(this, jsii.String("CrossRegionCertificate"), &dnsValidatedCertificateProps{
	domainName: jsii.String("hello.example.com"),
	hostedZone: myHostedZone,
	region: jsii.String("us-east-1"),
})
```

## Requesting private certificates

AWS Certificate Manager can create [private certificates](https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-request-private.html) issued by [Private Certificate Authority (PCA)](https://docs.aws.amazon.com/acm-pca/latest/userguide/PcaWelcome.html). Validation of private certificates is not necessary.

```go
import acmpca "github.com/aws/aws-cdk-go/awscdk"


acm.NewPrivateCertificate(this, jsii.String("PrivateCertificate"), &privateCertificateProps{
	domainName: jsii.String("test.example.com"),
	subjectAlternativeNames: []*string{
		jsii.String("cool.example.com"),
		jsii.String("test.example.net"),
	},
	 // optional
	certificateAuthority: acmpca.certificateAuthority.fromCertificateAuthorityArn(this, jsii.String("CA"), jsii.String("arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/023077d8-2bfa-4eb0-8f22-05c96deade77")),
})
```

## Requesting certificates without transparency logging

Transparency logging can be opted out of for AWS Certificate Manager certificates. See [opting out of certificate transparency logging](https://docs.aws.amazon.com/acm/latest/userguide/acm-bestpractices.html#best-practices-transparency) for limits.

```go
acm.NewCertificate(this, jsii.String("Certificate"), &certificateProps{
	domainName: jsii.String("test.example.com"),
	transparencyLoggingEnabled: jsii.Boolean(false),
})
```

## Importing

If you want to import an existing certificate, you can do so from its ARN:

```go
arn := "arn:aws:..."
certificate := acm.certificate.fromCertificateArn(this, jsii.String("Certificate"), arn)
```

## Sharing between Stacks

To share the certificate between stacks in the same CDK application, simply
pass the `Certificate` object between the stacks.

## Metrics

The `DaysToExpiry` metric is available via the `metricDaysToExpiry` method for
all certificates. This metric is emitted by AWS Certificates Manager once per
day until the certificate has effectively expired.

An alarm can be created to determine whether a certificate is soon due for
renewal ussing the following code:

```go
import cloudwatch "github.com/aws/aws-cdk-go/awscdk"

var myHostedZone hostedZone

certificate := acm.NewCertificate(this, jsii.String("Certificate"), &certificateProps{
	domainName: jsii.String("hello.example.com"),
	validation: acm.certificateValidation.fromDns(myHostedZone),
})
certificate.metricDaysToExpiry().createAlarm(this, jsii.String("Alarm"), &createAlarmOptions{
	comparisonOperator: cloudwatch.comparisonOperator_LESS_THAN_THRESHOLD,
	evaluationPeriods: jsii.Number(1),
	threshold: jsii.Number(45),
})
```
