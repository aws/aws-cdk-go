package awsacmpca


// Describes the certificate extensions to be added to the certificate signing request (CSR).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   csrExtensionsProperty := &CsrExtensionsProperty{
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
//   	SubjectInformationAccess: []interface{}{
//   		&AccessDescriptionProperty{
//   			AccessLocation: &GeneralNameProperty{
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
//   			AccessMethod: &AccessMethodProperty{
//   				AccessMethodType: jsii.String("accessMethodType"),
//   				CustomObjectIdentifier: jsii.String("customObjectIdentifier"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-csrextensions.html
//
type CfnCertificateAuthority_CsrExtensionsProperty struct {
	// Indicates the purpose of the certificate and of the key contained in the certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-csrextensions.html#cfn-acmpca-certificateauthority-csrextensions-keyusage
	//
	KeyUsage interface{} `field:"optional" json:"keyUsage" yaml:"keyUsage"`
	// For CA certificates, provides a path to additional information pertaining to the CA, such as revocation and policy.
	//
	// For more information, see [Subject Information Access](https://docs.aws.amazon.com/https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.2.2) in RFC 5280.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-csrextensions.html#cfn-acmpca-certificateauthority-csrextensions-subjectinformationaccess
	//
	SubjectInformationAccess interface{} `field:"optional" json:"subjectInformationAccess" yaml:"subjectInformationAccess"`
}

