package awscdkpipesalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkpipesalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2/internal"
)

// Dynamic variables that can be used in the input transformation.
//
// Example:
//   var sourceQueue Queue
//   var targetQueue Queue
//
//
//   targetInputTransformation := pipes.inputTransformation_FromObject(map[string]interface{}{
//   	"staticField": jsii.String("static value"),
//   	"dynamicField": pipes.DynamicInput_fromEventPath(jsii.String("$.body.payload")),
//   	"pipeVariable": pipes.DynamicInput_pipeName(),
//   })
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	PipeName: jsii.String("MyPipe"),
//   	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
//   	Target: awscdkpipestargetsalpha.NewSqsTarget(targetQueue, &SqsTargetParameters{
//   		InputTransformation: targetInputTransformation,
//   	}),
//   })
//
// Experimental.
type DynamicInput interface {
	awscdk.IResolvable
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// This may return an array with a single informational element indicating how
	// to get this property populated, if it was skipped for performance reasons.
	// Experimental.
	CreationStack() *[]*string
	// Human readable display hint about the event pattern.
	// Experimental.
	DisplayHint() *string
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context awscdk.IResolveContext) interface{}
	// Return a JSON representation of a dynamic input.
	// Experimental.
	ToJSON() *string
	// Return a string representation of a dynamic input.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DynamicInput
type jsiiProxy_DynamicInput struct {
	internal.Type__awscdkIResolvable
}

func (j *jsiiProxy_DynamicInput) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicInput) DisplayHint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayHint",
		&returns,
	)
	return returns
}


// Value from the event payload at jsonPath.
// Experimental.
func DynamicInput_FromEventPath(path *string) DynamicInput {
	_init_.Initialize()

	if err := validateDynamicInput_FromEventPathParameters(path); err != nil {
		panic(err)
	}
	var returns DynamicInput

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-pipes-alpha.DynamicInput",
		"fromEventPath",
		[]interface{}{path},
		&returns,
	)

	return returns
}

func DynamicInput_EnrichmentArn() DynamicInput {
	_init_.Initialize()
	var returns DynamicInput
	_jsii_.StaticGet(
		"@aws-cdk/aws-pipes-alpha.DynamicInput",
		"enrichmentArn",
		&returns,
	)
	return returns
}

func DynamicInput_Event() DynamicInput {
	_init_.Initialize()
	var returns DynamicInput
	_jsii_.StaticGet(
		"@aws-cdk/aws-pipes-alpha.DynamicInput",
		"event",
		&returns,
	)
	return returns
}

func DynamicInput_EventIngestionTime() DynamicInput {
	_init_.Initialize()
	var returns DynamicInput
	_jsii_.StaticGet(
		"@aws-cdk/aws-pipes-alpha.DynamicInput",
		"eventIngestionTime",
		&returns,
	)
	return returns
}

func DynamicInput_EventJson() DynamicInput {
	_init_.Initialize()
	var returns DynamicInput
	_jsii_.StaticGet(
		"@aws-cdk/aws-pipes-alpha.DynamicInput",
		"eventJson",
		&returns,
	)
	return returns
}

func DynamicInput_PipeArn() DynamicInput {
	_init_.Initialize()
	var returns DynamicInput
	_jsii_.StaticGet(
		"@aws-cdk/aws-pipes-alpha.DynamicInput",
		"pipeArn",
		&returns,
	)
	return returns
}

func DynamicInput_PipeName() DynamicInput {
	_init_.Initialize()
	var returns DynamicInput
	_jsii_.StaticGet(
		"@aws-cdk/aws-pipes-alpha.DynamicInput",
		"pipeName",
		&returns,
	)
	return returns
}

func DynamicInput_SourceArn() DynamicInput {
	_init_.Initialize()
	var returns DynamicInput
	_jsii_.StaticGet(
		"@aws-cdk/aws-pipes-alpha.DynamicInput",
		"sourceArn",
		&returns,
	)
	return returns
}

func DynamicInput_TargetArn() DynamicInput {
	_init_.Initialize()
	var returns DynamicInput
	_jsii_.StaticGet(
		"@aws-cdk/aws-pipes-alpha.DynamicInput",
		"targetArn",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DynamicInput) Resolve(context awscdk.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamicInput) ToJSON() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamicInput) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

