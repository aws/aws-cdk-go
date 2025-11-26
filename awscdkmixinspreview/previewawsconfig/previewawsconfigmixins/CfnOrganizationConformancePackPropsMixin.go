package previewawsconfigmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsconfig/previewawsconfigmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// OrganizationConformancePack deploys conformance packs across member accounts in an AWS Organizations .
//
// OrganizationConformancePack enables organization service access for `config-multiaccountsetup.amazonaws.com` through the `EnableAWSServiceAccess` action and creates a service linked role in the master account of your organization. The service linked role is created only when the role does not exist in the master account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnOrganizationConformancePackPropsMixin := awscdkmixinspreview.Mixins.NewCfnOrganizationConformancePackPropsMixin(&CfnOrganizationConformancePackMixinProps{
//   	ConformancePackInputParameters: []interface{}{
//   		&ConformancePackInputParameterProperty{
//   			ParameterName: jsii.String("parameterName"),
//   			ParameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   	DeliveryS3Bucket: jsii.String("deliveryS3Bucket"),
//   	DeliveryS3KeyPrefix: jsii.String("deliveryS3KeyPrefix"),
//   	ExcludedAccounts: []*string{
//   		jsii.String("excludedAccounts"),
//   	},
//   	OrganizationConformancePackName: jsii.String("organizationConformancePackName"),
//   	TemplateBody: jsii.String("templateBody"),
//   	TemplateS3Uri: jsii.String("templateS3Uri"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconformancepack.html
//
type CfnOrganizationConformancePackPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnOrganizationConformancePackMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnOrganizationConformancePackPropsMixin
type jsiiProxy_CfnOrganizationConformancePackPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnOrganizationConformancePackPropsMixin) Props() *CfnOrganizationConformancePackMixinProps {
	var returns *CfnOrganizationConformancePackMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationConformancePackPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Config::OrganizationConformancePack`.
func NewCfnOrganizationConformancePackPropsMixin(props *CfnOrganizationConformancePackMixinProps, options *mixins.CfnPropertyMixinOptions) CfnOrganizationConformancePackPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnOrganizationConformancePackPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnOrganizationConformancePackPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_config.mixins.CfnOrganizationConformancePackPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Config::OrganizationConformancePack`.
func NewCfnOrganizationConformancePackPropsMixin_Override(c CfnOrganizationConformancePackPropsMixin, props *CfnOrganizationConformancePackMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_config.mixins.CfnOrganizationConformancePackPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnOrganizationConformancePackPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOrganizationConformancePackPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_config.mixins.CfnOrganizationConformancePackPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnOrganizationConformancePackPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_config.mixins.CfnOrganizationConformancePackPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnOrganizationConformancePackPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnOrganizationConformancePackPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

