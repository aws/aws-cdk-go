package awskinesisanalytics


// A reference to a ApplicationCloudWatchLoggingOption resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationCloudWatchLoggingOptionV2Reference := &ApplicationCloudWatchLoggingOptionV2Reference{
//   	ApplicationCloudWatchLoggingOptionId: jsii.String("applicationCloudWatchLoggingOptionId"),
//   }
//
type ApplicationCloudWatchLoggingOptionV2Reference struct {
	// The Id of the ApplicationCloudWatchLoggingOption resource.
	ApplicationCloudWatchLoggingOptionId *string `field:"required" json:"applicationCloudWatchLoggingOptionId" yaml:"applicationCloudWatchLoggingOptionId"`
}

