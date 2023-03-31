package awssagemaker


// Specifies configuration for how an endpoint performs asynchronous inference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   asyncInferenceConfigProperty := &asyncInferenceConfigProperty{
//   	outputConfig: &asyncInferenceOutputConfigProperty{
//   		s3OutputPath: jsii.String("s3OutputPath"),
//
//   		// the properties below are optional
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   		notificationConfig: &asyncInferenceNotificationConfigProperty{
//   			errorTopic: jsii.String("errorTopic"),
//   			successTopic: jsii.String("successTopic"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	clientConfig: &asyncInferenceClientConfigProperty{
//   		maxConcurrentInvocationsPerInstance: jsii.Number(123),
//   	},
//   }
//
type CfnEndpointConfig_AsyncInferenceConfigProperty struct {
	// Specifies the configuration for asynchronous inference invocation outputs.
	OutputConfig interface{} `field:"required" json:"outputConfig" yaml:"outputConfig"`
	// Configures the behavior of the client used by SageMaker to interact with the model container during asynchronous inference.
	ClientConfig interface{} `field:"optional" json:"clientConfig" yaml:"clientConfig"`
}

