package previewawssagemakermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssagemaker/previewawssagemakermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an inference experiment using the configurations specified in the request.
//
// Use this API to setup and schedule an experiment to compare model variants on a Amazon SageMaker inference endpoint. For more information about inference experiments, see [Shadow tests](https://docs.aws.amazon.com/sagemaker/latest/dg/shadow-tests.html) .
//
// Amazon SageMaker begins your experiment at the scheduled time and routes traffic to your endpoint's model variants based on your specified configuration.
//
// While the experiment is in progress or after it has concluded, you can view metrics that compare your model variants. For more information, see [View, monitor, and edit shadow tests](https://docs.aws.amazon.com/sagemaker/latest/dg/shadow-tests-view-monitor-edit.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnInferenceExperimentPropsMixin := awscdkmixinspreview.Mixins.NewCfnInferenceExperimentPropsMixin(&CfnInferenceExperimentMixinProps{
//   	DataStorageConfig: &DataStorageConfigProperty{
//   		ContentType: &CaptureContentTypeHeaderProperty{
//   			CsvContentTypes: []*string{
//   				jsii.String("csvContentTypes"),
//   			},
//   			JsonContentTypes: []*string{
//   				jsii.String("jsonContentTypes"),
//   			},
//   		},
//   		Destination: jsii.String("destination"),
//   		KmsKey: jsii.String("kmsKey"),
//   	},
//   	Description: jsii.String("description"),
//   	DesiredState: jsii.String("desiredState"),
//   	EndpointName: jsii.String("endpointName"),
//   	KmsKey: jsii.String("kmsKey"),
//   	ModelVariants: []interface{}{
//   		&ModelVariantConfigProperty{
//   			InfrastructureConfig: &ModelInfrastructureConfigProperty{
//   				InfrastructureType: jsii.String("infrastructureType"),
//   				RealTimeInferenceConfig: &RealTimeInferenceConfigProperty{
//   					InstanceCount: jsii.Number(123),
//   					InstanceType: jsii.String("instanceType"),
//   				},
//   			},
//   			ModelName: jsii.String("modelName"),
//   			VariantName: jsii.String("variantName"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//   	Schedule: &InferenceExperimentScheduleProperty{
//   		EndTime: jsii.String("endTime"),
//   		StartTime: jsii.String("startTime"),
//   	},
//   	ShadowModeConfig: &ShadowModeConfigProperty{
//   		ShadowModelVariants: []interface{}{
//   			&ShadowModelVariantConfigProperty{
//   				SamplingPercentage: jsii.Number(123),
//   				ShadowModelVariantName: jsii.String("shadowModelVariantName"),
//   			},
//   		},
//   		SourceModelVariantName: jsii.String("sourceModelVariantName"),
//   	},
//   	StatusReason: jsii.String("statusReason"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-inferenceexperiment.html
//
type CfnInferenceExperimentPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnInferenceExperimentMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnInferenceExperimentPropsMixin
type jsiiProxy_CfnInferenceExperimentPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnInferenceExperimentPropsMixin) Props() *CfnInferenceExperimentMixinProps {
	var returns *CfnInferenceExperimentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInferenceExperimentPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SageMaker::InferenceExperiment`.
func NewCfnInferenceExperimentPropsMixin(props *CfnInferenceExperimentMixinProps, options *mixins.CfnPropertyMixinOptions) CfnInferenceExperimentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnInferenceExperimentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnInferenceExperimentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnInferenceExperimentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SageMaker::InferenceExperiment`.
func NewCfnInferenceExperimentPropsMixin_Override(c CfnInferenceExperimentPropsMixin, props *CfnInferenceExperimentMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnInferenceExperimentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnInferenceExperimentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInferenceExperimentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnInferenceExperimentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInferenceExperimentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnInferenceExperimentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnInferenceExperimentPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnInferenceExperimentPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

