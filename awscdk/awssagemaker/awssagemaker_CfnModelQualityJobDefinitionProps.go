package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnModelQualityJobDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var json interface{}
//
//   cfnModelQualityJobDefinitionProps := &cfnModelQualityJobDefinitionProps{
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
//   	modelQualityAppSpecification: &modelQualityAppSpecificationProperty{
//   		imageUri: jsii.String("imageUri"),
//   		problemType: jsii.String("problemType"),
//
//   		// the properties below are optional
//   		containerArguments: []*string{
//   			jsii.String("containerArguments"),
//   		},
//   		containerEntrypoint: []*string{
//   			jsii.String("containerEntrypoint"),
//   		},
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		postAnalyticsProcessorSourceUri: jsii.String("postAnalyticsProcessorSourceUri"),
//   		recordPreprocessorSourceUri: jsii.String("recordPreprocessorSourceUri"),
//   	},
//   	modelQualityJobInput: &modelQualityJobInputProperty{
//   		groundTruthS3Input: &monitoringGroundTruthS3InputProperty{
//   			s3Uri: jsii.String("s3Uri"),
//   		},
//
//   		// the properties below are optional
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
//   			endTimeOffset: jsii.String("endTimeOffset"),
//   			inferenceAttribute: jsii.String("inferenceAttribute"),
//   			probabilityAttribute: jsii.String("probabilityAttribute"),
//   			probabilityThresholdAttribute: jsii.Number(123),
//   			s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   			s3InputMode: jsii.String("s3InputMode"),
//   			startTimeOffset: jsii.String("startTimeOffset"),
//   		},
//   		endpointInput: &endpointInputProperty{
//   			endpointName: jsii.String("endpointName"),
//   			localPath: jsii.String("localPath"),
//
//   			// the properties below are optional
//   			endTimeOffset: jsii.String("endTimeOffset"),
//   			inferenceAttribute: jsii.String("inferenceAttribute"),
//   			probabilityAttribute: jsii.String("probabilityAttribute"),
//   			probabilityThresholdAttribute: jsii.Number(123),
//   			s3DataDistributionType: jsii.String("s3DataDistributionType"),
//   			s3InputMode: jsii.String("s3InputMode"),
//   			startTimeOffset: jsii.String("startTimeOffset"),
//   		},
//   	},
//   	modelQualityJobOutputConfig: &monitoringOutputConfigProperty{
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
//   	modelQualityBaselineConfig: &modelQualityBaselineConfigProperty{
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
type CfnModelQualityJobDefinitionProps struct {
	// Identifies the resources to deploy for a monitoring job.
	JobResources interface{} `field:"required" json:"jobResources" yaml:"jobResources"`
	// Container image configuration object for the monitoring job.
	ModelQualityAppSpecification interface{} `field:"required" json:"modelQualityAppSpecification" yaml:"modelQualityAppSpecification"`
	// A list of the inputs that are monitored.
	//
	// Currently endpoints are supported.
	ModelQualityJobInput interface{} `field:"required" json:"modelQualityJobInput" yaml:"modelQualityJobInput"`
	// The output configuration for monitoring jobs.
	ModelQualityJobOutputConfig interface{} `field:"required" json:"modelQualityJobOutputConfig" yaml:"modelQualityJobOutputConfig"`
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// `AWS::SageMaker::ModelQualityJobDefinition.EndpointName`.
	EndpointName *string `field:"optional" json:"endpointName" yaml:"endpointName"`
	// The name of the monitoring job definition.
	JobDefinitionName *string `field:"optional" json:"jobDefinitionName" yaml:"jobDefinitionName"`
	// Specifies the constraints and baselines for the monitoring job.
	ModelQualityBaselineConfig interface{} `field:"optional" json:"modelQualityBaselineConfig" yaml:"modelQualityBaselineConfig"`
	// Specifies the network configuration for the monitoring job.
	NetworkConfig interface{} `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// A time limit for how long the monitoring job is allowed to run before stopping.
	StoppingCondition interface{} `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

