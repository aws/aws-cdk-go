package awsstepfunctions


// Properties for defining a Pass state.
//
// Example:
//   // JSONata Pattern
//   sfn.Pass_Jsonata(this, jsii.String("JSONata Pattern"), &PassJsonataProps{
//   	Outputs: map[string]*string{
//   		"foo": jsii.String("bar"),
//   	},
//   })
//
//   // JSONPath Pattern
//   sfn.Pass_JsonPath(this, jsii.String("JSONPath Pattern"), &PassJsonPathProps{
//   	// The outputs does not exist in the props type
//   	// outputs: { foo: 'bar' },
//   	OutputPath: jsii.String("$.status"),
//   })
//
//   // Constructor (Legacy) Pattern
//   // Constructor (Legacy) Pattern
//   sfn.NewPass(this, jsii.String("Constructor Pattern"), &PassProps{
//   	QueryLanguage: sfn.QueryLanguage_JSONATA,
//   	 // or JSON_PATH
//   	// Both outputs and outputPath exist as prop types.
//   	Outputs: map[string]*string{
//   		"foo": jsii.String("bar"),
//   	},
//   	 // For JSONata
//   	// or
//   	OutputPath: jsii.String("$.status"),
//   })
//
type PassProps struct {
	// A comment describing this state.
	// Default: No comment.
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The name of the query language used by the state.
	//
	// If the state does not contain a `queryLanguage` field,
	// then it will use the query language specified in the top-level `queryLanguage` field.
	// Default: - JSONPath.
	//
	QueryLanguage QueryLanguage `field:"optional" json:"queryLanguage" yaml:"queryLanguage"`
	// Optional name for this state.
	// Default: - The construct ID will be used as state name.
	//
	StateName *string `field:"optional" json:"stateName" yaml:"stateName"`
	// Workflow variables to store in this step.
	//
	// Using workflow variables, you can store data in a step and retrieve that data in future steps.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/workflow-variables.html
	//
	// Default: - Not assign variables.
	//
	Assign *map[string]interface{} `field:"optional" json:"assign" yaml:"assign"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Default: $.
	//
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// JSONPath expression to select part of the state to be the output to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Default: $.
	//
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
	// Used to specify and transform output from the state.
	//
	// When specified, the value overrides the state output default.
	// The output field accepts any JSON value (object, array, string, number, boolean, null).
	// Any string value, including those inside objects or arrays,
	// will be evaluated as JSONata if surrounded by {% %} characters.
	// Output also accepts a JSONata expression directly.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-input-output-filtering.html
	//
	// Default: - $states.result or $states.errorOutput
	//
	Outputs interface{} `field:"optional" json:"outputs" yaml:"outputs"`
	// Parameters pass a collection of key-value pairs, either static values or JSONPath expressions that select from the input.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-parameters
	//
	// Default: No parameters.
	//
	Parameters *map[string]interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// If given, treat as the result of this operation.
	//
	// Can be used to inject or replace the current execution state.
	// Default: No injected result.
	//
	Result Result `field:"optional" json:"result" yaml:"result"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Default: $.
	//
	ResultPath *string `field:"optional" json:"resultPath" yaml:"resultPath"`
}

