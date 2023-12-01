package awssagemaker


// The specification for the inference component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inferenceComponentSpecificationProperty := &InferenceComponentSpecificationProperty{
//   	ComputeResourceRequirements: &InferenceComponentComputeResourceRequirementsProperty{
//   		MaxMemoryRequiredInMb: jsii.Number(123),
//   		MinMemoryRequiredInMb: jsii.Number(123),
//   		NumberOfAcceleratorDevicesRequired: jsii.Number(123),
//   		NumberOfCpuCoresRequired: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	Container: &InferenceComponentContainerSpecificationProperty{
//   		ArtifactUrl: jsii.String("artifactUrl"),
//   		DeployedImage: &DeployedImageProperty{
//   			ResolutionTime: jsii.String("resolutionTime"),
//   			ResolvedImage: jsii.String("resolvedImage"),
//   			SpecifiedImage: jsii.String("specifiedImage"),
//   		},
//   		Environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		Image: jsii.String("image"),
//   	},
//   	ModelName: jsii.String("modelName"),
//   	StartupParameters: &InferenceComponentStartupParametersProperty{
//   		ContainerStartupHealthCheckTimeoutInSeconds: jsii.Number(123),
//   		ModelDataDownloadTimeoutInSeconds: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentspecification.html
//
type CfnInferenceComponent_InferenceComponentSpecificationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentspecification.html#cfn-sagemaker-inferencecomponent-inferencecomponentspecification-computeresourcerequirements
	//
	ComputeResourceRequirements interface{} `field:"required" json:"computeResourceRequirements" yaml:"computeResourceRequirements"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentspecification.html#cfn-sagemaker-inferencecomponent-inferencecomponentspecification-container
	//
	Container interface{} `field:"optional" json:"container" yaml:"container"`
	// The name of the model to use with the inference component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentspecification.html#cfn-sagemaker-inferencecomponent-inferencecomponentspecification-modelname
	//
	ModelName *string `field:"optional" json:"modelName" yaml:"modelName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentspecification.html#cfn-sagemaker-inferencecomponent-inferencecomponentspecification-startupparameters
	//
	StartupParameters interface{} `field:"optional" json:"startupParameters" yaml:"startupParameters"`
}

