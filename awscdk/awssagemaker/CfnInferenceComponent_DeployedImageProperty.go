package awssagemaker


// Gets the Amazon EC2 Container Registry path of the docker image of the model that is hosted in this [ProductionVariant](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ProductionVariant.html) .
//
// If you used the `registry/repository[:tag]` form to specify the image path of the primary container when you created the model hosted in this `ProductionVariant` , the path resolves to a path of the form `registry/repository[@digest]` . A digest is a hash value that identifies a specific version of an image. For information about Amazon ECR paths, see [Pulling an Image](https://docs.aws.amazon.com//AmazonECR/latest/userguide/docker-pull-ecr-image.html) in the *Amazon ECR User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deployedImageProperty := &DeployedImageProperty{
//   	ResolutionTime: jsii.String("resolutionTime"),
//   	ResolvedImage: jsii.String("resolvedImage"),
//   	SpecifiedImage: jsii.String("specifiedImage"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-deployedimage.html
//
type CfnInferenceComponent_DeployedImageProperty struct {
	// The date and time when the image path for the model resolved to the `ResolvedImage`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-deployedimage.html#cfn-sagemaker-inferencecomponent-deployedimage-resolutiontime
	//
	ResolutionTime *string `field:"optional" json:"resolutionTime" yaml:"resolutionTime"`
	// The specific digest path of the image hosted in this `ProductionVariant` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-deployedimage.html#cfn-sagemaker-inferencecomponent-deployedimage-resolvedimage
	//
	ResolvedImage *string `field:"optional" json:"resolvedImage" yaml:"resolvedImage"`
	// The image path you specified when you created the model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-deployedimage.html#cfn-sagemaker-inferencecomponent-deployedimage-specifiedimage
	//
	SpecifiedImage *string `field:"optional" json:"specifiedImage" yaml:"specifiedImage"`
}

