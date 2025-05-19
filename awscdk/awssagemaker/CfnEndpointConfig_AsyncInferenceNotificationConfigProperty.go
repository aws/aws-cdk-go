package awssagemaker


// Specifies the configuration for notifications of inference results for asynchronous inference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   asyncInferenceNotificationConfigProperty := &AsyncInferenceNotificationConfigProperty{
//   	ErrorTopic: jsii.String("errorTopic"),
//   	IncludeInferenceResponseIn: []*string{
//   		jsii.String("includeInferenceResponseIn"),
//   	},
//   	SuccessTopic: jsii.String("successTopic"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-asyncinferencenotificationconfig.html
//
type CfnEndpointConfig_AsyncInferenceNotificationConfigProperty struct {
	// Amazon SNS topic to post a notification to when an inference fails.
	//
	// If no topic is provided, no notification is sent on failure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-asyncinferencenotificationconfig.html#cfn-sagemaker-endpointconfig-asyncinferencenotificationconfig-errortopic
	//
	ErrorTopic *string `field:"optional" json:"errorTopic" yaml:"errorTopic"`
	// The Amazon SNS topics where you want the inference response to be included.
	//
	// > The inference response is included only if the response size is less than or equal to 128 KB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-asyncinferencenotificationconfig.html#cfn-sagemaker-endpointconfig-asyncinferencenotificationconfig-includeinferenceresponsein
	//
	IncludeInferenceResponseIn *[]*string `field:"optional" json:"includeInferenceResponseIn" yaml:"includeInferenceResponseIn"`
	// Amazon SNS topic to post a notification to when an inference completes successfully.
	//
	// If no topic is provided, no notification is sent on success.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-asyncinferencenotificationconfig.html#cfn-sagemaker-endpointconfig-asyncinferencenotificationconfig-successtopic
	//
	SuccessTopic *string `field:"optional" json:"successTopic" yaml:"successTopic"`
}

