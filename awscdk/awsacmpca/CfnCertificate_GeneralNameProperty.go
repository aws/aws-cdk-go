package awsacmpca


// Describes an ASN.1 X.400 `GeneralName` as defined in [RFC 5280](https://docs.aws.amazon.com/https://datatracker.ietf.org/doc/html/rfc5280) . Only one of the following naming options should be provided. Providing more than one option results in an `InvalidArgsException` error.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   generalNameProperty := &GeneralNameProperty{
//   	DirectoryName: &SubjectProperty{
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
//   	DnsName: jsii.String("dnsName"),
//   	EdiPartyName: &EdiPartyNameProperty{
//   		NameAssigner: jsii.String("nameAssigner"),
//   		PartyName: jsii.String("partyName"),
//   	},
//   	IpAddress: jsii.String("ipAddress"),
//   	OtherName: &OtherNameProperty{
//   		TypeId: jsii.String("typeId"),
//   		Value: jsii.String("value"),
//   	},
//   	RegisteredId: jsii.String("registeredId"),
//   	Rfc822Name: jsii.String("rfc822Name"),
//   	UniformResourceIdentifier: jsii.String("uniformResourceIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-generalname.html
//
type CfnCertificate_GeneralNameProperty struct {
	// Contains information about the certificate subject.
	//
	// The certificate can be one issued by your private certificate authority (CA) or it can be your private CA certificate. The Subject field in the certificate identifies the entity that owns or controls the public key in the certificate. The entity can be a user, computer, device, or service. The Subject must contain an X.500 distinguished name (DN). A DN is a sequence of relative distinguished names (RDNs). The RDNs are separated by commas in the certificate. The DN must be unique for each entity, but your private CA can issue more than one certificate with the same DN to the same entity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-generalname.html#cfn-acmpca-certificate-generalname-directoryname
	//
	DirectoryName interface{} `field:"optional" json:"directoryName" yaml:"directoryName"`
	// Represents `GeneralName` as a DNS name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-generalname.html#cfn-acmpca-certificate-generalname-dnsname
	//
	DnsName *string `field:"optional" json:"dnsName" yaml:"dnsName"`
	// Represents `GeneralName` as an `EdiPartyName` object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-generalname.html#cfn-acmpca-certificate-generalname-edipartyname
	//
	EdiPartyName interface{} `field:"optional" json:"ediPartyName" yaml:"ediPartyName"`
	// Represents `GeneralName` as an IPv4 or IPv6 address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-generalname.html#cfn-acmpca-certificate-generalname-ipaddress
	//
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// Represents `GeneralName` using an `OtherName` object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-generalname.html#cfn-acmpca-certificate-generalname-othername
	//
	OtherName interface{} `field:"optional" json:"otherName" yaml:"otherName"`
	// Represents `GeneralName` as an object identifier (OID).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-generalname.html#cfn-acmpca-certificate-generalname-registeredid
	//
	RegisteredId *string `field:"optional" json:"registeredId" yaml:"registeredId"`
	// Represents `GeneralName` as an [RFC 822](https://docs.aws.amazon.com/https://datatracker.ietf.org/doc/html/rfc822) email address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-generalname.html#cfn-acmpca-certificate-generalname-rfc822name
	//
	Rfc822Name *string `field:"optional" json:"rfc822Name" yaml:"rfc822Name"`
	// Represents `GeneralName` as a URI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-generalname.html#cfn-acmpca-certificate-generalname-uniformresourceidentifier
	//
	UniformResourceIdentifier *string `field:"optional" json:"uniformResourceIdentifier" yaml:"uniformResourceIdentifier"`
}

