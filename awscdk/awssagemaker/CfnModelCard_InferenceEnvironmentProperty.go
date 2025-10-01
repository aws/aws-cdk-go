package awssagemaker


// An overview of a model's inference environment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inferenceEnvironmentProperty := &InferenceEnvironmentProperty{
//   	ContainerImage: []*string{
//   		jsii.String("containerImage"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-inferenceenvironment.html
//
type CfnModelCard_InferenceEnvironmentProperty struct {
	// The container used to run the inference environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-inferenceenvironment.html#cfn-sagemaker-modelcard-inferenceenvironment-containerimage
	//
	ContainerImage *[]*string `field:"optional" json:"containerImage" yaml:"containerImage"`
}

