package previewawssagemakermixins


// Specifies configuration for how an endpoint performs asynchronous inference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   asyncInferenceConfigProperty := &AsyncInferenceConfigProperty{
//   	ClientConfig: &AsyncInferenceClientConfigProperty{
//   		MaxConcurrentInvocationsPerInstance: jsii.Number(123),
//   	},
//   	OutputConfig: &AsyncInferenceOutputConfigProperty{
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   		NotificationConfig: &AsyncInferenceNotificationConfigProperty{
//   			ErrorTopic: jsii.String("errorTopic"),
//   			IncludeInferenceResponseIn: []*string{
//   				jsii.String("includeInferenceResponseIn"),
//   			},
//   			SuccessTopic: jsii.String("successTopic"),
//   		},
//   		S3FailurePath: jsii.String("s3FailurePath"),
//   		S3OutputPath: jsii.String("s3OutputPath"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-asyncinferenceconfig.html
//
type CfnEndpointConfigPropsMixin_AsyncInferenceConfigProperty struct {
	// Configures the behavior of the client used by SageMaker to interact with the model container during asynchronous inference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-asyncinferenceconfig.html#cfn-sagemaker-endpointconfig-asyncinferenceconfig-clientconfig
	//
	ClientConfig interface{} `field:"optional" json:"clientConfig" yaml:"clientConfig"`
	// Specifies the configuration for asynchronous inference invocation outputs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-asyncinferenceconfig.html#cfn-sagemaker-endpointconfig-asyncinferenceconfig-outputconfig
	//
	OutputConfig interface{} `field:"optional" json:"outputConfig" yaml:"outputConfig"`
}

