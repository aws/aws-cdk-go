package previewawssagemakermixins


// Defines the monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   monitoringJobDefinitionProperty := &MonitoringJobDefinitionProperty{
//   	BaselineConfig: &BaselineConfigProperty{
//   		ConstraintsResource: &ConstraintsResourceProperty{
//   			S3Uri: jsii.String("s3Uri"),
//   		},
//   		StatisticsResource: &StatisticsResourceProperty{
//   			S3Uri: jsii.String("s3Uri"),
//   		},
//   	},
//   	Environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	MonitoringAppSpecification: &MonitoringAppSpecificationProperty{
//   		ContainerArguments: []*string{
//   			jsii.String("containerArguments"),
//   		},
//   		ContainerEntrypoint: []*string{
//   			jsii.String("containerEntrypoint"),
//   		},
//   		ImageUri: jsii.String("imageUri"),
//   		PostAnalyticsProcessorSourceUri: jsii.String("postAnalyticsProcessorSourceUri"),
//   		RecordPreprocessorSourceUri: jsii.String("recordPreprocessorSourceUri"),
//   	},
//   	MonitoringInputs: []interface{}{
//   		&MonitoringInputProperty{
//   			BatchTransformInput: &BatchTransformInputProperty{
//   				DataCapturedDestinationS3Uri: jsii.String("dataCapturedDestinationS3Uri"),
//   				DatasetFormat: &DatasetFormatProperty{
//   					Csv: &CsvProperty{
//   						Header: jsii.Boolean(false),
//   					},
//   					Json: &JsonProperty{
//   						Line: jsii.Boolean(false),
//   					},
//   					Parquet: jsii.Boolean(false),
//   				},
//   				ExcludeFeaturesAttribute: jsii.String("excludeFeaturesAttribute"),
//   				LocalPath: jsii.String("localPath"),
//   				S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   				S3InputMode: jsii.String("s3InputMode"),
//   			},
//   			EndpointInput: &EndpointInputProperty{
//   				EndpointName: jsii.String("endpointName"),
//   				ExcludeFeaturesAttribute: jsii.String("excludeFeaturesAttribute"),
//   				LocalPath: jsii.String("localPath"),
//   				S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   				S3InputMode: jsii.String("s3InputMode"),
//   			},
//   		},
//   	},
//   	MonitoringOutputConfig: &MonitoringOutputConfigProperty{
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   		MonitoringOutputs: []interface{}{
//   			&MonitoringOutputProperty{
//   				S3Output: &S3OutputProperty{
//   					LocalPath: jsii.String("localPath"),
//   					S3UploadMode: jsii.String("s3UploadMode"),
//   					S3Uri: jsii.String("s3Uri"),
//   				},
//   			},
//   		},
//   	},
//   	MonitoringResources: &MonitoringResourcesProperty{
//   		ClusterConfig: &ClusterConfigProperty{
//   			InstanceCount: jsii.Number(123),
//   			InstanceType: jsii.String("instanceType"),
//   			VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   			VolumeSizeInGb: jsii.Number(123),
//   		},
//   	},
//   	NetworkConfig: &NetworkConfigProperty{
//   		EnableInterContainerTrafficEncryption: jsii.Boolean(false),
//   		EnableNetworkIsolation: jsii.Boolean(false),
//   		VpcConfig: &VpcConfigProperty{
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			Subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	StoppingCondition: &StoppingConditionProperty{
//   		MaxRuntimeInSeconds: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringjobdefinition.html
//
type CfnMonitoringSchedulePropsMixin_MonitoringJobDefinitionProperty struct {
	// Baseline configuration used to validate that the data conforms to the specified constraints and statistics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringjobdefinition.html#cfn-sagemaker-monitoringschedule-monitoringjobdefinition-baselineconfig
	//
	BaselineConfig interface{} `field:"optional" json:"baselineConfig" yaml:"baselineConfig"`
	// Sets the environment variables in the Docker container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringjobdefinition.html#cfn-sagemaker-monitoringschedule-monitoringjobdefinition-environment
	//
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// Configures the monitoring job to run a specified Docker container image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringjobdefinition.html#cfn-sagemaker-monitoringschedule-monitoringjobdefinition-monitoringappspecification
	//
	MonitoringAppSpecification interface{} `field:"optional" json:"monitoringAppSpecification" yaml:"monitoringAppSpecification"`
	// The array of inputs for the monitoring job.
	//
	// Currently we support monitoring an Amazon SageMaker AI Endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringjobdefinition.html#cfn-sagemaker-monitoringschedule-monitoringjobdefinition-monitoringinputs
	//
	MonitoringInputs interface{} `field:"optional" json:"monitoringInputs" yaml:"monitoringInputs"`
	// The array of outputs from the monitoring job to be uploaded to Amazon S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringjobdefinition.html#cfn-sagemaker-monitoringschedule-monitoringjobdefinition-monitoringoutputconfig
	//
	MonitoringOutputConfig interface{} `field:"optional" json:"monitoringOutputConfig" yaml:"monitoringOutputConfig"`
	// Identifies the resources, ML compute instances, and ML storage volumes to deploy for a monitoring job.
	//
	// In distributed processing, you specify more than one instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringjobdefinition.html#cfn-sagemaker-monitoringschedule-monitoringjobdefinition-monitoringresources
	//
	MonitoringResources interface{} `field:"optional" json:"monitoringResources" yaml:"monitoringResources"`
	// Specifies networking options for an monitoring job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringjobdefinition.html#cfn-sagemaker-monitoringschedule-monitoringjobdefinition-networkconfig
	//
	NetworkConfig interface{} `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker AI can assume to perform tasks on your behalf.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringjobdefinition.html#cfn-sagemaker-monitoringschedule-monitoringjobdefinition-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Specifies a time limit for how long the monitoring job is allowed to run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringjobdefinition.html#cfn-sagemaker-monitoringschedule-monitoringjobdefinition-stoppingcondition
	//
	StoppingCondition interface{} `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
}

