package previewawscloudformationmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscloudformation/previewawscloudformationmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::CloudFormation::LambdaHook` resource creates and activates a Lambda Hook.
//
// You can use a Lambda Hook to evaluate your resources before allowing stack operations. This resource forwards requests for resource evaluation to a Lambda function.
//
// For more information, see [Lambda Hooks](https://docs.aws.amazon.com/cloudformation-cli/latest/hooks-userguide/lambda-hooks.html) in the *CloudFormation Hooks User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLambdaHookPropsMixin := awscdkmixinspreview.Mixins.NewCfnLambdaHookPropsMixin(&CfnLambdaHookMixinProps{
//   	Alias: jsii.String("alias"),
//   	ExecutionRole: jsii.String("executionRole"),
//   	FailureMode: jsii.String("failureMode"),
//   	HookStatus: jsii.String("hookStatus"),
//   	LambdaFunction: jsii.String("lambdaFunction"),
//   	StackFilters: &StackFiltersProperty{
//   		FilteringCriteria: jsii.String("filteringCriteria"),
//   		StackNames: &StackNamesProperty{
//   			Exclude: []*string{
//   				jsii.String("exclude"),
//   			},
//   			Include: []*string{
//   				jsii.String("include"),
//   			},
//   		},
//   		StackRoles: &StackRolesProperty{
//   			Exclude: []*string{
//   				jsii.String("exclude"),
//   			},
//   			Include: []*string{
//   				jsii.String("include"),
//   			},
//   		},
//   	},
//   	TargetFilters: &TargetFiltersProperty{
//   		Actions: []*string{
//   			jsii.String("actions"),
//   		},
//   		InvocationPoints: []*string{
//   			jsii.String("invocationPoints"),
//   		},
//   		TargetNames: []*string{
//   			jsii.String("targetNames"),
//   		},
//   		Targets: []interface{}{
//   			&HookTargetProperty{
//   				Action: jsii.String("action"),
//   				InvocationPoint: jsii.String("invocationPoint"),
//   				TargetName: jsii.String("targetName"),
//   			},
//   		},
//   	},
//   	TargetOperations: []*string{
//   		jsii.String("targetOperations"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-lambdahook.html
//
type CfnLambdaHookPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnLambdaHookMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLambdaHookPropsMixin
type jsiiProxy_CfnLambdaHookPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnLambdaHookPropsMixin) Props() *CfnLambdaHookMixinProps {
	var returns *CfnLambdaHookMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLambdaHookPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CloudFormation::LambdaHook`.
func NewCfnLambdaHookPropsMixin(props *CfnLambdaHookMixinProps, options *mixins.CfnPropertyMixinOptions) CfnLambdaHookPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLambdaHookPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLambdaHookPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudformation.mixins.CfnLambdaHookPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CloudFormation::LambdaHook`.
func NewCfnLambdaHookPropsMixin_Override(c CfnLambdaHookPropsMixin, props *CfnLambdaHookMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudformation.mixins.CfnLambdaHookPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnLambdaHookPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLambdaHookPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cloudformation.mixins.CfnLambdaHookPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLambdaHookPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cloudformation.mixins.CfnLambdaHookPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLambdaHookPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnLambdaHookPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

