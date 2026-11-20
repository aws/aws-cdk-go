package awslambda


// A resolved OCI image platform descriptor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   platformProperty := &PlatformProperty{
//   	Architecture: jsii.String("architecture"),
//   	Digest: jsii.String("digest"),
//   	Os: jsii.String("os"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-runtimeversion-platform.html
//
type CfnRuntimeVersion_PlatformProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-runtimeversion-platform.html#cfn-lambda-runtimeversion-platform-architecture
	//
	Architecture *string `field:"required" json:"architecture" yaml:"architecture"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-runtimeversion-platform.html#cfn-lambda-runtimeversion-platform-digest
	//
	Digest *string `field:"required" json:"digest" yaml:"digest"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-runtimeversion-platform.html#cfn-lambda-runtimeversion-platform-os
	//
	Os *string `field:"required" json:"os" yaml:"os"`
}

