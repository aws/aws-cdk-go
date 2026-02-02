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
// Creates an access policy that grants the specified identity (IAM Identity Center user, IAM Identity Center group, or IAM user) access to the specified AWS IoT SiteWise Monitor portal or project resource.
//
// > Support for access policies that use an SSO Group as the identity is not supported at this time.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAccessPolicyPropsMixin := awscdkmixinspreview.Mixins.NewCfnAccessPolicyPropsMixin(&CfnAccessPolicyMixinProps{
//   	AccessPolicyIdentity: &AccessPolicyIdentityProperty{
//   		IamRole: &IamRoleProperty{
//   			Arn: jsii.String("arn"),
//   		},
//   		IamUser: &IamUserProperty{
//   			Arn: jsii.String("arn"),
//   		},
//   		User: &UserProperty{
//   			Id: jsii.String("id"),
//   		},
//   	},
//   	AccessPolicyPermission: jsii.String("accessPolicyPermission"),
//   	AccessPolicyResource: &AccessPolicyResourceProperty{
//   		Portal: &PortalProperty{
//   			Id: jsii.String("id"),
//   		},
//   		Project: &ProjectProperty{
//   			Id: jsii.String("id"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-accesspolicy.html
//
type CfnAccessPolicyPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAccessPolicyMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAccessPolicyPropsMixin
type jsiiProxy_CfnAccessPolicyPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAccessPolicyPropsMixin) Props() *CfnAccessPolicyMixinProps {
	var returns *CfnAccessPolicyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPolicyPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoTSiteWise::AccessPolicy`.
func NewCfnAccessPolicyPropsMixin(props *CfnAccessPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAccessPolicyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAccessPolicyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAccessPolicyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotsitewise.mixins.CfnAccessPolicyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoTSiteWise::AccessPolicy`.
func NewCfnAccessPolicyPropsMixin_Override(c CfnAccessPolicyPropsMixin, props *CfnAccessPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotsitewise.mixins.CfnAccessPolicyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAccessPolicyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAccessPolicyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_iotsitewise.mixins.CfnAccessPolicyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAccessPolicyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_iotsitewise.mixins.CfnAccessPolicyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAccessPolicyPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAccessPolicyPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

