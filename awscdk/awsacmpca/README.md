# AWS::ACMPCA Construct Library

This module is part of the [AWS Cloud Development Kit](https://github.com/aws/aws-cdk) project.

```go
import acmpca "github.com/aws/aws-cdk-go/awscdk"
```

## Certificate Authority

This package contains a `CertificateAuthority` class.
At the moment, you cannot create new Authorities using it,
but you can import existing ones using the `fromCertificateAuthorityArn` static method:

```go
certificateAuthority := acmpca.certificateAuthority.fromCertificateAuthorityArn(this, jsii.String("CA"), jsii.String("arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/023077d8-2bfa-4eb0-8f22-05c96deade77"))
```

## Low-level `Cfn*` classes

You can always use the low-level classes
(starting with `Cfn*`) to create resources like the Certificate Authority:

```go
cfnCertificateAuthority := acmpca.NewCfnCertificateAuthority(this, jsii.String("CA"), &cfnCertificateAuthorityProps{
	type: jsii.String("ROOT"),
	keyAlgorithm: jsii.String("RSA_2048"),
	signingAlgorithm: jsii.String("SHA256WITHRSA"),
	subject: &subjectProperty{
		country: jsii.String("US"),
		organization: jsii.String("string"),
		organizationalUnit: jsii.String("string"),
		distinguishedNameQualifier: jsii.String("string"),
		state: jsii.String("string"),
		commonName: jsii.String("123"),
		serialNumber: jsii.String("string"),
		locality: jsii.String("string"),
		title: jsii.String("string"),
		surname: jsii.String("string"),
		givenName: jsii.String("string"),
		initials: jsii.String("DG"),
		pseudonym: jsii.String("string"),
		generationQualifier: jsii.String("DBG"),
	},
})
```

If you need to pass the higher-level `ICertificateAuthority` somewhere,
you can get it from the lower-level `CfnCertificateAuthority` using the same `fromCertificateAuthorityArn` method:

```go
var cfnCertificateAuthority cfnCertificateAuthority


certificateAuthority := acmpca.certificateAuthority.fromCertificateAuthorityArn(this, jsii.String("CertificateAuthority"), cfnCertificateAuthority.attrArn)
```
