package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Helper functions to work with structures containing fields.
type FieldUtils interface {
}

// The jsii proxy struct for FieldUtils
type jsiiProxy_FieldUtils struct {
	_ byte // padding
}

// Returns whether the given task structure contains the TaskToken field anywhere.
//
// The field is considered included if the field itself or one of its containing
// fields occurs anywhere in the payload.
func FieldUtils_ContainsTaskToken(obj *map[string]interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.FieldUtils",
		"containsTaskToken",
		[]interface{}{obj},
		&returns,
	)

	return returns
}

// Return all JSON paths used in the given structure.
func FieldUtils_FindReferencedPaths(obj *map[string]interface{}) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.FieldUtils",
		"findReferencedPaths",
		[]interface{}{obj},
		&returns,
	)

	return returns
}

// Render a JSON structure containing fields to the right StepFunctions structure.
func FieldUtils_RenderObject(obj *map[string]interface{}) *map[string]interface{} {
	_init_.Initialize()

	var returns *map[string]interface{}

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.FieldUtils",
		"renderObject",
		[]interface{}{obj},
		&returns,
	)

	return returns
}

