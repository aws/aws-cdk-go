package interfacesawskinesisanalyticsv2


// A reference to a ApplicationCloudWatchLoggingOption resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationCloudWatchLoggingOptionReference := &ApplicationCloudWatchLoggingOptionReference{
//   	ApplicationCloudWatchLoggingOptionId: jsii.String("applicationCloudWatchLoggingOptionId"),
//   }
//
type ApplicationCloudWatchLoggingOptionReference struct {
	// The Id of the ApplicationCloudWatchLoggingOption resource.
	ApplicationCloudWatchLoggingOptionId *string `field:"required" json:"applicationCloudWatchLoggingOptionId" yaml:"applicationCloudWatchLoggingOptionId"`
}

