package awscodepipeline


// The Git repository file paths specified as filter criteria to start the pipeline.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gitFilePathFilterCriteriaProperty := &GitFilePathFilterCriteriaProperty{
//   	Excludes: []*string{
//   		jsii.String("excludes"),
//   	},
//   	Includes: []*string{
//   		jsii.String("includes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-gitfilepathfiltercriteria.html
//
type CfnPipeline_GitFilePathFilterCriteriaProperty struct {
	// The list of patterns of Git repository file paths that, when a commit is pushed, are to be excluded from starting the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-gitfilepathfiltercriteria.html#cfn-codepipeline-pipeline-gitfilepathfiltercriteria-excludes
	//
	Excludes *[]*string `field:"optional" json:"excludes" yaml:"excludes"`
	// The list of patterns of Git repository file paths that, when a commit is pushed, are to be included as criteria that starts the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-gitfilepathfiltercriteria.html#cfn-codepipeline-pipeline-gitfilepathfiltercriteria-includes
	//
	Includes *[]*string `field:"optional" json:"includes" yaml:"includes"`
}

