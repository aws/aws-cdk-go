package awskinesisanalyticsv2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationProps := &cfnApplicationProps{
//   	runtimeEnvironment: jsii.String("runtimeEnvironment"),
//   	serviceExecutionRole: jsii.String("serviceExecutionRole"),
//
//   	// the properties below are optional
//   	applicationConfiguration: &applicationConfigurationProperty{
//   		applicationCodeConfiguration: &applicationCodeConfigurationProperty{
//   			codeContent: &codeContentProperty{
//   				s3ContentLocation: &s3ContentLocationProperty{
//   					bucketArn: jsii.String("bucketArn"),
//   					fileKey: jsii.String("fileKey"),
//
//   					// the properties below are optional
//   					objectVersion: jsii.String("objectVersion"),
//   				},
//   				textContent: jsii.String("textContent"),
//   				zipFileContent: jsii.String("zipFileContent"),
//   			},
//   			codeContentType: jsii.String("codeContentType"),
//   		},
//   		applicationSnapshotConfiguration: &applicationSnapshotConfigurationProperty{
//   			snapshotsEnabled: jsii.Boolean(false),
//   		},
//   		environmentProperties: &environmentPropertiesProperty{
//   			propertyGroups: []interface{}{
//   				&propertyGroupProperty{
//   					propertyGroupId: jsii.String("propertyGroupId"),
//   					propertyMap: map[string]*string{
//   						"propertyMapKey": jsii.String("propertyMap"),
//   					},
//   				},
//   			},
//   		},
//   		flinkApplicationConfiguration: &flinkApplicationConfigurationProperty{
//   			checkpointConfiguration: &checkpointConfigurationProperty{
//   				configurationType: jsii.String("configurationType"),
//
//   				// the properties below are optional
//   				checkpointingEnabled: jsii.Boolean(false),
//   				checkpointInterval: jsii.Number(123),
//   				minPauseBetweenCheckpoints: jsii.Number(123),
//   			},
//   			monitoringConfiguration: &monitoringConfigurationProperty{
//   				configurationType: jsii.String("configurationType"),
//
//   				// the properties below are optional
//   				logLevel: jsii.String("logLevel"),
//   				metricsLevel: jsii.String("metricsLevel"),
//   			},
//   			parallelismConfiguration: &parallelismConfigurationProperty{
//   				configurationType: jsii.String("configurationType"),
//
//   				// the properties below are optional
//   				autoScalingEnabled: jsii.Boolean(false),
//   				parallelism: jsii.Number(123),
//   				parallelismPerKpu: jsii.Number(123),
//   			},
//   		},
//   		sqlApplicationConfiguration: &sqlApplicationConfigurationProperty{
//   			inputs: []interface{}{
//   				&inputProperty{
//   					inputSchema: &inputSchemaProperty{
//   						recordColumns: []interface{}{
//   							&recordColumnProperty{
//   								name: jsii.String("name"),
//   								sqlType: jsii.String("sqlType"),
//
//   								// the properties below are optional
//   								mapping: jsii.String("mapping"),
//   							},
//   						},
//   						recordFormat: &recordFormatProperty{
//   							recordFormatType: jsii.String("recordFormatType"),
//
//   							// the properties below are optional
//   							mappingParameters: &mappingParametersProperty{
//   								csvMappingParameters: &cSVMappingParametersProperty{
//   									recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   									recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   								},
//   								jsonMappingParameters: &jSONMappingParametersProperty{
//   									recordRowPath: jsii.String("recordRowPath"),
//   								},
//   							},
//   						},
//
//   						// the properties below are optional
//   						recordEncoding: jsii.String("recordEncoding"),
//   					},
//   					namePrefix: jsii.String("namePrefix"),
//
//   					// the properties below are optional
//   					inputParallelism: &inputParallelismProperty{
//   						count: jsii.Number(123),
//   					},
//   					inputProcessingConfiguration: &inputProcessingConfigurationProperty{
//   						inputLambdaProcessor: &inputLambdaProcessorProperty{
//   							resourceArn: jsii.String("resourceArn"),
//   						},
//   					},
//   					kinesisFirehoseInput: &kinesisFirehoseInputProperty{
//   						resourceArn: jsii.String("resourceArn"),
//   					},
//   					kinesisStreamsInput: &kinesisStreamsInputProperty{
//   						resourceArn: jsii.String("resourceArn"),
//   					},
//   				},
//   			},
//   		},
//   		vpcConfigurations: []interface{}{
//   			&vpcConfigurationProperty{
//   				securityGroupIds: []*string{
//   					jsii.String("securityGroupIds"),
//   				},
//   				subnetIds: []*string{
//   					jsii.String("subnetIds"),
//   				},
//   			},
//   		},
//   		zeppelinApplicationConfiguration: &zeppelinApplicationConfigurationProperty{
//   			catalogConfiguration: &catalogConfigurationProperty{
//   				glueDataCatalogConfiguration: &glueDataCatalogConfigurationProperty{
//   					databaseArn: jsii.String("databaseArn"),
//   				},
//   			},
//   			customArtifactsConfiguration: []interface{}{
//   				&customArtifactConfigurationProperty{
//   					artifactType: jsii.String("artifactType"),
//
//   					// the properties below are optional
//   					mavenReference: &mavenReferenceProperty{
//   						artifactId: jsii.String("artifactId"),
//   						groupId: jsii.String("groupId"),
//   						version: jsii.String("version"),
//   					},
//   					s3ContentLocation: &s3ContentLocationProperty{
//   						bucketArn: jsii.String("bucketArn"),
//   						fileKey: jsii.String("fileKey"),
//
//   						// the properties below are optional
//   						objectVersion: jsii.String("objectVersion"),
//   					},
//   				},
//   			},
//   			deployAsApplicationConfiguration: &deployAsApplicationConfigurationProperty{
//   				s3ContentLocation: &s3ContentBaseLocationProperty{
//   					bucketArn: jsii.String("bucketArn"),
//
//   					// the properties below are optional
//   					basePath: jsii.String("basePath"),
//   				},
//   			},
//   			monitoringConfiguration: &zeppelinMonitoringConfigurationProperty{
//   				logLevel: jsii.String("logLevel"),
//   			},
//   		},
//   	},
//   	applicationDescription: jsii.String("applicationDescription"),
//   	applicationMaintenanceConfiguration: &applicationMaintenanceConfigurationProperty{
//   		applicationMaintenanceWindowStartTime: jsii.String("applicationMaintenanceWindowStartTime"),
//   	},
//   	applicationMode: jsii.String("applicationMode"),
//   	applicationName: jsii.String("applicationName"),
//   	runConfiguration: &runConfigurationProperty{
//   		applicationRestoreConfiguration: &applicationRestoreConfigurationProperty{
//   			applicationRestoreType: jsii.String("applicationRestoreType"),
//
//   			// the properties below are optional
//   			snapshotName: jsii.String("snapshotName"),
//   		},
//   		flinkRunConfiguration: &flinkRunConfigurationProperty{
//   			allowNonRestoredState: jsii.Boolean(false),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnApplicationProps struct {
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

