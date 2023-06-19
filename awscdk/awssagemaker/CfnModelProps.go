package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnModel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var environment interface{}
//
//   cfnModelProps := &CfnModelProps{
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//
//   	// the properties below are optional
//   	Containers: []interface{}{
//   		&ContainerDefinitionProperty{
//   			ContainerHostname: jsii.String("containerHostname"),
//   			Environment: environment,
//   			Image: jsii.String("image"),
//   			ImageConfig: &ImageConfigProperty{
//   				RepositoryAccessMode: jsii.String("repositoryAccessMode"),
//
//   				// the properties below are optional
//   				RepositoryAuthConfig: &RepositoryAuthConfigProperty{
//   					RepositoryCredentialsProviderArn: jsii.String("repositoryCredentialsProviderArn"),
//   				},
//   			},
//   			InferenceSpecificationName: jsii.String("inferenceSpecificationName"),
//   			Mode: jsii.String("mode"),
//   			ModelDataUrl: jsii.String("modelDataUrl"),
//   			ModelPackageName: jsii.String("modelPackageName"),
//   			MultiModelConfig: &MultiModelConfigProperty{
//   				ModelCacheSetting: jsii.String("modelCacheSetting"),
//   			},
//   		},
//   	},
//   	EnableNetworkIsolation: jsii.Boolean(false),
//   	InferenceExecutionConfig: &InferenceExecutionConfigProperty{
//   		Mode: jsii.String("mode"),
//   	},
//   	ModelName: jsii.String("modelName"),
//   	PrimaryContainer: &ContainerDefinitionProperty{
//   		ContainerHostname: jsii.String("containerHostname"),
//   		Environment: environment,
//   		Image: jsii.String("image"),
//   		ImageConfig: &ImageConfigProperty{
//   			RepositoryAccessMode: jsii.String("repositoryAccessMode"),
//
//   			// the properties below are optional
//   			RepositoryAuthConfig: &RepositoryAuthConfigProperty{
//   				RepositoryCredentialsProviderArn: jsii.String("repositoryCredentialsProviderArn"),
//   			},
//   		},
//   		InferenceSpecificationName: jsii.String("inferenceSpecificationName"),
//   		Mode: jsii.String("mode"),
//   		ModelDataUrl: jsii.String("modelDataUrl"),
//   		ModelPackageName: jsii.String("modelPackageName"),
//   		MultiModelConfig: &MultiModelConfigProperty{
//   			ModelCacheSetting: jsii.String("modelCacheSetting"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcConfig: &VpcConfigProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   }
//
type CfnModelProps struct {
	// The Amazon Resource Name (ARN) of the IAM role that SageMaker can assume to access model artifacts and docker image for deployment on ML compute instances or for batch transform jobs.
	//
	// Deploying on ML compute instances is part of model hosting. For more information, see [SageMaker Roles](https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html) .
	//
	// > To be able to pass this role to SageMaker, the caller of this API must have the `iam:PassRole` permission.
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// Specifies the containers in the inference pipeline.
	Containers interface{} `field:"optional" json:"containers" yaml:"containers"`
	// Isolates the model container.
	//
	// No inbound or outbound network calls can be made to or from the model container.
	EnableNetworkIsolation interface{} `field:"optional" json:"enableNetworkIsolation" yaml:"enableNetworkIsolation"`
	// Specifies details of how containers in a multi-container endpoint are called.
	InferenceExecutionConfig interface{} `field:"optional" json:"inferenceExecutionConfig" yaml:"inferenceExecutionConfig"`
	// The name of the new model.
	ModelName *string `field:"optional" json:"modelName" yaml:"modelName"`
	// The location of the primary docker image containing inference code, associated artifacts, and custom environment map that the inference code uses when the model is deployed for predictions.
	PrimaryContainer interface{} `field:"optional" json:"primaryContainer" yaml:"primaryContainer"`
	// A list of key-value pairs to apply to this resource.
	//
	// For more information, see [Resource Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) and [Using Cost Allocation Tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what) in the *AWS Billing and Cost Management User Guide* .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A [VpcConfig](https://docs.aws.amazon.com/sagemaker/latest/dg/API_VpcConfig.html) object that specifies the VPC that you want your model to connect to. Control access to and from your model container by configuring the VPC. `VpcConfig` is used in hosting services and in batch transform. For more information, see [Protect Endpoints by Using an Amazon Virtual Private Cloud](https://docs.aws.amazon.com/sagemaker/latest/dg/host-vpc.html) and [Protect Data in Batch Transform Jobs by Using an Amazon Virtual Private Cloud](https://docs.aws.amazon.com/sagemaker/latest/dg/batch-vpc.html) .
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

