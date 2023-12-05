package awssagemaker


// Defines a container that provides the runtime environment for a model that you deploy with an inference component.
//
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
	// The Amazon S3 path where the model artifacts, which result from model training, are stored.
	//
	// This path must point to a single gzip compressed tar archive (.tar.gz suffix).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentcontainerspecification.html#cfn-sagemaker-inferencecomponent-inferencecomponentcontainerspecification-artifacturl
	//
	ArtifactUrl *string `field:"optional" json:"artifactUrl" yaml:"artifactUrl"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentcontainerspecification.html#cfn-sagemaker-inferencecomponent-inferencecomponentcontainerspecification-deployedimage
	//
	DeployedImage interface{} `field:"optional" json:"deployedImage" yaml:"deployedImage"`
	// The environment variables to set in the Docker container.
	//
	// Each key and value in the Environment string-to-string map can have length of up to 1024. We support up to 16 entries in the map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentcontainerspecification.html#cfn-sagemaker-inferencecomponent-inferencecomponentcontainerspecification-environment
	//
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// The Amazon Elastic Container Registry (Amazon ECR) path where the Docker image for the model is stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentcontainerspecification.html#cfn-sagemaker-inferencecomponent-inferencecomponentcontainerspecification-image
	//
	Image *string `field:"optional" json:"image" yaml:"image"`
}

