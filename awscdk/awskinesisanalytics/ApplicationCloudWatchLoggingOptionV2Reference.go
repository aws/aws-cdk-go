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
// Deprecated: use `aws-kinesisanalyticsv2` instead.
type ApplicationCloudWatchLoggingOptionV2Reference struct {
	// The Id of the ApplicationCloudWatchLoggingOption resource.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	ApplicationCloudWatchLoggingOptionId *string `field:"required" json:"applicationCloudWatchLoggingOptionId" yaml:"applicationCloudWatchLoggingOptionId"`
}

