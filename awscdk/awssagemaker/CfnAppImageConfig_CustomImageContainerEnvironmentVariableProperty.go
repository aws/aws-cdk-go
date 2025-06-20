package awssagemaker


// The environment variables to set in the container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customImageContainerEnvironmentVariableProperty := &CustomImageContainerEnvironmentVariableProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-customimagecontainerenvironmentvariable.html
//
type CfnAppImageConfig_CustomImageContainerEnvironmentVariableProperty struct {
	// The key that identifies a container environment variable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-customimagecontainerenvironmentvariable.html#cfn-sagemaker-appimageconfig-customimagecontainerenvironmentvariable-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value of the container environment variable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-customimagecontainerenvironmentvariable.html#cfn-sagemaker-appimageconfig-customimagecontainerenvironmentvariable-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

