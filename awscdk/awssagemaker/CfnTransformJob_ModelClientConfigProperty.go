package awssagemaker


// Configures the timeout and maximum number of retries for processing a transform job invocation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelClientConfigProperty := &ModelClientConfigProperty{
//   	InvocationsMaxRetries: jsii.Number(123),
//   	InvocationsTimeoutInSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-modelclientconfig.html
//
type CfnTransformJob_ModelClientConfigProperty struct {
	// The maximum number of retries when invocation requests are failing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-modelclientconfig.html#cfn-sagemaker-transformjob-modelclientconfig-invocationsmaxretries
	//
	InvocationsMaxRetries *float64 `field:"optional" json:"invocationsMaxRetries" yaml:"invocationsMaxRetries"`
	// The timeout value in seconds for an invocation request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-modelclientconfig.html#cfn-sagemaker-transformjob-modelclientconfig-invocationstimeoutinseconds
	//
	InvocationsTimeoutInSeconds *float64 `field:"optional" json:"invocationsTimeoutInSeconds" yaml:"invocationsTimeoutInSeconds"`
}

