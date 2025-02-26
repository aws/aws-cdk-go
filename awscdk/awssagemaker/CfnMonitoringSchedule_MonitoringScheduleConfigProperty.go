package awssagemaker


// Configures the monitoring schedule and defines the monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringScheduleConfigProperty := &MonitoringScheduleConfigProperty{
//   	MonitoringJobDefinition: &MonitoringJobDefinitionProperty{
//   		MonitoringAppSpecification: &MonitoringAppSpecificationProperty{
//   			ImageUri: jsii.String("imageUri"),
//
//   			// the properties below are optional
//   			ContainerArguments: []*string{
//   				jsii.String("containerArguments"),
//   			},
//   			ContainerEntrypoint: []*string{
//   				jsii.String("containerEntrypoint"),
//   			},
//   			PostAnalyticsProcessorSourceUri: jsii.String("postAnalyticsProcessorSourceUri"),
//   			RecordPreprocessorSourceUri: jsii.String("recordPreprocessorSourceUri"),
//   		},
//   		MonitoringInputs: []interface{}{
//   			&MonitoringInputProperty{
//   				BatchTransformInput: &BatchTransformInputProperty{
//   					DataCapturedDestinationS3Uri: jsii.String("dataCapturedDestinationS3Uri"),
//   					DatasetFormat: &DatasetFormatProperty{
//   						Csv: &CsvProperty{
//   							Header: jsii.Boolean(false),
//   						},
//   						Json: &JsonProperty{
//   							Line: jsii.Boolean(false),
//   						},
//   						Parquet: jsii.Boolean(false),
//   					},
//   					LocalPath: jsii.String("localPath"),
//
//   					// the properties below are optional
//   					ExcludeFeaturesAttribute: jsii.String("excludeFeaturesAttribute"),
//   					S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   					S3InputMode: jsii.String("s3InputMode"),
//   				},
//   				EndpointInput: &EndpointInputProperty{
//   					EndpointName: jsii.String("endpointName"),
//   					LocalPath: jsii.String("localPath"),
//
//   					// the properties below are optional
//   					ExcludeFeaturesAttribute: jsii.String("excludeFeaturesAttribute"),
//   					S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   					S3InputMode: jsii.String("s3InputMode"),
//   				},
//   			},
//   		},
//   		MonitoringOutputConfig: &MonitoringOutputConfigProperty{
//   			MonitoringOutputs: []interface{}{
//   				&MonitoringOutputProperty{
//   					S3Output: &S3OutputProperty{
//   						LocalPath: jsii.String("localPath"),
//   						S3Uri: jsii.String("s3Uri"),
//
//   						// the properties below are optional
//   						S3UploadMode: jsii.String("s3UploadMode"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			KmsKeyId: jsii.String("kmsKeyId"),
//   		},
//   		MonitoringResources: &MonitoringResourcesProperty{
//   			ClusterConfig: &ClusterConfigProperty{
//   				InstanceCount: jsii.Number(123),
//   				InstanceType: jsii.String("instanceType"),
//   				VolumeSizeInGb: jsii.Number(123),
//
//   				// the properties below are optional
//   				VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   			},
//   		},
//   		RoleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		BaselineConfig: &BaselineConfigProperty{
//   			ConstraintsResource: &ConstraintsResourceProperty{
//   				S3Uri: jsii.String("s3Uri"),
//   			},
//   			StatisticsResource: &StatisticsResourceProperty{
//   				S3Uri: jsii.String("s3Uri"),
//   			},
//   		},
//   		Environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		NetworkConfig: &NetworkConfigProperty{
//   			EnableInterContainerTrafficEncryption: jsii.Boolean(false),
//   			EnableNetworkIsolation: jsii.Boolean(false),
//   			VpcConfig: &VpcConfigProperty{
//   				SecurityGroupIds: []*string{
//   					jsii.String("securityGroupIds"),
//   				},
//   				Subnets: []*string{
//   					jsii.String("subnets"),
//   				},
//   			},
//   		},
//   		StoppingCondition: &StoppingConditionProperty{
//   			MaxRuntimeInSeconds: jsii.Number(123),
//   		},
//   	},
//   	MonitoringJobDefinitionName: jsii.String("monitoringJobDefinitionName"),
//   	MonitoringType: jsii.String("monitoringType"),
//   	ScheduleConfig: &ScheduleConfigProperty{
//   		ScheduleExpression: jsii.String("scheduleExpression"),
//
//   		// the properties below are optional
//   		DataAnalysisEndTime: jsii.String("dataAnalysisEndTime"),
//   		DataAnalysisStartTime: jsii.String("dataAnalysisStartTime"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringscheduleconfig.html
//
type CfnMonitoringSchedule_MonitoringScheduleConfigProperty struct {
	// Defines the monitoring job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringscheduleconfig.html#cfn-sagemaker-monitoringschedule-monitoringscheduleconfig-monitoringjobdefinition
	//
	MonitoringJobDefinition interface{} `field:"optional" json:"monitoringJobDefinition" yaml:"monitoringJobDefinition"`
	// The name of the monitoring job definition to schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringscheduleconfig.html#cfn-sagemaker-monitoringschedule-monitoringscheduleconfig-monitoringjobdefinitionname
	//
	MonitoringJobDefinitionName *string `field:"optional" json:"monitoringJobDefinitionName" yaml:"monitoringJobDefinitionName"`
	// The type of the monitoring job definition to schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringscheduleconfig.html#cfn-sagemaker-monitoringschedule-monitoringscheduleconfig-monitoringtype
	//
	MonitoringType *string `field:"optional" json:"monitoringType" yaml:"monitoringType"`
	// Configures the monitoring schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringscheduleconfig.html#cfn-sagemaker-monitoringschedule-monitoringscheduleconfig-scheduleconfig
	//
	ScheduleConfig interface{} `field:"optional" json:"scheduleConfig" yaml:"scheduleConfig"`
}

