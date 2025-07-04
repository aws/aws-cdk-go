package awssagemaker


// The configuration used to run the application image container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerConfigProperty := &ContainerConfigProperty{
//   	ContainerArguments: []*string{
//   		jsii.String("containerArguments"),
//   	},
//   	ContainerEntrypoint: []*string{
//   		jsii.String("containerEntrypoint"),
//   	},
//   	ContainerEnvironmentVariables: []interface{}{
//   		&CustomImageContainerEnvironmentVariableProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-containerconfig.html
//
type CfnAppImageConfig_ContainerConfigProperty struct {
	// The arguments for the container when you're running the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-containerconfig.html#cfn-sagemaker-appimageconfig-containerconfig-containerarguments
	//
	ContainerArguments *[]*string `field:"optional" json:"containerArguments" yaml:"containerArguments"`
	// The entrypoint used to run the application in the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-containerconfig.html#cfn-sagemaker-appimageconfig-containerconfig-containerentrypoint
	//
	ContainerEntrypoint *[]*string `field:"optional" json:"containerEntrypoint" yaml:"containerEntrypoint"`
	// The environment variables to set in the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-containerconfig.html#cfn-sagemaker-appimageconfig-containerconfig-containerenvironmentvariables
	//
	ContainerEnvironmentVariables interface{} `field:"optional" json:"containerEnvironmentVariables" yaml:"containerEnvironmentVariables"`
}

