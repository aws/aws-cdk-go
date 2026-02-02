package previewawssagemakermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssagemaker/previewawssagemakermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An Amazon SageMaker processing job that is used to analyze data and evaluate models.
//
// For more information, see [Process Data and Evaluate Models](https://docs.aws.amazon.com/sagemaker/latest/dg/processing-job.html) .
//
// Also, note the following details specific to processing jobs created using CloudFormation stacks:
//
// - When you delete a CloudFormation stack with a processing job resource, the processing job is stopped using the [StopProcessingJob](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_StopProcessingJob.html) API but not deleted. Any tags associated with the processing job are deleted using the [DeleteTags](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DeleteTags.html) API.
// - If any part of your CloudFormation stack deployment fails and a rollback initiates, processing jobs with a specified `ProcessingJobName` value might cause the stack to become stuck in a failed state. This occurs because during a rollback, CloudFormation attempts to recreate the stack resources. Processing job names must be unique, so when CloudFormation attempts to recreate a processing job using the already defined name, this results in an `AlreadyExists` error. To prevent this, we recommend that you don't specify the optional `ProcessingJobName` property, thereby allowing SageMaker to auto-generate a unique name for your processing job. This ensures successful stack rollbacks when necessary. If you must use custom job names, you have to manually modify the `ProcessingJobName` and redeploy the stack to recover from a failed rollback.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnProcessingJobPropsMixin := awscdkmixinspreview.Mixins.NewCfnProcessingJobPropsMixin(&CfnProcessingJobMixinProps{
//   	AppSpecification: &AppSpecificationProperty{
//   		ContainerArguments: []*string{
//   			jsii.String("containerArguments"),
//   		},
//   		ContainerEntrypoint: []*string{
//   			jsii.String("containerEntrypoint"),
//   		},
//   		ImageUri: jsii.String("imageUri"),
//   	},
//   	Environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	ExperimentConfig: &ExperimentConfigProperty{
//   		ExperimentName: jsii.String("experimentName"),
//   		RunName: jsii.String("runName"),
//   		TrialComponentDisplayName: jsii.String("trialComponentDisplayName"),
//   		TrialName: jsii.String("trialName"),
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
//   	ProcessingInputs: []interface{}{
//   		&ProcessingInputsObjectProperty{
//   			AppManaged: jsii.Boolean(false),
//   			DatasetDefinition: &DatasetDefinitionProperty{
//   				AthenaDatasetDefinition: &AthenaDatasetDefinitionProperty{
//   					Catalog: jsii.String("catalog"),
//   					Database: jsii.String("database"),
//   					KmsKeyId: jsii.String("kmsKeyId"),
//   					OutputCompression: jsii.String("outputCompression"),
//   					OutputFormat: jsii.String("outputFormat"),
//   					OutputS3Uri: jsii.String("outputS3Uri"),
//   					QueryString: jsii.String("queryString"),
//   					WorkGroup: jsii.String("workGroup"),
//   				},
//   				DataDistributionType: jsii.String("dataDistributionType"),
//   				InputMode: jsii.String("inputMode"),
//   				LocalPath: jsii.String("localPath"),
//   				RedshiftDatasetDefinition: &RedshiftDatasetDefinitionProperty{
//   					ClusterId: jsii.String("clusterId"),
//   					ClusterRoleArn: jsii.String("clusterRoleArn"),
//   					Database: jsii.String("database"),
//   					DbUser: jsii.String("dbUser"),
//   					KmsKeyId: jsii.String("kmsKeyId"),
//   					OutputCompression: jsii.String("outputCompression"),
//   					OutputFormat: jsii.String("outputFormat"),
//   					OutputS3Uri: jsii.String("outputS3Uri"),
//   					QueryString: jsii.String("queryString"),
//   				},
//   			},
//   			InputName: jsii.String("inputName"),
//   			S3Input: &S3InputProperty{
//   				LocalPath: jsii.String("localPath"),
//   				S3CompressionType: jsii.String("s3CompressionType"),
//   				S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   				S3DataType: jsii.String("s3DataType"),
//   				S3InputMode: jsii.String("s3InputMode"),
//   				S3Uri: jsii.String("s3Uri"),
//   			},
//   		},
//   	},
//   	ProcessingJobName: jsii.String("processingJobName"),
//   	ProcessingOutputConfig: &ProcessingOutputConfigProperty{
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   		Outputs: []interface{}{
//   			&ProcessingOutputsObjectProperty{
//   				AppManaged: jsii.Boolean(false),
//   				FeatureStoreOutput: &FeatureStoreOutputProperty{
//   					FeatureGroupName: jsii.String("featureGroupName"),
//   				},
//   				OutputName: jsii.String("outputName"),
//   				S3Output: &S3OutputProperty{
//   					LocalPath: jsii.String("localPath"),
//   					S3UploadMode: jsii.String("s3UploadMode"),
//   					S3Uri: jsii.String("s3Uri"),
//   				},
//   			},
//   		},
//   	},
//   	ProcessingResources: &ProcessingResourcesProperty{
//   		ClusterConfig: &ClusterConfigProperty{
//   			InstanceCount: jsii.Number(123),
//   			InstanceType: jsii.String("instanceType"),
//   			VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   			VolumeSizeInGb: jsii.Number(123),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-processingjob.html
//
type CfnProcessingJobPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnProcessingJobMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnProcessingJobPropsMixin
type jsiiProxy_CfnProcessingJobPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnProcessingJobPropsMixin) Props() *CfnProcessingJobMixinProps {
	var returns *CfnProcessingJobMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcessingJobPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SageMaker::ProcessingJob`.
func NewCfnProcessingJobPropsMixin(props *CfnProcessingJobMixinProps, options *mixins.CfnPropertyMixinOptions) CfnProcessingJobPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnProcessingJobPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnProcessingJobPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnProcessingJobPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SageMaker::ProcessingJob`.
func NewCfnProcessingJobPropsMixin_Override(c CfnProcessingJobPropsMixin, props *CfnProcessingJobMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnProcessingJobPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnProcessingJobPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnProcessingJobPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnProcessingJobPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnProcessingJobPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnProcessingJobPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnProcessingJobPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnProcessingJobPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

