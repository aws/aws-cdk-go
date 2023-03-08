package awsacmpca


// Provides access information used by the `authorityInfoAccess` and `subjectInfoAccess` extensions described in [RFC 5280](https://docs.aws.amazon.com/https://datatracker.ietf.org/doc/html/rfc5280) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessDescriptionProperty := &accessDescriptionProperty{
//   	accessLocation: &generalNameProperty{
//   		directoryName: &subjectProperty{
//   			commonName: jsii.String("commonName"),
//   			country: jsii.String("country"),
//   			customAttributes: []interface{}{
//   				&customAttributeProperty{
//   					objectIdentifier: jsii.String("objectIdentifier"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			distinguishedNameQualifier: jsii.String("distinguishedNameQualifier"),
//   			generationQualifier: jsii.String("generationQualifier"),
//   			givenName: jsii.String("givenName"),
//   			initials: jsii.String("initials"),
//   			locality: jsii.String("locality"),
//   			organization: jsii.String("organization"),
//   			organizationalUnit: jsii.String("organizationalUnit"),
//   			pseudonym: jsii.String("pseudonym"),
//   			serialNumber: jsii.String("serialNumber"),
//   			state: jsii.String("state"),
//   			surname: jsii.String("surname"),
//   			title: jsii.String("title"),
//   		},
//   		dnsName: jsii.String("dnsName"),
//   		ediPartyName: &ediPartyNameProperty{
//   			nameAssigner: jsii.String("nameAssigner"),
//   			partyName: jsii.String("partyName"),
//   		},
//   		ipAddress: jsii.String("ipAddress"),
//   		otherName: &otherNameProperty{
//   			typeId: jsii.String("typeId"),
//   			value: jsii.String("value"),
//   		},
//   		registeredId: jsii.String("registeredId"),
//   		rfc822Name: jsii.String("rfc822Name"),
//   		uniformResourceIdentifier: jsii.String("uniformResourceIdentifier"),
//   	},
//   	accessMethod: &accessMethodProperty{
//   		accessMethodType: jsii.String("accessMethodType"),
//   		customObjectIdentifier: jsii.String("customObjectIdentifier"),
//   	},
//   }
//
type CfnCertificateAuthority_AccessDescriptionProperty struct {
	// The location of `AccessDescription` information.
	AccessLocation interface{} `field:"required" json:"accessLocation" yaml:"accessLocation"`
	// The type and format of `AccessDescription` information.
	AccessMethod interface{} `field:"required" json:"accessMethod" yaml:"accessMethod"`
}

