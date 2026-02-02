package previewawsiotsitewisemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsiotsitewise/previewawsiotsitewisemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// > The AWS IoT SiteWise Monitor feature will no longer be open to new customers starting November 7, 2025 .
//
// If you would like to use the AWS IoT SiteWise Monitor feature, sign up prior to that date. Existing customers can continue to use the service as normal. For more information, see [AWS IoT SiteWise Monitor availability change](https://docs.aws.amazon.com/iot-sitewise/latest/appguide/iotsitewise-monitor-availability-change.html) .
//
// Creates a portal, which can contain projects and dashboards. AWS IoT SiteWise Monitor uses IAM Identity Center or IAM to authenticate portal users and manage user permissions.
//
// > Before you can sign in to a new portal, you must add at least one identity to that portal. For more information, see [Adding or removing portal administrators](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/administer-portals.html#portal-change-admins) in the *AWS IoT SiteWise User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var alarms interface{}
//
//   cfnPortalPropsMixin := awscdkmixinspreview.Mixins.NewCfnPortalPropsMixin(&CfnPortalMixinProps{
//   	Alarms: alarms,
//   	NotificationSenderEmail: jsii.String("notificationSenderEmail"),
//   	PortalAuthMode: jsii.String("portalAuthMode"),
//   	PortalContactEmail: jsii.String("portalContactEmail"),
//   	PortalDescription: jsii.String("portalDescription"),
//   	PortalName: jsii.String("portalName"),
//   	PortalType: jsii.String("portalType"),
//   	PortalTypeConfiguration: map[string]interface{}{
//   		"portalTypeConfigurationKey": &PortalTypeEntryProperty{
//   			"portalTools": []*string{
//   				jsii.String("portalTools"),
//   			},
//   		},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-portal.html
//
type CfnPortalPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnPortalMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPortalPropsMixin
type jsiiProxy_CfnPortalPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPortalPropsMixin) Props() *CfnPortalMixinProps {
	var returns *CfnPortalMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPortalPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoTSiteWise::Portal`.
func NewCfnPortalPropsMixin(props *CfnPortalMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPortalPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPortalPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPortalPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotsitewise.mixins.CfnPortalPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoTSiteWise::Portal`.
func NewCfnPortalPropsMixin_Override(c CfnPortalPropsMixin, props *CfnPortalMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotsitewise.mixins.CfnPortalPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPortalPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPortalPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_iotsitewise.mixins.CfnPortalPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPortalPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_iotsitewise.mixins.CfnPortalPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPortalPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnPortalPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

