package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnModelExplainabilityJobDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var json interface{}
//
//   cfnModelExplainabilityJobDefinitionProps := &cfnModelExplainabilityJobDefinitionProps{
//   	jobResources: &monitoringResourcesProperty{
//   		clusterConfig: &clusterConfigProperty{
//   			instanceCount: jsii.Number(123),
//   			instanceType: jsii.String("instanceType"),
//   			volumeSizeInGb: jsii.Number(123),
//
//   			// the properties below are optional
//   			volumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   		},
//   	},
//   	modelExplainabilityAppSpecification: &modelExplainabilityAppSpecificationProperty{
//   		configUri: jsii.String("configUri"),
//   		imageUri: jsii.String("imageUri"),
//
//   		// the properties below are optional
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   	},
//   	modelExplainabilityJobInput: &modelExplainabilityJobInputProperty{
//   		batchTransformInput: &batchTransformInputProperty{
//   			dataCapturedDestinationS3Uri: jsii.String("dataCapturedDestinationS3Uri"),
//   			datasetFormat: &datasetFormatProperty{
//   				csv: &csvProperty{
//   					header: jsii.Boolean(false),
//   				},
//   				json: json,
//   				parquet: jsii.Boolean(false),
//   			},
//   			localPath: jsii.String("localPath"),
//
//   			// the properties below are optional
//   			featuresAttribute: jsii.String("featuresAttribute"),
//   			inferenceAttribute: jsii.String("inferenceAttribute"),
//   			probabilityAttribute: jsii.String("probabilityAttribute"),
//   			s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   			s3InputMode: jsii.String("s3InputMode"),
//   		},
//   		endpointInput: &endpointInputProperty{
//   			endpointName: jsii.String("endpointName"),
//   			localPath: jsii.String("localPath"),
//
//   			// the properties below are optional
//   			featuresAttribute: jsii.String("featuresAttribute"),
//   			inferenceAttribute: jsii.String("inferenceAttribute"),
//   			probabilityAttribute: jsii.String("probabilityAttribute"),
//   			s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   			s3InputMode: jsii.String("s3InputMode"),
//   		},
//   	},
//   	modelExplainabilityJobOutputConfig: &monitoringOutputConfigProperty{
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
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	endpointName: jsii.String("endpointName"),
//   	jobDefinitionName: jsii.String("jobDefinitionName"),
//   	modelExplainabilityBaselineConfig: &modelExplainabilityBaselineConfigProperty{
//   		baseliningJobName: jsii.String("baseliningJobName"),
//   		constraintsResource: &constraintsResourceProperty{
//   			s3Uri: jsii.String("s3Uri"),
//   		},
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
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnModelExplainabilityJobDefinitionProps struct {
	// Identifies the resources to deploy for a monitoring job.
	JobResources interface{} `field:"required" json:"jobResources" yaml:"jobResources"`
	// Configures the model explainability job to run a specified Docker container image.
	ModelExplainabilityAppSpecification interface{} `field:"required" json:"modelExplainabilityAppSpecification" yaml:"modelExplainabilityAppSpecification"`
	// Inputs for the model explainability job.
	ModelExplainabilityJobInput interface{} `field:"required" json:"modelExplainabilityJobInput" yaml:"modelExplainabilityJobInput"`
	// The output configuration for monitoring jobs.
	ModelExplainabilityJobOutputConfig interface{} `field:"required" json:"modelExplainabilityJobOutputConfig" yaml:"modelExplainabilityJobOutputConfig"`
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// `AWS::SageMaker::ModelExplainabilityJobDefinition.EndpointName`.
	EndpointName *string `field:"optional" json:"endpointName" yaml:"endpointName"`
	// The name of the model explainability job definition.
	//
	// The name must be unique within an AWS Region in the AWS account.
	JobDefinitionName *string `field:"optional" json:"jobDefinitionName" yaml:"jobDefinitionName"`
	// The baseline configuration for a model explainability job.
	ModelExplainabilityBaselineConfig interface{} `field:"optional" json:"modelExplainabilityBaselineConfig" yaml:"modelExplainabilityBaselineConfig"`
	// Networking options for a model explainability job.
	NetworkConfig interface{} `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// `AWS::SageMaker::ModelExplainabilityJobDefinition.StoppingCondition`.
	StoppingCondition interface{} `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

