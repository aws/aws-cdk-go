package mixinsawscodebuild


// `SourceAuth` is a property of the [AWS CodeBuild Project Source](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-source.html) property type that specifies authorization settings for AWS CodeBuild to access the source code to be built.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sourceAuthProperty := &SourceAuthProperty{
//   	Resource: jsii.String("resource"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-sourceauth.html
//
type CfnProjectPropsMixin_SourceAuthProperty struct {
	// The resource value that applies to the specified authorization type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-sourceauth.html#cfn-codebuild-project-sourceauth-resource
	//
	Resource *string `field:"optional" json:"resource" yaml:"resource"`
	// The authorization type to use.
	//
	// Valid options are OAUTH, CODECONNECTIONS, or SECRETS_MANAGER.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-sourceauth.html#cfn-codebuild-project-sourceauth-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

