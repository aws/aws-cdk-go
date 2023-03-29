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
certificateAuthority := acmpca.CertificateAuthority_FromCertificateAuthorityArn(this, jsii.String("CA"), jsii.String("arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/023077d8-2bfa-4eb0-8f22-05c96deade77"))
```

## Low-level `Cfn*` classes

You can always use the low-level classes
(starting with `Cfn*`) to create resources like the Certificate Authority:

```go
cfnCertificateAuthority := acmpca.NewCfnCertificateAuthority(this, jsii.String("CA"), &CfnCertificateAuthorityProps{
	Type: jsii.String("ROOT"),
	KeyAlgorithm: jsii.String("RSA_2048"),
	SigningAlgorithm: jsii.String("SHA256WITHRSA"),
	Subject: &SubjectProperty{
		Country: jsii.String("US"),
		Organization: jsii.String("string"),
		OrganizationalUnit: jsii.String("string"),
		DistinguishedNameQualifier: jsii.String("string"),
		State: jsii.String("string"),
		CommonName: jsii.String("123"),
		SerialNumber: jsii.String("string"),
		Locality: jsii.String("string"),
		Title: jsii.String("string"),
		Surname: jsii.String("string"),
		GivenName: jsii.String("string"),
		Initials: jsii.String("DG"),
		Pseudonym: jsii.String("string"),
		GenerationQualifier: jsii.String("DBG"),
	},
})
```

If you need to pass the higher-level `ICertificateAuthority` somewhere,
you can get it from the lower-level `CfnCertificateAuthority` using the same `fromCertificateAuthorityArn` method:

```go
var cfnCertificateAuthority cfnCertificateAuthority


certificateAuthority := acmpca.CertificateAuthority_FromCertificateAuthorityArn(this, jsii.String("CertificateAuthority"), cfnCertificateAuthority.AttrArn)
```
