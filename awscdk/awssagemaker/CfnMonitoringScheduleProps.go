package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnMonitoringSchedule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMonitoringScheduleProps := &CfnMonitoringScheduleProps{
//   	MonitoringScheduleConfig: &MonitoringScheduleConfigProperty{
//   		MonitoringJobDefinition: &MonitoringJobDefinitionProperty{
//   			MonitoringAppSpecification: &MonitoringAppSpecificationProperty{
//   				ImageUri: jsii.String("imageUri"),
//
//   				// the properties below are optional
//   				ContainerArguments: []*string{
//   					jsii.String("containerArguments"),
//   				},
//   				ContainerEntrypoint: []*string{
//   					jsii.String("containerEntrypoint"),
//   				},
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
//   						LocalPath: jsii.String("localPath"),
//
//   						// the properties below are optional
//   						ExcludeFeaturesAttribute: jsii.String("excludeFeaturesAttribute"),
//   						S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   						S3InputMode: jsii.String("s3InputMode"),
//   					},
//   					EndpointInput: &EndpointInputProperty{
//   						EndpointName: jsii.String("endpointName"),
//   						LocalPath: jsii.String("localPath"),
//
//   						// the properties below are optional
//   						ExcludeFeaturesAttribute: jsii.String("excludeFeaturesAttribute"),
//   						S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   						S3InputMode: jsii.String("s3InputMode"),
//   					},
//   				},
//   			},
//   			MonitoringOutputConfig: &MonitoringOutputConfigProperty{
//   				MonitoringOutputs: []interface{}{
//   					&MonitoringOutputProperty{
//   						S3Output: &S3OutputProperty{
//   							LocalPath: jsii.String("localPath"),
//   							S3Uri: jsii.String("s3Uri"),
//
//   							// the properties below are optional
//   							S3UploadMode: jsii.String("s3UploadMode"),
//   						},
//   					},
//   				},
//
//   				// the properties below are optional
//   				KmsKeyId: jsii.String("kmsKeyId"),
//   			},
//   			MonitoringResources: &MonitoringResourcesProperty{
//   				ClusterConfig: &ClusterConfigProperty{
//   					InstanceCount: jsii.Number(123),
//   					InstanceType: jsii.String("instanceType"),
//   					VolumeSizeInGb: jsii.Number(123),
//
//   					// the properties below are optional
//   					VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   				},
//   			},
//   			RoleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
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
//   			StoppingCondition: &StoppingConditionProperty{
//   				MaxRuntimeInSeconds: jsii.Number(123),
//   			},
//   		},
//   		MonitoringJobDefinitionName: jsii.String("monitoringJobDefinitionName"),
//   		MonitoringType: jsii.String("monitoringType"),
//   		ScheduleConfig: &ScheduleConfigProperty{
//   			ScheduleExpression: jsii.String("scheduleExpression"),
//
//   			// the properties below are optional
//   			DataAnalysisEndTime: jsii.String("dataAnalysisEndTime"),
//   			DataAnalysisStartTime: jsii.String("dataAnalysisStartTime"),
//   		},
//   	},
//   	MonitoringScheduleName: jsii.String("monitoringScheduleName"),
//
//   	// the properties below are optional
//   	EndpointName: jsii.String("endpointName"),
//   	FailureReason: jsii.String("failureReason"),
//   	LastMonitoringExecutionSummary: &MonitoringExecutionSummaryProperty{
//   		CreationTime: jsii.String("creationTime"),
//   		LastModifiedTime: jsii.String("lastModifiedTime"),
//   		MonitoringExecutionStatus: jsii.String("monitoringExecutionStatus"),
//   		MonitoringScheduleName: jsii.String("monitoringScheduleName"),
//   		ScheduledTime: jsii.String("scheduledTime"),
//
//   		// the properties below are optional
//   		EndpointName: jsii.String("endpointName"),
//   		FailureReason: jsii.String("failureReason"),
//   		ProcessingJobArn: jsii.String("processingJobArn"),
//   	},
//   	MonitoringScheduleStatus: jsii.String("monitoringScheduleStatus"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-monitoringschedule.html
//
type CfnMonitoringScheduleProps struct {
	// The configuration object that specifies the monitoring schedule and defines the monitoring job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-monitoringschedule.html#cfn-sagemaker-monitoringschedule-monitoringscheduleconfig
	//
	MonitoringScheduleConfig interface{} `field:"required" json:"monitoringScheduleConfig" yaml:"monitoringScheduleConfig"`
	// The name of the monitoring schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-monitoringschedule.html#cfn-sagemaker-monitoringschedule-monitoringschedulename
	//
	MonitoringScheduleName *string `field:"required" json:"monitoringScheduleName" yaml:"monitoringScheduleName"`
	// The name of the endpoint using the monitoring schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-monitoringschedule.html#cfn-sagemaker-monitoringschedule-endpointname
	//
	EndpointName *string `field:"optional" json:"endpointName" yaml:"endpointName"`
	// Contains the reason a monitoring job failed, if it failed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-monitoringschedule.html#cfn-sagemaker-monitoringschedule-failurereason
	//
	FailureReason *string `field:"optional" json:"failureReason" yaml:"failureReason"`
	// Describes metadata on the last execution to run, if there was one.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-monitoringschedule.html#cfn-sagemaker-monitoringschedule-lastmonitoringexecutionsummary
	//
	LastMonitoringExecutionSummary interface{} `field:"optional" json:"lastMonitoringExecutionSummary" yaml:"lastMonitoringExecutionSummary"`
	// The status of the monitoring schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-monitoringschedule.html#cfn-sagemaker-monitoringschedule-monitoringschedulestatus
	//
	MonitoringScheduleStatus *string `field:"optional" json:"monitoringScheduleStatus" yaml:"monitoringScheduleStatus"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-monitoringschedule.html#cfn-sagemaker-monitoringschedule-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

