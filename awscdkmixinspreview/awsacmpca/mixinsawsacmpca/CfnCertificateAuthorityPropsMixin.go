package mixinsawsacmpca

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsacmpca/mixinsawsacmpca/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use the `AWS::ACMPCA::CertificateAuthority` resource to create a private CA.
//
// Once the CA exists, you can use the `AWS::ACMPCA::Certificate` resource to issue a new CA certificate. Alternatively, you can issue a CA certificate using an on-premises CA, and then use the `AWS::ACMPCA::CertificateAuthorityActivation` resource to import the new CA certificate and activate the CA.
//
// > Before removing a `AWS::ACMPCA::CertificateAuthority` resource from the CloudFormation stack, disable the affected CA. Otherwise, the action will fail. You can disable the CA by removing its associated `AWS::ACMPCA::CertificateAuthorityActivation` resource from CloudFormation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCertificateAuthorityPropsMixin := awscdkmixinspreview.Mixins.NewCfnCertificateAuthorityPropsMixin(&CfnCertificateAuthorityMixinProps{
//   	CsrExtensions: &CsrExtensionsProperty{
//   		KeyUsage: &KeyUsageProperty{
//   			CrlSign: jsii.Boolean(false),
//   			DataEncipherment: jsii.Boolean(false),
//   			DecipherOnly: jsii.Boolean(false),
//   			DigitalSignature: jsii.Boolean(false),
//   			EncipherOnly: jsii.Boolean(false),
//   			KeyAgreement: jsii.Boolean(false),
//   			KeyCertSign: jsii.Boolean(false),
//   			KeyEncipherment: jsii.Boolean(false),
//   			NonRepudiation: jsii.Boolean(false),
//   		},
//   		SubjectInformationAccess: []interface{}{
//   			&AccessDescriptionProperty{
//   				AccessLocation: &GeneralNameProperty{
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
//   				AccessMethod: &AccessMethodProperty{
//   					AccessMethodType: jsii.String("accessMethodType"),
//   					CustomObjectIdentifier: jsii.String("customObjectIdentifier"),
//   				},
//   			},
//   		},
//   	},
//   	KeyAlgorithm: jsii.String("keyAlgorithm"),
//   	KeyStorageSecurityStandard: jsii.String("keyStorageSecurityStandard"),
//   	RevocationConfiguration: &RevocationConfigurationProperty{
//   		CrlConfiguration: &CrlConfigurationProperty{
//   			CrlDistributionPointExtensionConfiguration: &CrlDistributionPointExtensionConfigurationProperty{
//   				OmitExtension: jsii.Boolean(false),
//   			},
//   			CrlType: jsii.String("crlType"),
//   			CustomCname: jsii.String("customCname"),
//   			CustomPath: jsii.String("customPath"),
//   			Enabled: jsii.Boolean(false),
//   			ExpirationInDays: jsii.Number(123),
//   			S3BucketName: jsii.String("s3BucketName"),
//   			S3ObjectAcl: jsii.String("s3ObjectAcl"),
//   		},
//   		OcspConfiguration: &OcspConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			OcspCustomCname: jsii.String("ocspCustomCname"),
//   		},
//   	},
//   	SigningAlgorithm: jsii.String("signingAlgorithm"),
//   	Subject: &SubjectProperty{
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
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   	UsageMode: jsii.String("usageMode"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html
//
type CfnCertificateAuthorityPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCertificateAuthorityMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCertificateAuthorityPropsMixin
type jsiiProxy_CfnCertificateAuthorityPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCertificateAuthorityPropsMixin) Props() *CfnCertificateAuthorityMixinProps {
	var returns *CfnCertificateAuthorityMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificateAuthorityPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ACMPCA::CertificateAuthority`.
func NewCfnCertificateAuthorityPropsMixin(props *CfnCertificateAuthorityMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCertificateAuthorityPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCertificateAuthorityPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCertificateAuthorityPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_acmpca.mixins.CfnCertificateAuthorityPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ACMPCA::CertificateAuthority`.
func NewCfnCertificateAuthorityPropsMixin_Override(c CfnCertificateAuthorityPropsMixin, props *CfnCertificateAuthorityMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_acmpca.mixins.CfnCertificateAuthorityPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCertificateAuthorityPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCertificateAuthorityPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_acmpca.mixins.CfnCertificateAuthorityPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCertificateAuthorityPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_acmpca.mixins.CfnCertificateAuthorityPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCertificateAuthorityPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnCertificateAuthorityPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

