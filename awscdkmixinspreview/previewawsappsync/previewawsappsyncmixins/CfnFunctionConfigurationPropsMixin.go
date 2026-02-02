package previewawsappsyncmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsappsync/previewawsappsyncmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::AppSync::FunctionConfiguration` resource defines the functions in GraphQL APIs to perform certain operations.
//
// You can use pipeline resolvers to attach functions. For more information, see [Pipeline Resolvers](https://docs.aws.amazon.com/appsync/latest/devguide/pipeline-resolvers.html) in the *AWS AppSync Developer Guide* .
//
// > When you submit an update, AWS CloudFormation updates resources based on differences between what you submit and the stack's current template. To cause this resource to be updated you must change a property value for this resource in the CloudFormation template. Changing the Amazon S3 file content without changing a property value will not result in an update operation.
// >
// > See [Update Behaviors of Stack Resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html) in the *AWS CloudFormation User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFunctionConfigurationPropsMixin := awscdkmixinspreview.Mixins.NewCfnFunctionConfigurationPropsMixin(&CfnFunctionConfigurationMixinProps{
//   	ApiId: jsii.String("apiId"),
//   	Code: jsii.String("code"),
//   	CodeS3Location: jsii.String("codeS3Location"),
//   	DataSourceName: jsii.String("dataSourceName"),
//   	Description: jsii.String("description"),
//   	FunctionVersion: jsii.String("functionVersion"),
//   	MaxBatchSize: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	RequestMappingTemplate: jsii.String("requestMappingTemplate"),
//   	RequestMappingTemplateS3Location: jsii.String("requestMappingTemplateS3Location"),
//   	ResponseMappingTemplate: jsii.String("responseMappingTemplate"),
//   	ResponseMappingTemplateS3Location: jsii.String("responseMappingTemplateS3Location"),
//   	Runtime: &AppSyncRuntimeProperty{
//   		Name: jsii.String("name"),
//   		RuntimeVersion: jsii.String("runtimeVersion"),
//   	},
//   	SyncConfig: &SyncConfigProperty{
//   		ConflictDetection: jsii.String("conflictDetection"),
//   		ConflictHandler: jsii.String("conflictHandler"),
//   		LambdaConflictHandlerConfig: &LambdaConflictHandlerConfigProperty{
//   			LambdaConflictHandlerArn: jsii.String("lambdaConflictHandlerArn"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-functionconfiguration.html
//
type CfnFunctionConfigurationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnFunctionConfigurationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFunctionConfigurationPropsMixin
type jsiiProxy_CfnFunctionConfigurationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnFunctionConfigurationPropsMixin) Props() *CfnFunctionConfigurationMixinProps {
	var returns *CfnFunctionConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfigurationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AppSync::FunctionConfiguration`.
func NewCfnFunctionConfigurationPropsMixin(props *CfnFunctionConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnFunctionConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnFunctionConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFunctionConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appsync.mixins.CfnFunctionConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AppSync::FunctionConfiguration`.
func NewCfnFunctionConfigurationPropsMixin_Override(c CfnFunctionConfigurationPropsMixin, props *CfnFunctionConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appsync.mixins.CfnFunctionConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnFunctionConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFunctionConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_appsync.mixins.CfnFunctionConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFunctionConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_appsync.mixins.CfnFunctionConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFunctionConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnFunctionConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

