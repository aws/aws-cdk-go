package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inferenceComponentContainerSpecificationProperty := &InferenceComponentContainerSpecificationProperty{
//   	ArtifactUrl: jsii.String("artifactUrl"),
//   	DeployedImage: &DeployedImageProperty{
//   		ResolutionTime: jsii.String("resolutionTime"),
//   		ResolvedImage: jsii.String("resolvedImage"),
//   		SpecifiedImage: jsii.String("specifiedImage"),
//   	},
//   	Environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	Image: jsii.String("image"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentcontainerspecification.html
//
type CfnInferenceComponent_InferenceComponentContainerSpecificationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentcontainerspecification.html#cfn-sagemaker-inferencecomponent-inferencecomponentcontainerspecification-artifacturl
	//
	ArtifactUrl *string `field:"optional" json:"artifactUrl" yaml:"artifactUrl"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentcontainerspecification.html#cfn-sagemaker-inferencecomponent-inferencecomponentcontainerspecification-deployedimage
	//
	DeployedImage interface{} `field:"optional" json:"deployedImage" yaml:"deployedImage"`
	// Environment variables to specify on the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentcontainerspecification.html#cfn-sagemaker-inferencecomponent-inferencecomponentcontainerspecification-environment
	//
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// The image to use for the container that will be materialized for the inference component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentcontainerspecification.html#cfn-sagemaker-inferencecomponent-inferencecomponentcontainerspecification-image
	//
	Image *string `field:"optional" json:"image" yaml:"image"`
}

