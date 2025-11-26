package previewawssagemakermixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnModelQualityJobDefinitionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnModelQualityJobDefinitionMixinProps := &CfnModelQualityJobDefinitionMixinProps{
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
//   	ModelQualityAppSpecification: &ModelQualityAppSpecificationProperty{
//   		ContainerArguments: []*string{
//   			jsii.String("containerArguments"),
//   		},
//   		ContainerEntrypoint: []*string{
//   			jsii.String("containerEntrypoint"),
//   		},
//   		Environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		ImageUri: jsii.String("imageUri"),
//   		PostAnalyticsProcessorSourceUri: jsii.String("postAnalyticsProcessorSourceUri"),
//   		ProblemType: jsii.String("problemType"),
//   		RecordPreprocessorSourceUri: jsii.String("recordPreprocessorSourceUri"),
//   	},
//   	ModelQualityBaselineConfig: &ModelQualityBaselineConfigProperty{
//   		BaseliningJobName: jsii.String("baseliningJobName"),
//   		ConstraintsResource: &ConstraintsResourceProperty{
//   			S3Uri: jsii.String("s3Uri"),
//   		},
//   	},
//   	ModelQualityJobInput: &ModelQualityJobInputProperty{
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
//   			EndTimeOffset: jsii.String("endTimeOffset"),
//   			InferenceAttribute: jsii.String("inferenceAttribute"),
//   			LocalPath: jsii.String("localPath"),
//   			ProbabilityAttribute: jsii.String("probabilityAttribute"),
//   			ProbabilityThresholdAttribute: jsii.Number(123),
//   			S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   			S3InputMode: jsii.String("s3InputMode"),
//   			StartTimeOffset: jsii.String("startTimeOffset"),
//   		},
//   		EndpointInput: &EndpointInputProperty{
//   			EndpointName: jsii.String("endpointName"),
//   			EndTimeOffset: jsii.String("endTimeOffset"),
//   			InferenceAttribute: jsii.String("inferenceAttribute"),
//   			LocalPath: jsii.String("localPath"),
//   			ProbabilityAttribute: jsii.String("probabilityAttribute"),
//   			ProbabilityThresholdAttribute: jsii.Number(123),
//   			S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   			S3InputMode: jsii.String("s3InputMode"),
//   			StartTimeOffset: jsii.String("startTimeOffset"),
//   		},
//   		GroundTruthS3Input: &MonitoringGroundTruthS3InputProperty{
//   			S3Uri: jsii.String("s3Uri"),
//   		},
//   	},
//   	ModelQualityJobOutputConfig: &MonitoringOutputConfigProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html
//
type CfnModelQualityJobDefinitionMixinProps struct {
	// The name of the endpoint used to run the monitoring job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-endpointname
	//
	EndpointName *string `field:"optional" json:"endpointName" yaml:"endpointName"`
	// The name of the monitoring job definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-jobdefinitionname
	//
	JobDefinitionName *string `field:"optional" json:"jobDefinitionName" yaml:"jobDefinitionName"`
	// Identifies the resources to deploy for a monitoring job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-jobresources
	//
	JobResources interface{} `field:"optional" json:"jobResources" yaml:"jobResources"`
	// Container image configuration object for the monitoring job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-modelqualityappspecification
	//
	ModelQualityAppSpecification interface{} `field:"optional" json:"modelQualityAppSpecification" yaml:"modelQualityAppSpecification"`
	// Specifies the constraints and baselines for the monitoring job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-modelqualitybaselineconfig
	//
	ModelQualityBaselineConfig interface{} `field:"optional" json:"modelQualityBaselineConfig" yaml:"modelQualityBaselineConfig"`
	// A list of the inputs that are monitored.
	//
	// Currently endpoints are supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-modelqualityjobinput
	//
	ModelQualityJobInput interface{} `field:"optional" json:"modelQualityJobInput" yaml:"modelQualityJobInput"`
	// The output configuration for monitoring jobs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-modelqualityjoboutputconfig
	//
	ModelQualityJobOutputConfig interface{} `field:"optional" json:"modelQualityJobOutputConfig" yaml:"modelQualityJobOutputConfig"`
	// Specifies the network configuration for the monitoring job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-networkconfig
	//
	NetworkConfig interface{} `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// A time limit for how long the monitoring job is allowed to run before stopping.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-stoppingcondition
	//
	StoppingCondition interface{} `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelqualityjobdefinition.html#cfn-sagemaker-modelqualityjobdefinition-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

