package awssagemaker


// Specifies the serverless configuration for an endpoint variant.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverlessConfigProperty := &ServerlessConfigProperty{
//   	MaxConcurrency: jsii.Number(123),
//   	MemorySizeInMb: jsii.Number(123),
//
//   	// the properties below are optional
//   	ProvisionedConcurrency: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-serverlessconfig.html
//
type CfnEndpointConfig_ServerlessConfigProperty struct {
	// The maximum number of concurrent invocations your serverless endpoint can process.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-serverlessconfig.html#cfn-sagemaker-endpointconfig-serverlessconfig-maxconcurrency
	//
	MaxConcurrency *float64 `field:"required" json:"maxConcurrency" yaml:"maxConcurrency"`
	// The memory size of your serverless endpoint.
	//
	// Valid values are in 1 GB increments: 1024 MB, 2048 MB, 3072 MB, 4096 MB, 5120 MB, or 6144 MB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-serverlessconfig.html#cfn-sagemaker-endpointconfig-serverlessconfig-memorysizeinmb
	//
	MemorySizeInMb *float64 `field:"required" json:"memorySizeInMb" yaml:"memorySizeInMb"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-serverlessconfig.html#cfn-sagemaker-endpointconfig-serverlessconfig-provisionedconcurrency
	//
	ProvisionedConcurrency *float64 `field:"optional" json:"provisionedConcurrency" yaml:"provisionedConcurrency"`
}

