package previewawsdatapipelinemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsdatapipeline/previewawsdatapipelinemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The AWS::DataPipeline::Pipeline resource specifies a data pipeline that you can use to automate the movement and transformation of data.
//
// > AWS Data Pipeline is no longer available to new customers. Existing customers of AWS Data Pipeline can continue to use the service as normal. [Learn more](https://docs.aws.amazon.com/big-data/migrate-workloads-from-aws-data-pipeline/)
//
// In each pipeline, you define pipeline objects, such as activities, schedules, data nodes, and resources.
//
// The `AWS::DataPipeline::Pipeline` resource adds tasks, schedules, and preconditions to the specified pipeline. You can use `PutPipelineDefinition` to populate a new pipeline.
//
// `PutPipelineDefinition` also validates the configuration as it adds it to the pipeline. Changes to the pipeline are saved unless one of the following validation errors exist in the pipeline.
//
// - An object is missing a name or identifier field.
// - A string or reference field is empty.
// - The number of objects in the pipeline exceeds the allowed maximum number of objects.
// - The pipeline is in a FINISHED state.
//
// Pipeline object definitions are passed to the [PutPipelineDefinition](https://docs.aws.amazon.com/datapipeline/latest/APIReference/API_PutPipelineDefinition.html) action and returned by the [GetPipelineDefinition](https://docs.aws.amazon.com/datapipeline/latest/APIReference/API_GetPipelineDefinition.html) action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPipelinePropsMixin := awscdkmixinspreview.Mixins.NewCfnPipelinePropsMixin(&CfnPipelineMixinProps{
//   	Activate: jsii.Boolean(false),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	ParameterObjects: []interface{}{
//   		&ParameterObjectProperty{
//   			Attributes: []interface{}{
//   				&ParameterAttributeProperty{
//   					Key: jsii.String("key"),
//   					StringValue: jsii.String("stringValue"),
//   				},
//   			},
//   			Id: jsii.String("id"),
//   		},
//   	},
//   	ParameterValues: []interface{}{
//   		&ParameterValueProperty{
//   			Id: jsii.String("id"),
//   			StringValue: jsii.String("stringValue"),
//   		},
//   	},
//   	PipelineObjects: []interface{}{
//   		&PipelineObjectProperty{
//   			Fields: []interface{}{
//   				&FieldProperty{
//   					Key: jsii.String("key"),
//   					RefValue: jsii.String("refValue"),
//   					StringValue: jsii.String("stringValue"),
//   				},
//   			},
//   			Id: jsii.String("id"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	PipelineTags: []PipelineTagProperty{
//   		&PipelineTagProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html
//
type CfnPipelinePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnPipelineMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPipelinePropsMixin
type jsiiProxy_CfnPipelinePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPipelinePropsMixin) Props() *CfnPipelineMixinProps {
	var returns *CfnPipelineMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipelinePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DataPipeline::Pipeline`.
func NewCfnPipelinePropsMixin(props *CfnPipelineMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPipelinePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPipelinePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPipelinePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_datapipeline.mixins.CfnPipelinePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DataPipeline::Pipeline`.
func NewCfnPipelinePropsMixin_Override(c CfnPipelinePropsMixin, props *CfnPipelineMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_datapipeline.mixins.CfnPipelinePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPipelinePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPipelinePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_datapipeline.mixins.CfnPipelinePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPipelinePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_datapipeline.mixins.CfnPipelinePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPipelinePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnPipelinePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

