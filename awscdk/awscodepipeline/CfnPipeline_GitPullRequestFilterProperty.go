package awscodepipeline


// The event criteria for the pull request trigger configuration, such as the lists of branches or file paths to include and exclude.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gitPullRequestFilterProperty := &GitPullRequestFilterProperty{
//   	Branches: &GitBranchFilterCriteriaProperty{
//   		Excludes: []*string{
//   			jsii.String("excludes"),
//   		},
//   		Includes: []*string{
//   			jsii.String("includes"),
//   		},
//   	},
//   	Events: []*string{
//   		jsii.String("events"),
//   	},
//   	FilePaths: &GitFilePathFilterCriteriaProperty{
//   		Excludes: []*string{
//   			jsii.String("excludes"),
//   		},
//   		Includes: []*string{
//   			jsii.String("includes"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-gitpullrequestfilter.html
//
type CfnPipeline_GitPullRequestFilterProperty struct {
	// The field that specifies to filter on branches for the pull request trigger configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-gitpullrequestfilter.html#cfn-codepipeline-pipeline-gitpullrequestfilter-branches
	//
	Branches interface{} `field:"optional" json:"branches" yaml:"branches"`
	// The field that specifies which pull request events to filter on (opened, updated, closed) for the trigger configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-gitpullrequestfilter.html#cfn-codepipeline-pipeline-gitpullrequestfilter-events
	//
	Events *[]*string `field:"optional" json:"events" yaml:"events"`
	// The field that specifies to filter on file paths for the pull request trigger configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-gitpullrequestfilter.html#cfn-codepipeline-pipeline-gitpullrequestfilter-filepaths
	//
	FilePaths interface{} `field:"optional" json:"filePaths" yaml:"filePaths"`
}

