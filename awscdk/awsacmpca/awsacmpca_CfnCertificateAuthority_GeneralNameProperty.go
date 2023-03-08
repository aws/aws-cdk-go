package awsacmpca


// Describes an ASN.1 X.400 `GeneralName` as defined in [RFC 5280](https://docs.aws.amazon.com/https://datatracker.ietf.org/doc/html/rfc5280) . Only one of the following naming options should be provided. Providing more than one option results in an `InvalidArgsException` error.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   generalNameProperty := &generalNameProperty{
//   	directoryName: &subjectProperty{
//   		commonName: jsii.String("commonName"),
//   		country: jsii.String("country"),
//   		customAttributes: []interface{}{
//   			&customAttributeProperty{
//   				objectIdentifier: jsii.String("objectIdentifier"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		distinguishedNameQualifier: jsii.String("distinguishedNameQualifier"),
//   		generationQualifier: jsii.String("generationQualifier"),
//   		givenName: jsii.String("givenName"),
//   		initials: jsii.String("initials"),
//   		locality: jsii.String("locality"),
//   		organization: jsii.String("organization"),
//   		organizationalUnit: jsii.String("organizationalUnit"),
//   		pseudonym: jsii.String("pseudonym"),
//   		serialNumber: jsii.String("serialNumber"),
//   		state: jsii.String("state"),
//   		surname: jsii.String("surname"),
//   		title: jsii.String("title"),
//   	},
//   	dnsName: jsii.String("dnsName"),
//   	ediPartyName: &ediPartyNameProperty{
//   		nameAssigner: jsii.String("nameAssigner"),
//   		partyName: jsii.String("partyName"),
//   	},
//   	ipAddress: jsii.String("ipAddress"),
//   	otherName: &otherNameProperty{
//   		typeId: jsii.String("typeId"),
//   		value: jsii.String("value"),
//   	},
//   	registeredId: jsii.String("registeredId"),
//   	rfc822Name: jsii.String("rfc822Name"),
//   	uniformResourceIdentifier: jsii.String("uniformResourceIdentifier"),
//   }
//
type CfnCertificateAuthority_GeneralNameProperty struct {
	// Contains information about the certificate subject.
	//
	// The certificate can be one issued by your private certificate authority (CA) or it can be your private CA certificate. The Subject field in the certificate identifies the entity that owns or controls the public key in the certificate. The entity can be a user, computer, device, or service. The Subject must contain an X.500 distinguished name (DN). A DN is a sequence of relative distinguished names (RDNs). The RDNs are separated by commas in the certificate. The DN must be unique for each entity, but your private CA can issue more than one certificate with the same DN to the same entity.
	DirectoryName interface{} `field:"optional" json:"directoryName" yaml:"directoryName"`
	// Represents `GeneralName` as a DNS name.
	DnsName *string `field:"optional" json:"dnsName" yaml:"dnsName"`
	// Represents `GeneralName` as an `EdiPartyName` object.
	EdiPartyName interface{} `field:"optional" json:"ediPartyName" yaml:"ediPartyName"`
	// Represents `GeneralName` as an IPv4 or IPv6 address.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// Represents `GeneralName` using an `OtherName` object.
	OtherName interface{} `field:"optional" json:"otherName" yaml:"otherName"`
	// Represents `GeneralName` as an object identifier (OID).
	RegisteredId *string `field:"optional" json:"registeredId" yaml:"registeredId"`
	// Represents `GeneralName` as an [RFC 822](https://docs.aws.amazon.com/https://datatracker.ietf.org/doc/html/rfc822) email address.
	Rfc822Name *string `field:"optional" json:"rfc822Name" yaml:"rfc822Name"`
	// Represents `GeneralName` as a URI.
	UniformResourceIdentifier *string `field:"optional" json:"uniformResourceIdentifier" yaml:"uniformResourceIdentifier"`
}

