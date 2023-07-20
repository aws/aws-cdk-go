package awsacmpca


// Properties for defining a `CfnCertificate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCertificateProps := &CfnCertificateProps{
//   	CertificateAuthorityArn: jsii.String("certificateAuthorityArn"),
//   	CertificateSigningRequest: jsii.String("certificateSigningRequest"),
//   	SigningAlgorithm: jsii.String("signingAlgorithm"),
//   	Validity: &ValidityProperty{
//   		Type: jsii.String("type"),
//   		Value: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	ApiPassthrough: &ApiPassthroughProperty{
//   		Extensions: &ExtensionsProperty{
//   			CertificatePolicies: []interface{}{
//   				&PolicyInformationProperty{
//   					CertPolicyId: jsii.String("certPolicyId"),
//
//   					// the properties below are optional
//   					PolicyQualifiers: []interface{}{
//   						&PolicyQualifierInfoProperty{
//   							PolicyQualifierId: jsii.String("policyQualifierId"),
//   							Qualifier: &QualifierProperty{
//   								CpsUri: jsii.String("cpsUri"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			CustomExtensions: []interface{}{
//   				&CustomExtensionProperty{
//   					ObjectIdentifier: jsii.String("objectIdentifier"),
//   					Value: jsii.String("value"),
//
//   					// the properties below are optional
//   					Critical: jsii.Boolean(false),
//   				},
//   			},
//   			ExtendedKeyUsage: []interface{}{
//   				&ExtendedKeyUsageProperty{
//   					ExtendedKeyUsageObjectIdentifier: jsii.String("extendedKeyUsageObjectIdentifier"),
//   					ExtendedKeyUsageType: jsii.String("extendedKeyUsageType"),
//   				},
//   			},
//   			KeyUsage: &KeyUsageProperty{
//   				CrlSign: jsii.Boolean(false),
//   				DataEncipherment: jsii.Boolean(false),
//   				DecipherOnly: jsii.Boolean(false),
//   				DigitalSignature: jsii.Boolean(false),
//   				EncipherOnly: jsii.Boolean(false),
//   				KeyAgreement: jsii.Boolean(false),
//   				KeyCertSign: jsii.Boolean(false),
//   				KeyEncipherment: jsii.Boolean(false),
//   				NonRepudiation: jsii.Boolean(false),
//   			},
//   			SubjectAlternativeNames: []interface{}{
//   				&GeneralNameProperty{
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
//   			},
//   		},
//   		Subject: &SubjectProperty{
//   			CommonName: jsii.String("commonName"),
//   			Country: jsii.String("country"),
//   			CustomAttributes: []interface{}{
//   				&CustomAttributeProperty{
//   					ObjectIdentifier: jsii.String("objectIdentifier"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			DistinguishedNameQualifier: jsii.String("distinguishedNameQualifier"),
//   			GenerationQualifier: jsii.String("generationQualifier"),
//   			GivenName: jsii.String("givenName"),
//   			Initials: jsii.String("initials"),
//   			Locality: jsii.String("locality"),
//   			Organization: jsii.String("organization"),
//   			OrganizationalUnit: jsii.String("organizationalUnit"),
//   			Pseudonym: jsii.String("pseudonym"),
//   			SerialNumber: jsii.String("serialNumber"),
//   			State: jsii.String("state"),
//   			Surname: jsii.String("surname"),
//   			Title: jsii.String("title"),
//   		},
//   	},
//   	TemplateArn: jsii.String("templateArn"),
//   	ValidityNotBefore: &ValidityProperty{
//   		Type: jsii.String("type"),
//   		Value: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificate.html
//
type CfnCertificateProps struct {
	// The Amazon Resource Name (ARN) for the private CA issues the certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificate.html#cfn-acmpca-certificate-certificateauthorityarn
	//
	CertificateAuthorityArn *string `field:"required" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// The certificate signing request (CSR) for the certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificate.html#cfn-acmpca-certificate-certificatesigningrequest
	//
	CertificateSigningRequest *string `field:"required" json:"certificateSigningRequest" yaml:"certificateSigningRequest"`
	// The name of the algorithm that will be used to sign the certificate to be issued.
	//
	// This parameter should not be confused with the `SigningAlgorithm` parameter used to sign a CSR in the `CreateCertificateAuthority` action.
	//
	// > The specified signing algorithm family (RSA or ECDSA) must match the algorithm family of the CA's secret key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificate.html#cfn-acmpca-certificate-signingalgorithm
	//
	SigningAlgorithm *string `field:"required" json:"signingAlgorithm" yaml:"signingAlgorithm"`
	// The period of time during which the certificate will be valid.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificate.html#cfn-acmpca-certificate-validity
	//
	Validity interface{} `field:"required" json:"validity" yaml:"validity"`
	// Specifies X.509 certificate information to be included in the issued certificate. An `APIPassthrough` or `APICSRPassthrough` template variant must be selected, or else this parameter is ignored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificate.html#cfn-acmpca-certificate-apipassthrough
	//
	ApiPassthrough interface{} `field:"optional" json:"apiPassthrough" yaml:"apiPassthrough"`
	// Specifies a custom configuration template to use when issuing a certificate.
	//
	// If this parameter is not provided, AWS Private CA defaults to the `EndEntityCertificate/V1` template. For more information about AWS Private CA templates, see [Using Templates](https://docs.aws.amazon.com/privateca/latest/userguide/UsingTemplates.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificate.html#cfn-acmpca-certificate-templatearn
	//
	TemplateArn *string `field:"optional" json:"templateArn" yaml:"templateArn"`
	// Information describing the start of the validity period of the certificate.
	//
	// This parameter sets the “Not Before" date for the certificate.
	//
	// By default, when issuing a certificate, AWS Private CA sets the "Not Before" date to the issuance time minus 60 minutes. This compensates for clock inconsistencies across computer systems. The `ValidityNotBefore` parameter can be used to customize the “Not Before” value.
	//
	// Unlike the `Validity` parameter, the `ValidityNotBefore` parameter is optional.
	//
	// The `ValidityNotBefore` value is expressed as an explicit date and time, using the `Validity` type value `ABSOLUTE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificate.html#cfn-acmpca-certificate-validitynotbefore
	//
	ValidityNotBefore interface{} `field:"optional" json:"validityNotBefore" yaml:"validityNotBefore"`
}

