package awscodebuild


// `SourceAuth` is a property of the [AWS CodeBuild Project Source](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-source.html) property type that specifies authorization settings for AWS CodeBuild to access the source code to be built.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceAuthProperty := &SourceAuthProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Resource: jsii.String("resource"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-sourceauth.html
//
type CfnProject_SourceAuthProperty struct {
	// The authorization type to use.
	//
	// Valid options are OAUTH, CODECONNECTIONS, or SECRETS_MANAGER.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-sourceauth.html#cfn-codebuild-project-sourceauth-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The resource value that applies to the specified authorization type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-sourceauth.html#cfn-codebuild-project-sourceauth-resource
	//
	Resource *string `field:"optional" json:"resource" yaml:"resource"`
}

