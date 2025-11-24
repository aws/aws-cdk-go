package mixinsawsinspectorv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsinspectorv2/mixinsawsinspectorv2/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a code security integration with a source code repository provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCodeSecurityIntegrationPropsMixin := awscdkmixinspreview.Mixins.NewCfnCodeSecurityIntegrationPropsMixin(&CfnCodeSecurityIntegrationMixinProps{
//   	CreateIntegrationDetails: &CreateDetailsProperty{
//   		GitlabSelfManaged: &CreateGitLabSelfManagedIntegrationDetailProperty{
//   			AccessToken: jsii.String("accessToken"),
//   			InstanceUrl: jsii.String("instanceUrl"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	Type: jsii.String("type"),
//   	UpdateIntegrationDetails: &UpdateDetailsProperty{
//   		Github: &UpdateGitHubIntegrationDetailProperty{
//   			Code: jsii.String("code"),
//   			InstallationId: jsii.String("installationId"),
//   		},
//   		GitlabSelfManaged: &UpdateGitLabSelfManagedIntegrationDetailProperty{
//   			AuthCode: jsii.String("authCode"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspectorv2-codesecurityintegration.html
//
type CfnCodeSecurityIntegrationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCodeSecurityIntegrationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCodeSecurityIntegrationPropsMixin
type jsiiProxy_CfnCodeSecurityIntegrationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCodeSecurityIntegrationPropsMixin) Props() *CfnCodeSecurityIntegrationMixinProps {
	var returns *CfnCodeSecurityIntegrationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeSecurityIntegrationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::InspectorV2::CodeSecurityIntegration`.
func NewCfnCodeSecurityIntegrationPropsMixin(props *CfnCodeSecurityIntegrationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCodeSecurityIntegrationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCodeSecurityIntegrationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCodeSecurityIntegrationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnCodeSecurityIntegrationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::InspectorV2::CodeSecurityIntegration`.
func NewCfnCodeSecurityIntegrationPropsMixin_Override(c CfnCodeSecurityIntegrationPropsMixin, props *CfnCodeSecurityIntegrationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnCodeSecurityIntegrationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCodeSecurityIntegrationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCodeSecurityIntegrationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnCodeSecurityIntegrationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCodeSecurityIntegrationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnCodeSecurityIntegrationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCodeSecurityIntegrationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnCodeSecurityIntegrationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

