package awscodepipelineactions


// A list that allows you to specify, or override, the source revision for a pipeline execution that's being started.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceRevision := &SourceRevision{
//   	ActionName: jsii.String("actionName"),
//   	RevisionType: awscdk.Aws_codepipeline_actions.RevisionType_COMMIT_ID,
//   	RevisionValue: jsii.String("revisionValue"),
//   }
//
type SourceRevision struct {
	// The name of the action where the override will be applied.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The type of source revision, based on the source provider.
	RevisionType RevisionType `field:"required" json:"revisionType" yaml:"revisionType"`
	// The source revision, or version of your source artifact, with the changes that you want to run in the pipeline execution.
	RevisionValue *string `field:"required" json:"revisionValue" yaml:"revisionValue"`
}

