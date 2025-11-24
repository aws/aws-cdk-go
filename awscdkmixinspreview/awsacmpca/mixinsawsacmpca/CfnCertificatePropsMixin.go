package mixinsawsacmpca

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsacmpca/mixinsawsacmpca/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::ACMPCA::Certificate` resource is used to issue a certificate using your private certificate authority.
//
// For more information, see the [IssueCertificate](https://docs.aws.amazon.com/privateca/latest/APIReference/API_IssueCertificate.html) action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCertificatePropsMixin := awscdkmixinspreview.Mixins.NewCfnCertificatePropsMixin(&CfnCertificateMixinProps{
//   	ApiPassthrough: &ApiPassthroughProperty{
//   		Extensions: &ExtensionsProperty{
//   			CertificatePolicies: []interface{}{
//   				&PolicyInformationProperty{
//   					CertPolicyId: jsii.String("certPolicyId"),
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
//   					Critical: jsii.Boolean(false),
//   					ObjectIdentifier: jsii.String("objectIdentifier"),
//   					Value: jsii.String("value"),
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
//   	CertificateAuthorityArn: jsii.String("certificateAuthorityArn"),
//   	CertificateSigningRequest: jsii.String("certificateSigningRequest"),
//   	SigningAlgorithm: jsii.String("signingAlgorithm"),
//   	TemplateArn: jsii.String("templateArn"),
//   	Validity: &ValidityProperty{
//   		Type: jsii.String("type"),
//   		Value: jsii.Number(123),
//   	},
//   	ValidityNotBefore: &ValidityProperty{
//   		Type: jsii.String("type"),
//   		Value: jsii.Number(123),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificate.html
//
type CfnCertificatePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCertificateMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCertificatePropsMixin
type jsiiProxy_CfnCertificatePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCertificatePropsMixin) Props() *CfnCertificateMixinProps {
	var returns *CfnCertificateMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificatePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ACMPCA::Certificate`.
func NewCfnCertificatePropsMixin(props *CfnCertificateMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCertificatePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCertificatePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCertificatePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_acmpca.mixins.CfnCertificatePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ACMPCA::Certificate`.
func NewCfnCertificatePropsMixin_Override(c CfnCertificatePropsMixin, props *CfnCertificateMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_acmpca.mixins.CfnCertificatePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCertificatePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCertificatePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_acmpca.mixins.CfnCertificatePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCertificatePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_acmpca.mixins.CfnCertificatePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCertificatePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCertificatePropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

