package awsimagebuilder


// When you create an image or container recipe with Image Builder , you can add the build or test components that your image pipeline uses to create the final image.
//
// You must have at least one build component to create a recipe, but test components are not required. Your pipeline runs tests after it builds the image, to ensure that the target image is functional and can be used reliably for launching Amazon EC2 instances.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageTestsConfigurationProperty := &imageTestsConfigurationProperty{
//   	imageTestsEnabled: jsii.Boolean(false),
//   	timeoutMinutes: jsii.Number(123),
//   }
//
type CfnImagePipeline_ImageTestsConfigurationProperty struct {
	// Defines if tests should be executed when building this image.
	//
	// For example, `true` or `false` .
	ImageTestsEnabled interface{} `field:"optional" json:"imageTestsEnabled" yaml:"imageTestsEnabled"`
	// The maximum time in minutes that tests are permitted to run.
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
}

