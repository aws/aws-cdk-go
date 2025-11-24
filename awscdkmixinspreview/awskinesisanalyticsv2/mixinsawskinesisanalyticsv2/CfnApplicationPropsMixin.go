package mixinsawskinesisanalyticsv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awskinesisanalyticsv2/mixinsawskinesisanalyticsv2/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an Amazon Kinesis Data Analytics application.
//
// For information about creating a Kinesis Data Analytics application, see [Creating an Application](https://docs.aws.amazon.com/managed-flink/latest/java/getting-started.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationPropsMixin := awscdkmixinspreview.Mixins.NewCfnApplicationPropsMixin(&CfnApplicationMixinProps{
//   	ApplicationConfiguration: &ApplicationConfigurationProperty{
//   		ApplicationCodeConfiguration: &ApplicationCodeConfigurationProperty{
//   			CodeContent: &CodeContentProperty{
//   				S3ContentLocation: &S3ContentLocationProperty{
//   					BucketArn: jsii.String("bucketArn"),
//   					FileKey: jsii.String("fileKey"),
//   					ObjectVersion: jsii.String("objectVersion"),
//   				},
//   				TextContent: jsii.String("textContent"),
//   				ZipFileContent: jsii.String("zipFileContent"),
//   			},
//   			CodeContentType: jsii.String("codeContentType"),
//   		},
//   		ApplicationEncryptionConfiguration: &ApplicationEncryptionConfigurationProperty{
//   			KeyId: jsii.String("keyId"),
//   			KeyType: jsii.String("keyType"),
//   		},
//   		ApplicationSnapshotConfiguration: &ApplicationSnapshotConfigurationProperty{
//   			SnapshotsEnabled: jsii.Boolean(false),
//   		},
//   		ApplicationSystemRollbackConfiguration: &ApplicationSystemRollbackConfigurationProperty{
//   			RollbackEnabled: jsii.Boolean(false),
//   		},
//   		EnvironmentProperties: &EnvironmentPropertiesProperty{
//   			PropertyGroups: []interface{}{
//   				&PropertyGroupProperty{
//   					PropertyGroupId: jsii.String("propertyGroupId"),
//   					PropertyMap: map[string]*string{
//   						"propertyMapKey": jsii.String("propertyMap"),
//   					},
//   				},
//   			},
//   		},
//   		FlinkApplicationConfiguration: &FlinkApplicationConfigurationProperty{
//   			CheckpointConfiguration: &CheckpointConfigurationProperty{
//   				CheckpointingEnabled: jsii.Boolean(false),
//   				CheckpointInterval: jsii.Number(123),
//   				ConfigurationType: jsii.String("configurationType"),
//   				MinPauseBetweenCheckpoints: jsii.Number(123),
//   			},
//   			MonitoringConfiguration: &MonitoringConfigurationProperty{
//   				ConfigurationType: jsii.String("configurationType"),
//   				LogLevel: jsii.String("logLevel"),
//   				MetricsLevel: jsii.String("metricsLevel"),
//   			},
//   			ParallelismConfiguration: &ParallelismConfigurationProperty{
//   				AutoScalingEnabled: jsii.Boolean(false),
//   				ConfigurationType: jsii.String("configurationType"),
//   				Parallelism: jsii.Number(123),
//   				ParallelismPerKpu: jsii.Number(123),
//   			},
//   		},
//   		SqlApplicationConfiguration: &SqlApplicationConfigurationProperty{
//   			Inputs: []interface{}{
//   				&InputProperty{
//   					InputParallelism: &InputParallelismProperty{
//   						Count: jsii.Number(123),
//   					},
//   					InputProcessingConfiguration: &InputProcessingConfigurationProperty{
//   						InputLambdaProcessor: &InputLambdaProcessorProperty{
//   							ResourceArn: jsii.String("resourceArn"),
//   						},
//   					},
//   					InputSchema: &InputSchemaProperty{
//   						RecordColumns: []interface{}{
//   							&RecordColumnProperty{
//   								Mapping: jsii.String("mapping"),
//   								Name: jsii.String("name"),
//   								SqlType: jsii.String("sqlType"),
//   							},
//   						},
//   						RecordEncoding: jsii.String("recordEncoding"),
//   						RecordFormat: &RecordFormatProperty{
//   							MappingParameters: &MappingParametersProperty{
//   								CsvMappingParameters: &CSVMappingParametersProperty{
//   									RecordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   									RecordRowDelimiter: jsii.String("recordRowDelimiter"),
//   								},
//   								JsonMappingParameters: &JSONMappingParametersProperty{
//   									RecordRowPath: jsii.String("recordRowPath"),
//   								},
//   							},
//   							RecordFormatType: jsii.String("recordFormatType"),
//   						},
//   					},
//   					KinesisFirehoseInput: &KinesisFirehoseInputProperty{
//   						ResourceArn: jsii.String("resourceArn"),
//   					},
//   					KinesisStreamsInput: &KinesisStreamsInputProperty{
//   						ResourceArn: jsii.String("resourceArn"),
//   					},
//   					NamePrefix: jsii.String("namePrefix"),
//   				},
//   			},
//   		},
//   		VpcConfigurations: []interface{}{
//   			&VpcConfigurationProperty{
//   				SecurityGroupIds: []*string{
//   					jsii.String("securityGroupIds"),
//   				},
//   				SubnetIds: []*string{
//   					jsii.String("subnetIds"),
//   				},
//   			},
//   		},
//   		ZeppelinApplicationConfiguration: &ZeppelinApplicationConfigurationProperty{
//   			CatalogConfiguration: &CatalogConfigurationProperty{
//   				GlueDataCatalogConfiguration: &GlueDataCatalogConfigurationProperty{
//   					DatabaseArn: jsii.String("databaseArn"),
//   				},
//   			},
//   			CustomArtifactsConfiguration: []interface{}{
//   				&CustomArtifactConfigurationProperty{
//   					ArtifactType: jsii.String("artifactType"),
//   					MavenReference: &MavenReferenceProperty{
//   						ArtifactId: jsii.String("artifactId"),
//   						GroupId: jsii.String("groupId"),
//   						Version: jsii.String("version"),
//   					},
//   					S3ContentLocation: &S3ContentLocationProperty{
//   						BucketArn: jsii.String("bucketArn"),
//   						FileKey: jsii.String("fileKey"),
//   						ObjectVersion: jsii.String("objectVersion"),
//   					},
//   				},
//   			},
//   			DeployAsApplicationConfiguration: &DeployAsApplicationConfigurationProperty{
//   				S3ContentLocation: &S3ContentBaseLocationProperty{
//   					BasePath: jsii.String("basePath"),
//   					BucketArn: jsii.String("bucketArn"),
//   				},
//   			},
//   			MonitoringConfiguration: &ZeppelinMonitoringConfigurationProperty{
//   				LogLevel: jsii.String("logLevel"),
//   			},
//   		},
//   	},
//   	ApplicationDescription: jsii.String("applicationDescription"),
//   	ApplicationMaintenanceConfiguration: &ApplicationMaintenanceConfigurationProperty{
//   		ApplicationMaintenanceWindowStartTime: jsii.String("applicationMaintenanceWindowStartTime"),
//   	},
//   	ApplicationMode: jsii.String("applicationMode"),
//   	ApplicationName: jsii.String("applicationName"),
//   	RunConfiguration: &RunConfigurationProperty{
//   		ApplicationRestoreConfiguration: &ApplicationRestoreConfigurationProperty{
//   			ApplicationRestoreType: jsii.String("applicationRestoreType"),
//   			SnapshotName: jsii.String("snapshotName"),
//   		},
//   		FlinkRunConfiguration: &FlinkRunConfigurationProperty{
//   			AllowNonRestoredState: jsii.Boolean(false),
//   		},
//   	},
//   	RuntimeEnvironment: jsii.String("runtimeEnvironment"),
//   	ServiceExecutionRole: jsii.String("serviceExecutionRole"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-application.html
//
type CfnApplicationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnApplicationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnApplicationPropsMixin
type jsiiProxy_CfnApplicationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnApplicationPropsMixin) Props() *CfnApplicationMixinProps {
	var returns *CfnApplicationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::KinesisAnalyticsV2::Application`.
func NewCfnApplicationPropsMixin(props *CfnApplicationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnApplicationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnApplicationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnApplicationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::KinesisAnalyticsV2::Application`.
func NewCfnApplicationPropsMixin_Override(c CfnApplicationPropsMixin, props *CfnApplicationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnApplicationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApplicationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApplicationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApplicationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnApplicationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

