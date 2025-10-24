package awscodepipelineactions


// The types of revision for a pipeline execution.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   targetPipeline := codepipeline.Pipeline_FromPipelineArn(this, jsii.String("Pipeline"), jsii.String("arn:aws:codepipeline:us-east-1:123456789012:InvokePipelineAction")) // If targetPipeline is not created by cdk, import from arn.
//   pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("stageName"),
//   	Actions: []IAction{
//   		cpactions.NewPipelineInvokeAction(&PipelineInvokeActionProps{
//   			ActionName: jsii.String("Invoke"),
//   			TargetPipeline: *TargetPipeline,
//   			Variables: []Variable{
//   				&Variable{
//   					Name: jsii.String("name1"),
//   					Value: jsii.String("value1"),
//   				},
//   			},
//   			SourceRevisions: []SourceRevision{
//   				&SourceRevision{
//   					ActionName: jsii.String("Source"),
//   					RevisionType: cpactions.RevisionType_S3_OBJECT_VERSION_ID,
//   					RevisionValue: jsii.String("testRevisionValue"),
//   				},
//   			},
//   		}),
//   	},
//   })
//
type RevisionType string

const (
	// The revision type is a commit id.
	RevisionType_COMMIT_ID RevisionType = "COMMIT_ID"
	// The revision type is an image digest.
	RevisionType_IMAGE_DIGEST RevisionType = "IMAGE_DIGEST"
	// The revision type is an s3 object version id.
	RevisionType_S3_OBJECT_VERSION_ID RevisionType = "S3_OBJECT_VERSION_ID"
	// The revision type is an s3 object version key.
	RevisionType_S3_OBJECT_KEY RevisionType = "S3_OBJECT_KEY"
)

