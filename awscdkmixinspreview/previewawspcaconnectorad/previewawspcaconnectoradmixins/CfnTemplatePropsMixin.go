package previewawspcaconnectoradmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawspcaconnectorad/previewawspcaconnectoradmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an Active Directory compatible certificate template.
//
// The connectors issues certificates using these templates based on the requester’s Active Directory group membership.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTemplatePropsMixin := awscdkmixinspreview.Mixins.NewCfnTemplatePropsMixin(&CfnTemplateMixinProps{
//   	ConnectorArn: jsii.String("connectorArn"),
//   	Definition: &TemplateDefinitionProperty{
//   		TemplateV2: &TemplateV2Property{
//   			CertificateValidity: &CertificateValidityProperty{
//   				RenewalPeriod: &ValidityPeriodProperty{
//   					Period: jsii.Number(123),
//   					PeriodType: jsii.String("periodType"),
//   				},
//   				ValidityPeriod: &ValidityPeriodProperty{
//   					Period: jsii.Number(123),
//   					PeriodType: jsii.String("periodType"),
//   				},
//   			},
//   			EnrollmentFlags: &EnrollmentFlagsV2Property{
//   				EnableKeyReuseOnNtTokenKeysetStorageFull: jsii.Boolean(false),
//   				IncludeSymmetricAlgorithms: jsii.Boolean(false),
//   				NoSecurityExtension: jsii.Boolean(false),
//   				RemoveInvalidCertificateFromPersonalStore: jsii.Boolean(false),
//   				UserInteractionRequired: jsii.Boolean(false),
//   			},
//   			Extensions: &ExtensionsV2Property{
//   				ApplicationPolicies: &ApplicationPoliciesProperty{
//   					Critical: jsii.Boolean(false),
//   					Policies: []interface{}{
//   						&ApplicationPolicyProperty{
//   							PolicyObjectIdentifier: jsii.String("policyObjectIdentifier"),
//   							PolicyType: jsii.String("policyType"),
//   						},
//   					},
//   				},
//   				KeyUsage: &KeyUsageProperty{
//   					Critical: jsii.Boolean(false),
//   					UsageFlags: &KeyUsageFlagsProperty{
//   						DataEncipherment: jsii.Boolean(false),
//   						DigitalSignature: jsii.Boolean(false),
//   						KeyAgreement: jsii.Boolean(false),
//   						KeyEncipherment: jsii.Boolean(false),
//   						NonRepudiation: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			GeneralFlags: &GeneralFlagsV2Property{
//   				AutoEnrollment: jsii.Boolean(false),
//   				MachineType: jsii.Boolean(false),
//   			},
//   			PrivateKeyAttributes: &PrivateKeyAttributesV2Property{
//   				CryptoProviders: []*string{
//   					jsii.String("cryptoProviders"),
//   				},
//   				KeySpec: jsii.String("keySpec"),
//   				MinimalKeyLength: jsii.Number(123),
//   			},
//   			PrivateKeyFlags: &PrivateKeyFlagsV2Property{
//   				ClientVersion: jsii.String("clientVersion"),
//   				ExportableKey: jsii.Boolean(false),
//   				StrongKeyProtectionRequired: jsii.Boolean(false),
//   			},
//   			SubjectNameFlags: &SubjectNameFlagsV2Property{
//   				RequireCommonName: jsii.Boolean(false),
//   				RequireDirectoryPath: jsii.Boolean(false),
//   				RequireDnsAsCn: jsii.Boolean(false),
//   				RequireEmail: jsii.Boolean(false),
//   				SanRequireDirectoryGuid: jsii.Boolean(false),
//   				SanRequireDns: jsii.Boolean(false),
//   				SanRequireDomainDns: jsii.Boolean(false),
//   				SanRequireEmail: jsii.Boolean(false),
//   				SanRequireSpn: jsii.Boolean(false),
//   				SanRequireUpn: jsii.Boolean(false),
//   			},
//   			SupersededTemplates: []*string{
//   				jsii.String("supersededTemplates"),
//   			},
//   		},
//   		TemplateV3: &TemplateV3Property{
//   			CertificateValidity: &CertificateValidityProperty{
//   				RenewalPeriod: &ValidityPeriodProperty{
//   					Period: jsii.Number(123),
//   					PeriodType: jsii.String("periodType"),
//   				},
//   				ValidityPeriod: &ValidityPeriodProperty{
//   					Period: jsii.Number(123),
//   					PeriodType: jsii.String("periodType"),
//   				},
//   			},
//   			EnrollmentFlags: &EnrollmentFlagsV3Property{
//   				EnableKeyReuseOnNtTokenKeysetStorageFull: jsii.Boolean(false),
//   				IncludeSymmetricAlgorithms: jsii.Boolean(false),
//   				NoSecurityExtension: jsii.Boolean(false),
//   				RemoveInvalidCertificateFromPersonalStore: jsii.Boolean(false),
//   				UserInteractionRequired: jsii.Boolean(false),
//   			},
//   			Extensions: &ExtensionsV3Property{
//   				ApplicationPolicies: &ApplicationPoliciesProperty{
//   					Critical: jsii.Boolean(false),
//   					Policies: []interface{}{
//   						&ApplicationPolicyProperty{
//   							PolicyObjectIdentifier: jsii.String("policyObjectIdentifier"),
//   							PolicyType: jsii.String("policyType"),
//   						},
//   					},
//   				},
//   				KeyUsage: &KeyUsageProperty{
//   					Critical: jsii.Boolean(false),
//   					UsageFlags: &KeyUsageFlagsProperty{
//   						DataEncipherment: jsii.Boolean(false),
//   						DigitalSignature: jsii.Boolean(false),
//   						KeyAgreement: jsii.Boolean(false),
//   						KeyEncipherment: jsii.Boolean(false),
//   						NonRepudiation: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			GeneralFlags: &GeneralFlagsV3Property{
//   				AutoEnrollment: jsii.Boolean(false),
//   				MachineType: jsii.Boolean(false),
//   			},
//   			HashAlgorithm: jsii.String("hashAlgorithm"),
//   			PrivateKeyAttributes: &PrivateKeyAttributesV3Property{
//   				Algorithm: jsii.String("algorithm"),
//   				CryptoProviders: []*string{
//   					jsii.String("cryptoProviders"),
//   				},
//   				KeySpec: jsii.String("keySpec"),
//   				KeyUsageProperty: &KeyUsagePropertyProperty{
//   					PropertyFlags: &KeyUsagePropertyFlagsProperty{
//   						Decrypt: jsii.Boolean(false),
//   						KeyAgreement: jsii.Boolean(false),
//   						Sign: jsii.Boolean(false),
//   					},
//   					PropertyType: jsii.String("propertyType"),
//   				},
//   				MinimalKeyLength: jsii.Number(123),
//   			},
//   			PrivateKeyFlags: &PrivateKeyFlagsV3Property{
//   				ClientVersion: jsii.String("clientVersion"),
//   				ExportableKey: jsii.Boolean(false),
//   				RequireAlternateSignatureAlgorithm: jsii.Boolean(false),
//   				StrongKeyProtectionRequired: jsii.Boolean(false),
//   			},
//   			SubjectNameFlags: &SubjectNameFlagsV3Property{
//   				RequireCommonName: jsii.Boolean(false),
//   				RequireDirectoryPath: jsii.Boolean(false),
//   				RequireDnsAsCn: jsii.Boolean(false),
//   				RequireEmail: jsii.Boolean(false),
//   				SanRequireDirectoryGuid: jsii.Boolean(false),
//   				SanRequireDns: jsii.Boolean(false),
//   				SanRequireDomainDns: jsii.Boolean(false),
//   				SanRequireEmail: jsii.Boolean(false),
//   				SanRequireSpn: jsii.Boolean(false),
//   				SanRequireUpn: jsii.Boolean(false),
//   			},
//   			SupersededTemplates: []*string{
//   				jsii.String("supersededTemplates"),
//   			},
//   		},
//   		TemplateV4: &TemplateV4Property{
//   			CertificateValidity: &CertificateValidityProperty{
//   				RenewalPeriod: &ValidityPeriodProperty{
//   					Period: jsii.Number(123),
//   					PeriodType: jsii.String("periodType"),
//   				},
//   				ValidityPeriod: &ValidityPeriodProperty{
//   					Period: jsii.Number(123),
//   					PeriodType: jsii.String("periodType"),
//   				},
//   			},
//   			EnrollmentFlags: &EnrollmentFlagsV4Property{
//   				EnableKeyReuseOnNtTokenKeysetStorageFull: jsii.Boolean(false),
//   				IncludeSymmetricAlgorithms: jsii.Boolean(false),
//   				NoSecurityExtension: jsii.Boolean(false),
//   				RemoveInvalidCertificateFromPersonalStore: jsii.Boolean(false),
//   				UserInteractionRequired: jsii.Boolean(false),
//   			},
//   			Extensions: &ExtensionsV4Property{
//   				ApplicationPolicies: &ApplicationPoliciesProperty{
//   					Critical: jsii.Boolean(false),
//   					Policies: []interface{}{
//   						&ApplicationPolicyProperty{
//   							PolicyObjectIdentifier: jsii.String("policyObjectIdentifier"),
//   							PolicyType: jsii.String("policyType"),
//   						},
//   					},
//   				},
//   				KeyUsage: &KeyUsageProperty{
//   					Critical: jsii.Boolean(false),
//   					UsageFlags: &KeyUsageFlagsProperty{
//   						DataEncipherment: jsii.Boolean(false),
//   						DigitalSignature: jsii.Boolean(false),
//   						KeyAgreement: jsii.Boolean(false),
//   						KeyEncipherment: jsii.Boolean(false),
//   						NonRepudiation: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			GeneralFlags: &GeneralFlagsV4Property{
//   				AutoEnrollment: jsii.Boolean(false),
//   				MachineType: jsii.Boolean(false),
//   			},
//   			HashAlgorithm: jsii.String("hashAlgorithm"),
//   			PrivateKeyAttributes: &PrivateKeyAttributesV4Property{
//   				Algorithm: jsii.String("algorithm"),
//   				CryptoProviders: []*string{
//   					jsii.String("cryptoProviders"),
//   				},
//   				KeySpec: jsii.String("keySpec"),
//   				KeyUsageProperty: &KeyUsagePropertyProperty{
//   					PropertyFlags: &KeyUsagePropertyFlagsProperty{
//   						Decrypt: jsii.Boolean(false),
//   						KeyAgreement: jsii.Boolean(false),
//   						Sign: jsii.Boolean(false),
//   					},
//   					PropertyType: jsii.String("propertyType"),
//   				},
//   				MinimalKeyLength: jsii.Number(123),
//   			},
//   			PrivateKeyFlags: &PrivateKeyFlagsV4Property{
//   				ClientVersion: jsii.String("clientVersion"),
//   				ExportableKey: jsii.Boolean(false),
//   				RequireAlternateSignatureAlgorithm: jsii.Boolean(false),
//   				RequireSameKeyRenewal: jsii.Boolean(false),
//   				StrongKeyProtectionRequired: jsii.Boolean(false),
//   				UseLegacyProvider: jsii.Boolean(false),
//   			},
//   			SubjectNameFlags: &SubjectNameFlagsV4Property{
//   				RequireCommonName: jsii.Boolean(false),
//   				RequireDirectoryPath: jsii.Boolean(false),
//   				RequireDnsAsCn: jsii.Boolean(false),
//   				RequireEmail: jsii.Boolean(false),
//   				SanRequireDirectoryGuid: jsii.Boolean(false),
//   				SanRequireDns: jsii.Boolean(false),
//   				SanRequireDomainDns: jsii.Boolean(false),
//   				SanRequireEmail: jsii.Boolean(false),
//   				SanRequireSpn: jsii.Boolean(false),
//   				SanRequireUpn: jsii.Boolean(false),
//   			},
//   			SupersededTemplates: []*string{
//   				jsii.String("supersededTemplates"),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	ReenrollAllCertificateHolders: jsii.Boolean(false),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorad-template.html
//
type CfnTemplatePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnTemplateMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTemplatePropsMixin
type jsiiProxy_CfnTemplatePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnTemplatePropsMixin) Props() *CfnTemplateMixinProps {
	var returns *CfnTemplateMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplatePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::PCAConnectorAD::Template`.
func NewCfnTemplatePropsMixin(props *CfnTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTemplatePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTemplatePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTemplatePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pcaconnectorad.mixins.CfnTemplatePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::PCAConnectorAD::Template`.
func NewCfnTemplatePropsMixin_Override(c CfnTemplatePropsMixin, props *CfnTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pcaconnectorad.mixins.CfnTemplatePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnTemplatePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTemplatePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_pcaconnectorad.mixins.CfnTemplatePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTemplatePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_pcaconnectorad.mixins.CfnTemplatePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTemplatePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnTemplatePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

