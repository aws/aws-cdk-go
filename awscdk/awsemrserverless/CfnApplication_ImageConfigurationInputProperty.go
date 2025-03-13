package awsemrserverless


// The image configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageConfigurationInputProperty := &ImageConfigurationInputProperty{
//   	ImageUri: jsii.String("imageUri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-imageconfigurationinput.html
//
type CfnApplication_ImageConfigurationInputProperty struct {
	// The URI of an image in the Amazon ECR registry.
	//
	// This field is required when you create a new application. If you leave this field blank in an update, Amazon EMR will remove the image configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-imageconfigurationinput.html#cfn-emrserverless-application-imageconfigurationinput-imageuri
	//
	ImageUri *string `field:"optional" json:"imageUri" yaml:"imageUri"`
}

