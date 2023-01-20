package awsstepfunctions


// The type of task input.
type InputType string

const (
	// Use a literal string This might be a JSON-encoded object, or just text.
	//
	// valid JSON text: standalone, quote-delimited strings; objects; arrays; numbers; Boolean values; and null.
	//
	// example: `literal string`
	// example: {"json": "encoded"}.
	InputType_TEXT InputType = "TEXT"
	// Use an object which may contain Data and Context fields as object values, if desired.
	//
	// example:
	// {
	//   literal: 'literal',
	//   SomeInput: sfn.JsonPath.stringAt('$.someField')
	// }.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-contextobject.html
	//
	InputType_OBJECT InputType = "OBJECT"
)

