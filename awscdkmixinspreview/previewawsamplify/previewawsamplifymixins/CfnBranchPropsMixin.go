package previewawsamplifymixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsamplify/previewawsamplifymixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The AWS::Amplify::Branch resource specifies a new branch within an app.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBranchPropsMixin := awscdkmixinspreview.Mixins.NewCfnBranchPropsMixin(&CfnBranchMixinProps{
//   	AppId: jsii.String("appId"),
//   	Backend: &BackendProperty{
//   		StackArn: jsii.String("stackArn"),
//   	},
//   	BasicAuthConfig: &BasicAuthConfigProperty{
//   		EnableBasicAuth: jsii.Boolean(false),
//   		Password: jsii.String("password"),
//   		Username: jsii.String("username"),
//   	},
//   	BranchName: jsii.String("branchName"),
//   	BuildSpec: jsii.String("buildSpec"),
//   	ComputeRoleArn: jsii.String("computeRoleArn"),
//   	Description: jsii.String("description"),
//   	EnableAutoBuild: jsii.Boolean(false),
//   	EnablePerformanceMode: jsii.Boolean(false),
//   	EnablePullRequestPreview: jsii.Boolean(false),
//   	EnableSkewProtection: jsii.Boolean(false),
//   	EnvironmentVariables: []interface{}{
//   		&EnvironmentVariableProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Framework: jsii.String("framework"),
//   	PullRequestEnvironmentName: jsii.String("pullRequestEnvironmentName"),
//   	Stage: jsii.String("stage"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplify-branch.html
//
type CfnBranchPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnBranchMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnBranchPropsMixin
type jsiiProxy_CfnBranchPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnBranchPropsMixin) Props() *CfnBranchMixinProps {
	var returns *CfnBranchMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBranchPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Amplify::Branch`.
func NewCfnBranchPropsMixin(props *CfnBranchMixinProps, options *mixins.CfnPropertyMixinOptions) CfnBranchPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnBranchPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnBranchPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_amplify.mixins.CfnBranchPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Amplify::Branch`.
func NewCfnBranchPropsMixin_Override(c CfnBranchPropsMixin, props *CfnBranchMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_amplify.mixins.CfnBranchPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnBranchPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnBranchPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_amplify.mixins.CfnBranchPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBranchPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_amplify.mixins.CfnBranchPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnBranchPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnBranchPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

