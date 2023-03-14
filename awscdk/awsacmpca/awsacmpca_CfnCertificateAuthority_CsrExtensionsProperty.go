package awsacmpca


// Describes the certificate extensions to be added to the certificate signing request (CSR).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   csrExtensionsProperty := &csrExtensionsProperty{
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
//   	subjectInformationAccess: []interface{}{
//   		&accessDescriptionProperty{
//   			accessLocation: &generalNameProperty{
//   				directoryName: &subjectProperty{
//   					commonName: jsii.String("commonName"),
//   					country: jsii.String("country"),
//   					customAttributes: []interface{}{
//   						&customAttributeProperty{
//   							objectIdentifier: jsii.String("objectIdentifier"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					distinguishedNameQualifier: jsii.String("distinguishedNameQualifier"),
//   					generationQualifier: jsii.String("generationQualifier"),
//   					givenName: jsii.String("givenName"),
//   					initials: jsii.String("initials"),
//   					locality: jsii.String("locality"),
//   					organization: jsii.String("organization"),
//   					organizationalUnit: jsii.String("organizationalUnit"),
//   					pseudonym: jsii.String("pseudonym"),
//   					serialNumber: jsii.String("serialNumber"),
//   					state: jsii.String("state"),
//   					surname: jsii.String("surname"),
//   					title: jsii.String("title"),
//   				},
//   				dnsName: jsii.String("dnsName"),
//   				ediPartyName: &ediPartyNameProperty{
//   					nameAssigner: jsii.String("nameAssigner"),
//   					partyName: jsii.String("partyName"),
//   				},
//   				ipAddress: jsii.String("ipAddress"),
//   				otherName: &otherNameProperty{
//   					typeId: jsii.String("typeId"),
//   					value: jsii.String("value"),
//   				},
//   				registeredId: jsii.String("registeredId"),
//   				rfc822Name: jsii.String("rfc822Name"),
//   				uniformResourceIdentifier: jsii.String("uniformResourceIdentifier"),
//   			},
//   			accessMethod: &accessMethodProperty{
//   				accessMethodType: jsii.String("accessMethodType"),
//   				customObjectIdentifier: jsii.String("customObjectIdentifier"),
//   			},
//   		},
//   	},
//   }
//
type CfnCertificateAuthority_CsrExtensionsProperty struct {
	// Indicates the purpose of the certificate and of the key contained in the certificate.
	KeyUsage interface{} `field:"optional" json:"keyUsage" yaml:"keyUsage"`
	// For CA certificates, provides a path to additional information pertaining to the CA, such as revocation and policy.
	//
	// For more information, see [Subject Information Access](https://docs.aws.amazon.com/https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.2.2) in RFC 5280.
	SubjectInformationAccess interface{} `field:"optional" json:"subjectInformationAccess" yaml:"subjectInformationAccess"`
}

