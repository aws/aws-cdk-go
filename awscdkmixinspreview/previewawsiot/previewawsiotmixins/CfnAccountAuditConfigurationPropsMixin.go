package previewawsiotmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsiot/previewawsiotmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use the `AWS::IoT::AccountAuditConfiguration` resource to configure or reconfigure the Device Defender audit settings for your account.
//
// Settings include how audit notifications are sent and which audit checks are enabled or disabled. For API reference, see [UpdateAccountAuditConfiguration](https://docs.aws.amazon.com/iot/latest/apireference/API_UpdateAccountAuditConfiguration.html) and for detailed information on all available audit checks, see [Audit checks](https://docs.aws.amazon.com/iot/latest/developerguide/device-defender-audit-checks.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAccountAuditConfigurationPropsMixin := awscdkmixinspreview.Mixins.NewCfnAccountAuditConfigurationPropsMixin(&CfnAccountAuditConfigurationMixinProps{
//   	AccountId: jsii.String("accountId"),
//   	AuditCheckConfigurations: &AuditCheckConfigurationsProperty{
//   		AuthenticatedCognitoRoleOverlyPermissiveCheck: &AuditCheckConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		CaCertificateExpiringCheck: &AuditCheckConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		CaCertificateKeyQualityCheck: &AuditCheckConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		ConflictingClientIdsCheck: &AuditCheckConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		DeviceCertificateAgeCheck: &DeviceCertAgeAuditCheckConfigurationProperty{
//   			Configuration: &CertAgeCheckCustomConfigurationProperty{
//   				CertAgeThresholdInDays: jsii.String("certAgeThresholdInDays"),
//   			},
//   			Enabled: jsii.Boolean(false),
//   		},
//   		DeviceCertificateExpiringCheck: &AuditCheckConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		DeviceCertificateKeyQualityCheck: &AuditCheckConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		DeviceCertificateSharedCheck: &AuditCheckConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		IntermediateCaRevokedForActiveDeviceCertificatesCheck: &AuditCheckConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		IotPolicyOverlyPermissiveCheck: &AuditCheckConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		IoTPolicyPotentialMisConfigurationCheck: &AuditCheckConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		IotRoleAliasAllowsAccessToUnusedServicesCheck: &AuditCheckConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		IotRoleAliasOverlyPermissiveCheck: &AuditCheckConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		LoggingDisabledCheck: &AuditCheckConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		RevokedCaCertificateStillActiveCheck: &AuditCheckConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		RevokedDeviceCertificateStillActiveCheck: &AuditCheckConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		UnauthenticatedCognitoRoleOverlyPermissiveCheck: &AuditCheckConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   	},
//   	AuditNotificationTargetConfigurations: &AuditNotificationTargetConfigurationsProperty{
//   		Sns: &AuditNotificationTargetProperty{
//   			Enabled: jsii.Boolean(false),
//   			RoleArn: jsii.String("roleArn"),
//   			TargetArn: jsii.String("targetArn"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-accountauditconfiguration.html
//
type CfnAccountAuditConfigurationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAccountAuditConfigurationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAccountAuditConfigurationPropsMixin
type jsiiProxy_CfnAccountAuditConfigurationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAccountAuditConfigurationPropsMixin) Props() *CfnAccountAuditConfigurationMixinProps {
	var returns *CfnAccountAuditConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccountAuditConfigurationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoT::AccountAuditConfiguration`.
func NewCfnAccountAuditConfigurationPropsMixin(props *CfnAccountAuditConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAccountAuditConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAccountAuditConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAccountAuditConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnAccountAuditConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoT::AccountAuditConfiguration`.
func NewCfnAccountAuditConfigurationPropsMixin_Override(c CfnAccountAuditConfigurationPropsMixin, props *CfnAccountAuditConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnAccountAuditConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAccountAuditConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAccountAuditConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnAccountAuditConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAccountAuditConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnAccountAuditConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAccountAuditConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAccountAuditConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

