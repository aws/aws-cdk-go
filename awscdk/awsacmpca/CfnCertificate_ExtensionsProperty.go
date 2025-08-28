package awsacmpca


// Contains X.509 extension information for a certificate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   extensionsProperty := &ExtensionsProperty{
//   	CertificatePolicies: []interface{}{
//   		&PolicyInformationProperty{
//   			CertPolicyId: jsii.String("certPolicyId"),
//
//   			// the properties below are optional
//   			PolicyQualifiers: []interface{}{
//   				&PolicyQualifierInfoProperty{
//   					PolicyQualifierId: jsii.String("policyQualifierId"),
//   					Qualifier: &QualifierProperty{
//   						CpsUri: jsii.String("cpsUri"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	CustomExtensions: []interface{}{
//   		&CustomExtensionProperty{
//   			ObjectIdentifier: jsii.String("objectIdentifier"),
//   			Value: jsii.String("value"),
//
//   			// the properties below are optional
//   			Critical: jsii.Boolean(false),
//   		},
//   	},
//   	ExtendedKeyUsage: []interface{}{
//   		&ExtendedKeyUsageProperty{
//   			ExtendedKeyUsageObjectIdentifier: jsii.String("extendedKeyUsageObjectIdentifier"),
//   			ExtendedKeyUsageType: jsii.String("extendedKeyUsageType"),
//   		},
//   	},
//   	KeyUsage: &KeyUsageProperty{
//   		CrlSign: jsii.Boolean(false),
//   		DataEncipherment: jsii.Boolean(false),
//   		DecipherOnly: jsii.Boolean(false),
//   		DigitalSignature: jsii.Boolean(false),
//   		EncipherOnly: jsii.Boolean(false),
//   		KeyAgreement: jsii.Boolean(false),
//   		KeyCertSign: jsii.Boolean(false),
//   		KeyEncipherment: jsii.Boolean(false),
//   		NonRepudiation: jsii.Boolean(false),
//   	},
//   	SubjectAlternativeNames: []interface{}{
//   		&GeneralNameProperty{
//   			DirectoryName: &SubjectProperty{
//   				CommonName: jsii.String("commonName"),
//   				Country: jsii.String("country"),
//   				CustomAttributes: []interface{}{
//   					&CustomAttributeProperty{
//   						ObjectIdentifier: jsii.String("objectIdentifier"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				DistinguishedNameQualifier: jsii.String("distinguishedNameQualifier"),
//   				GenerationQualifier: jsii.String("generationQualifier"),
//   				GivenName: jsii.String("givenName"),
//   				Initials: jsii.String("initials"),
//   				Locality: jsii.String("locality"),
//   				Organization: jsii.String("organization"),
//   				OrganizationalUnit: jsii.String("organizationalUnit"),
//   				Pseudonym: jsii.String("pseudonym"),
//   				SerialNumber: jsii.String("serialNumber"),
//   				State: jsii.String("state"),
//   				Surname: jsii.String("surname"),
//   				Title: jsii.String("title"),
//   			},
//   			DnsName: jsii.String("dnsName"),
//   			EdiPartyName: &EdiPartyNameProperty{
//   				NameAssigner: jsii.String("nameAssigner"),
//   				PartyName: jsii.String("partyName"),
//   			},
//   			IpAddress: jsii.String("ipAddress"),
//   			OtherName: &OtherNameProperty{
//   				TypeId: jsii.String("typeId"),
//   				Value: jsii.String("value"),
//   			},
//   			RegisteredId: jsii.String("registeredId"),
//   			Rfc822Name: jsii.String("rfc822Name"),
//   			UniformResourceIdentifier: jsii.String("uniformResourceIdentifier"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-extensions.html
//
type CfnCertificate_ExtensionsProperty struct {
	// Contains a sequence of one or more policy information terms, each of which consists of an object identifier (OID) and optional qualifiers.
	//
	// For more information, see NIST's definition of [Object Identifier (OID)](https://docs.aws.amazon.com/https://csrc.nist.gov/glossary/term/Object_Identifier) .
	//
	// In an end-entity certificate, these terms indicate the policy under which the certificate was issued and the purposes for which it may be used. In a CA certificate, these terms limit the set of policies for certification paths that include this certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-extensions.html#cfn-acmpca-certificate-extensions-certificatepolicies
	//
	CertificatePolicies interface{} `field:"optional" json:"certificatePolicies" yaml:"certificatePolicies"`
	// Contains a sequence of one or more X.509 extensions, each of which consists of an object identifier (OID), a base64-encoded value, and the critical flag. For more information, see the [Global OID reference database.](https://docs.aws.amazon.com/https://oidref.com/2.5.29).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-extensions.html#cfn-acmpca-certificate-extensions-customextensions
	//
	CustomExtensions interface{} `field:"optional" json:"customExtensions" yaml:"customExtensions"`
	// Specifies additional purposes for which the certified public key may be used other than basic purposes indicated in the `KeyUsage` extension.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-extensions.html#cfn-acmpca-certificate-extensions-extendedkeyusage
	//
	ExtendedKeyUsage interface{} `field:"optional" json:"extendedKeyUsage" yaml:"extendedKeyUsage"`
	// Defines one or more purposes for which the key contained in the certificate can be used.
	//
	// Default value for each option is false.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-extensions.html#cfn-acmpca-certificate-extensions-keyusage
	//
	KeyUsage interface{} `field:"optional" json:"keyUsage" yaml:"keyUsage"`
	// The subject alternative name extension allows identities to be bound to the subject of the certificate.
	//
	// These identities may be included in addition to or in place of the identity in the subject field of the certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-extensions.html#cfn-acmpca-certificate-extensions-subjectalternativenames
	//
	SubjectAlternativeNames interface{} `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}

