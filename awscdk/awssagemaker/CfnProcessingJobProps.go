package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnProcessingJob`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProcessingJobProps := &CfnProcessingJobProps{
//   	AppSpecification: &AppSpecificationProperty{
//   		ImageUri: jsii.String("imageUri"),
//
//   		// the properties below are optional
//   		ContainerArguments: []*string{
//   			jsii.String("containerArguments"),
//   		},
//   		ContainerEntrypoint: []*string{
//   			jsii.String("containerEntrypoint"),
//   		},
//   	},
//   	ProcessingResources: &ProcessingResourcesProperty{
//   		ClusterConfig: &ClusterConfigProperty{
//   			InstanceCount: jsii.Number(123),
//   			InstanceType: jsii.String("instanceType"),
//   			VolumeSizeInGb: jsii.Number(123),
//
//   			// the properties below are optional
//   			VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	Environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	ExperimentConfig: &ExperimentConfigProperty{
//   		ExperimentName: jsii.String("experimentName"),
//   		RunName: jsii.String("runName"),
//   		TrialComponentDisplayName: jsii.String("trialComponentDisplayName"),
//   		TrialName: jsii.String("trialName"),
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
//   	ProcessingInputs: []interface{}{
//   		&ProcessingInputsObjectProperty{
//   			InputName: jsii.String("inputName"),
//
//   			// the properties below are optional
//   			AppManaged: jsii.Boolean(false),
//   			DatasetDefinition: &DatasetDefinitionProperty{
//   				AthenaDatasetDefinition: &AthenaDatasetDefinitionProperty{
//   					Catalog: jsii.String("catalog"),
//   					Database: jsii.String("database"),
//   					OutputFormat: jsii.String("outputFormat"),
//   					OutputS3Uri: jsii.String("outputS3Uri"),
//   					QueryString: jsii.String("queryString"),
//
//   					// the properties below are optional
//   					KmsKeyId: jsii.String("kmsKeyId"),
//   					OutputCompression: jsii.String("outputCompression"),
//   					WorkGroup: jsii.String("workGroup"),
//   				},
//   				DataDistributionType: jsii.String("dataDistributionType"),
//   				InputMode: jsii.String("inputMode"),
//   				LocalPath: jsii.String("localPath"),
//   				RedshiftDatasetDefinition: &RedshiftDatasetDefinitionProperty{
//   					ClusterId: jsii.String("clusterId"),
//   					ClusterRoleArn: jsii.String("clusterRoleArn"),
//   					Database: jsii.String("database"),
//   					DbUser: jsii.String("dbUser"),
//   					OutputFormat: jsii.String("outputFormat"),
//   					OutputS3Uri: jsii.String("outputS3Uri"),
//   					QueryString: jsii.String("queryString"),
//
//   					// the properties below are optional
//   					KmsKeyId: jsii.String("kmsKeyId"),
//   					OutputCompression: jsii.String("outputCompression"),
//   				},
//   			},
//   			S3Input: &S3InputProperty{
//   				S3DataType: jsii.String("s3DataType"),
//   				S3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				LocalPath: jsii.String("localPath"),
//   				S3CompressionType: jsii.String("s3CompressionType"),
//   				S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   				S3InputMode: jsii.String("s3InputMode"),
//   			},
//   		},
//   	},
//   	ProcessingJobName: jsii.String("processingJobName"),
//   	ProcessingOutputConfig: &ProcessingOutputConfigProperty{
//   		Outputs: []interface{}{
//   			&ProcessingOutputsObjectProperty{
//   				OutputName: jsii.String("outputName"),
//
//   				// the properties below are optional
//   				AppManaged: jsii.Boolean(false),
//   				FeatureStoreOutput: &FeatureStoreOutputProperty{
//   					FeatureGroupName: jsii.String("featureGroupName"),
//   				},
//   				S3Output: &S3OutputProperty{
//   					S3UploadMode: jsii.String("s3UploadMode"),
//   					S3Uri: jsii.String("s3Uri"),
//
//   					// the properties below are optional
//   					LocalPath: jsii.String("localPath"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	StoppingCondition: &StoppingConditionProperty{
//   		MaxRuntimeInSeconds: jsii.Number(123),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-processingjob.html
//
type CfnProcessingJobProps struct {
	// Configuration to run a processing job in a specified container image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-processingjob.html#cfn-sagemaker-processingjob-appspecification
	//
	AppSpecification interface{} `field:"required" json:"appSpecification" yaml:"appSpecification"`
	// Identifies the resources, ML compute instances, and ML storage volumes to deploy for a processing job.
	//
	// In distributed training, you specify more than one instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-processingjob.html#cfn-sagemaker-processingjob-processingresources
	//
	ProcessingResources interface{} `field:"required" json:"processingResources" yaml:"processingResources"`
	// The ARN of the role used to create the processing job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-processingjob.html#cfn-sagemaker-processingjob-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Sets the environment variables in the Docker container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-processingjob.html#cfn-sagemaker-processingjob-environment
	//
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// Associates a SageMaker job as a trial component with an experiment and trial.
	//
	// Specified when you call the [CreateProcessingJob](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateProcessingJob.html) API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-processingjob.html#cfn-sagemaker-processingjob-experimentconfig
	//
	ExperimentConfig interface{} `field:"optional" json:"experimentConfig" yaml:"experimentConfig"`
	// Networking options for a job, such as network traffic encryption between containers, whether to allow inbound and outbound network calls to and from containers, and the VPC subnets and security groups to use for VPC-enabled jobs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-processingjob.html#cfn-sagemaker-processingjob-networkconfig
	//
	NetworkConfig interface{} `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// List of input configurations for the processing job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-processingjob.html#cfn-sagemaker-processingjob-processinginputs
	//
	ProcessingInputs interface{} `field:"optional" json:"processingInputs" yaml:"processingInputs"`
	// The name of the processing job.
	//
	// If you don't provide a job name, then a unique name is automatically created for the job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-processingjob.html#cfn-sagemaker-processingjob-processingjobname
	//
	ProcessingJobName *string `field:"optional" json:"processingJobName" yaml:"processingJobName"`
	// Contains information about the output location for the compiled model and the target device that the model runs on.
	//
	// `TargetDevice` and `TargetPlatform` are mutually exclusive, so you need to choose one between the two to specify your target device or platform. If you cannot find your device you want to use from the `TargetDevice` list, use `TargetPlatform` to describe the platform of your edge device and `CompilerOptions` if there are specific settings that are required or recommended to use for particular TargetPlatform.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-processingjob.html#cfn-sagemaker-processingjob-processingoutputconfig
	//
	ProcessingOutputConfig interface{} `field:"optional" json:"processingOutputConfig" yaml:"processingOutputConfig"`
	// Configures conditions under which the processing job should be stopped, such as how long the processing job has been running.
	//
	// After the condition is met, the processing job is stopped.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-processingjob.html#cfn-sagemaker-processingjob-stoppingcondition
	//
	StoppingCondition interface{} `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
	// An array of key-value pairs.
	//
	// For more information, see [Using Cost Allocation Tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-whatURL) in the *AWS Billing and Cost Management User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-processingjob.html#cfn-sagemaker-processingjob-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

