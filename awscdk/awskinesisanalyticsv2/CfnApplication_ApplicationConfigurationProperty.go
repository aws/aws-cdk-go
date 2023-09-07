package awskinesisanalyticsv2


// Specifies the creation parameters for a Managed Service for Apache Flink application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationConfigurationProperty := &ApplicationConfigurationProperty{
//   	ApplicationCodeConfiguration: &ApplicationCodeConfigurationProperty{
//   		CodeContent: &CodeContentProperty{
//   			S3ContentLocation: &S3ContentLocationProperty{
//   				BucketArn: jsii.String("bucketArn"),
//   				FileKey: jsii.String("fileKey"),
//
//   				// the properties below are optional
//   				ObjectVersion: jsii.String("objectVersion"),
//   			},
//   			TextContent: jsii.String("textContent"),
//   			ZipFileContent: jsii.String("zipFileContent"),
//   		},
//   		CodeContentType: jsii.String("codeContentType"),
//   	},
//   	ApplicationSnapshotConfiguration: &ApplicationSnapshotConfigurationProperty{
//   		SnapshotsEnabled: jsii.Boolean(false),
//   	},
//   	EnvironmentProperties: &EnvironmentPropertiesProperty{
//   		PropertyGroups: []interface{}{
//   			&PropertyGroupProperty{
//   				PropertyGroupId: jsii.String("propertyGroupId"),
//   				PropertyMap: map[string]*string{
//   					"propertyMapKey": jsii.String("propertyMap"),
//   				},
//   			},
//   		},
//   	},
//   	FlinkApplicationConfiguration: &FlinkApplicationConfigurationProperty{
//   		CheckpointConfiguration: &CheckpointConfigurationProperty{
//   			ConfigurationType: jsii.String("configurationType"),
//
//   			// the properties below are optional
//   			CheckpointingEnabled: jsii.Boolean(false),
//   			CheckpointInterval: jsii.Number(123),
//   			MinPauseBetweenCheckpoints: jsii.Number(123),
//   		},
//   		MonitoringConfiguration: &MonitoringConfigurationProperty{
//   			ConfigurationType: jsii.String("configurationType"),
//
//   			// the properties below are optional
//   			LogLevel: jsii.String("logLevel"),
//   			MetricsLevel: jsii.String("metricsLevel"),
//   		},
//   		ParallelismConfiguration: &ParallelismConfigurationProperty{
//   			ConfigurationType: jsii.String("configurationType"),
//
//   			// the properties below are optional
//   			AutoScalingEnabled: jsii.Boolean(false),
//   			Parallelism: jsii.Number(123),
//   			ParallelismPerKpu: jsii.Number(123),
//   		},
//   	},
//   	SqlApplicationConfiguration: &SqlApplicationConfigurationProperty{
//   		Inputs: []interface{}{
//   			&InputProperty{
//   				InputSchema: &InputSchemaProperty{
//   					RecordColumns: []interface{}{
//   						&RecordColumnProperty{
//   							Name: jsii.String("name"),
//   							SqlType: jsii.String("sqlType"),
//
//   							// the properties below are optional
//   							Mapping: jsii.String("mapping"),
//   						},
//   					},
//   					RecordFormat: &RecordFormatProperty{
//   						RecordFormatType: jsii.String("recordFormatType"),
//
//   						// the properties below are optional
//   						MappingParameters: &MappingParametersProperty{
//   							CsvMappingParameters: &CSVMappingParametersProperty{
//   								RecordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   								RecordRowDelimiter: jsii.String("recordRowDelimiter"),
//   							},
//   							JsonMappingParameters: &JSONMappingParametersProperty{
//   								RecordRowPath: jsii.String("recordRowPath"),
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					RecordEncoding: jsii.String("recordEncoding"),
//   				},
//   				NamePrefix: jsii.String("namePrefix"),
//
//   				// the properties below are optional
//   				InputParallelism: &InputParallelismProperty{
//   					Count: jsii.Number(123),
//   				},
//   				InputProcessingConfiguration: &InputProcessingConfigurationProperty{
//   					InputLambdaProcessor: &InputLambdaProcessorProperty{
//   						ResourceArn: jsii.String("resourceArn"),
//   					},
//   				},
//   				KinesisFirehoseInput: &KinesisFirehoseInputProperty{
//   					ResourceArn: jsii.String("resourceArn"),
//   				},
//   				KinesisStreamsInput: &KinesisStreamsInputProperty{
//   					ResourceArn: jsii.String("resourceArn"),
//   				},
//   			},
//   		},
//   	},
//   	VpcConfigurations: []interface{}{
//   		&VpcConfigurationProperty{
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			SubnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//   		},
//   	},
//   	ZeppelinApplicationConfiguration: &ZeppelinApplicationConfigurationProperty{
//   		CatalogConfiguration: &CatalogConfigurationProperty{
//   			GlueDataCatalogConfiguration: &GlueDataCatalogConfigurationProperty{
//   				DatabaseArn: jsii.String("databaseArn"),
//   			},
//   		},
//   		CustomArtifactsConfiguration: []interface{}{
//   			&CustomArtifactConfigurationProperty{
//   				ArtifactType: jsii.String("artifactType"),
//
//   				// the properties below are optional
//   				MavenReference: &MavenReferenceProperty{
//   					ArtifactId: jsii.String("artifactId"),
//   					GroupId: jsii.String("groupId"),
//   					Version: jsii.String("version"),
//   				},
//   				S3ContentLocation: &S3ContentLocationProperty{
//   					BucketArn: jsii.String("bucketArn"),
//   					FileKey: jsii.String("fileKey"),
//
//   					// the properties below are optional
//   					ObjectVersion: jsii.String("objectVersion"),
//   				},
//   			},
//   		},
//   		DeployAsApplicationConfiguration: &DeployAsApplicationConfigurationProperty{
//   			S3ContentLocation: &S3ContentBaseLocationProperty{
//   				BucketArn: jsii.String("bucketArn"),
//
//   				// the properties below are optional
//   				BasePath: jsii.String("basePath"),
//   			},
//   		},
//   		MonitoringConfiguration: &ZeppelinMonitoringConfigurationProperty{
//   			LogLevel: jsii.String("logLevel"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationconfiguration.html
//
type CfnApplication_ApplicationConfigurationProperty struct {
	// The code location and type parameters for a Managed Service for Apache Flink application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationconfiguration.html#cfn-kinesisanalyticsv2-application-applicationconfiguration-applicationcodeconfiguration
	//
	ApplicationCodeConfiguration interface{} `field:"optional" json:"applicationCodeConfiguration" yaml:"applicationCodeConfiguration"`
	// Describes whether snapshots are enabled for a Managed Service for Apache Flink application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationconfiguration.html#cfn-kinesisanalyticsv2-application-applicationconfiguration-applicationsnapshotconfiguration
	//
	ApplicationSnapshotConfiguration interface{} `field:"optional" json:"applicationSnapshotConfiguration" yaml:"applicationSnapshotConfiguration"`
	// Describes execution properties for a Managed Service for Apache Flink application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationconfiguration.html#cfn-kinesisanalyticsv2-application-applicationconfiguration-environmentproperties
	//
	EnvironmentProperties interface{} `field:"optional" json:"environmentProperties" yaml:"environmentProperties"`
	// The creation and update parameters for a Managed Service for Apache Flink application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationconfiguration.html#cfn-kinesisanalyticsv2-application-applicationconfiguration-flinkapplicationconfiguration
	//
	FlinkApplicationConfiguration interface{} `field:"optional" json:"flinkApplicationConfiguration" yaml:"flinkApplicationConfiguration"`
	// The creation and update parameters for a SQL-based Managed Service for Apache Flink application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationconfiguration.html#cfn-kinesisanalyticsv2-application-applicationconfiguration-sqlapplicationconfiguration
	//
	SqlApplicationConfiguration interface{} `field:"optional" json:"sqlApplicationConfiguration" yaml:"sqlApplicationConfiguration"`
	// The array of descriptions of VPC configurations available to the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationconfiguration.html#cfn-kinesisanalyticsv2-application-applicationconfiguration-vpcconfigurations
	//
	VpcConfigurations interface{} `field:"optional" json:"vpcConfigurations" yaml:"vpcConfigurations"`
	// The configuration parameters for a Kinesis Data Analytics Studio notebook.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationconfiguration.html#cfn-kinesisanalyticsv2-application-applicationconfiguration-zeppelinapplicationconfiguration
	//
	ZeppelinApplicationConfiguration interface{} `field:"optional" json:"zeppelinApplicationConfiguration" yaml:"zeppelinApplicationConfiguration"`
}

