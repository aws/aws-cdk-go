package awscdkkinesisanalyticsflinkalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisanalytics"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// The return type of `ApplicationCode.bind`. This represents CloudFormation configuration and an s3 bucket holding the Flink application JAR file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import kinesisanalytics_flink_alpha "github.com/aws/aws-cdk-go/awscdkkinesisanalyticsflinkalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   applicationCodeConfig := &ApplicationCodeConfig{
//   	ApplicationCodeConfigurationProperty: &ApplicationConfigurationProperty{
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
//   	Bucket: bucket,
//   }
//
// Experimental.
type ApplicationCodeConfig struct {
	// Low-level Cloudformation ApplicationConfigurationProperty.
	// Experimental.
	ApplicationCodeConfigurationProperty *awskinesisanalytics.CfnApplicationV2_ApplicationConfigurationProperty `field:"required" json:"applicationCodeConfigurationProperty" yaml:"applicationCodeConfigurationProperty"`
	// S3 Bucket that stores the Flink application code.
	// Experimental.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
}

