package awscodepipeline


// Represents information about the output of an action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputArtifactProperty := &OutputArtifactProperty{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Files: []*string{
//   		jsii.String("files"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-outputartifact.html
//
type CfnPipeline_OutputArtifactProperty struct {
	// The name of the output of an artifact, such as "My App".
	//
	// The output artifact name must exactly match the input artifact declared for a downstream action. However, the downstream action's input artifact does not have to be the next action in strict sequence from the action that provided the output artifact. Actions in parallel can declare different output artifacts, which are in turn consumed by different following actions.
	//
	// Output artifact names must be unique within a pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-outputartifact.html#cfn-codepipeline-pipeline-outputartifact-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The files that you want to associate with the output artifact that will be exported from the compute action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-outputartifact.html#cfn-codepipeline-pipeline-outputartifact-files
	//
	Files *[]*string `field:"optional" json:"files" yaml:"files"`
}

