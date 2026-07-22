package awssagemaker

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::SageMaker::Algorithm.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnAlgorithmPropsMixin := awscdkcfnpropertymixins.Aws_sagemaker.NewCfnAlgorithmPropsMixin(&CfnAlgorithmMixinProps{
//   	AlgorithmDescription: jsii.String("algorithmDescription"),
//   	AlgorithmName: jsii.String("algorithmName"),
//   	CertifyForMarketplace: jsii.Boolean(false),
//   	InferenceSpecification: &InferenceSpecificationProperty{
//   		Containers: []interface{}{
//   			&ModelPackageContainerDefinitionProperty{
//   				ContainerHostname: jsii.String("containerHostname"),
//   				Environment: map[string]*string{
//   					"environmentKey": jsii.String("environment"),
//   				},
//   				Framework: jsii.String("framework"),
//   				FrameworkVersion: jsii.String("frameworkVersion"),
//   				Image: jsii.String("image"),
//   				ImageDigest: jsii.String("imageDigest"),
//   				IsCheckpoint: jsii.Boolean(false),
//   				ModelInput: &ModelInputProperty{
//   					DataInputConfig: jsii.String("dataInputConfig"),
//   				},
//   				NearestModelName: jsii.String("nearestModelName"),
//   			},
//   		},
//   		SupportedContentTypes: []*string{
//   			jsii.String("supportedContentTypes"),
//   		},
//   		SupportedRealtimeInferenceInstanceTypes: []*string{
//   			jsii.String("supportedRealtimeInferenceInstanceTypes"),
//   		},
//   		SupportedResponseMimeTypes: []*string{
//   			jsii.String("supportedResponseMimeTypes"),
//   		},
//   		SupportedTransformInstanceTypes: []*string{
//   			jsii.String("supportedTransformInstanceTypes"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrainingSpecification: &TrainingSpecificationProperty{
//   		MetricDefinitions: []interface{}{
//   			&MetricDefinitionProperty{
//   				Name: jsii.String("name"),
//   				Regex: jsii.String("regex"),
//   			},
//   		},
//   		SupportedHyperParameters: []interface{}{
//   			&HyperParameterSpecificationProperty{
//   				DefaultValue: jsii.String("defaultValue"),
//   				Description: jsii.String("description"),
//   				IsRequired: jsii.Boolean(false),
//   				IsTunable: jsii.Boolean(false),
//   				Name: jsii.String("name"),
//   				Range: &ParameterRangeProperty{
//   					CategoricalParameterRangeSpecification: &CategoricalParameterRangeSpecificationProperty{
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					ContinuousParameterRangeSpecification: &ContinuousParameterRangeSpecificationProperty{
//   						MaxValue: jsii.String("maxValue"),
//   						MinValue: jsii.String("minValue"),
//   					},
//   					IntegerParameterRangeSpecification: &IntegerParameterRangeSpecificationProperty{
//   						MaxValue: jsii.String("maxValue"),
//   						MinValue: jsii.String("minValue"),
//   					},
//   				},
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		SupportedTrainingInstanceTypes: []*string{
//   			jsii.String("supportedTrainingInstanceTypes"),
//   		},
//   		SupportedTuningJobObjectiveMetrics: []interface{}{
//   			&HyperParameterTuningJobObjectiveProperty{
//   				MetricName: jsii.String("metricName"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		SupportsDistributedTraining: jsii.Boolean(false),
//   		TrainingChannels: []interface{}{
//   			&ChannelSpecificationProperty{
//   				Description: jsii.String("description"),
//   				IsRequired: jsii.Boolean(false),
//   				Name: jsii.String("name"),
//   				SupportedCompressionTypes: []*string{
//   					jsii.String("supportedCompressionTypes"),
//   				},
//   				SupportedContentTypes: []*string{
//   					jsii.String("supportedContentTypes"),
//   				},
//   				SupportedInputModes: []*string{
//   					jsii.String("supportedInputModes"),
//   				},
//   			},
//   		},
//   		TrainingImage: jsii.String("trainingImage"),
//   		TrainingImageDigest: jsii.String("trainingImageDigest"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-algorithm.html
//
type CfnAlgorithmPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnAlgorithmMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAlgorithmPropsMixin
type jsiiProxy_CfnAlgorithmPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnAlgorithmPropsMixin) Props() *CfnAlgorithmMixinProps {
	var returns *CfnAlgorithmMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlgorithmPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SageMaker::Algorithm`.
func NewCfnAlgorithmPropsMixin(props *CfnAlgorithmMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnAlgorithmPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAlgorithmPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAlgorithmPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnAlgorithmPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SageMaker::Algorithm`.
func NewCfnAlgorithmPropsMixin_Override(c CfnAlgorithmPropsMixin, props *CfnAlgorithmMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnAlgorithmPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnAlgorithmPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAlgorithmPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnAlgorithmPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAlgorithmPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnAlgorithmPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAlgorithmPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAlgorithmPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

