package awssagemaker


// Details about the resources to deploy with this inference component, including the model, container, and compute resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inferenceComponentSpecificationProperty := &InferenceComponentSpecificationProperty{
//   	BaseInferenceComponentName: jsii.String("baseInferenceComponentName"),
//   	ComputeResourceRequirements: &InferenceComponentComputeResourceRequirementsProperty{
//   		MaxMemoryRequiredInMb: jsii.Number(123),
//   		MinMemoryRequiredInMb: jsii.Number(123),
//   		NumberOfAcceleratorDevicesRequired: jsii.Number(123),
//   		NumberOfCpuCoresRequired: jsii.Number(123),
//   	},
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
	// The name of an existing inference component that is to contain the inference component that you're creating with your request.
	//
	// Specify this parameter only if your request is meant to create an adapter inference component. An adapter inference component contains the path to an adapter model. The purpose of the adapter model is to tailor the inference output of a base foundation model, which is hosted by the base inference component. The adapter inference component uses the compute resources that you assigned to the base inference component.
	//
	// When you create an adapter inference component, use the `Container` parameter to specify the location of the adapter artifacts. In the parameter value, use the `ArtifactUrl` parameter of the `InferenceComponentContainerSpecification` data type.
	//
	// Before you can create an adapter inference component, you must have an existing inference component that contains the foundation model that you want to adapt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentspecification.html#cfn-sagemaker-inferencecomponent-inferencecomponentspecification-baseinferencecomponentname
	//
	BaseInferenceComponentName *string `field:"optional" json:"baseInferenceComponentName" yaml:"baseInferenceComponentName"`
	// The compute resources allocated to run the model, plus any adapter models, that you assign to the inference component.
	//
	// Omit this parameter if your request is meant to create an adapter inference component. An adapter inference component is loaded by a base inference component, and it uses the compute resources of the base inference component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentspecification.html#cfn-sagemaker-inferencecomponent-inferencecomponentspecification-computeresourcerequirements
	//
	ComputeResourceRequirements interface{} `field:"optional" json:"computeResourceRequirements" yaml:"computeResourceRequirements"`
	// Defines a container that provides the runtime environment for a model that you deploy with an inference component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentspecification.html#cfn-sagemaker-inferencecomponent-inferencecomponentspecification-container
	//
	Container interface{} `field:"optional" json:"container" yaml:"container"`
	// The name of an existing SageMaker AI model object in your account that you want to deploy with the inference component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentspecification.html#cfn-sagemaker-inferencecomponent-inferencecomponentspecification-modelname
	//
	ModelName *string `field:"optional" json:"modelName" yaml:"modelName"`
	// Settings that take effect while the model container starts up.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentspecification.html#cfn-sagemaker-inferencecomponent-inferencecomponentspecification-startupparameters
	//
	StartupParameters interface{} `field:"optional" json:"startupParameters" yaml:"startupParameters"`
}

