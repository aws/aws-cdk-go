package awsacmpca


// Provides access information used by the `authorityInfoAccess` and `subjectInfoAccess` extensions described in [RFC 5280](https://docs.aws.amazon.com/https://datatracker.ietf.org/doc/html/rfc5280) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessDescriptionProperty := &AccessDescriptionProperty{
//   	AccessLocation: &GeneralNameProperty{
//   		DirectoryName: &SubjectProperty{
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
//   		DnsName: jsii.String("dnsName"),
//   		EdiPartyName: &EdiPartyNameProperty{
//   			NameAssigner: jsii.String("nameAssigner"),
//   			PartyName: jsii.String("partyName"),
//   		},
//   		IpAddress: jsii.String("ipAddress"),
//   		OtherName: &OtherNameProperty{
//   			TypeId: jsii.String("typeId"),
//   			Value: jsii.String("value"),
//   		},
//   		RegisteredId: jsii.String("registeredId"),
//   		Rfc822Name: jsii.String("rfc822Name"),
//   		UniformResourceIdentifier: jsii.String("uniformResourceIdentifier"),
//   	},
//   	AccessMethod: &AccessMethodProperty{
//   		AccessMethodType: jsii.String("accessMethodType"),
//   		CustomObjectIdentifier: jsii.String("customObjectIdentifier"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-accessdescription.html
//
type CfnCertificateAuthority_AccessDescriptionProperty struct {
	// The location of `AccessDescription` information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-accessdescription.html#cfn-acmpca-certificateauthority-accessdescription-accesslocation
	//
	AccessLocation interface{} `field:"required" json:"accessLocation" yaml:"accessLocation"`
	// The type and format of `AccessDescription` information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-accessdescription.html#cfn-acmpca-certificateauthority-accessdescription-accessmethod
	//
	AccessMethod interface{} `field:"required" json:"accessMethod" yaml:"accessMethod"`
}

