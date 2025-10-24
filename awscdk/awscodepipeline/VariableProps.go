package awscodepipeline


// Properties of pipeline-level variable.
//
// Example:
//   var sourceAction S3SourceAction
//   var sourceOutput Artifact
//   var deployBucket Bucket
//
//
//   // Pipeline-level variable
//   variable := codepipeline.NewVariable(&VariableProps{
//   	VariableName: jsii.String("bucket-var"),
//   	Description: jsii.String("description"),
//   	DefaultValue: jsii.String("sample"),
//   })
//
//   codepipeline.NewPipeline(this, jsii.String("Pipeline"), &PipelineProps{
//   	PipelineType: codepipeline.PipelineType_V2,
//   	Variables: []Variable{
//   		variable,
//   	},
//   	Stages: []StageProps{
//   		&StageProps{
//   			StageName: jsii.String("Source"),
//   			Actions: []IAction{
//   				sourceAction,
//   			},
//   		},
//   		&StageProps{
//   			StageName: jsii.String("Deploy"),
//   			Actions: []IAction{
//   				codepipeline_actions.NewS3DeployAction(&S3DeployActionProps{
//   					ActionName: jsii.String("DeployAction"),
//   					// can reference the variables
//   					ObjectKey: fmt.Sprintf("%v.txt", variable.Reference()),
//   					Input: sourceOutput,
//   					Bucket: deployBucket,
//   				}),
//   			},
//   		},
//   	},
//   })
//
type VariableProps struct {
	// The name of a pipeline-level variable.
	VariableName *string `field:"required" json:"variableName" yaml:"variableName"`
	// The default value of a pipeline-level variable.
	// Default: - No default value.
	//
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// The description of a pipeline-level variable.
	//
	// It's used to add additional context
	// about the variable, and not being used at time when pipeline executes.
	// Default: - No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

