package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnInferenceComponent`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInferenceComponentProps := &CfnInferenceComponentProps{
//   	EndpointName: jsii.String("endpointName"),
//   	Specification: &InferenceComponentSpecificationProperty{
//   		BaseInferenceComponentName: jsii.String("baseInferenceComponentName"),
//   		ComputeResourceRequirements: &InferenceComponentComputeResourceRequirementsProperty{
//   			MaxMemoryRequiredInMb: jsii.Number(123),
//   			MinMemoryRequiredInMb: jsii.Number(123),
//   			NumberOfAcceleratorDevicesRequired: jsii.Number(123),
//   			NumberOfCpuCoresRequired: jsii.Number(123),
//   		},
//   		Container: &InferenceComponentContainerSpecificationProperty{
//   			ArtifactUrl: jsii.String("artifactUrl"),
//   			DeployedImage: &DeployedImageProperty{
//   				ResolutionTime: jsii.String("resolutionTime"),
//   				ResolvedImage: jsii.String("resolvedImage"),
//   				SpecifiedImage: jsii.String("specifiedImage"),
//   			},
//   			Environment: map[string]*string{
//   				"environmentKey": jsii.String("environment"),
//   			},
//   			Image: jsii.String("image"),
//   		},
//   		ModelName: jsii.String("modelName"),
//   		StartupParameters: &InferenceComponentStartupParametersProperty{
//   			ContainerStartupHealthCheckTimeoutInSeconds: jsii.Number(123),
//   			ModelDataDownloadTimeoutInSeconds: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	EndpointArn: jsii.String("endpointArn"),
//   	InferenceComponentName: jsii.String("inferenceComponentName"),
//   	RuntimeConfig: &InferenceComponentRuntimeConfigProperty{
//   		CopyCount: jsii.Number(123),
//   		CurrentCopyCount: jsii.Number(123),
//   		DesiredCopyCount: jsii.Number(123),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VariantName: jsii.String("variantName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-inferencecomponent.html
//
type CfnInferenceComponentProps struct {
	// The name of the endpoint that hosts the inference component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-inferencecomponent.html#cfn-sagemaker-inferencecomponent-endpointname
	//
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// The specification for the inference component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-inferencecomponent.html#cfn-sagemaker-inferencecomponent-specification
	//
	Specification interface{} `field:"required" json:"specification" yaml:"specification"`
	// The Amazon Resource Name (ARN) of the endpoint that hosts the inference component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-inferencecomponent.html#cfn-sagemaker-inferencecomponent-endpointarn
	//
	EndpointArn *string `field:"optional" json:"endpointArn" yaml:"endpointArn"`
	// The name of the inference component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-inferencecomponent.html#cfn-sagemaker-inferencecomponent-inferencecomponentname
	//
	InferenceComponentName *string `field:"optional" json:"inferenceComponentName" yaml:"inferenceComponentName"`
	// The runtime config for the inference component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-inferencecomponent.html#cfn-sagemaker-inferencecomponent-runtimeconfig
	//
	RuntimeConfig interface{} `field:"optional" json:"runtimeConfig" yaml:"runtimeConfig"`
	// An array of tags to apply to the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-inferencecomponent.html#cfn-sagemaker-inferencecomponent-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name of the production variant that hosts the inference component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-inferencecomponent.html#cfn-sagemaker-inferencecomponent-variantname
	//
	VariantName *string `field:"optional" json:"variantName" yaml:"variantName"`
}

