package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A collection of validation results.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var validationResult validationResult
//
//   validationResults := cdk.NewValidationResults([]validationResult{
//   	validationResult,
//   })
//
type ValidationResults interface {
	IsSuccess() *bool
	Results() *[]ValidationResult
	SetResults(val *[]ValidationResult)
	Collect(result ValidationResult)
	ErrorTreeList() *string
	// Wrap up all validation results into a single tree node.
	//
	// If there are failures in the collection, add a message, otherwise
	// return a success.
	Wrap(message *string) ValidationResult
}

// The jsii proxy struct for ValidationResults
type jsiiProxy_ValidationResults struct {
	_ byte // padding
}

func (j *jsiiProxy_ValidationResults) IsSuccess() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSuccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidationResults) Results() *[]ValidationResult {
	var returns *[]ValidationResult
	_jsii_.Get(
		j,
		"results",
		&returns,
	)
	return returns
}


func NewValidationResults(results *[]ValidationResult) ValidationResults {
	_init_.Initialize()

	j := jsiiProxy_ValidationResults{}

	_jsii_.Create(
		"aws-cdk-lib.ValidationResults",
		[]interface{}{results},
		&j,
	)

	return &j
}

func NewValidationResults_Override(v ValidationResults, results *[]ValidationResult) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.ValidationResults",
		[]interface{}{results},
		v,
	)
}

func (j *jsiiProxy_ValidationResults)SetResults(val *[]ValidationResult) {
	if err := j.validateSetResultsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"results",
		val,
	)
}

func (v *jsiiProxy_ValidationResults) Collect(result ValidationResult) {
	if err := v.validateCollectParameters(result); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"collect",
		[]interface{}{result},
	)
}

func (v *jsiiProxy_ValidationResults) ErrorTreeList() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"errorTreeList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidationResults) Wrap(message *string) ValidationResult {
	if err := v.validateWrapParameters(message); err != nil {
		panic(err)
	}
	var returns ValidationResult

	_jsii_.Invoke(
		v,
		"wrap",
		[]interface{}{message},
		&returns,
	)

	return returns
}

