package awskinesisanalyticsv2


// Specifies the creation parameters for a Kinesis Data Analytics application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationConfigurationProperty := &applicationConfigurationProperty{
//   	applicationCodeConfiguration: &applicationCodeConfigurationProperty{
//   		codeContent: &codeContentProperty{
//   			s3ContentLocation: &s3ContentLocationProperty{
//   				bucketArn: jsii.String("bucketArn"),
//   				fileKey: jsii.String("fileKey"),
//
//   				// the properties below are optional
//   				objectVersion: jsii.String("objectVersion"),
//   			},
//   			textContent: jsii.String("textContent"),
//   			zipFileContent: jsii.String("zipFileContent"),
//   		},
//   		codeContentType: jsii.String("codeContentType"),
//   	},
//   	applicationSnapshotConfiguration: &applicationSnapshotConfigurationProperty{
//   		snapshotsEnabled: jsii.Boolean(false),
//   	},
//   	environmentProperties: &environmentPropertiesProperty{
//   		propertyGroups: []interface{}{
//   			&propertyGroupProperty{
//   				propertyGroupId: jsii.String("propertyGroupId"),
//   				propertyMap: map[string]*string{
//   					"propertyMapKey": jsii.String("propertyMap"),
//   				},
//   			},
//   		},
//   	},
//   	flinkApplicationConfiguration: &flinkApplicationConfigurationProperty{
//   		checkpointConfiguration: &checkpointConfigurationProperty{
//   			configurationType: jsii.String("configurationType"),
//
//   			// the properties below are optional
//   			checkpointingEnabled: jsii.Boolean(false),
//   			checkpointInterval: jsii.Number(123),
//   			minPauseBetweenCheckpoints: jsii.Number(123),
//   		},
//   		monitoringConfiguration: &monitoringConfigurationProperty{
//   			configurationType: jsii.String("configurationType"),
//
//   			// the properties below are optional
//   			logLevel: jsii.String("logLevel"),
//   			metricsLevel: jsii.String("metricsLevel"),
//   		},
//   		parallelismConfiguration: &parallelismConfigurationProperty{
//   			configurationType: jsii.String("configurationType"),
//
//   			// the properties below are optional
//   			autoScalingEnabled: jsii.Boolean(false),
//   			parallelism: jsii.Number(123),
//   			parallelismPerKpu: jsii.Number(123),
//   		},
//   	},
//   	sqlApplicationConfiguration: &sqlApplicationConfigurationProperty{
//   		inputs: []interface{}{
//   			&inputProperty{
//   				inputSchema: &inputSchemaProperty{
//   					recordColumns: []interface{}{
//   						&recordColumnProperty{
//   							name: jsii.String("name"),
//   							sqlType: jsii.String("sqlType"),
//
//   							// the properties below are optional
//   							mapping: jsii.String("mapping"),
//   						},
//   					},
//   					recordFormat: &recordFormatProperty{
//   						recordFormatType: jsii.String("recordFormatType"),
//
//   						// the properties below are optional
//   						mappingParameters: &mappingParametersProperty{
//   							csvMappingParameters: &cSVMappingParametersProperty{
//   								recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   								recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   							},
//   							jsonMappingParameters: &jSONMappingParametersProperty{
//   								recordRowPath: jsii.String("recordRowPath"),
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					recordEncoding: jsii.String("recordEncoding"),
//   				},
//   				namePrefix: jsii.String("namePrefix"),
//
//   				// the properties below are optional
//   				inputParallelism: &inputParallelismProperty{
//   					count: jsii.Number(123),
//   				},
//   				inputProcessingConfiguration: &inputProcessingConfigurationProperty{
//   					inputLambdaProcessor: &inputLambdaProcessorProperty{
//   						resourceArn: jsii.String("resourceArn"),
//   					},
//   				},
//   				kinesisFirehoseInput: &kinesisFirehoseInputProperty{
//   					resourceArn: jsii.String("resourceArn"),
//   				},
//   				kinesisStreamsInput: &kinesisStreamsInputProperty{
//   					resourceArn: jsii.String("resourceArn"),
//   				},
//   			},
//   		},
//   	},
//   	vpcConfigurations: []interface{}{
//   		&vpcConfigurationProperty{
//   			securityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			subnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//   		},
//   	},
//   	zeppelinApplicationConfiguration: &zeppelinApplicationConfigurationProperty{
//   		catalogConfiguration: &catalogConfigurationProperty{
//   			glueDataCatalogConfiguration: &glueDataCatalogConfigurationProperty{
//   				databaseArn: jsii.String("databaseArn"),
//   			},
//   		},
//   		customArtifactsConfiguration: []interface{}{
//   			&customArtifactConfigurationProperty{
//   				artifactType: jsii.String("artifactType"),
//
//   				// the properties below are optional
//   				mavenReference: &mavenReferenceProperty{
//   					artifactId: jsii.String("artifactId"),
//   					groupId: jsii.String("groupId"),
//   					version: jsii.String("version"),
//   				},
//   				s3ContentLocation: &s3ContentLocationProperty{
//   					bucketArn: jsii.String("bucketArn"),
//   					fileKey: jsii.String("fileKey"),
//
//   					// the properties below are optional
//   					objectVersion: jsii.String("objectVersion"),
//   				},
//   			},
//   		},
//   		deployAsApplicationConfiguration: &deployAsApplicationConfigurationProperty{
//   			s3ContentLocation: &s3ContentBaseLocationProperty{
//   				bucketArn: jsii.String("bucketArn"),
//
//   				// the properties below are optional
//   				basePath: jsii.String("basePath"),
//   			},
//   		},
//   		monitoringConfiguration: &zeppelinMonitoringConfigurationProperty{
//   			logLevel: jsii.String("logLevel"),
//   		},
//   	},
//   }
//
type CfnApplication_ApplicationConfigurationProperty struct {
	// The code location and type parameters for a Flink-based Kinesis Data Analytics application.
	ApplicationCodeConfiguration interface{} `field:"optional" json:"applicationCodeConfiguration" yaml:"applicationCodeConfiguration"`
	// Describes whether snapshots are enabled for a Flink-based Kinesis Data Analytics application.
	ApplicationSnapshotConfiguration interface{} `field:"optional" json:"applicationSnapshotConfiguration" yaml:"applicationSnapshotConfiguration"`
	// Describes execution properties for a Flink-based Kinesis Data Analytics application.
	EnvironmentProperties interface{} `field:"optional" json:"environmentProperties" yaml:"environmentProperties"`
	// The creation and update parameters for a Flink-based Kinesis Data Analytics application.
	FlinkApplicationConfiguration interface{} `field:"optional" json:"flinkApplicationConfiguration" yaml:"flinkApplicationConfiguration"`
	// The creation and update parameters for a SQL-based Kinesis Data Analytics application.
	SqlApplicationConfiguration interface{} `field:"optional" json:"sqlApplicationConfiguration" yaml:"sqlApplicationConfiguration"`
	// `CfnApplication.ApplicationConfigurationProperty.VpcConfigurations`.
	VpcConfigurations interface{} `field:"optional" json:"vpcConfigurations" yaml:"vpcConfigurations"`
	// The configuration parameters for a Kinesis Data Analytics Studio notebook.
	ZeppelinApplicationConfiguration interface{} `field:"optional" json:"zeppelinApplicationConfiguration" yaml:"zeppelinApplicationConfiguration"`
}

