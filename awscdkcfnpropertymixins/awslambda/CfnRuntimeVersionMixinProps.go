package awslambda


// Properties for CfnRuntimeVersionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnRuntimeVersionMixinProps := &CfnRuntimeVersionMixinProps{
//   	ImageUri: jsii.String("imageUri"),
//   	Name: jsii.String("name"),
//   	ReleaseNotes: jsii.String("releaseNotes"),
//   	RuntimeName: jsii.String("runtimeName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-runtimeversion.html
//
type CfnRuntimeVersionMixinProps struct {
	// The URI of the OCI image to publish as this runtime version.
	//
	// Not returned in Read.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-runtimeversion.html#cfn-lambda-runtimeversion-imageuri
	//
	ImageUri *string `field:"optional" json:"imageUri" yaml:"imageUri"`
	// The version identifier (1-128 chars, alphanumeric plus _.-).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-runtimeversion.html#cfn-lambda-runtimeversion-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Optional release notes for this runtime version.
	//
	// Updatable; removing it from the template clears the release notes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-runtimeversion.html#cfn-lambda-runtimeversion-releasenotes
	//
	ReleaseNotes *string `field:"optional" json:"releaseNotes" yaml:"releaseNotes"`
	// The name of the parent runtime.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-runtimeversion.html#cfn-lambda-runtimeversion-runtimename
	//
	RuntimeName *string `field:"optional" json:"runtimeName" yaml:"runtimeName"`
}

