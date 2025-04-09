package awscodepipeline


// The Git tags specified as filter criteria for whether a Git tag repository event will start the pipeline.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gitTagFilterCriteriaProperty := &GitTagFilterCriteriaProperty{
//   	Excludes: []*string{
//   		jsii.String("excludes"),
//   	},
//   	Includes: []*string{
//   		jsii.String("includes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-gittagfiltercriteria.html
//
type CfnPipeline_GitTagFilterCriteriaProperty struct {
	// The list of patterns of Git tags that, when pushed, are to be excluded from starting the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-gittagfiltercriteria.html#cfn-codepipeline-pipeline-gittagfiltercriteria-excludes
	//
	Excludes *[]*string `field:"optional" json:"excludes" yaml:"excludes"`
	// The list of patterns of Git tags that, when pushed, are to be included as criteria that starts the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-gittagfiltercriteria.html#cfn-codepipeline-pipeline-gittagfiltercriteria-includes
	//
	Includes *[]*string `field:"optional" json:"includes" yaml:"includes"`
}

