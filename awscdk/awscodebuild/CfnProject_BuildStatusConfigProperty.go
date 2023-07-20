package awscodebuild


// Contains information that defines how the AWS CodeBuild build project reports the build status to the source provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   buildStatusConfigProperty := &BuildStatusConfigProperty{
//   	Context: jsii.String("context"),
//   	TargetUrl: jsii.String("targetUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-buildstatusconfig.html
//
type CfnProject_BuildStatusConfigProperty struct {
	// Specifies the context of the build status CodeBuild sends to the source provider.
	//
	// The usage of this parameter depends on the source provider.
	//
	// - **Bitbucket** - This parameter is used for the `name` parameter in the Bitbucket commit status. For more information, see [build](https://docs.aws.amazon.com/https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Bworkspace%7D/%7Brepo_slug%7D/commit/%7Bnode%7D/statuses/build) in the Bitbucket API documentation.
	// - **GitHub/GitHub Enterprise Server** - This parameter is used for the `context` parameter in the GitHub commit status. For more information, see [Create a commit status](https://docs.aws.amazon.com/https://developer.github.com/v3/repos/statuses/#create-a-commit-status) in the GitHub developer guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-buildstatusconfig.html#cfn-codebuild-project-buildstatusconfig-context
	//
	Context *string `field:"optional" json:"context" yaml:"context"`
	// Specifies the target url of the build status CodeBuild sends to the source provider.
	//
	// The usage of this parameter depends on the source provider.
	//
	// - **Bitbucket** - This parameter is used for the `url` parameter in the Bitbucket commit status. For more information, see [build](https://docs.aws.amazon.com/https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Bworkspace%7D/%7Brepo_slug%7D/commit/%7Bnode%7D/statuses/build) in the Bitbucket API documentation.
	// - **GitHub/GitHub Enterprise Server** - This parameter is used for the `target_url` parameter in the GitHub commit status. For more information, see [Create a commit status](https://docs.aws.amazon.com/https://developer.github.com/v3/repos/statuses/#create-a-commit-status) in the GitHub developer guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-buildstatusconfig.html#cfn-codebuild-project-buildstatusconfig-targeturl
	//
	TargetUrl *string `field:"optional" json:"targetUrl" yaml:"targetUrl"`
}

