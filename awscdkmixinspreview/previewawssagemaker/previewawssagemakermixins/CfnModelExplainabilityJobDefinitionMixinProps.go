package previewawssagemakermixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnModelExplainabilityJobDefinitionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnModelExplainabilityJobDefinitionMixinProps := &CfnModelExplainabilityJobDefinitionMixinProps{
//   	EndpointName: jsii.String("endpointName"),
//   	JobDefinitionName: jsii.String("jobDefinitionName"),
//   	JobResources: &MonitoringResourcesProperty{
//   		ClusterConfig: &ClusterConfigProperty{
//   			InstanceCount: jsii.Number(123),
//   			InstanceType: jsii.String("instanceType"),
//   			VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   			VolumeSizeInGb: jsii.Number(123),
//   		},
//   	},
//   	ModelExplainabilityAppSpecification: &ModelExplainabilityAppSpecificationProperty{
//   		ConfigUri: jsii.String("configUri"),
//   		Environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		ImageUri: jsii.String("imageUri"),
//   	},
//   	ModelExplainabilityBaselineConfig: &ModelExplainabilityBaselineConfigProperty{
//   		BaseliningJobName: jsii.String("baseliningJobName"),
//   		ConstraintsResource: &ConstraintsResourceProperty{
//   			S3Uri: jsii.String("s3Uri"),
//   		},
//   	},
//   	ModelExplainabilityJobInput: &ModelExplainabilityJobInputProperty{
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
//   			FeaturesAttribute: jsii.String("featuresAttribute"),
//   			InferenceAttribute: jsii.String("inferenceAttribute"),
//   			LocalPath: jsii.String("localPath"),
//   			ProbabilityAttribute: jsii.String("probabilityAttribute"),
//   			S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   			S3InputMode: jsii.String("s3InputMode"),
//   		},
//   		EndpointInput: &EndpointInputProperty{
//   			EndpointName: jsii.String("endpointName"),
//   			FeaturesAttribute: jsii.String("featuresAttribute"),
//   			InferenceAttribute: jsii.String("inferenceAttribute"),
//   			LocalPath: jsii.String("localPath"),
//   			ProbabilityAttribute: jsii.String("probabilityAttribute"),
//   			S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   			S3InputMode: jsii.String("s3InputMode"),
//   		},
//   	},
//   	ModelExplainabilityJobOutputConfig: &MonitoringOutputConfigProperty{
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
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelexplainabilityjobdefinition.html
//
type CfnModelExplainabilityJobDefinitionMixinProps struct {
	// The name of the endpoint used to run the monitoring job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelexplainabilityjobdefinition.html#cfn-sagemaker-modelexplainabilityjobdefinition-endpointname
	//
	EndpointName *string `field:"optional" json:"endpointName" yaml:"endpointName"`
	// The name of the model explainability job definition.
	//
	// The name must be unique within an AWS Region in the AWS account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelexplainabilityjobdefinition.html#cfn-sagemaker-modelexplainabilityjobdefinition-jobdefinitionname
	//
	JobDefinitionName *string `field:"optional" json:"jobDefinitionName" yaml:"jobDefinitionName"`
	// Identifies the resources to deploy for a monitoring job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelexplainabilityjobdefinition.html#cfn-sagemaker-modelexplainabilityjobdefinition-jobresources
	//
	JobResources interface{} `field:"optional" json:"jobResources" yaml:"jobResources"`
	// Configures the model explainability job to run a specified Docker container image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelexplainabilityjobdefinition.html#cfn-sagemaker-modelexplainabilityjobdefinition-modelexplainabilityappspecification
	//
	ModelExplainabilityAppSpecification interface{} `field:"optional" json:"modelExplainabilityAppSpecification" yaml:"modelExplainabilityAppSpecification"`
	// The baseline configuration for a model explainability job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelexplainabilityjobdefinition.html#cfn-sagemaker-modelexplainabilityjobdefinition-modelexplainabilitybaselineconfig
	//
	ModelExplainabilityBaselineConfig interface{} `field:"optional" json:"modelExplainabilityBaselineConfig" yaml:"modelExplainabilityBaselineConfig"`
	// Inputs for the model explainability job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelexplainabilityjobdefinition.html#cfn-sagemaker-modelexplainabilityjobdefinition-modelexplainabilityjobinput
	//
	ModelExplainabilityJobInput interface{} `field:"optional" json:"modelExplainabilityJobInput" yaml:"modelExplainabilityJobInput"`
	// The output configuration for monitoring jobs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelexplainabilityjobdefinition.html#cfn-sagemaker-modelexplainabilityjobdefinition-modelexplainabilityjoboutputconfig
	//
	ModelExplainabilityJobOutputConfig interface{} `field:"optional" json:"modelExplainabilityJobOutputConfig" yaml:"modelExplainabilityJobOutputConfig"`
	// Networking options for a model explainability job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelexplainabilityjobdefinition.html#cfn-sagemaker-modelexplainabilityjobdefinition-networkconfig
	//
	NetworkConfig interface{} `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelexplainabilityjobdefinition.html#cfn-sagemaker-modelexplainabilityjobdefinition-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// A time limit for how long the monitoring job is allowed to run before stopping.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelexplainabilityjobdefinition.html#cfn-sagemaker-modelexplainabilityjobdefinition-stoppingcondition
	//
	StoppingCondition interface{} `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelexplainabilityjobdefinition.html#cfn-sagemaker-modelexplainabilityjobdefinition-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

