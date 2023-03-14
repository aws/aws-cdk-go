package awsemrserverless


// The image configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageConfigurationInputProperty := &imageConfigurationInputProperty{
//   	imageUri: jsii.String("imageUri"),
//   }
//
type CfnApplication_ImageConfigurationInputProperty struct {
	// The URI of an image in the Amazon ECR registry.
	//
	// This field is required when you create a new application. If you leave this field blank in an update, Amazon EMR will remove the image configuration.
	ImageUri *string `field:"optional" json:"imageUri" yaml:"imageUri"`
}

