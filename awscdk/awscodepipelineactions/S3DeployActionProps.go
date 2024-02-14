package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Construction properties of the `S3DeployAction S3 deploy Action`.
//
// Example:
//   var sourceAction s3SourceAction
//   var sourceOutput artifact
//   var deployBucket bucket
//
//
//   codepipeline.NewPipeline(this, jsii.String("Pipeline"), &PipelineProps{
//   	Stages: []stageProps{
//   		&stageProps{
//   			StageName: jsii.String("Source"),
//   			Actions: []iAction{
//   				sourceAction,
//   			},
//   		},
//   		&stageProps{
//   			StageName: jsii.String("Deploy"),
//   			Actions: []*iAction{
//   				codepipeline_actions.NewS3DeployAction(&S3DeployActionProps{
//   					ActionName: jsii.String("DeployAction"),
//   					// can reference the variables
//   					ObjectKey: fmt.Sprintf("%v.txt", sourceAction.variables.VersionId),
//   					Input: sourceOutput,
//   					Bucket: deployBucket,
//   				}),
//   			},
//   		},
//   	},
//   })
//
type S3DeployActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Default: 1.
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Default: - a name will be generated, based on the stage and action names,
	// if any of the action's variables were referenced - otherwise,
	// no namespace will be set.
	//
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your `IAction.bind`
	// method in the `ActionBindOptions.role` property.
	// Default: a new Role will be generated.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The Amazon S3 bucket that is the deploy target.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// The input Artifact to deploy to Amazon S3.
	Input awscodepipeline.Artifact `field:"required" json:"input" yaml:"input"`
	// The specified canned ACL to objects deployed to Amazon S3.
	//
	// This overwrites any existing ACL that was applied to the object.
	// Default: - the original object ACL.
	//
	AccessControl awss3.BucketAccessControl `field:"optional" json:"accessControl" yaml:"accessControl"`
	// The caching behavior for requests/responses for objects in the bucket.
	//
	// The final cache control property will be the result of joining all of the provided array elements with a comma
	// (plus a space after the comma).
	// Default: - none, decided by the HTTP client.
	//
	CacheControl *[]CacheControl `field:"optional" json:"cacheControl" yaml:"cacheControl"`
	// The AWS KMS encryption key for the host bucket.
	//
	// The encryptionKey parameter encrypts uploaded artifacts with the provided AWS KMS key.
	// Default: - none.
	//
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Should the deploy action extract the artifact before deploying to Amazon S3.
	// Default: true.
	//
	Extract *bool `field:"optional" json:"extract" yaml:"extract"`
	// The key of the target object.
	//
	// This is required if extract is false.
	ObjectKey *string `field:"optional" json:"objectKey" yaml:"objectKey"`
}

