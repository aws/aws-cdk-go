package awssagemaker


// Configures the monitoring schedule and defines the monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var json interface{}
//
//   monitoringScheduleConfigProperty := &monitoringScheduleConfigProperty{
//   	monitoringJobDefinition: &monitoringJobDefinitionProperty{
//   		monitoringAppSpecification: &monitoringAppSpecificationProperty{
//   			imageUri: jsii.String("imageUri"),
//
//   			// the properties below are optional
//   			containerArguments: []*string{
//   				jsii.String("containerArguments"),
//   			},
//   			containerEntrypoint: []*string{
//   				jsii.String("containerEntrypoint"),
//   			},
//   			postAnalyticsProcessorSourceUri: jsii.String("postAnalyticsProcessorSourceUri"),
//   			recordPreprocessorSourceUri: jsii.String("recordPreprocessorSourceUri"),
//   		},
//   		monitoringInputs: []interface{}{
//   			&monitoringInputProperty{
//   				batchTransformInput: &batchTransformInputProperty{
//   					dataCapturedDestinationS3Uri: jsii.String("dataCapturedDestinationS3Uri"),
//   					datasetFormat: &datasetFormatProperty{
//   						csv: &csvProperty{
//   							header: jsii.Boolean(false),
//   						},
//   						json: json,
//   						parquet: jsii.Boolean(false),
//   					},
//   					localPath: jsii.String("localPath"),
//
//   					// the properties below are optional
//   					s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   					s3InputMode: jsii.String("s3InputMode"),
//   				},
//   				endpointInput: &endpointInputProperty{
//   					endpointName: jsii.String("endpointName"),
//   					localPath: jsii.String("localPath"),
//
//   					// the properties below are optional
//   					s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   					s3InputMode: jsii.String("s3InputMode"),
//   				},
//   			},
//   		},
//   		monitoringOutputConfig: &monitoringOutputConfigProperty{
//   			monitoringOutputs: []interface{}{
//   				&monitoringOutputProperty{
//   					s3Output: &s3OutputProperty{
//   						localPath: jsii.String("localPath"),
//   						s3Uri: jsii.String("s3Uri"),
//
//   						// the properties below are optional
//   						s3UploadMode: jsii.String("s3UploadMode"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			kmsKeyId: jsii.String("kmsKeyId"),
//   		},
//   		monitoringResources: &monitoringResourcesProperty{
//   			clusterConfig: &clusterConfigProperty{
//   				instanceCount: jsii.Number(123),
//   				instanceType: jsii.String("instanceType"),
//   				volumeSizeInGb: jsii.Number(123),
//
//   				// the properties below are optional
//   				volumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   			},
//   		},
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		baselineConfig: &baselineConfigProperty{
//   			constraintsResource: &constraintsResourceProperty{
//   				s3Uri: jsii.String("s3Uri"),
//   			},
//   			statisticsResource: &statisticsResourceProperty{
//   				s3Uri: jsii.String("s3Uri"),
//   			},
//   		},
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		networkConfig: &networkConfigProperty{
//   			enableInterContainerTrafficEncryption: jsii.Boolean(false),
//   			enableNetworkIsolation: jsii.Boolean(false),
//   			vpcConfig: &vpcConfigProperty{
//   				securityGroupIds: []*string{
//   					jsii.String("securityGroupIds"),
//   				},
//   				subnets: []*string{
//   					jsii.String("subnets"),
//   				},
//   			},
//   		},
//   		stoppingCondition: &stoppingConditionProperty{
//   			maxRuntimeInSeconds: jsii.Number(123),
//   		},
//   	},
//   	monitoringJobDefinitionName: jsii.String("monitoringJobDefinitionName"),
//   	monitoringType: jsii.String("monitoringType"),
//   	scheduleConfig: &scheduleConfigProperty{
//   		scheduleExpression: jsii.String("scheduleExpression"),
//   	},
//   }
//
type CfnMonitoringSchedule_MonitoringScheduleConfigProperty struct {
	// Defines the monitoring job.
	MonitoringJobDefinition interface{} `field:"optional" json:"monitoringJobDefinition" yaml:"monitoringJobDefinition"`
	// The name of the monitoring job definition to schedule.
	MonitoringJobDefinitionName *string `field:"optional" json:"monitoringJobDefinitionName" yaml:"monitoringJobDefinitionName"`
	// The type of the monitoring job definition to schedule.
	MonitoringType *string `field:"optional" json:"monitoringType" yaml:"monitoringType"`
	// Configures the monitoring schedule.
	ScheduleConfig interface{} `field:"optional" json:"scheduleConfig" yaml:"scheduleConfig"`
}

