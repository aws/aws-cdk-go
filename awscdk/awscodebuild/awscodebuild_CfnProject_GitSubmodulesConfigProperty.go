package awscodebuild


// `GitSubmodulesConfig` is a property of the [AWS CodeBuild Project Source](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-source.html) property type that specifies information about the Git submodules configuration for the build project.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gitSubmodulesConfigProperty := &gitSubmodulesConfigProperty{
//   	fetchSubmodules: jsii.Boolean(false),
//   }
//
type CfnProject_GitSubmodulesConfigProperty struct {
	// Set to true to fetch Git submodules for your AWS CodeBuild build project.
	FetchSubmodules interface{} `field:"required" json:"fetchSubmodules" yaml:"fetchSubmodules"`
}

