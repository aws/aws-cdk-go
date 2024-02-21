package awscodepipeline


// The event criteria that specify when a specified repository event will start the pipeline for the specified trigger configuration, such as the lists of Git tags to include and exclude.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gitPushFilterProperty := &GitPushFilterProperty{
//   	Branches: &GitBranchFilterCriteriaProperty{
//   		Excludes: []*string{
//   			jsii.String("excludes"),
//   		},
//   		Includes: []*string{
//   			jsii.String("includes"),
//   		},
//   	},
//   	FilePaths: &GitFilePathFilterCriteriaProperty{
//   		Excludes: []*string{
//   			jsii.String("excludes"),
//   		},
//   		Includes: []*string{
//   			jsii.String("includes"),
//   		},
//   	},
//   	Tags: &GitTagFilterCriteriaProperty{
//   		Excludes: []*string{
//   			jsii.String("excludes"),
//   		},
//   		Includes: []*string{
//   			jsii.String("includes"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-gitpushfilter.html
//
type CfnPipeline_GitPushFilterProperty struct {
	// The field that specifies to filter on branches for the push trigger configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-gitpushfilter.html#cfn-codepipeline-pipeline-gitpushfilter-branches
	//
	Branches interface{} `field:"optional" json:"branches" yaml:"branches"`
	// The field that specifies to filter on file paths for the push trigger configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-gitpushfilter.html#cfn-codepipeline-pipeline-gitpushfilter-filepaths
	//
	FilePaths interface{} `field:"optional" json:"filePaths" yaml:"filePaths"`
	// The field that contains the details for the Git tags trigger configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-gitpushfilter.html#cfn-codepipeline-pipeline-gitpushfilter-tags
	//
	Tags *CfnPipeline_GitTagFilterCriteriaProperty `field:"optional" json:"tags" yaml:"tags"`
}

