package previewawscloudformationmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscloudformation/previewawscloudformationmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::CloudFormation::GuardHook` resource creates and activates a Guard Hook.
//
// Using the Guard domain specific language (DSL), you can author Guard Hooks to evaluate your resources before allowing stack operations.
//
// For more information, see [Guard Hooks](https://docs.aws.amazon.com/cloudformation-cli/latest/hooks-userguide/guard-hooks.html) in the *CloudFormation Hooks User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGuardHookPropsMixin := awscdkmixinspreview.Mixins.NewCfnGuardHookPropsMixin(&CfnGuardHookMixinProps{
//   	Alias: jsii.String("alias"),
//   	ExecutionRole: jsii.String("executionRole"),
//   	FailureMode: jsii.String("failureMode"),
//   	HookStatus: jsii.String("hookStatus"),
//   	LogBucket: jsii.String("logBucket"),
//   	Options: &OptionsProperty{
//   		InputParams: &S3LocationProperty{
//   			Uri: jsii.String("uri"),
//   			VersionId: jsii.String("versionId"),
//   		},
//   	},
//   	RuleLocation: &S3LocationProperty{
//   		Uri: jsii.String("uri"),
//   		VersionId: jsii.String("versionId"),
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-guardhook.html
//
type CfnGuardHookPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnGuardHookMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnGuardHookPropsMixin
type jsiiProxy_CfnGuardHookPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnGuardHookPropsMixin) Props() *CfnGuardHookMixinProps {
	var returns *CfnGuardHookMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGuardHookPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CloudFormation::GuardHook`.
func NewCfnGuardHookPropsMixin(props *CfnGuardHookMixinProps, options *mixins.CfnPropertyMixinOptions) CfnGuardHookPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnGuardHookPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnGuardHookPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudformation.mixins.CfnGuardHookPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CloudFormation::GuardHook`.
func NewCfnGuardHookPropsMixin_Override(c CfnGuardHookPropsMixin, props *CfnGuardHookMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudformation.mixins.CfnGuardHookPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnGuardHookPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnGuardHookPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cloudformation.mixins.CfnGuardHookPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGuardHookPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cloudformation.mixins.CfnGuardHookPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGuardHookPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnGuardHookPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

