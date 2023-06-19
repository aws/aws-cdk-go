package pipelines

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// Deployment of a single `Stage`.
//
// A `Stage` consists of one or more `Stacks`, which will be
// deployed in dependency order.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var stack stack
//   var stage stage
//   var step step
//
//   stageDeployment := awscdk.Pipelines.StageDeployment_FromStage(stage, &StageDeploymentProps{
//   	Post: []*step{
//   		step,
//   	},
//   	Pre: []*step{
//   		step,
//   	},
//   	StackSteps: []stackSteps{
//   		&stackSteps{
//   			Stack: stack,
//
//   			// the properties below are optional
//   			ChangeSet: []*step{
//   				step,
//   			},
//   			Post: []*step{
//   				step,
//   			},
//   			Pre: []*step{
//   				step,
//   			},
//   		},
//   	},
//   	StageName: jsii.String("stageName"),
//   })
//
// Experimental.
type StageDeployment interface {
	// Additional steps that are run after all of the stacks in the stage.
	// Experimental.
	Post() *[]Step
	// Additional steps that are run before any of the stacks in the stage.
	// Experimental.
	Pre() *[]Step
	// The stacks deployed in this stage.
	// Experimental.
	Stacks() *[]StackDeployment
	// Instructions for additional steps that are run at stack level.
	// Experimental.
	StackSteps() *[]*StackSteps
	// The display name of this stage.
	// Experimental.
	StageName() *string
	// Add an additional step to run after all of the stacks in this stage.
	// Experimental.
	AddPost(steps ...Step)
	// Add an additional step to run before any of the stacks in this stage.
	// Experimental.
	AddPre(steps ...Step)
}

// The jsii proxy struct for StageDeployment
type jsiiProxy_StageDeployment struct {
	_ byte // padding
}

func (j *jsiiProxy_StageDeployment) Post() *[]Step {
	var returns *[]Step
	_jsii_.Get(
		j,
		"post",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageDeployment) Pre() *[]Step {
	var returns *[]Step
	_jsii_.Get(
		j,
		"pre",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageDeployment) Stacks() *[]StackDeployment {
	var returns *[]StackDeployment
	_jsii_.Get(
		j,
		"stacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageDeployment) StackSteps() *[]*StackSteps {
	var returns *[]*StackSteps
	_jsii_.Get(
		j,
		"stackSteps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageDeployment) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}


// Create a new `StageDeployment` from a `Stage`.
//
// Synthesizes the target stage, and deployes the stacks found inside
// in dependency order.
// Experimental.
func StageDeployment_FromStage(stage awscdk.Stage, props *StageDeploymentProps) StageDeployment {
	_init_.Initialize()

	if err := validateStageDeployment_FromStageParameters(stage, props); err != nil {
		panic(err)
	}
	var returns StageDeployment

	_jsii_.StaticInvoke(
		"monocdk.pipelines.StageDeployment",
		"fromStage",
		[]interface{}{stage, props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StageDeployment) AddPost(steps ...Step) {
	args := []interface{}{}
	for _, a := range steps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		s,
		"addPost",
		args,
	)
}

func (s *jsiiProxy_StageDeployment) AddPre(steps ...Step) {
	args := []interface{}{}
	for _, a := range steps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		s,
		"addPre",
		args,
	)
}

