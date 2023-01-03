package awsacmpca


// Contains X.509 certificate information to be placed in an issued certificate. An `APIPassthrough` or `APICSRPassthrough` template variant must be selected, or else this parameter is ignored.
//
// If conflicting or duplicate certificate information is supplied from other sources, ACM Private CA applies [order of operation rules](https://docs.aws.amazon.com/acm-pca/latest/userguide/UsingTemplates.html#template-order-of-operations) to determine what information is used.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiPassthroughProperty := &apiPassthroughProperty{
//   	extensions: &extensionsProperty{
//   		certificatePolicies: []interface{}{
//   			&policyInformationProperty{
//   				certPolicyId: jsii.String("certPolicyId"),
//
//   				// the properties below are optional
//   				policyQualifiers: []interface{}{
//   					&policyQualifierInfoProperty{
//   						policyQualifierId: jsii.String("policyQualifierId"),
//   						qualifier: &qualifierProperty{
//   							cpsUri: jsii.String("cpsUri"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		customExtensions: []interface{}{
//   			&customExtensionProperty{
//   				objectIdentifier: jsii.String("objectIdentifier"),
//   				value: jsii.String("value"),
//
//   				// the properties below are optional
//   				critical: jsii.Boolean(false),
//   			},
//   		},
//   		extendedKeyUsage: []interface{}{
//   			&extendedKeyUsageProperty{
//   				extendedKeyUsageObjectIdentifier: jsii.String("extendedKeyUsageObjectIdentifier"),
//   				extendedKeyUsageType: jsii.String("extendedKeyUsageType"),
//   			},
//   		},
//   		keyUsage: &keyUsageProperty{
//   			crlSign: jsii.Boolean(false),
//   			dataEncipherment: jsii.Boolean(false),
//   			decipherOnly: jsii.Boolean(false),
//   			digitalSignature: jsii.Boolean(false),
//   			encipherOnly: jsii.Boolean(false),
//   			keyAgreement: jsii.Boolean(false),
//   			keyCertSign: jsii.Boolean(false),
//   			keyEncipherment: jsii.Boolean(false),
//   			nonRepudiation: jsii.Boolean(false),
//   		},
//   		subjectAlternativeNames: []interface{}{
//   			&generalNameProperty{
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
//   		},
//   	},
//   	subject: &subjectProperty{
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
//   }
//
type CfnCertificate_ApiPassthroughProperty struct {
	// Specifies X.509 extension information for a certificate.
	Extensions interface{} `field:"optional" json:"extensions" yaml:"extensions"`
	// Contains information about the certificate subject.
	//
	// The Subject field in the certificate identifies the entity that owns or controls the public key in the certificate. The entity can be a user, computer, device, or service. The Subject must contain an X.500 distinguished name (DN). A DN is a sequence of relative distinguished names (RDNs). The RDNs are separated by commas in the certificate.
	Subject interface{} `field:"optional" json:"subject" yaml:"subject"`
}

