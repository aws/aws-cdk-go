package pipelines

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A set of files traveling through the deployment pipeline.
//
// Individual steps in the pipeline produce or consume
// `FileSet`s.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   type myJenkinsStep struct {
//   	step
//   }
//
//   func newMyJenkinsStep(provider jenkinsProvider, input fileSet) *myJenkinsStep {
//   	this := &myJenkinsStep{}
//   	pipelines.NewStep_Override(this, jsii.String("MyJenkinsStep"))
//
//   	// This is necessary if your step accepts parameters, like environment variables,
//   	// that may contain outputs from other steps. It doesn't matter what the
//   	// structure is, as long as it contains the values that may contain outputs.
//   	this.DiscoverReferencedOutputs(map[string]map[string]interface{}{
//   		"env": map[string]interface{}{
//   		},
//   	})
//   	return this
//   }
//
//   func (this *myJenkinsStep) produceAction(stage iStage, options produceActionOptions) codePipelineActionFactoryResult {
//
//   	// This is where you control what type of Action gets added to the
//   	// CodePipeline
//   	*stage.AddAction(cpactions.NewJenkinsAction(&JenkinsActionProps{
//   		// Copy 'actionName' and 'runOrder' from the options
//   		ActionName: options.ActionName,
//   		RunOrder: options.RunOrder,
//
//   		// Jenkins-specific configuration
//   		Type: cpactions.JenkinsActionType_TEST,
//   		JenkinsProvider: this.provider,
//   		ProjectName: jsii.String("MyJenkinsProject"),
//
//   		// Translate the FileSet into a codepipeline.Artifact
//   		Inputs: []artifact{
//   			options.Artifacts.ToCodePipeline(this.input),
//   		},
//   	}))
//
//   	return &codePipelineActionFactoryResult{
//   		RunOrdersConsumed: jsii.Number(1),
//   	}
//   }
//
type FileSet interface {
	IFileSetProducer
	// Human-readable descriptor for this file set (does not need to be unique).
	Id() *string
	// The primary output of a file set producer.
	//
	// The primary output of a FileSet is itself.
	PrimaryOutput() FileSet
	// The Step that produces this FileSet.
	Producer() Step
	// Mark the given Step as the producer for this FileSet.
	//
	// This method can only be called once.
	ProducedBy(producer Step)
	// Return a string representation of this FileSet.
	ToString() *string
}

// The jsii proxy struct for FileSet
type jsiiProxy_FileSet struct {
	jsiiProxy_IFileSetProducer
}

func (j *jsiiProxy_FileSet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSet) PrimaryOutput() FileSet {
	var returns FileSet
	_jsii_.Get(
		j,
		"primaryOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSet) Producer() Step {
	var returns Step
	_jsii_.Get(
		j,
		"producer",
		&returns,
	)
	return returns
}


func NewFileSet(id *string, producer Step) FileSet {
	_init_.Initialize()

	if err := validateNewFileSetParameters(id); err != nil {
		panic(err)
	}
	j := jsiiProxy_FileSet{}

	_jsii_.Create(
		"aws-cdk-lib.pipelines.FileSet",
		[]interface{}{id, producer},
		&j,
	)

	return &j
}

func NewFileSet_Override(f FileSet, id *string, producer Step) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.pipelines.FileSet",
		[]interface{}{id, producer},
		f,
	)
}

func (f *jsiiProxy_FileSet) ProducedBy(producer Step) {
	_jsii_.InvokeVoid(
		f,
		"producedBy",
		[]interface{}{producer},
	)
}

func (f *jsiiProxy_FileSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

