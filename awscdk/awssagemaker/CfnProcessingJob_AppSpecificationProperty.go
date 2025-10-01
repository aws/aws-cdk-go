package awssagemaker


// Configuration to run a processing job in a specified container image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appSpecificationProperty := &AppSpecificationProperty{
//   	ImageUri: jsii.String("imageUri"),
//
//   	// the properties below are optional
//   	ContainerArguments: []*string{
//   		jsii.String("containerArguments"),
//   	},
//   	ContainerEntrypoint: []*string{
//   		jsii.String("containerEntrypoint"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-appspecification.html
//
type CfnProcessingJob_AppSpecificationProperty struct {
	// The container image to be run by the processing job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-appspecification.html#cfn-sagemaker-processingjob-appspecification-imageuri
	//
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
	// The arguments for a container used to run a processing job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-appspecification.html#cfn-sagemaker-processingjob-appspecification-containerarguments
	//
	ContainerArguments *[]*string `field:"optional" json:"containerArguments" yaml:"containerArguments"`
	// The entrypoint for a container used to run a processing job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-appspecification.html#cfn-sagemaker-processingjob-appspecification-containerentrypoint
	//
	ContainerEntrypoint *[]*string `field:"optional" json:"containerEntrypoint" yaml:"containerEntrypoint"`
}

