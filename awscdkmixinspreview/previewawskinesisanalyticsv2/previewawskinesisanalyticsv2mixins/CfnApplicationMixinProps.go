package previewawskinesisanalyticsv2mixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnApplicationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationMixinProps := &CfnApplicationMixinProps{
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-application.html
//
type CfnApplicationMixinProps struct {
	// Use this parameter to configure the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-application.html#cfn-kinesisanalyticsv2-application-applicationconfiguration
	//
	ApplicationConfiguration interface{} `field:"optional" json:"applicationConfiguration" yaml:"applicationConfiguration"`
	// The description of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-application.html#cfn-kinesisanalyticsv2-application-applicationdescription
	//
	// Default: - "".
	//
	ApplicationDescription *string `field:"optional" json:"applicationDescription" yaml:"applicationDescription"`
	// Specifies the maintenance window parameters for a Kinesis Data Analytics application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-application.html#cfn-kinesisanalyticsv2-application-applicationmaintenanceconfiguration
	//
	ApplicationMaintenanceConfiguration interface{} `field:"optional" json:"applicationMaintenanceConfiguration" yaml:"applicationMaintenanceConfiguration"`
	// To create a Kinesis Data Analytics Studio notebook, you must set the mode to `INTERACTIVE` .
	//
	// However, for a Kinesis Data Analytics for Apache Flink application, the mode is optional.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-application.html#cfn-kinesisanalyticsv2-application-applicationmode
	//
	ApplicationMode *string `field:"optional" json:"applicationMode" yaml:"applicationMode"`
	// The name of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-application.html#cfn-kinesisanalyticsv2-application-applicationname
	//
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
	// Describes the starting parameters for an Managed Service for Apache Flink application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-application.html#cfn-kinesisanalyticsv2-application-runconfiguration
	//
	RunConfiguration interface{} `field:"optional" json:"runConfiguration" yaml:"runConfiguration"`
	// The runtime environment for the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-application.html#cfn-kinesisanalyticsv2-application-runtimeenvironment
	//
	RuntimeEnvironment *string `field:"optional" json:"runtimeEnvironment" yaml:"runtimeEnvironment"`
	// Specifies the IAM role that the application uses to access external resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-application.html#cfn-kinesisanalyticsv2-application-serviceexecutionrole
	//
	ServiceExecutionRole *string `field:"optional" json:"serviceExecutionRole" yaml:"serviceExecutionRole"`
	// A list of one or more tags to assign to the application.
	//
	// A tag is a key-value pair that identifies an application. Note that the maximum number of application tags includes system tags. The maximum number of user-defined application tags is 50.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-application.html#cfn-kinesisanalyticsv2-application-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

