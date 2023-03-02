package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

// Construction properties of the {@link S3DeployAction S3 deploy Action}.
//
// Example:
//   sourceOutput := codepipeline.NewArtifact()
//   targetBucket := s3.NewBucket(this, jsii.String("MyBucket"))
//
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   deployAction := codepipeline_actions.NewS3DeployAction(&s3DeployActionProps{
//   	actionName: jsii.String("S3Deploy"),
//   	bucket: targetBucket,
//   	input: sourceOutput,
//   })
//   deployStage := pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("Deploy"),
//   	actions: []iAction{
//   		deployAction,
//   	},
//   })
//
// Experimental.
type S3DeployActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	// Experimental.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Experimental.
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Experimental.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The Amazon S3 bucket that is the deploy target.
	// Experimental.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// The input Artifact to deploy to Amazon S3.
	// Experimental.
	Input awscodepipeline.Artifact `field:"required" json:"input" yaml:"input"`
	// The specified canned ACL to objects deployed to Amazon S3.
	//
	// This overwrites any existing ACL that was applied to the object.
	// Experimental.
	AccessControl awss3.BucketAccessControl `field:"optional" json:"accessControl" yaml:"accessControl"`
	// The caching behavior for requests/responses for objects in the bucket.
	//
	// The final cache control property will be the result of joining all of the provided array elements with a comma
	// (plus a space after the comma).
	// Experimental.
	CacheControl *[]CacheControl `field:"optional" json:"cacheControl" yaml:"cacheControl"`
	// Should the deploy action extract the artifact before deploying to Amazon S3.
	// Experimental.
	Extract *bool `field:"optional" json:"extract" yaml:"extract"`
	// The key of the target object.
	//
	// This is required if extract is false.
	// Experimental.
	ObjectKey *string `field:"optional" json:"objectKey" yaml:"objectKey"`
}

