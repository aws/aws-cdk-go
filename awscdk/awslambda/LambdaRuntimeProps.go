package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaRuntimeProps := &LambdaRuntimeProps{
//   	BundlingDockerImage: jsii.String("bundlingDockerImage"),
//   	IsVariable: jsii.Boolean(false),
//   	SupportsCodeGuruProfiling: jsii.Boolean(false),
//   	SupportsInlineCode: jsii.Boolean(false),
//   	SupportsSnapStart: jsii.Boolean(false),
//   }
//
type LambdaRuntimeProps struct {
	// The Docker image name to be used for bundling in this runtime.
	// Default: - the latest docker image "amazon/public.ecr.aws/sam/build-<runtime>" from https://gallery.ecr.aws
	//
	BundlingDockerImage *string `field:"optional" json:"bundlingDockerImage" yaml:"bundlingDockerImage"`
	// Whether the runtime enum is meant to change over time, IE NODEJS_LATEST.
	// Default: false.
	//
	IsVariable *bool `field:"optional" json:"isVariable" yaml:"isVariable"`
	// Whether this runtime is integrated with and supported for profiling using Amazon CodeGuru Profiler.
	// Default: false.
	//
	SupportsCodeGuruProfiling *bool `field:"optional" json:"supportsCodeGuruProfiling" yaml:"supportsCodeGuruProfiling"`
	// Whether the ``ZipFile`` (aka inline code) property can be used with this runtime.
	// Default: false.
	//
	SupportsInlineCode *bool `field:"optional" json:"supportsInlineCode" yaml:"supportsInlineCode"`
	// Whether this runtime supports SnapStart.
	// Default: false.
	//
	SupportsSnapStart *bool `field:"optional" json:"supportsSnapStart" yaml:"supportsSnapStart"`
}

