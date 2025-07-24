package awscodepipeline

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Pipeline-Level variable.
//
// Example:
//   var sourceAction s3SourceAction
//   var sourceOutput artifact
//   var deployBucket bucket
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
//   	Variables: []variable{
//   		variable,
//   	},
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
//   					ObjectKey: fmt.Sprintf("%v.txt", variable.Reference()),
//   					Input: sourceOutput,
//   					Bucket: deployBucket,
//   				}),
//   			},
//   		},
//   	},
//   })
//
type Variable interface {
	// The name of a pipeline-level variable.
	VariableName() *string
	// Reference the variable name at Pipeline actions.
	//
	// Returns: The variable name in a format that can be referenced at Pipeline actions.
	Reference() *string
}

// The jsii proxy struct for Variable
type jsiiProxy_Variable struct {
	_ byte // padding
}

func (j *jsiiProxy_Variable) VariableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"variableName",
		&returns,
	)
	return returns
}


func NewVariable(props *VariableProps) Variable {
	_init_.Initialize()

	if err := validateNewVariableParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Variable{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline.Variable",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewVariable_Override(v Variable, props *VariableProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline.Variable",
		[]interface{}{props},
		v,
	)
}

func (v *jsiiProxy_Variable) Reference() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"reference",
		nil, // no parameters
		&returns,
	)

	return returns
}

