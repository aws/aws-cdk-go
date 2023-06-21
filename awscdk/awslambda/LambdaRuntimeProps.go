package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaRuntimeProps := &LambdaRuntimeProps{
//   	BundlingDockerImage: jsii.String("bundlingDockerImage"),
//   	SupportsCodeGuruProfiling: jsii.Boolean(false),
//   	SupportsInlineCode: jsii.Boolean(false),
//   }
//
type LambdaRuntimeProps struct {
	// The Docker image name to be used for bundling in this runtime.
	BundlingDockerImage *string `field:"optional" json:"bundlingDockerImage" yaml:"bundlingDockerImage"`
	// Whether this runtime is integrated with and supported for profiling using Amazon CodeGuru Profiler.
	SupportsCodeGuruProfiling *bool `field:"optional" json:"supportsCodeGuruProfiling" yaml:"supportsCodeGuruProfiling"`
	// Whether the ``ZipFile`` (aka inline code) property can be used with this runtime.
	SupportsInlineCode *bool `field:"optional" json:"supportsInlineCode" yaml:"supportsInlineCode"`
}

