package awssagemaker


// Specifies the configuration for notifications of inference results for asynchronous inference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   asyncInferenceNotificationConfigProperty := &asyncInferenceNotificationConfigProperty{
//   	errorTopic: jsii.String("errorTopic"),
//   	successTopic: jsii.String("successTopic"),
//   }
//
type CfnEndpointConfig_AsyncInferenceNotificationConfigProperty struct {
	// Amazon SNS topic to post a notification to when an inference fails.
	//
	// If no topic is provided, no notification is sent on failure.
	ErrorTopic *string `field:"optional" json:"errorTopic" yaml:"errorTopic"`
	// Amazon SNS topic to post a notification to when an inference completes successfully.
	//
	// If no topic is provided, no notification is sent on success.
	SuccessTopic *string `field:"optional" json:"successTopic" yaml:"successTopic"`
}

