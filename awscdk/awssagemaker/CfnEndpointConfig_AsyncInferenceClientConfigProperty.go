package awssagemaker


// Configures the behavior of the client used by SageMaker to interact with the model container during asynchronous inference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   asyncInferenceClientConfigProperty := &AsyncInferenceClientConfigProperty{
//   	MaxConcurrentInvocationsPerInstance: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-asyncinferenceclientconfig.html
//
type CfnEndpointConfig_AsyncInferenceClientConfigProperty struct {
	// The maximum number of concurrent requests sent by the SageMaker client to the model container.
	//
	// If no value is provided, SageMaker will choose an optimal value for you.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-asyncinferenceclientconfig.html#cfn-sagemaker-endpointconfig-asyncinferenceclientconfig-maxconcurrentinvocationsperinstance
	//
	MaxConcurrentInvocationsPerInstance *float64 `field:"optional" json:"maxConcurrentInvocationsPerInstance" yaml:"maxConcurrentInvocationsPerInstance"`
}

