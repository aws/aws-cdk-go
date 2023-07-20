package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnModelBiasJobDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnModelBiasJobDefinitionProps := &CfnModelBiasJobDefinitionProps{
//   	JobResources: &MonitoringResourcesProperty{
//   		ClusterConfig: &ClusterConfigProperty{
//   			InstanceCount: jsii.Number(123),
//   			InstanceType: jsii.String("instanceType"),
//   			VolumeSizeInGb: jsii.Number(123),
//
//   			// the properties below are optional
//   			VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   		},
//   	},
//   	ModelBiasAppSpecification: &ModelBiasAppSpecificationProperty{
//   		ConfigUri: jsii.String("configUri"),
//   		ImageUri: jsii.String("imageUri"),
//
//   		// the properties below are optional
//   		Environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   	},
//   	ModelBiasJobInput: &ModelBiasJobInputProperty{
//   		GroundTruthS3Input: &MonitoringGroundTruthS3InputProperty{
//   			S3Uri: jsii.String("s3Uri"),
//   		},
//
//   		// the properties below are optional
//   		BatchTransformInput: &BatchTransformInputProperty{
//   			DataCapturedDestinationS3Uri: jsii.String("dataCapturedDestinationS3Uri"),
//   			DatasetFormat: &DatasetFormatProperty{
//   				Csv: &CsvProperty{
//   					Header: jsii.Boolean(false),
//   				},
//   				Json: &JsonProperty{
//   					Line: jsii.Boolean(false),
//   				},
//   				Parquet: jsii.Boolean(false),
//   			},
//   			LocalPath: jsii.String("localPath"),
//
//   			// the properties below are optional
//   			EndTimeOffset: jsii.String("endTimeOffset"),
//   			FeaturesAttribute: jsii.String("featuresAttribute"),
//   			InferenceAttribute: jsii.String("inferenceAttribute"),
//   			ProbabilityAttribute: jsii.String("probabilityAttribute"),
//   			ProbabilityThresholdAttribute: jsii.Number(123),
//   			S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   			S3InputMode: jsii.String("s3InputMode"),
//   			StartTimeOffset: jsii.String("startTimeOffset"),
//   		},
//   		EndpointInput: &EndpointInputProperty{
//   			EndpointName: jsii.String("endpointName"),
//   			LocalPath: jsii.String("localPath"),
//
//   			// the properties below are optional
//   			EndTimeOffset: jsii.String("endTimeOffset"),
//   			FeaturesAttribute: jsii.String("featuresAttribute"),
//   			InferenceAttribute: jsii.String("inferenceAttribute"),
//   			ProbabilityAttribute: jsii.String("probabilityAttribute"),
//   			ProbabilityThresholdAttribute: jsii.Number(123),
//   			S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   			S3InputMode: jsii.String("s3InputMode"),
//   			StartTimeOffset: jsii.String("startTimeOffset"),
//   		},
//   	},
//   	ModelBiasJobOutputConfig: &MonitoringOutputConfigProperty{
//   		MonitoringOutputs: []interface{}{
//   			&MonitoringOutputProperty{
//   				S3Output: &S3OutputProperty{
//   					LocalPath: jsii.String("localPath"),
//   					S3Uri: jsii.String("s3Uri"),
//
//   					// the properties below are optional
//   					S3UploadMode: jsii.String("s3UploadMode"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	EndpointName: jsii.String("endpointName"),
//   	JobDefinitionName: jsii.String("jobDefinitionName"),
//   	ModelBiasBaselineConfig: &ModelBiasBaselineConfigProperty{
//   		BaseliningJobName: jsii.String("baseliningJobName"),
//   		ConstraintsResource: &ConstraintsResourceProperty{
//   			S3Uri: jsii.String("s3Uri"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelbiasjobdefinition.html
//
type CfnModelBiasJobDefinitionProps struct {
	// Identifies the resources to deploy for a monitoring job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelbiasjobdefinition.html#cfn-sagemaker-modelbiasjobdefinition-jobresources
	//
	JobResources interface{} `field:"required" json:"jobResources" yaml:"jobResources"`
	// Configures the model bias job to run a specified Docker container image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelbiasjobdefinition.html#cfn-sagemaker-modelbiasjobdefinition-modelbiasappspecification
	//
	ModelBiasAppSpecification interface{} `field:"required" json:"modelBiasAppSpecification" yaml:"modelBiasAppSpecification"`
	// Inputs for the model bias job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelbiasjobdefinition.html#cfn-sagemaker-modelbiasjobdefinition-modelbiasjobinput
	//
	ModelBiasJobInput interface{} `field:"required" json:"modelBiasJobInput" yaml:"modelBiasJobInput"`
	// The output configuration for monitoring jobs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelbiasjobdefinition.html#cfn-sagemaker-modelbiasjobdefinition-modelbiasjoboutputconfig
	//
	ModelBiasJobOutputConfig interface{} `field:"required" json:"modelBiasJobOutputConfig" yaml:"modelBiasJobOutputConfig"`
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelbiasjobdefinition.html#cfn-sagemaker-modelbiasjobdefinition-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The name of the endpoint used to run the monitoring job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelbiasjobdefinition.html#cfn-sagemaker-modelbiasjobdefinition-endpointname
	//
	EndpointName *string `field:"optional" json:"endpointName" yaml:"endpointName"`
	// The name of the bias job definition.
	//
	// The name must be unique within an AWS Region in the AWS account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelbiasjobdefinition.html#cfn-sagemaker-modelbiasjobdefinition-jobdefinitionname
	//
	JobDefinitionName *string `field:"optional" json:"jobDefinitionName" yaml:"jobDefinitionName"`
	// The baseline configuration for a model bias job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelbiasjobdefinition.html#cfn-sagemaker-modelbiasjobdefinition-modelbiasbaselineconfig
	//
	ModelBiasBaselineConfig interface{} `field:"optional" json:"modelBiasBaselineConfig" yaml:"modelBiasBaselineConfig"`
	// Networking options for a model bias job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelbiasjobdefinition.html#cfn-sagemaker-modelbiasjobdefinition-networkconfig
	//
	NetworkConfig interface{} `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// A time limit for how long the monitoring job is allowed to run before stopping.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelbiasjobdefinition.html#cfn-sagemaker-modelbiasjobdefinition-stoppingcondition
	//
	StoppingCondition interface{} `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelbiasjobdefinition.html#cfn-sagemaker-modelbiasjobdefinition-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

