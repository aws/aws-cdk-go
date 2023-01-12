package pipelines

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Multiple stages that are deployed in parallel.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var pipeline codePipeline
//
//   europeWave := pipeline.addWave(jsii.String("Europe"))
//   europeWave.addStage(NewMyApplicationStage(this, jsii.String("Ireland"), &stageProps{
//   	env: &environment{
//   		region: jsii.String("eu-west-1"),
//   	},
//   }))
//   europeWave.addStage(NewMyApplicationStage(this, jsii.String("Germany"), &stageProps{
//   	env: &environment{
//   		region: jsii.String("eu-central-1"),
//   	},
//   }))
//
type Wave interface {
	// Identifier for this Wave.
	Id() *string
	// Additional steps that are run after all of the stages in the wave.
	Post() *[]Step
	// Additional steps that are run before any of the stages in the wave.
	Pre() *[]Step
	// The stages that are deployed in this wave.
	Stages() *[]StageDeployment
	// Add an additional step to run after all of the stages in this wave.
	AddPost(steps ...Step)
	// Add an additional step to run before any of the stages in this wave.
	AddPre(steps ...Step)
	// Add a Stage to this wave.
	//
	// It will be deployed in parallel with all other stages in this
	// wave.
	AddStage(stage awscdk.Stage, options *AddStageOpts) StageDeployment
}

// The jsii proxy struct for Wave
type jsiiProxy_Wave struct {
	_ byte // padding
}

func (j *jsiiProxy_Wave) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wave) Post() *[]Step {
	var returns *[]Step
	_jsii_.Get(
		j,
		"post",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wave) Pre() *[]Step {
	var returns *[]Step
	_jsii_.Get(
		j,
		"pre",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wave) Stages() *[]StageDeployment {
	var returns *[]StageDeployment
	_jsii_.Get(
		j,
		"stages",
		&returns,
	)
	return returns
}


func NewWave(id *string, props *WaveProps) Wave {
	_init_.Initialize()

	if err := validateNewWaveParameters(id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wave{}

	_jsii_.Create(
		"aws-cdk-lib.pipelines.Wave",
		[]interface{}{id, props},
		&j,
	)

	return &j
}

func NewWave_Override(w Wave, id *string, props *WaveProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.pipelines.Wave",
		[]interface{}{id, props},
		w,
	)
}

func (w *jsiiProxy_Wave) AddPost(steps ...Step) {
	args := []interface{}{}
	for _, a := range steps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		w,
		"addPost",
		args,
	)
}

func (w *jsiiProxy_Wave) AddPre(steps ...Step) {
	args := []interface{}{}
	for _, a := range steps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		w,
		"addPre",
		args,
	)
}

func (w *jsiiProxy_Wave) AddStage(stage awscdk.Stage, options *AddStageOpts) StageDeployment {
	if err := w.validateAddStageParameters(stage, options); err != nil {
		panic(err)
	}
	var returns StageDeployment

	_jsii_.Invoke(
		w,
		"addStage",
		[]interface{}{stage, options},
		&returns,
	)

	return returns
}

