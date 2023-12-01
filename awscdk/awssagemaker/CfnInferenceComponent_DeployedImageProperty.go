package awssagemaker


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-deployedimage.html#cfn-sagemaker-inferencecomponent-deployedimage-resolutiontime
	//
	ResolutionTime *string `field:"optional" json:"resolutionTime" yaml:"resolutionTime"`
	// The image to use for the container that will be materialized for the inference component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-deployedimage.html#cfn-sagemaker-inferencecomponent-deployedimage-resolvedimage
	//
	ResolvedImage *string `field:"optional" json:"resolvedImage" yaml:"resolvedImage"`
	// The image to use for the container that will be materialized for the inference component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-deployedimage.html#cfn-sagemaker-inferencecomponent-deployedimage-specifiedimage
	//
	SpecifiedImage *string `field:"optional" json:"specifiedImage" yaml:"specifiedImage"`
}

