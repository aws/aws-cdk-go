package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnMonitoringSchedule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var json interface{}
//
//   cfnMonitoringScheduleProps := &cfnMonitoringScheduleProps{
//   	monitoringScheduleConfig: &monitoringScheduleConfigProperty{
//   		monitoringJobDefinition: &monitoringJobDefinitionProperty{
//   			monitoringAppSpecification: &monitoringAppSpecificationProperty{
//   				imageUri: jsii.String("imageUri"),
//
//   				// the properties below are optional
//   				containerArguments: []*string{
//   					jsii.String("containerArguments"),
//   				},
//   				containerEntrypoint: []*string{
//   					jsii.String("containerEntrypoint"),
//   				},
//   				postAnalyticsProcessorSourceUri: jsii.String("postAnalyticsProcessorSourceUri"),
//   				recordPreprocessorSourceUri: jsii.String("recordPreprocessorSourceUri"),
//   			},
//   			monitoringInputs: []interface{}{
//   				&monitoringInputProperty{
//   					batchTransformInput: &batchTransformInputProperty{
//   						dataCapturedDestinationS3Uri: jsii.String("dataCapturedDestinationS3Uri"),
//   						datasetFormat: &datasetFormatProperty{
//   							csv: &csvProperty{
//   								header: jsii.Boolean(false),
//   							},
//   							json: json,
//   							parquet: jsii.Boolean(false),
//   						},
//   						localPath: jsii.String("localPath"),
//
//   						// the properties below are optional
//   						s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   						s3InputMode: jsii.String("s3InputMode"),
//   					},
//   					endpointInput: &endpointInputProperty{
//   						endpointName: jsii.String("endpointName"),
//   						localPath: jsii.String("localPath"),
//
//   						// the properties below are optional
//   						s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   						s3InputMode: jsii.String("s3InputMode"),
//   					},
//   				},
//   			},
//   			monitoringOutputConfig: &monitoringOutputConfigProperty{
//   				monitoringOutputs: []interface{}{
//   					&monitoringOutputProperty{
//   						s3Output: &s3OutputProperty{
//   							localPath: jsii.String("localPath"),
//   							s3Uri: jsii.String("s3Uri"),
//
//   							// the properties below are optional
//   							s3UploadMode: jsii.String("s3UploadMode"),
//   						},
//   					},
//   				},
//
//   				// the properties below are optional
//   				kmsKeyId: jsii.String("kmsKeyId"),
//   			},
//   			monitoringResources: &monitoringResourcesProperty{
//   				clusterConfig: &clusterConfigProperty{
//   					instanceCount: jsii.Number(123),
//   					instanceType: jsii.String("instanceType"),
//   					volumeSizeInGb: jsii.Number(123),
//
//   					// the properties below are optional
//   					volumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   				},
//   			},
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			baselineConfig: &baselineConfigProperty{
//   				constraintsResource: &constraintsResourceProperty{
//   					s3Uri: jsii.String("s3Uri"),
//   				},
//   				statisticsResource: &statisticsResourceProperty{
//   					s3Uri: jsii.String("s3Uri"),
//   				},
//   			},
//   			environment: map[string]*string{
//   				"environmentKey": jsii.String("environment"),
//   			},
//   			networkConfig: &networkConfigProperty{
//   				enableInterContainerTrafficEncryption: jsii.Boolean(false),
//   				enableNetworkIsolation: jsii.Boolean(false),
//   				vpcConfig: &vpcConfigProperty{
//   					securityGroupIds: []*string{
//   						jsii.String("securityGroupIds"),
//   					},
//   					subnets: []*string{
//   						jsii.String("subnets"),
//   					},
//   				},
//   			},
//   			stoppingCondition: &stoppingConditionProperty{
//   				maxRuntimeInSeconds: jsii.Number(123),
//   			},
//   		},
//   		monitoringJobDefinitionName: jsii.String("monitoringJobDefinitionName"),
//   		monitoringType: jsii.String("monitoringType"),
//   		scheduleConfig: &scheduleConfigProperty{
//   			scheduleExpression: jsii.String("scheduleExpression"),
//   		},
//   	},
//   	monitoringScheduleName: jsii.String("monitoringScheduleName"),
//
//   	// the properties below are optional
//   	endpointName: jsii.String("endpointName"),
//   	failureReason: jsii.String("failureReason"),
//   	lastMonitoringExecutionSummary: &monitoringExecutionSummaryProperty{
//   		creationTime: jsii.String("creationTime"),
//   		lastModifiedTime: jsii.String("lastModifiedTime"),
//   		monitoringExecutionStatus: jsii.String("monitoringExecutionStatus"),
//   		monitoringScheduleName: jsii.String("monitoringScheduleName"),
//   		scheduledTime: jsii.String("scheduledTime"),
//
//   		// the properties below are optional
//   		endpointName: jsii.String("endpointName"),
//   		failureReason: jsii.String("failureReason"),
//   		processingJobArn: jsii.String("processingJobArn"),
//   	},
//   	monitoringScheduleStatus: jsii.String("monitoringScheduleStatus"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnMonitoringScheduleProps struct {
	// The configuration object that specifies the monitoring schedule and defines the monitoring job.
	MonitoringScheduleConfig interface{} `field:"required" json:"monitoringScheduleConfig" yaml:"monitoringScheduleConfig"`
	// The name of the monitoring schedule.
	MonitoringScheduleName *string `field:"required" json:"monitoringScheduleName" yaml:"monitoringScheduleName"`
	// The name of the endpoint using the monitoring schedule.
	EndpointName *string `field:"optional" json:"endpointName" yaml:"endpointName"`
	// Contains the reason a monitoring job failed, if it failed.
	FailureReason *string `field:"optional" json:"failureReason" yaml:"failureReason"`
	// Describes metadata on the last execution to run, if there was one.
	LastMonitoringExecutionSummary interface{} `field:"optional" json:"lastMonitoringExecutionSummary" yaml:"lastMonitoringExecutionSummary"`
	// The status of the monitoring schedule.
	MonitoringScheduleStatus *string `field:"optional" json:"monitoringScheduleStatus" yaml:"monitoringScheduleStatus"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

