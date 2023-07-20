package awsacmpca

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCertificateAuthority`.
//
// Example:
//   cfnCertificateAuthority := acmpca.NewCfnCertificateAuthority(this, jsii.String("CA"), &CfnCertificateAuthorityProps{
//   	Type: jsii.String("ROOT"),
//   	KeyAlgorithm: jsii.String("RSA_2048"),
//   	SigningAlgorithm: jsii.String("SHA256WITHRSA"),
//   	Subject: &SubjectProperty{
//   		Country: jsii.String("US"),
//   		Organization: jsii.String("string"),
//   		OrganizationalUnit: jsii.String("string"),
//   		DistinguishedNameQualifier: jsii.String("string"),
//   		State: jsii.String("string"),
//   		CommonName: jsii.String("123"),
//   		SerialNumber: jsii.String("string"),
//   		Locality: jsii.String("string"),
//   		Title: jsii.String("string"),
//   		Surname: jsii.String("string"),
//   		GivenName: jsii.String("string"),
//   		Initials: jsii.String("DG"),
//   		Pseudonym: jsii.String("string"),
//   		GenerationQualifier: jsii.String("DBG"),
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html
//
type CfnCertificateAuthorityProps struct {
	// Type of the public key algorithm and size, in bits, of the key pair that your CA creates when it issues a certificate.
	//
	// When you create a subordinate CA, you must use a key algorithm supported by the parent CA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html#cfn-acmpca-certificateauthority-keyalgorithm
	//
	KeyAlgorithm *string `field:"required" json:"keyAlgorithm" yaml:"keyAlgorithm"`
	// Name of the algorithm your private CA uses to sign certificate requests.
	//
	// This parameter should not be confused with the `SigningAlgorithm` parameter used to sign certificates when they are issued.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html#cfn-acmpca-certificateauthority-signingalgorithm
	//
	SigningAlgorithm *string `field:"required" json:"signingAlgorithm" yaml:"signingAlgorithm"`
	// Structure that contains X.500 distinguished name information for your private CA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html#cfn-acmpca-certificateauthority-subject
	//
	Subject interface{} `field:"required" json:"subject" yaml:"subject"`
	// Type of your private CA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html#cfn-acmpca-certificateauthority-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Specifies information to be added to the extension section of the certificate signing request (CSR).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html#cfn-acmpca-certificateauthority-csrextensions
	//
	CsrExtensions interface{} `field:"optional" json:"csrExtensions" yaml:"csrExtensions"`
	// Specifies a cryptographic key management compliance standard used for handling CA keys.
	//
	// Default: FIPS_140_2_LEVEL_3_OR_HIGHER
	//
	// > Some AWS Regions do not support the default. When creating a CA in these Regions, you must provide `FIPS_140_2_LEVEL_2_OR_HIGHER` as the argument for `KeyStorageSecurityStandard` . Failure to do this results in an `InvalidArgsException` with the message, "A certificate authority cannot be created in this region with the specified security standard."
	// >
	// > For information about security standard support in various Regions, see [Storage and security compliance of AWS Private CA private keys](https://docs.aws.amazon.com/privateca/latest/userguide/data-protection.html#private-keys) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html#cfn-acmpca-certificateauthority-keystoragesecuritystandard
	//
	KeyStorageSecurityStandard *string `field:"optional" json:"keyStorageSecurityStandard" yaml:"keyStorageSecurityStandard"`
	// Certificate revocation information used by the [CreateCertificateAuthority](https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html) and [UpdateCertificateAuthority](https://docs.aws.amazon.com/privateca/latest/APIReference/API_UpdateCertificateAuthority.html) actions. Your private certificate authority (CA) can configure Online Certificate Status Protocol (OCSP) support and/or maintain a certificate revocation list (CRL). OCSP returns validation information about certificates as requested by clients, and a CRL contains an updated list of certificates revoked by your CA. For more information, see [RevokeCertificate](https://docs.aws.amazon.com/privateca/latest/APIReference/API_RevokeCertificate.html) in the *AWS Private CA API Reference* and [Setting up a certificate revocation method](https://docs.aws.amazon.com/privateca/latest/userguide/revocation-setup.html) in the *AWS Private CA User Guide* .
	//
	// > The following requirements apply to revocation configurations.
	// >
	// > - A configuration disabling CRLs or OCSP must contain only the `Enabled=False` parameter, and will fail if other parameters such as `CustomCname` or `ExpirationInDays` are included.
	// > - In a CRL configuration, the `S3BucketName` parameter must conform to the [Amazon S3 bucket naming rules](https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html) .
	// > - A configuration containing a custom Canonical Name (CNAME) parameter for CRLs or OCSP must conform to [RFC2396](https://docs.aws.amazon.com/https://www.ietf.org/rfc/rfc2396.txt) restrictions on the use of special characters in a CNAME.
	// > - In a CRL or OCSP configuration, the value of a CNAME parameter must not include a protocol prefix such as "http://" or "https://".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html#cfn-acmpca-certificateauthority-revocationconfiguration
	//
	RevocationConfiguration interface{} `field:"optional" json:"revocationConfiguration" yaml:"revocationConfiguration"`
	// Key-value pairs that will be attached to the new private CA.
	//
	// You can associate up to 50 tags with a private CA. For information using tags with IAM to manage permissions, see [Controlling Access Using IAM Tags](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_iam-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html#cfn-acmpca-certificateauthority-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specifies whether the CA issues general-purpose certificates that typically require a revocation mechanism, or short-lived certificates that may optionally omit revocation because they expire quickly.
	//
	// Short-lived certificate validity is limited to seven days.
	//
	// The default value is GENERAL_PURPOSE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html#cfn-acmpca-certificateauthority-usagemode
	//
	UsageMode *string `field:"optional" json:"usageMode" yaml:"usageMode"`
}

