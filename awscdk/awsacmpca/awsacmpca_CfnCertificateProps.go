package awsacmpca


// Properties for defining a `CfnCertificate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCertificateProps := &cfnCertificateProps{
//   	certificateAuthorityArn: jsii.String("certificateAuthorityArn"),
//   	certificateSigningRequest: jsii.String("certificateSigningRequest"),
//   	signingAlgorithm: jsii.String("signingAlgorithm"),
//   	validity: &validityProperty{
//   		type: jsii.String("type"),
//   		value: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	apiPassthrough: &apiPassthroughProperty{
//   		extensions: &extensionsProperty{
//   			certificatePolicies: []interface{}{
//   				&policyInformationProperty{
//   					certPolicyId: jsii.String("certPolicyId"),
//
//   					// the properties below are optional
//   					policyQualifiers: []interface{}{
//   						&policyQualifierInfoProperty{
//   							policyQualifierId: jsii.String("policyQualifierId"),
//   							qualifier: &qualifierProperty{
//   								cpsUri: jsii.String("cpsUri"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			customExtensions: []interface{}{
//   				&customExtensionProperty{
//   					objectIdentifier: jsii.String("objectIdentifier"),
//   					value: jsii.String("value"),
//
//   					// the properties below are optional
//   					critical: jsii.Boolean(false),
//   				},
//   			},
//   			extendedKeyUsage: []interface{}{
//   				&extendedKeyUsageProperty{
//   					extendedKeyUsageObjectIdentifier: jsii.String("extendedKeyUsageObjectIdentifier"),
//   					extendedKeyUsageType: jsii.String("extendedKeyUsageType"),
//   				},
//   			},
//   			keyUsage: &keyUsageProperty{
//   				crlSign: jsii.Boolean(false),
//   				dataEncipherment: jsii.Boolean(false),
//   				decipherOnly: jsii.Boolean(false),
//   				digitalSignature: jsii.Boolean(false),
//   				encipherOnly: jsii.Boolean(false),
//   				keyAgreement: jsii.Boolean(false),
//   				keyCertSign: jsii.Boolean(false),
//   				keyEncipherment: jsii.Boolean(false),
//   				nonRepudiation: jsii.Boolean(false),
//   			},
//   			subjectAlternativeNames: []interface{}{
//   				&generalNameProperty{
//   					directoryName: &subjectProperty{
//   						commonName: jsii.String("commonName"),
//   						country: jsii.String("country"),
//   						customAttributes: []interface{}{
//   							&customAttributeProperty{
//   								objectIdentifier: jsii.String("objectIdentifier"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   						distinguishedNameQualifier: jsii.String("distinguishedNameQualifier"),
//   						generationQualifier: jsii.String("generationQualifier"),
//   						givenName: jsii.String("givenName"),
//   						initials: jsii.String("initials"),
//   						locality: jsii.String("locality"),
//   						organization: jsii.String("organization"),
//   						organizationalUnit: jsii.String("organizationalUnit"),
//   						pseudonym: jsii.String("pseudonym"),
//   						serialNumber: jsii.String("serialNumber"),
//   						state: jsii.String("state"),
//   						surname: jsii.String("surname"),
//   						title: jsii.String("title"),
//   					},
//   					dnsName: jsii.String("dnsName"),
//   					ediPartyName: &ediPartyNameProperty{
//   						nameAssigner: jsii.String("nameAssigner"),
//   						partyName: jsii.String("partyName"),
//   					},
//   					ipAddress: jsii.String("ipAddress"),
//   					otherName: &otherNameProperty{
//   						typeId: jsii.String("typeId"),
//   						value: jsii.String("value"),
//   					},
//   					registeredId: jsii.String("registeredId"),
//   					rfc822Name: jsii.String("rfc822Name"),
//   					uniformResourceIdentifier: jsii.String("uniformResourceIdentifier"),
//   				},
//   			},
//   		},
//   		subject: &subjectProperty{
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
//   	},
//   	templateArn: jsii.String("templateArn"),
//   	validityNotBefore: &validityProperty{
//   		type: jsii.String("type"),
//   		value: jsii.Number(123),
//   	},
//   }
//
type CfnCertificateProps struct {
	// The Amazon Resource Name (ARN) for the private CA issues the certificate.
	CertificateAuthorityArn *string `field:"required" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// The certificate signing request (CSR) for the certificate.
	CertificateSigningRequest *string `field:"required" json:"certificateSigningRequest" yaml:"certificateSigningRequest"`
	// The name of the algorithm that will be used to sign the certificate to be issued.
	//
	// This parameter should not be confused with the `SigningAlgorithm` parameter used to sign a CSR in the `CreateCertificateAuthority` action.
	//
	// > The specified signing algorithm family (RSA or ECDSA) must match the algorithm family of the CA's secret key.
	SigningAlgorithm *string `field:"required" json:"signingAlgorithm" yaml:"signingAlgorithm"`
	// The period of time during which the certificate will be valid.
	Validity interface{} `field:"required" json:"validity" yaml:"validity"`
	// Specifies X.509 certificate information to be included in the issued certificate. An `APIPassthrough` or `APICSRPassthrough` template variant must be selected, or else this parameter is ignored.
	ApiPassthrough interface{} `field:"optional" json:"apiPassthrough" yaml:"apiPassthrough"`
	// Specifies a custom configuration template to use when issuing a certificate.
	//
	// If this parameter is not provided, ACM Private CA defaults to the `EndEntityCertificate/V1` template. For more information about ACM Private CA templates, see [Using Templates](https://docs.aws.amazon.com/acm-pca/latest/userguide/UsingTemplates.html) .
	TemplateArn *string `field:"optional" json:"templateArn" yaml:"templateArn"`
	// Information describing the start of the validity period of the certificate.
	//
	// This parameter sets the “Not Before" date for the certificate.
	//
	// By default, when issuing a certificate, ACM Private CA sets the "Not Before" date to the issuance time minus 60 minutes. This compensates for clock inconsistencies across computer systems. The `ValidityNotBefore` parameter can be used to customize the “Not Before” value.
	//
	// Unlike the `Validity` parameter, the `ValidityNotBefore` parameter is optional.
	//
	// The `ValidityNotBefore` value is expressed as an explicit date and time, using the `Validity` type value `ABSOLUTE` .
	ValidityNotBefore interface{} `field:"optional" json:"validityNotBefore" yaml:"validityNotBefore"`
}

