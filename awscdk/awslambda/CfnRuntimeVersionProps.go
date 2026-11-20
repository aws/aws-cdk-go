package awslambda


// Properties for defining a `CfnRuntimeVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRuntimeVersionProps := &CfnRuntimeVersionProps{
//   	Name: jsii.String("name"),
//   	RuntimeName: jsii.String("runtimeName"),
//
//   	// the properties below are optional
//   	ImageUri: jsii.String("imageUri"),
//   	ReleaseNotes: jsii.String("releaseNotes"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-runtimeversion.html
//
type CfnRuntimeVersionProps struct {
	// The version identifier (1-128 chars, alphanumeric plus _.-).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-runtimeversion.html#cfn-lambda-runtimeversion-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the parent runtime.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-runtimeversion.html#cfn-lambda-runtimeversion-runtimename
	//
	RuntimeName *string `field:"required" json:"runtimeName" yaml:"runtimeName"`
	// The URI of the OCI image to publish as this runtime version.
	//
	// Not returned in Read.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-runtimeversion.html#cfn-lambda-runtimeversion-imageuri
	//
	ImageUri *string `field:"optional" json:"imageUri" yaml:"imageUri"`
	// Optional release notes for this runtime version.
	//
	// Updatable; removing it from the template clears the release notes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-runtimeversion.html#cfn-lambda-runtimeversion-releasenotes
	//
	ReleaseNotes *string `field:"optional" json:"releaseNotes" yaml:"releaseNotes"`
}

