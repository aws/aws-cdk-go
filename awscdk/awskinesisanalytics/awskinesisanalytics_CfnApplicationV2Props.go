package awskinesisanalytics

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnApplicationV2`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationV2Props := &CfnApplicationV2Props{
//   	RuntimeEnvironment: jsii.String("runtimeEnvironment"),
//   	ServiceExecutionRole: jsii.String("serviceExecutionRole"),
//
//   	// the properties below are optional
//   	ApplicationConfiguration: &ApplicationConfigurationProperty{
//   		ApplicationCodeConfiguration: &ApplicationCodeConfigurationProperty{
//   			CodeContent: &CodeContentProperty{
//   				S3ContentLocation: &S3ContentLocationProperty{
//   					BucketArn: jsii.String("bucketArn"),
//   					FileKey: jsii.String("fileKey"),
//
//   					// the properties below are optional
//   					ObjectVersion: jsii.String("objectVersion"),
//   				},
//   				TextContent: jsii.String("textContent"),
//   				ZipFileContent: jsii.String("zipFileContent"),
//   			},
//   			CodeContentType: jsii.String("codeContentType"),
//   		},
//   		ApplicationSnapshotConfiguration: &ApplicationSnapshotConfigurationProperty{
//   			SnapshotsEnabled: jsii.Boolean(false),
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
//   				ConfigurationType: jsii.String("configurationType"),
//
//   				// the properties below are optional
//   				CheckpointingEnabled: jsii.Boolean(false),
//   				CheckpointInterval: jsii.Number(123),
//   				MinPauseBetweenCheckpoints: jsii.Number(123),
//   			},
//   			MonitoringConfiguration: &MonitoringConfigurationProperty{
//   				ConfigurationType: jsii.String("configurationType"),
//
//   				// the properties below are optional
//   				LogLevel: jsii.String("logLevel"),
//   				MetricsLevel: jsii.String("metricsLevel"),
//   			},
//   			ParallelismConfiguration: &ParallelismConfigurationProperty{
//   				ConfigurationType: jsii.String("configurationType"),
//
//   				// the properties below are optional
//   				AutoScalingEnabled: jsii.Boolean(false),
//   				Parallelism: jsii.Number(123),
//   				ParallelismPerKpu: jsii.Number(123),
//   			},
//   		},
//   		SqlApplicationConfiguration: &SqlApplicationConfigurationProperty{
//   			Inputs: []interface{}{
//   				&InputProperty{
//   					InputSchema: &InputSchemaProperty{
//   						RecordColumns: []interface{}{
//   							&RecordColumnProperty{
//   								Name: jsii.String("name"),
//   								SqlType: jsii.String("sqlType"),
//
//   								// the properties below are optional
//   								Mapping: jsii.String("mapping"),
//   							},
//   						},
//   						RecordFormat: &RecordFormatProperty{
//   							RecordFormatType: jsii.String("recordFormatType"),
//
//   							// the properties below are optional
//   							MappingParameters: &MappingParametersProperty{
//   								CsvMappingParameters: &CSVMappingParametersProperty{
//   									RecordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   									RecordRowDelimiter: jsii.String("recordRowDelimiter"),
//   								},
//   								JsonMappingParameters: &JSONMappingParametersProperty{
//   									RecordRowPath: jsii.String("recordRowPath"),
//   								},
//   							},
//   						},
//
//   						// the properties below are optional
//   						RecordEncoding: jsii.String("recordEncoding"),
//   					},
//   					NamePrefix: jsii.String("namePrefix"),
//
//   					// the properties below are optional
//   					InputParallelism: &InputParallelismProperty{
//   						Count: jsii.Number(123),
//   					},
//   					InputProcessingConfiguration: &InputProcessingConfigurationProperty{
//   						InputLambdaProcessor: &InputLambdaProcessorProperty{
//   							ResourceArn: jsii.String("resourceArn"),
//   						},
//   					},
//   					KinesisFirehoseInput: &KinesisFirehoseInputProperty{
//   						ResourceArn: jsii.String("resourceArn"),
//   					},
//   					KinesisStreamsInput: &KinesisStreamsInputProperty{
//   						ResourceArn: jsii.String("resourceArn"),
//   					},
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
//
//   					// the properties below are optional
//   					MavenReference: &MavenReferenceProperty{
//   						ArtifactId: jsii.String("artifactId"),
//   						GroupId: jsii.String("groupId"),
//   						Version: jsii.String("version"),
//   					},
//   					S3ContentLocation: &S3ContentLocationProperty{
//   						BucketArn: jsii.String("bucketArn"),
//   						FileKey: jsii.String("fileKey"),
//
//   						// the properties below are optional
//   						ObjectVersion: jsii.String("objectVersion"),
//   					},
//   				},
//   			},
//   			DeployAsApplicationConfiguration: &DeployAsApplicationConfigurationProperty{
//   				S3ContentLocation: &S3ContentBaseLocationProperty{
//   					BucketArn: jsii.String("bucketArn"),
//
//   					// the properties below are optional
//   					BasePath: jsii.String("basePath"),
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
//
//   			// the properties below are optional
//   			SnapshotName: jsii.String("snapshotName"),
//   		},
//   		FlinkRunConfiguration: &FlinkRunConfigurationProperty{
//   			AllowNonRestoredState: jsii.Boolean(false),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnApplicationV2Props struct {
	// The runtime environment for the application.
	RuntimeEnvironment *string `field:"required" json:"runtimeEnvironment" yaml:"runtimeEnvironment"`
	// Specifies the IAM role that the application uses to access external resources.
	ServiceExecutionRole *string `field:"required" json:"serviceExecutionRole" yaml:"serviceExecutionRole"`
	// Use this parameter to configure the application.
	ApplicationConfiguration interface{} `field:"optional" json:"applicationConfiguration" yaml:"applicationConfiguration"`
	// The description of the application.
	ApplicationDescription *string `field:"optional" json:"applicationDescription" yaml:"applicationDescription"`
	// `AWS::KinesisAnalyticsV2::Application.ApplicationMaintenanceConfiguration`.
	ApplicationMaintenanceConfiguration interface{} `field:"optional" json:"applicationMaintenanceConfiguration" yaml:"applicationMaintenanceConfiguration"`
	// To create a Kinesis Data Analytics Studio notebook, you must set the mode to `INTERACTIVE` .
	//
	// However, for a Kinesis Data Analytics for Apache Flink application, the mode is optional.
	ApplicationMode *string `field:"optional" json:"applicationMode" yaml:"applicationMode"`
	// The name of the application.
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
	// `AWS::KinesisAnalyticsV2::Application.RunConfiguration`.
	RunConfiguration interface{} `field:"optional" json:"runConfiguration" yaml:"runConfiguration"`
	// A list of one or more tags to assign to the application.
	//
	// A tag is a key-value pair that identifies an application. Note that the maximum number of application tags includes system tags. The maximum number of user-defined application tags is 50.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

