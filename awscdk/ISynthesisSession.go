package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/cxapi"
)

// Represents a single session of synthesis.
//
// Passed into `Construct.synthesize()` methods.
type ISynthesisSession interface {
	// Cloud assembly builder.
	Assembly() cxapi.CloudAssemblyBuilder
	SetAssembly(a cxapi.CloudAssemblyBuilder)
	// The output directory for this synthesis session.
	Outdir() *string
	SetOutdir(o *string)
	// Whether the stack should be validated after synthesis to check for error metadata.
	// Default: - false.
	//
	ValidateOnSynth() *bool
	SetValidateOnSynth(v *bool)
}

// The jsii proxy for ISynthesisSession
type jsiiProxy_ISynthesisSession struct {
	_ byte // padding
}

func (j *jsiiProxy_ISynthesisSession) Assembly() cxapi.CloudAssemblyBuilder {
	var returns cxapi.CloudAssemblyBuilder
	_jsii_.Get(
		j,
		"assembly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISynthesisSession)SetAssembly(val cxapi.CloudAssemblyBuilder) {
	if err := j.validateSetAssemblyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assembly",
		val,
	)
}

func (j *jsiiProxy_ISynthesisSession) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISynthesisSession)SetOutdir(val *string) {
	if err := j.validateSetOutdirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outdir",
		val,
	)
}

func (j *jsiiProxy_ISynthesisSession) ValidateOnSynth() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"validateOnSynth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISynthesisSession)SetValidateOnSynth(val *bool) {
	_jsii_.Set(
		j,
		"validateOnSynth",
		val,
	)
}

