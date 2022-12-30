package awskinesisanalyticsflink

import (
	"github.com/aws/aws-cdk-go/awscdk/awskinesisanalytics"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

// The return type of {@link ApplicationCode.bind}. This represents CloudFormation configuration and an s3 bucket holding the Flink application JAR file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   applicationCodeConfig := &applicationCodeConfig{
//   	applicationCodeConfigurationProperty: &applicationConfigurationProperty{
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
//   	bucket: bucket,
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

