package awssagemaker


// Defines the monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var json interface{}
//
//   monitoringJobDefinitionProperty := &monitoringJobDefinitionProperty{
//   	monitoringAppSpecification: &monitoringAppSpecificationProperty{
//   		imageUri: jsii.String("imageUri"),
//
//   		// the properties below are optional
//   		containerArguments: []*string{
//   			jsii.String("containerArguments"),
//   		},
//   		containerEntrypoint: []*string{
//   			jsii.String("containerEntrypoint"),
//   		},
//   		postAnalyticsProcessorSourceUri: jsii.String("postAnalyticsProcessorSourceUri"),
//   		recordPreprocessorSourceUri: jsii.String("recordPreprocessorSourceUri"),
//   	},
//   	monitoringInputs: []interface{}{
//   		&monitoringInputProperty{
//   			batchTransformInput: &batchTransformInputProperty{
//   				dataCapturedDestinationS3Uri: jsii.String("dataCapturedDestinationS3Uri"),
//   				datasetFormat: &datasetFormatProperty{
//   					csv: &csvProperty{
//   						header: jsii.Boolean(false),
//   					},
//   					json: json,
//   					parquet: jsii.Boolean(false),
//   				},
//   				localPath: jsii.String("localPath"),
//
//   				// the properties below are optional
//   				s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   				s3InputMode: jsii.String("s3InputMode"),
//   			},
//   			endpointInput: &endpointInputProperty{
//   				endpointName: jsii.String("endpointName"),
//   				localPath: jsii.String("localPath"),
//
//   				// the properties below are optional
//   				s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   				s3InputMode: jsii.String("s3InputMode"),
//   			},
//   		},
//   	},
//   	monitoringOutputConfig: &monitoringOutputConfigProperty{
//   		monitoringOutputs: []interface{}{
//   			&monitoringOutputProperty{
//   				s3Output: &s3OutputProperty{
//   					localPath: jsii.String("localPath"),
//   					s3Uri: jsii.String("s3Uri"),
//
//   					// the properties below are optional
//   					s3UploadMode: jsii.String("s3UploadMode"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	monitoringResources: &monitoringResourcesProperty{
//   		clusterConfig: &clusterConfigProperty{
//   			instanceCount: jsii.Number(123),
//   			instanceType: jsii.String("instanceType"),
//   			volumeSizeInGb: jsii.Number(123),
//
//   			// the properties below are optional
//   			volumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   		},
//   	},
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	baselineConfig: &baselineConfigProperty{
//   		constraintsResource: &constraintsResourceProperty{
//   			s3Uri: jsii.String("s3Uri"),
//   		},
//   		statisticsResource: &statisticsResourceProperty{
//   			s3Uri: jsii.String("s3Uri"),
//   		},
//   	},
//   	environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	networkConfig: &networkConfigProperty{
//   		enableInterContainerTrafficEncryption: jsii.Boolean(false),
//   		enableNetworkIsolation: jsii.Boolean(false),
//   		vpcConfig: &vpcConfigProperty{
//   			securityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
//   	},
//   	stoppingCondition: &stoppingConditionProperty{
//   		maxRuntimeInSeconds: jsii.Number(123),
//   	},
//   }
//
type CfnMonitoringSchedule_MonitoringJobDefinitionProperty struct {
	// Configures the monitoring job to run a specified Docker container image.
	MonitoringAppSpecification interface{} `field:"required" json:"monitoringAppSpecification" yaml:"monitoringAppSpecification"`
	// The array of inputs for the monitoring job.
	//
	// Currently we support monitoring an Amazon SageMaker Endpoint.
	MonitoringInputs interface{} `field:"required" json:"monitoringInputs" yaml:"monitoringInputs"`
	// The array of outputs from the monitoring job to be uploaded to Amazon Simple Storage Service (Amazon S3).
	MonitoringOutputConfig interface{} `field:"required" json:"monitoringOutputConfig" yaml:"monitoringOutputConfig"`
	// Identifies the resources, ML compute instances, and ML storage volumes to deploy for a monitoring job.
	//
	// In distributed processing, you specify more than one instance.
	MonitoringResources interface{} `field:"required" json:"monitoringResources" yaml:"monitoringResources"`
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Baseline configuration used to validate that the data conforms to the specified constraints and statistics.
	BaselineConfig interface{} `field:"optional" json:"baselineConfig" yaml:"baselineConfig"`
	// Sets the environment variables in the Docker container.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// Specifies networking options for an monitoring job.
	NetworkConfig interface{} `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// Specifies a time limit for how long the monitoring job is allowed to run.
	StoppingCondition interface{} `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
}

