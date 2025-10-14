package awsimagebuilder


// The logging configuration that's defined for the image.
//
// Image Builder uses the defined settings to direct execution log output during image creation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageLoggingConfigurationProperty := &ImageLoggingConfigurationProperty{
//   	LogGroupName: jsii.String("logGroupName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-imageloggingconfiguration.html
//
type CfnImage_ImageLoggingConfigurationProperty struct {
	// The log group name that Image Builder uses for image creation.
	//
	// If not specified, the log group name defaults to `/aws/imagebuilder/image-name` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-imageloggingconfiguration.html#cfn-imagebuilder-image-imageloggingconfiguration-loggroupname
	//
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
}

