package awsacmpca


// ASN1 subject for the certificate authority.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subjectProperty := &subjectProperty{
//   	commonName: jsii.String("commonName"),
//   	country: jsii.String("country"),
//   	customAttributes: []interface{}{
//   		&customAttributeProperty{
//   			objectIdentifier: jsii.String("objectIdentifier"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	distinguishedNameQualifier: jsii.String("distinguishedNameQualifier"),
//   	generationQualifier: jsii.String("generationQualifier"),
//   	givenName: jsii.String("givenName"),
//   	initials: jsii.String("initials"),
//   	locality: jsii.String("locality"),
//   	organization: jsii.String("organization"),
//   	organizationalUnit: jsii.String("organizationalUnit"),
//   	pseudonym: jsii.String("pseudonym"),
//   	serialNumber: jsii.String("serialNumber"),
//   	state: jsii.String("state"),
//   	surname: jsii.String("surname"),
//   	title: jsii.String("title"),
//   }
//
type CfnCertificateAuthority_SubjectProperty struct {
	// Fully qualified domain name (FQDN) associated with the certificate subject.
	CommonName *string `field:"optional" json:"commonName" yaml:"commonName"`
	// Two-digit code that specifies the country in which the certificate subject located.
	Country *string `field:"optional" json:"country" yaml:"country"`
	// Contains a sequence of one or more X.500 relative distinguished names (RDNs), each of which consists of an object identifier (OID) and a value. For more information, see NIST’s definition of [Object Identifier (OID)](https://docs.aws.amazon.com/https://csrc.nist.gov/glossary/term/Object_Identifier) .
	//
	// > Custom attributes cannot be used in combination with standard attributes.
	CustomAttributes interface{} `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// Disambiguating information for the certificate subject.
	DistinguishedNameQualifier *string `field:"optional" json:"distinguishedNameQualifier" yaml:"distinguishedNameQualifier"`
	// Typically a qualifier appended to the name of an individual.
	//
	// Examples include Jr. for junior, Sr. for senior, and III for third.
	GenerationQualifier *string `field:"optional" json:"generationQualifier" yaml:"generationQualifier"`
	// First name.
	GivenName *string `field:"optional" json:"givenName" yaml:"givenName"`
	// Concatenation that typically contains the first letter of the GivenName, the first letter of the middle name if one exists, and the first letter of the SurName.
	Initials *string `field:"optional" json:"initials" yaml:"initials"`
	// The locality (such as a city or town) in which the certificate subject is located.
	Locality *string `field:"optional" json:"locality" yaml:"locality"`
	// Legal name of the organization with which the certificate subject is affiliated.
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// A subdivision or unit of the organization (such as sales or finance) with which the certificate subject is affiliated.
	OrganizationalUnit *string `field:"optional" json:"organizationalUnit" yaml:"organizationalUnit"`
	// Typically a shortened version of a longer GivenName.
	//
	// For example, Jonathan is often shortened to John. Elizabeth is often shortened to Beth, Liz, or Eliza.
	Pseudonym *string `field:"optional" json:"pseudonym" yaml:"pseudonym"`
	// The certificate serial number.
	SerialNumber *string `field:"optional" json:"serialNumber" yaml:"serialNumber"`
	// State in which the subject of the certificate is located.
	State *string `field:"optional" json:"state" yaml:"state"`
	// Family name.
	Surname *string `field:"optional" json:"surname" yaml:"surname"`
	// A personal title such as Mr.
	Title *string `field:"optional" json:"title" yaml:"title"`
}

