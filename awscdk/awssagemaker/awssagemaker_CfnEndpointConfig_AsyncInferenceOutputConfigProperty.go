package awssagemaker


// Specifies the configuration for asynchronous inference invocation outputs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   asyncInferenceOutputConfigProperty := &asyncInferenceOutputConfigProperty{
//   	s3OutputPath: jsii.String("s3OutputPath"),
//
//   	// the properties below are optional
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	notificationConfig: &asyncInferenceNotificationConfigProperty{
//   		errorTopic: jsii.String("errorTopic"),
//   		successTopic: jsii.String("successTopic"),
//   	},
//   }
//
type CfnEndpointConfig_AsyncInferenceOutputConfigProperty struct {
	// The Amazon S3 location to upload inference responses to.
	S3OutputPath *string `field:"required" json:"s3OutputPath" yaml:"s3OutputPath"`
	// The AWS Key Management Service ( AWS KMS) key that Amazon SageMaker uses to encrypt the asynchronous inference output in Amazon S3.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Specifies the configuration for notifications of inference results for asynchronous inference.
	NotificationConfig interface{} `field:"optional" json:"notificationConfig" yaml:"notificationConfig"`
}

