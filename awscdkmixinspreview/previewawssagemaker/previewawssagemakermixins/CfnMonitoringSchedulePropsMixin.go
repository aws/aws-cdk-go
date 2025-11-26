package previewawssagemakermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssagemaker/previewawssagemakermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::SageMaker::MonitoringSchedule` resource is an Amazon SageMaker resource type that regularly starts SageMaker processing Jobs to monitor the data captured for a SageMaker endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMonitoringSchedulePropsMixin := awscdkmixinspreview.Mixins.NewCfnMonitoringSchedulePropsMixin(&CfnMonitoringScheduleMixinProps{
//   	EndpointName: jsii.String("endpointName"),
//   	FailureReason: jsii.String("failureReason"),
//   	LastMonitoringExecutionSummary: &MonitoringExecutionSummaryProperty{
//   		CreationTime: jsii.String("creationTime"),
//   		EndpointName: jsii.String("endpointName"),
//   		FailureReason: jsii.String("failureReason"),
//   		LastModifiedTime: jsii.String("lastModifiedTime"),
//   		MonitoringExecutionStatus: jsii.String("monitoringExecutionStatus"),
//   		MonitoringScheduleName: jsii.String("monitoringScheduleName"),
//   		ProcessingJobArn: jsii.String("processingJobArn"),
//   		ScheduledTime: jsii.String("scheduledTime"),
//   	},
//   	MonitoringScheduleConfig: &MonitoringScheduleConfigProperty{
//   		MonitoringJobDefinition: &MonitoringJobDefinitionProperty{
//   			BaselineConfig: &BaselineConfigProperty{
//   				ConstraintsResource: &ConstraintsResourceProperty{
//   					S3Uri: jsii.String("s3Uri"),
//   				},
//   				StatisticsResource: &StatisticsResourceProperty{
//   					S3Uri: jsii.String("s3Uri"),
//   				},
//   			},
//   			Environment: map[string]*string{
//   				"environmentKey": jsii.String("environment"),
//   			},
//   			MonitoringAppSpecification: &MonitoringAppSpecificationProperty{
//   				ContainerArguments: []*string{
//   					jsii.String("containerArguments"),
//   				},
//   				ContainerEntrypoint: []*string{
//   					jsii.String("containerEntrypoint"),
//   				},
//   				ImageUri: jsii.String("imageUri"),
//   				PostAnalyticsProcessorSourceUri: jsii.String("postAnalyticsProcessorSourceUri"),
//   				RecordPreprocessorSourceUri: jsii.String("recordPreprocessorSourceUri"),
//   			},
//   			MonitoringInputs: []interface{}{
//   				&MonitoringInputProperty{
//   					BatchTransformInput: &BatchTransformInputProperty{
//   						DataCapturedDestinationS3Uri: jsii.String("dataCapturedDestinationS3Uri"),
//   						DatasetFormat: &DatasetFormatProperty{
//   							Csv: &CsvProperty{
//   								Header: jsii.Boolean(false),
//   							},
//   							Json: &JsonProperty{
//   								Line: jsii.Boolean(false),
//   							},
//   							Parquet: jsii.Boolean(false),
//   						},
//   						ExcludeFeaturesAttribute: jsii.String("excludeFeaturesAttribute"),
//   						LocalPath: jsii.String("localPath"),
//   						S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   						S3InputMode: jsii.String("s3InputMode"),
//   					},
//   					EndpointInput: &EndpointInputProperty{
//   						EndpointName: jsii.String("endpointName"),
//   						ExcludeFeaturesAttribute: jsii.String("excludeFeaturesAttribute"),
//   						LocalPath: jsii.String("localPath"),
//   						S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   						S3InputMode: jsii.String("s3InputMode"),
//   					},
//   				},
//   			},
//   			MonitoringOutputConfig: &MonitoringOutputConfigProperty{
//   				KmsKeyId: jsii.String("kmsKeyId"),
//   				MonitoringOutputs: []interface{}{
//   					&MonitoringOutputProperty{
//   						S3Output: &S3OutputProperty{
//   							LocalPath: jsii.String("localPath"),
//   							S3UploadMode: jsii.String("s3UploadMode"),
//   							S3Uri: jsii.String("s3Uri"),
//   						},
//   					},
//   				},
//   			},
//   			MonitoringResources: &MonitoringResourcesProperty{
//   				ClusterConfig: &ClusterConfigProperty{
//   					InstanceCount: jsii.Number(123),
//   					InstanceType: jsii.String("instanceType"),
//   					VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   					VolumeSizeInGb: jsii.Number(123),
//   				},
//   			},
//   			NetworkConfig: &NetworkConfigProperty{
//   				EnableInterContainerTrafficEncryption: jsii.Boolean(false),
//   				EnableNetworkIsolation: jsii.Boolean(false),
//   				VpcConfig: &VpcConfigProperty{
//   					SecurityGroupIds: []*string{
//   						jsii.String("securityGroupIds"),
//   					},
//   					Subnets: []*string{
//   						jsii.String("subnets"),
//   					},
//   				},
//   			},
//   			RoleArn: jsii.String("roleArn"),
//   			StoppingCondition: &StoppingConditionProperty{
//   				MaxRuntimeInSeconds: jsii.Number(123),
//   			},
//   		},
//   		MonitoringJobDefinitionName: jsii.String("monitoringJobDefinitionName"),
//   		MonitoringType: jsii.String("monitoringType"),
//   		ScheduleConfig: &ScheduleConfigProperty{
//   			DataAnalysisEndTime: jsii.String("dataAnalysisEndTime"),
//   			DataAnalysisStartTime: jsii.String("dataAnalysisStartTime"),
//   			ScheduleExpression: jsii.String("scheduleExpression"),
//   		},
//   	},
//   	MonitoringScheduleName: jsii.String("monitoringScheduleName"),
//   	MonitoringScheduleStatus: jsii.String("monitoringScheduleStatus"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-monitoringschedule.html
//
type CfnMonitoringSchedulePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnMonitoringScheduleMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMonitoringSchedulePropsMixin
type jsiiProxy_CfnMonitoringSchedulePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnMonitoringSchedulePropsMixin) Props() *CfnMonitoringScheduleMixinProps {
	var returns *CfnMonitoringScheduleMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMonitoringSchedulePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SageMaker::MonitoringSchedule`.
func NewCfnMonitoringSchedulePropsMixin(props *CfnMonitoringScheduleMixinProps, options *mixins.CfnPropertyMixinOptions) CfnMonitoringSchedulePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMonitoringSchedulePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMonitoringSchedulePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnMonitoringSchedulePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SageMaker::MonitoringSchedule`.
func NewCfnMonitoringSchedulePropsMixin_Override(c CfnMonitoringSchedulePropsMixin, props *CfnMonitoringScheduleMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnMonitoringSchedulePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnMonitoringSchedulePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMonitoringSchedulePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnMonitoringSchedulePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMonitoringSchedulePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnMonitoringSchedulePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMonitoringSchedulePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnMonitoringSchedulePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

