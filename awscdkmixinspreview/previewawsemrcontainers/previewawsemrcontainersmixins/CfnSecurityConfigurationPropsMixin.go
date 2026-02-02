package previewawsemrcontainersmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsemrcontainers/previewawsemrcontainersmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Schema of AWS::EMRContainers::SecurityConfiguration Type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSecurityConfigurationPropsMixin := awscdkmixinspreview.Mixins.NewCfnSecurityConfigurationPropsMixin(&CfnSecurityConfigurationMixinProps{
//   	ContainerProvider: &ContainerProviderProperty{
//   		Id: jsii.String("id"),
//   		Info: &ContainerInfoProperty{
//   			EksInfo: &EksInfoProperty{
//   				Namespace: jsii.String("namespace"),
//   			},
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	Name: jsii.String("name"),
//   	SecurityConfigurationData: &SecurityConfigurationDataProperty{
//   		AuthenticationConfiguration: &AuthenticationConfigurationProperty{
//   			IamConfiguration: map[string]*string{
//   				"systemRole": jsii.String("systemRole"),
//   			},
//   			IdentityCenterConfiguration: &IdentityCenterConfigurationProperty{
//   				EnableIdentityCenter: jsii.Boolean(false),
//   				IdentityCenterApplicationAssignmentRequired: jsii.Boolean(false),
//   				IdentityCenterInstanceArn: jsii.String("identityCenterInstanceArn"),
//   			},
//   		},
//   		AuthorizationConfiguration: &AuthorizationConfigurationProperty{
//   			LakeFormationConfiguration: &LakeFormationConfigurationProperty{
//   				AuthorizedSessionTagValue: jsii.String("authorizedSessionTagValue"),
//   				QueryAccessControlEnabled: jsii.Boolean(false),
//   				QueryEngineRoleArn: jsii.String("queryEngineRoleArn"),
//   				SecureNamespaceInfo: &SecureNamespaceInfoProperty{
//   					ClusterId: jsii.String("clusterId"),
//   					Namespace: jsii.String("namespace"),
//   				},
//   			},
//   		},
//   		EncryptionConfiguration: &EncryptionConfigurationProperty{
//   			AtRestEncryptionConfiguration: &AtRestEncryptionConfigurationProperty{
//   				LocalDiskEncryptionConfiguration: &LocalDiskEncryptionConfigurationProperty{
//   					AwsKmsKeyId: jsii.String("awsKmsKeyId"),
//   					EncryptionKeyProviderType: jsii.String("encryptionKeyProviderType"),
//   				},
//   				S3EncryptionConfiguration: &S3EncryptionConfigurationProperty{
//   					EncryptionOption: jsii.String("encryptionOption"),
//   					KmsKeyId: jsii.String("kmsKeyId"),
//   				},
//   			},
//   			InTransitEncryptionConfiguration: &InTransitEncryptionConfigurationProperty{
//   				TlsCertificateConfiguration: &TLSCertificateConfigurationProperty{
//   					CertificateProviderType: jsii.String("certificateProviderType"),
//   					PrivateKeySecretArn: jsii.String("privateKeySecretArn"),
//   					PublicKeySecretArn: jsii.String("publicKeySecretArn"),
//   				},
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrcontainers-securityconfiguration.html
//
type CfnSecurityConfigurationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSecurityConfigurationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSecurityConfigurationPropsMixin
type jsiiProxy_CfnSecurityConfigurationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSecurityConfigurationPropsMixin) Props() *CfnSecurityConfigurationMixinProps {
	var returns *CfnSecurityConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityConfigurationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EMRContainers::SecurityConfiguration`.
func NewCfnSecurityConfigurationPropsMixin(props *CfnSecurityConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSecurityConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSecurityConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSecurityConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnSecurityConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EMRContainers::SecurityConfiguration`.
func NewCfnSecurityConfigurationPropsMixin_Override(c CfnSecurityConfigurationPropsMixin, props *CfnSecurityConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnSecurityConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSecurityConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSecurityConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnSecurityConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSecurityConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnSecurityConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSecurityConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnSecurityConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

