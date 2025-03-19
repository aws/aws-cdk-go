package awscodepipeline


// The Git repository branches specified as filter criteria to start the pipeline.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gitBranchFilterCriteriaProperty := &GitBranchFilterCriteriaProperty{
//   	Excludes: []*string{
//   		jsii.String("excludes"),
//   	},
//   	Includes: []*string{
//   		jsii.String("includes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-gitbranchfiltercriteria.html
//
type CfnPipeline_GitBranchFilterCriteriaProperty struct {
	// The list of patterns of Git branches that, when a commit is pushed, are to be excluded from starting the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-gitbranchfiltercriteria.html#cfn-codepipeline-pipeline-gitbranchfiltercriteria-excludes
	//
	Excludes *[]*string `field:"optional" json:"excludes" yaml:"excludes"`
	// The list of patterns of Git branches that, when a commit is pushed, are to be included as criteria that starts the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-gitbranchfiltercriteria.html#cfn-codepipeline-pipeline-gitbranchfiltercriteria-includes
	//
	Includes *[]*string `field:"optional" json:"includes" yaml:"includes"`
}

