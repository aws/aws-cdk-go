package awsacmpca


// Contains X.509 extension information for a certificate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   extensionsProperty := &extensionsProperty{
//   	certificatePolicies: []interface{}{
//   		&policyInformationProperty{
//   			certPolicyId: jsii.String("certPolicyId"),
//
//   			// the properties below are optional
//   			policyQualifiers: []interface{}{
//   				&policyQualifierInfoProperty{
//   					policyQualifierId: jsii.String("policyQualifierId"),
//   					qualifier: &qualifierProperty{
//   						cpsUri: jsii.String("cpsUri"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	customExtensions: []interface{}{
//   		&customExtensionProperty{
//   			objectIdentifier: jsii.String("objectIdentifier"),
//   			value: jsii.String("value"),
//
//   			// the properties below are optional
//   			critical: jsii.Boolean(false),
//   		},
//   	},
//   	extendedKeyUsage: []interface{}{
//   		&extendedKeyUsageProperty{
//   			extendedKeyUsageObjectIdentifier: jsii.String("extendedKeyUsageObjectIdentifier"),
//   			extendedKeyUsageType: jsii.String("extendedKeyUsageType"),
//   		},
//   	},
//   	keyUsage: &keyUsageProperty{
//   		crlSign: jsii.Boolean(false),
//   		dataEncipherment: jsii.Boolean(false),
//   		decipherOnly: jsii.Boolean(false),
//   		digitalSignature: jsii.Boolean(false),
//   		encipherOnly: jsii.Boolean(false),
//   		keyAgreement: jsii.Boolean(false),
//   		keyCertSign: jsii.Boolean(false),
//   		keyEncipherment: jsii.Boolean(false),
//   		nonRepudiation: jsii.Boolean(false),
//   	},
//   	subjectAlternativeNames: []interface{}{
//   		&generalNameProperty{
//   			directoryName: &subjectProperty{
//   				commonName: jsii.String("commonName"),
//   				country: jsii.String("country"),
//   				customAttributes: []interface{}{
//   					&customAttributeProperty{
//   						objectIdentifier: jsii.String("objectIdentifier"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				distinguishedNameQualifier: jsii.String("distinguishedNameQualifier"),
//   				generationQualifier: jsii.String("generationQualifier"),
//   				givenName: jsii.String("givenName"),
//   				initials: jsii.String("initials"),
//   				locality: jsii.String("locality"),
//   				organization: jsii.String("organization"),
//   				organizationalUnit: jsii.String("organizationalUnit"),
//   				pseudonym: jsii.String("pseudonym"),
//   				serialNumber: jsii.String("serialNumber"),
//   				state: jsii.String("state"),
//   				surname: jsii.String("surname"),
//   				title: jsii.String("title"),
//   			},
//   			dnsName: jsii.String("dnsName"),
//   			ediPartyName: &ediPartyNameProperty{
//   				nameAssigner: jsii.String("nameAssigner"),
//   				partyName: jsii.String("partyName"),
//   			},
//   			ipAddress: jsii.String("ipAddress"),
//   			otherName: &otherNameProperty{
//   				typeId: jsii.String("typeId"),
//   				value: jsii.String("value"),
//   			},
//   			registeredId: jsii.String("registeredId"),
//   			rfc822Name: jsii.String("rfc822Name"),
//   			uniformResourceIdentifier: jsii.String("uniformResourceIdentifier"),
//   		},
//   	},
//   }
//
type CfnCertificate_ExtensionsProperty struct {
	// Contains a sequence of one or more policy information terms, each of which consists of an object identifier (OID) and optional qualifiers.
	//
	// For more information, see NIST's definition of [Object Identifier (OID)](https://docs.aws.amazon.com/https://csrc.nist.gov/glossary/term/Object_Identifier) .
	//
	// In an end-entity certificate, these terms indicate the policy under which the certificate was issued and the purposes for which it may be used. In a CA certificate, these terms limit the set of policies for certification paths that include this certificate.
	CertificatePolicies interface{} `field:"optional" json:"certificatePolicies" yaml:"certificatePolicies"`
	// Contains a sequence of one or more X.509 extensions, each of which consists of an object identifier (OID), a base64-encoded value, and the critical flag. For more information, see the [Global OID reference database.](https://docs.aws.amazon.com/https://oidref.com/2.5.29).
	//
	// > The OID value of a [CustomExtension](https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CustomExtension.html) must not match the OID of a predefined extension.
	CustomExtensions interface{} `field:"optional" json:"customExtensions" yaml:"customExtensions"`
	// Specifies additional purposes for which the certified public key may be used other than basic purposes indicated in the `KeyUsage` extension.
	ExtendedKeyUsage interface{} `field:"optional" json:"extendedKeyUsage" yaml:"extendedKeyUsage"`
	// Defines one or more purposes for which the key contained in the certificate can be used.
	//
	// Default value for each option is false.
	KeyUsage interface{} `field:"optional" json:"keyUsage" yaml:"keyUsage"`
	// The subject alternative name extension allows identities to be bound to the subject of the certificate.
	//
	// These identities may be included in addition to or in place of the identity in the subject field of the certificate.
	SubjectAlternativeNames interface{} `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}

