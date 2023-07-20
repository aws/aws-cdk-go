package awsacmpca


// Contains X.509 certificate information to be placed in an issued certificate. An `APIPassthrough` or `APICSRPassthrough` template variant must be selected, or else this parameter is ignored.
//
// If conflicting or duplicate certificate information is supplied from other sources, AWS Private CA applies [order of operation rules](https://docs.aws.amazon.com/privateca/latest/userguide/UsingTemplates.html#template-order-of-operations) to determine what information is used.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiPassthroughProperty := &ApiPassthroughProperty{
//   	Extensions: &ExtensionsProperty{
//   		CertificatePolicies: []interface{}{
//   			&PolicyInformationProperty{
//   				CertPolicyId: jsii.String("certPolicyId"),
//
//   				// the properties below are optional
//   				PolicyQualifiers: []interface{}{
//   					&PolicyQualifierInfoProperty{
//   						PolicyQualifierId: jsii.String("policyQualifierId"),
//   						Qualifier: &QualifierProperty{
//   							CpsUri: jsii.String("cpsUri"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		CustomExtensions: []interface{}{
//   			&CustomExtensionProperty{
//   				ObjectIdentifier: jsii.String("objectIdentifier"),
//   				Value: jsii.String("value"),
//
//   				// the properties below are optional
//   				Critical: jsii.Boolean(false),
//   			},
//   		},
//   		ExtendedKeyUsage: []interface{}{
//   			&ExtendedKeyUsageProperty{
//   				ExtendedKeyUsageObjectIdentifier: jsii.String("extendedKeyUsageObjectIdentifier"),
//   				ExtendedKeyUsageType: jsii.String("extendedKeyUsageType"),
//   			},
//   		},
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
//   		SubjectAlternativeNames: []interface{}{
//   			&GeneralNameProperty{
//   				DirectoryName: &SubjectProperty{
//   					CommonName: jsii.String("commonName"),
//   					Country: jsii.String("country"),
//   					CustomAttributes: []interface{}{
//   						&CustomAttributeProperty{
//   							ObjectIdentifier: jsii.String("objectIdentifier"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					DistinguishedNameQualifier: jsii.String("distinguishedNameQualifier"),
//   					GenerationQualifier: jsii.String("generationQualifier"),
//   					GivenName: jsii.String("givenName"),
//   					Initials: jsii.String("initials"),
//   					Locality: jsii.String("locality"),
//   					Organization: jsii.String("organization"),
//   					OrganizationalUnit: jsii.String("organizationalUnit"),
//   					Pseudonym: jsii.String("pseudonym"),
//   					SerialNumber: jsii.String("serialNumber"),
//   					State: jsii.String("state"),
//   					Surname: jsii.String("surname"),
//   					Title: jsii.String("title"),
//   				},
//   				DnsName: jsii.String("dnsName"),
//   				EdiPartyName: &EdiPartyNameProperty{
//   					NameAssigner: jsii.String("nameAssigner"),
//   					PartyName: jsii.String("partyName"),
//   				},
//   				IpAddress: jsii.String("ipAddress"),
//   				OtherName: &OtherNameProperty{
//   					TypeId: jsii.String("typeId"),
//   					Value: jsii.String("value"),
//   				},
//   				RegisteredId: jsii.String("registeredId"),
//   				Rfc822Name: jsii.String("rfc822Name"),
//   				UniformResourceIdentifier: jsii.String("uniformResourceIdentifier"),
//   			},
//   		},
//   	},
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-apipassthrough.html
//
type CfnCertificate_ApiPassthroughProperty struct {
	// Specifies X.509 extension information for a certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-apipassthrough.html#cfn-acmpca-certificate-apipassthrough-extensions
	//
	Extensions interface{} `field:"optional" json:"extensions" yaml:"extensions"`
	// Contains information about the certificate subject.
	//
	// The Subject field in the certificate identifies the entity that owns or controls the public key in the certificate. The entity can be a user, computer, device, or service. The Subject must contain an X.500 distinguished name (DN). A DN is a sequence of relative distinguished names (RDNs). The RDNs are separated by commas in the certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-apipassthrough.html#cfn-acmpca-certificate-apipassthrough-subject
	//
	Subject interface{} `field:"optional" json:"subject" yaml:"subject"`
}

