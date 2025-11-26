package previewawssagemakermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssagemaker/previewawssagemakermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a definition for a job that monitors model quality and drift.
//
// For information about model monitor, see [Amazon SageMaker Model Monitor](https://docs.aws.amazon.com/sagemaker/latest/dg/model-monitor.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnModelQualityJobDefinitionPropsMixin := awscdkmixinspreview.Mixins.NewCfnModelQualityJobDefinitionPropsMixin(&CfnModelQualityJobDefinitionMixinProps{
//   	EndpointName: jsii.String("endpointName"),
//   	JobDefinitionName: jsii.String("jobDefinitionName"),
//   	JobResources: &MonitoringResourcesProperty{
//   		ClusterConfig: &ClusterConfigProperty{
//   			InstanceCount: jsii.Number(123),
//   			InstanceType: jsii.String("instanceType"),
//   			VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   			VolumeSizeInGb: jsii.Number(123),
//   		},
//   	},
//   	ModelQualityAppSpecification: &ModelQualityAppSpecificationProperty{
//   		ContainerArguments: []*string{
//   			jsii.String("containerArguments"),
//   		},
//   		ContainerEntrypoint: []*string{
//   			jsii.String("containerEntrypoint"),
//   		},
//   		Environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		ImageUri: jsii.String("imageUri"),
//   		PostAnalyticsProcessorSourceUri: jsii.String("postAnalyticsProcessorSourceUri"),
//   		ProblemType: jsii.String("problemType"),
//   		RecordPreprocessorSourceUri: jsii.String("recordPreprocessorSourceUri"),
//   	},
//   	ModelQualityBaselineConfig: &ModelQualityBaselineConfigProperty{
//   		BaseliningJobName: jsii.String("baseliningJobName"),
//   		ConstraintsResource: &ConstraintsResourceProperty{
//   			S3Uri: jsii.String("s3Uri"),
//   		},
//   	},
//   	ModelQualityJobInput: &ModelQualityJobInputProperty{
//   		BatchTransformInput: &BatchTransformInputProperty{
//   			DataCapturedDestinationS3Uri: jsii.String("dataCapturedDestinationS3Uri"),
//   			DatasetFormat: &DatasetFormatProperty{
//   				Csv: &CsvProperty{
//   					Header: jsii.Boolean(false),
//   				},
//   				Json: &JsonProperty{
//   					Line: jsii.Boolean(false),
//   				},
//   				Parquet: jsii.Boolean(false),
//   			},
//   			EndTimeOffset: jsii.String("endTimeOffset"),
//   			InferenceAttribute: jsii.String("inferenceAttribute"),
//   			LocalPath: jsii.String("localPath"),
//   			ProbabilityAttribute: jsii.String("probabilityAttribute"),
//   			ProbabilityThresholdAttribute: jsii.Number(123),
//   			S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   			S3InputMode: jsii.String("s3InputMode"),
//   			StartTimeOffset: jsii.String("startTimeOffset"),
//   		},
//   		EndpointInput: &EndpointInputProperty{
//   			EndpointName: jsii.String("endpointName"),
//   			EndTimeOffset: jsii.String("endTimeOffset"),
//   			InferenceAttribute: jsii.String("inferenceAttribute"),
//   			LocalPath: jsii.String("localPath"),
//   			ProbabilityAttribute: jsii.String("probabilityAttribute"),
//   			ProbabilityThresholdAttribute: jsii.Number(123),
//   			S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   			S3InputMode: jsii.String("s3InputMode"),
//   			StartTimeOffset: jsii.String("startTimeOffset"),
//   		},
//   		GroundTruthS3Input: &MonitoringGroundTruthS3InputProperty{
//   			S3Uri: jsii.String("s3Uri"),
//   		},
//   	},
//   	ModelQualityJobOutputConfig: &MonitoringOutputConfigProperty{
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   		MonitoringOutputs: []interface{}{
//   			&MonitoringOutputProperty{
//   				S3Output: &S3OutputProperty{
//   					LocalPath: jsii.String("localPath"),
//   					S3UploadMode: jsii.String("s3UploadMode"),
//   					S3Uri: jsii.String("s3Uri"),
//   				},
//   			},
//   		},
//   	},
//   	NetworkConfig: &NetworkConfigProperty{
//   		EnableInterContainerTrafficEncryption: jsii.Boolean(false),
//   		EnableNetworkIsolation: jsii.Boolean(false),
//   		VpcConfig: &VpcConfigProperty{
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			Subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	StoppingCondition: &StoppingConditionProperty{
//   		MaxRuntimeInSeconds: jsii.Number(123),
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html
//
type CfnModelQualityJobDefinitionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnModelQualityJobDefinitionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnModelQualityJobDefinitionPropsMixin
type jsiiProxy_CfnModelQualityJobDefinitionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnModelQualityJobDefinitionPropsMixin) Props() *CfnModelQualityJobDefinitionMixinProps {
	var returns *CfnModelQualityJobDefinitionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelQualityJobDefinitionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SageMaker::ModelQualityJobDefinition`.
func NewCfnModelQualityJobDefinitionPropsMixin(props *CfnModelQualityJobDefinitionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnModelQualityJobDefinitionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnModelQualityJobDefinitionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnModelQualityJobDefinitionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelQualityJobDefinitionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SageMaker::ModelQualityJobDefinition`.
func NewCfnModelQualityJobDefinitionPropsMixin_Override(c CfnModelQualityJobDefinitionPropsMixin, props *CfnModelQualityJobDefinitionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelQualityJobDefinitionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnModelQualityJobDefinitionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnModelQualityJobDefinitionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelQualityJobDefinitionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnModelQualityJobDefinitionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelQualityJobDefinitionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnModelQualityJobDefinitionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnModelQualityJobDefinitionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

