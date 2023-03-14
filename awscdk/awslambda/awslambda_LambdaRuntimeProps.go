package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaRuntimeProps := &lambdaRuntimeProps{
//   	bundlingDockerImage: jsii.String("bundlingDockerImage"),
//   	supportsCodeGuruProfiling: jsii.Boolean(false),
//   	supportsInlineCode: jsii.Boolean(false),
//   }
//
// Experimental.
type LambdaRuntimeProps struct {
	// The Docker image name to be used for bundling in this runtime.
	// Experimental.
	BundlingDockerImage *string `field:"optional" json:"bundlingDockerImage" yaml:"bundlingDockerImage"`
	// Whether this runtime is integrated with and supported for profiling using Amazon CodeGuru Profiler.
	// Experimental.
	SupportsCodeGuruProfiling *bool `field:"optional" json:"supportsCodeGuruProfiling" yaml:"supportsCodeGuruProfiling"`
	// Whether the ``ZipFile`` (aka inline code) property can be used with this runtime.
	// Experimental.
	SupportsInlineCode *bool `field:"optional" json:"supportsInlineCode" yaml:"supportsInlineCode"`
}

