package awssagemaker


// Specifies the configuration for asynchronous inference invocation outputs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   asyncInferenceOutputConfigProperty := &AsyncInferenceOutputConfigProperty{
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	NotificationConfig: &AsyncInferenceNotificationConfigProperty{
//   		ErrorTopic: jsii.String("errorTopic"),
//   		IncludeInferenceResponseIn: []*string{
//   			jsii.String("includeInferenceResponseIn"),
//   		},
//   		SuccessTopic: jsii.String("successTopic"),
//   	},
//   	S3FailurePath: jsii.String("s3FailurePath"),
//   	S3OutputPath: jsii.String("s3OutputPath"),
//   }
//
type CfnEndpointConfig_AsyncInferenceOutputConfigProperty struct {
	// The AWS Key Management Service ( AWS KMS) key that Amazon SageMaker uses to encrypt the asynchronous inference output in Amazon S3.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Specifies the configuration for notifications of inference results for asynchronous inference.
	NotificationConfig interface{} `field:"optional" json:"notificationConfig" yaml:"notificationConfig"`
	// `CfnEndpointConfig.AsyncInferenceOutputConfigProperty.S3FailurePath`.
	S3FailurePath *string `field:"optional" json:"s3FailurePath" yaml:"s3FailurePath"`
	// The Amazon S3 location to upload inference responses to.
	S3OutputPath *string `field:"optional" json:"s3OutputPath" yaml:"s3OutputPath"`
}

