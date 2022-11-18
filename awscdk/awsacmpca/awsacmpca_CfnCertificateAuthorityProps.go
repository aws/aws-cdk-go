package awsacmpca

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnCertificateAuthority`.
//
// Example:
//   cfnCertificateAuthority := acmpca.NewCfnCertificateAuthority(this, jsii.String("CA"), &cfnCertificateAuthorityProps{
//   	type: jsii.String("ROOT"),
//   	keyAlgorithm: jsii.String("RSA_2048"),
//   	signingAlgorithm: jsii.String("SHA256WITHRSA"),
//   	subject: &subjectProperty{
//   		country: jsii.String("US"),
//   		organization: jsii.String("string"),
//   		organizationalUnit: jsii.String("string"),
//   		distinguishedNameQualifier: jsii.String("string"),
//   		state: jsii.String("string"),
//   		commonName: jsii.String("123"),
//   		serialNumber: jsii.String("string"),
//   		locality: jsii.String("string"),
//   		title: jsii.String("string"),
//   		surname: jsii.String("string"),
//   		givenName: jsii.String("string"),
//   		initials: jsii.String("DG"),
//   		pseudonym: jsii.String("string"),
//   		generationQualifier: jsii.String("DBG"),
//   	},
//   })
//
type CfnCertificateAuthorityProps struct {
	// Type of the public key algorithm and size, in bits, of the key pair that your CA creates when it issues a certificate.
	//
	// When you create a subordinate CA, you must use a key algorithm supported by the parent CA.
	KeyAlgorithm *string `field:"required" json:"keyAlgorithm" yaml:"keyAlgorithm"`
	// Name of the algorithm your private CA uses to sign certificate requests.
	//
	// This parameter should not be confused with the `SigningAlgorithm` parameter used to sign certificates when they are issued.
	SigningAlgorithm *string `field:"required" json:"signingAlgorithm" yaml:"signingAlgorithm"`
	// Structure that contains X.500 distinguished name information for your private CA.
	Subject interface{} `field:"required" json:"subject" yaml:"subject"`
	// Type of your private CA.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Specifies information to be added to the extension section of the certificate signing request (CSR).
	CsrExtensions interface{} `field:"optional" json:"csrExtensions" yaml:"csrExtensions"`
	// Specifies a cryptographic key management compliance standard used for handling CA keys.
	//
	// Default: FIPS_140_2_LEVEL_3_OR_HIGHER
	//
	// *Note:* `FIPS_140_2_LEVEL_3_OR_HIGHER` is not supported in the following Regions:
	//
	// - ap-northeast-3
	// - ap-southeast-3
	//
	// When creating a CA in these Regions, you must provide `FIPS_140_2_LEVEL_2_OR_HIGHER` as the argument for `KeyStorageSecurityStandard` . Failure to do this results in an `InvalidArgsException` with the message, "A certificate authority cannot be created in this region with the specified security standard."
	KeyStorageSecurityStandard *string `field:"optional" json:"keyStorageSecurityStandard" yaml:"keyStorageSecurityStandard"`
	// Information about the certificate revocation list (CRL) created and maintained by your private CA.
	//
	// Certificate revocation information used by the CreateCertificateAuthority and UpdateCertificateAuthority actions. Your certificate authority can create and maintain a certificate revocation list (CRL). A CRL contains information about certificates that have been revoked.
	RevocationConfiguration interface{} `field:"optional" json:"revocationConfiguration" yaml:"revocationConfiguration"`
	// Key-value pairs that will be attached to the new private CA.
	//
	// You can associate up to 50 tags with a private CA. For information using tags with IAM to manage permissions, see [Controlling Access Using IAM Tags](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_iam-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// `AWS::ACMPCA::CertificateAuthority.UsageMode`.
	UsageMode *string `field:"optional" json:"usageMode" yaml:"usageMode"`
}

