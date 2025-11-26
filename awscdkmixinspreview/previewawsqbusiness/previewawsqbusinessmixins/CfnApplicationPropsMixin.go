package previewawsqbusinessmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsqbusiness/previewawsqbusinessmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an Amazon Q Business application.
//
// > There are new tiers for Amazon Q Business. Not all features in Amazon Q Business Pro are also available in Amazon Q Business Lite. For information on what's included in Amazon Q Business Lite and what's included in Amazon Q Business Pro, see [Amazon Q Business tiers](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/tiers.html#user-sub-tiers) . You must use the Amazon Q Business console to assign subscription tiers to users.
// >
// > An Amazon Q Apps service linked role will be created if it's absent in the AWS account when `QAppsConfiguration` is enabled in the request. For more information, see [Using service-linked roles for Q Apps](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/using-service-linked-roles-qapps.html) .
// >
// > When you create an application, Amazon Q Business may securely transmit data for processing from your selected AWS region, but within your geography. For more information, see [Cross region inference in Amazon Q Business](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/cross-region-inference.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationPropsMixin := awscdkmixinspreview.Mixins.NewCfnApplicationPropsMixin(&CfnApplicationMixinProps{
//   	AttachmentsConfiguration: &AttachmentsConfigurationProperty{
//   		AttachmentsControlMode: jsii.String("attachmentsControlMode"),
//   	},
//   	AutoSubscriptionConfiguration: &AutoSubscriptionConfigurationProperty{
//   		AutoSubscribe: jsii.String("autoSubscribe"),
//   		DefaultSubscriptionType: jsii.String("defaultSubscriptionType"),
//   	},
//   	ClientIdsForOidc: []*string{
//   		jsii.String("clientIdsForOidc"),
//   	},
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	IamIdentityProviderArn: jsii.String("iamIdentityProviderArn"),
//   	IdentityCenterInstanceArn: jsii.String("identityCenterInstanceArn"),
//   	IdentityType: jsii.String("identityType"),
//   	PersonalizationConfiguration: &PersonalizationConfigurationProperty{
//   		PersonalizationControlMode: jsii.String("personalizationControlMode"),
//   	},
//   	QAppsConfiguration: &QAppsConfigurationProperty{
//   		QAppsControlMode: jsii.String("qAppsControlMode"),
//   	},
//   	QuickSightConfiguration: &QuickSightConfigurationProperty{
//   		ClientNamespace: jsii.String("clientNamespace"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-application.html
//
type CfnApplicationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnApplicationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnApplicationPropsMixin
type jsiiProxy_CfnApplicationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnApplicationPropsMixin) Props() *CfnApplicationMixinProps {
	var returns *CfnApplicationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::QBusiness::Application`.
func NewCfnApplicationPropsMixin(props *CfnApplicationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnApplicationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnApplicationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnApplicationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::QBusiness::Application`.
func NewCfnApplicationPropsMixin_Override(c CfnApplicationPropsMixin, props *CfnApplicationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnApplicationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApplicationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApplicationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApplicationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnApplicationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

