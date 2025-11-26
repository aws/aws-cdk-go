package previewawsacmpcamixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnCertificateAuthorityPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCertificateAuthorityMixinProps := &CfnCertificateAuthorityMixinProps{
//   	CsrExtensions: &CsrExtensionsProperty{
//   		KeyUsage: &KeyUsageProperty{
//   			CrlSign: jsii.Boolean(false),
//   			DataEncipherment: jsii.Boolean(false),
//   			DecipherOnly: jsii.Boolean(false),
//   			DigitalSignature: jsii.Boolean(false),
//   			EncipherOnly: jsii.Boolean(false),
//   			KeyAgreement: jsii.Boolean(false),
//   			KeyCertSign: jsii.Boolean(false),
//   			KeyEncipherment: jsii.Boolean(false),
//   			NonRepudiation: jsii.Boolean(false),
//   		},
//   		SubjectInformationAccess: []interface{}{
//   			&AccessDescriptionProperty{
//   				AccessLocation: &GeneralNameProperty{
//   					DirectoryName: &SubjectProperty{
//   						CommonName: jsii.String("commonName"),
//   						Country: jsii.String("country"),
//   						CustomAttributes: []interface{}{
//   							&CustomAttributeProperty{
//   								ObjectIdentifier: jsii.String("objectIdentifier"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						DistinguishedNameQualifier: jsii.String("distinguishedNameQualifier"),
//   						GenerationQualifier: jsii.String("generationQualifier"),
//   						GivenName: jsii.String("givenName"),
//   						Initials: jsii.String("initials"),
//   						Locality: jsii.String("locality"),
//   						Organization: jsii.String("organization"),
//   						OrganizationalUnit: jsii.String("organizationalUnit"),
//   						Pseudonym: jsii.String("pseudonym"),
//   						SerialNumber: jsii.String("serialNumber"),
//   						State: jsii.String("state"),
//   						Surname: jsii.String("surname"),
//   						Title: jsii.String("title"),
//   					},
//   					DnsName: jsii.String("dnsName"),
//   					EdiPartyName: &EdiPartyNameProperty{
//   						NameAssigner: jsii.String("nameAssigner"),
//   						PartyName: jsii.String("partyName"),
//   					},
//   					IpAddress: jsii.String("ipAddress"),
//   					OtherName: &OtherNameProperty{
//   						TypeId: jsii.String("typeId"),
//   						Value: jsii.String("value"),
//   					},
//   					RegisteredId: jsii.String("registeredId"),
//   					Rfc822Name: jsii.String("rfc822Name"),
//   					UniformResourceIdentifier: jsii.String("uniformResourceIdentifier"),
//   				},
//   				AccessMethod: &AccessMethodProperty{
//   					AccessMethodType: jsii.String("accessMethodType"),
//   					CustomObjectIdentifier: jsii.String("customObjectIdentifier"),
//   				},
//   			},
//   		},
//   	},
//   	KeyAlgorithm: jsii.String("keyAlgorithm"),
//   	KeyStorageSecurityStandard: jsii.String("keyStorageSecurityStandard"),
//   	RevocationConfiguration: &RevocationConfigurationProperty{
//   		CrlConfiguration: &CrlConfigurationProperty{
//   			CrlDistributionPointExtensionConfiguration: &CrlDistributionPointExtensionConfigurationProperty{
//   				OmitExtension: jsii.Boolean(false),
//   			},
//   			CrlType: jsii.String("crlType"),
//   			CustomCname: jsii.String("customCname"),
//   			CustomPath: jsii.String("customPath"),
//   			Enabled: jsii.Boolean(false),
//   			ExpirationInDays: jsii.Number(123),
//   			S3BucketName: jsii.String("s3BucketName"),
//   			S3ObjectAcl: jsii.String("s3ObjectAcl"),
//   		},
//   		OcspConfiguration: &OcspConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			OcspCustomCname: jsii.String("ocspCustomCname"),
//   		},
//   	},
//   	SigningAlgorithm: jsii.String("signingAlgorithm"),
//   	Subject: &SubjectProperty{
//   		CommonName: jsii.String("commonName"),
//   		Country: jsii.String("country"),
//   		CustomAttributes: []interface{}{
//   			&CustomAttributeProperty{
//   				ObjectIdentifier: jsii.String("objectIdentifier"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		DistinguishedNameQualifier: jsii.String("distinguishedNameQualifier"),
//   		GenerationQualifier: jsii.String("generationQualifier"),
//   		GivenName: jsii.String("givenName"),
//   		Initials: jsii.String("initials"),
//   		Locality: jsii.String("locality"),
//   		Organization: jsii.String("organization"),
//   		OrganizationalUnit: jsii.String("organizationalUnit"),
//   		Pseudonym: jsii.String("pseudonym"),
//   		SerialNumber: jsii.String("serialNumber"),
//   		State: jsii.String("state"),
//   		Surname: jsii.String("surname"),
//   		Title: jsii.String("title"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   	UsageMode: jsii.String("usageMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html
//
type CfnCertificateAuthorityMixinProps struct {
	// Specifies information to be added to the extension section of the certificate signing request (CSR).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html#cfn-acmpca-certificateauthority-csrextensions
	//
	CsrExtensions interface{} `field:"optional" json:"csrExtensions" yaml:"csrExtensions"`
	// Type of the public key algorithm and size, in bits, of the key pair that your CA creates when it issues a certificate.
	//
	// When you create a subordinate CA, you must use a key algorithm supported by the parent CA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html#cfn-acmpca-certificateauthority-keyalgorithm
	//
	KeyAlgorithm *string `field:"optional" json:"keyAlgorithm" yaml:"keyAlgorithm"`
	// Specifies a cryptographic key management compliance standard for handling and protecting CA keys.
	//
	// Default: FIPS_140_2_LEVEL_3_OR_HIGHER
	//
	// > Some AWS Regions don't support the default value. When you create a CA in these Regions, you must use `CCPC_LEVEL_1_OR_HIGHER` for the `KeyStorageSecurityStandard` parameter. If you don't, the operation returns an `InvalidArgsException` with this message: "A certificate authority cannot be created in this region with the specified security standard."
	// >
	// > For information about security standard support in different AWS Regions, see [Storage and security compliance of AWS Private CA private keys](https://docs.aws.amazon.com/privateca/latest/userguide/data-protection.html#private-keys) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html#cfn-acmpca-certificateauthority-keystoragesecuritystandard
	//
	KeyStorageSecurityStandard *string `field:"optional" json:"keyStorageSecurityStandard" yaml:"keyStorageSecurityStandard"`
	// Information about the Online Certificate Status Protocol (OCSP) configuration or certificate revocation list (CRL) created and maintained by your private CA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html#cfn-acmpca-certificateauthority-revocationconfiguration
	//
	RevocationConfiguration interface{} `field:"optional" json:"revocationConfiguration" yaml:"revocationConfiguration"`
	// Name of the algorithm your private CA uses to sign certificate requests.
	//
	// This parameter should not be confused with the `SigningAlgorithm` parameter used to sign certificates when they are issued.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html#cfn-acmpca-certificateauthority-signingalgorithm
	//
	SigningAlgorithm *string `field:"optional" json:"signingAlgorithm" yaml:"signingAlgorithm"`
	// Structure that contains X.500 distinguished name information for your private CA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html#cfn-acmpca-certificateauthority-subject
	//
	Subject interface{} `field:"optional" json:"subject" yaml:"subject"`
	// Key-value pairs that will be attached to the new private CA.
	//
	// You can associate up to 50 tags with a private CA. For information using tags with IAM to manage permissions, see [Controlling Access Using IAM Tags](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_iam-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html#cfn-acmpca-certificateauthority-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Type of your private CA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html#cfn-acmpca-certificateauthority-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Specifies whether the CA issues general-purpose certificates that typically require a revocation mechanism, or short-lived certificates that may optionally omit revocation because they expire quickly.
	//
	// Short-lived certificate validity is limited to seven days.
	//
	// The default value is GENERAL_PURPOSE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html#cfn-acmpca-certificateauthority-usagemode
	//
	UsageMode *string `field:"optional" json:"usageMode" yaml:"usageMode"`
}

