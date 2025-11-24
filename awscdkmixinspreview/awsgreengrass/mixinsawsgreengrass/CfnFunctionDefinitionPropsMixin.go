package mixinsawsgreengrass

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsgreengrass/mixinsawsgreengrass/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Greengrass::FunctionDefinition` resource represents a function definition for AWS IoT Greengrass .
//
// Function definitions are used to organize your function definition versions.
//
// Function definitions can reference multiple function definition versions. All function definition versions must be associated with a function definition. Each function definition version can contain one or more functions.
//
// > When you create a function definition, you can optionally include an initial function definition version. To associate a function definition version later, create an [`AWS::Greengrass::FunctionDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-functiondefinitionversion.html) resource and specify the ID of this function definition.
// >
// > After you create the function definition version that contains the functions you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var tags interface{}
//   var variables interface{}
//
//   cfnFunctionDefinitionPropsMixin := awscdkmixinspreview.Mixins.NewCfnFunctionDefinitionPropsMixin(&CfnFunctionDefinitionMixinProps{
//   	InitialVersion: &FunctionDefinitionVersionProperty{
//   		DefaultConfig: &DefaultConfigProperty{
//   			Execution: &ExecutionProperty{
//   				IsolationMode: jsii.String("isolationMode"),
//   				RunAs: &RunAsProperty{
//   					Gid: jsii.Number(123),
//   					Uid: jsii.Number(123),
//   				},
//   			},
//   		},
//   		Functions: []interface{}{
//   			&FunctionProperty{
//   				FunctionArn: jsii.String("functionArn"),
//   				FunctionConfiguration: &FunctionConfigurationProperty{
//   					EncodingType: jsii.String("encodingType"),
//   					Environment: &EnvironmentProperty{
//   						AccessSysfs: jsii.Boolean(false),
//   						Execution: &ExecutionProperty{
//   							IsolationMode: jsii.String("isolationMode"),
//   							RunAs: &RunAsProperty{
//   								Gid: jsii.Number(123),
//   								Uid: jsii.Number(123),
//   							},
//   						},
//   						ResourceAccessPolicies: []interface{}{
//   							&ResourceAccessPolicyProperty{
//   								Permission: jsii.String("permission"),
//   								ResourceId: jsii.String("resourceId"),
//   							},
//   						},
//   						Variables: variables,
//   					},
//   					ExecArgs: jsii.String("execArgs"),
//   					Executable: jsii.String("executable"),
//   					MemorySize: jsii.Number(123),
//   					Pinned: jsii.Boolean(false),
//   					Timeout: jsii.Number(123),
//   				},
//   				Id: jsii.String("id"),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Tags: tags,
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-functiondefinition.html
//
type CfnFunctionDefinitionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnFunctionDefinitionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFunctionDefinitionPropsMixin
type jsiiProxy_CfnFunctionDefinitionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnFunctionDefinitionPropsMixin) Props() *CfnFunctionDefinitionMixinProps {
	var returns *CfnFunctionDefinitionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionDefinitionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Greengrass::FunctionDefinition`.
func NewCfnFunctionDefinitionPropsMixin(props *CfnFunctionDefinitionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnFunctionDefinitionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnFunctionDefinitionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFunctionDefinitionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_greengrass.mixins.CfnFunctionDefinitionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Greengrass::FunctionDefinition`.
func NewCfnFunctionDefinitionPropsMixin_Override(c CfnFunctionDefinitionPropsMixin, props *CfnFunctionDefinitionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_greengrass.mixins.CfnFunctionDefinitionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnFunctionDefinitionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFunctionDefinitionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_greengrass.mixins.CfnFunctionDefinitionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFunctionDefinitionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_greengrass.mixins.CfnFunctionDefinitionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFunctionDefinitionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnFunctionDefinitionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

