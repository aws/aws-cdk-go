package awsstepfunctions


// Properties for defining a Pass state.
//
// Example:
//   // Makes the current JSON state { ..., "subObject": { "hello": "world" } }
//   pass := sfn.NewPass(this, jsii.String("Add Hello World"), &passProps{
//   	result: sfn.result.fromObject(map[string]interface{}{
//   		"hello": jsii.String("world"),
//   	}),
//   	resultPath: jsii.String("$.subObject"),
//   })
//
//   // Set the next state
//   nextState := sfn.NewPass(this, jsii.String("NextState"))
//   pass.next(nextState)
//
type PassProps struct {
	// An optional description for this state.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// JSONPath expression to select part of the state to be the output to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
	// Parameters pass a collection of key-value pairs, either static values or JSONPath expressions that select from the input.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-parameters
	//
	Parameters *map[string]interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// If given, treat as the result of this operation.
	//
	// Can be used to inject or replace the current execution state.
	Result Result `field:"optional" json:"result" yaml:"result"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	ResultPath *string `field:"optional" json:"resultPath" yaml:"resultPath"`
}

