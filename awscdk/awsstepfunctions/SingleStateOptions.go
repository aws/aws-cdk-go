package awsstepfunctions


// Options for creating a single state.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var arguments_ interface{}
//   var assign interface{}
//   var outputs interface{}
//   var parameters interface{}
//   var resultSelector interface{}
//
//   singleStateOptions := &SingleStateOptions{
//   	Arguments: map[string]interface{}{
//   		"argumentsKey": arguments_,
//   	},
//   	Assign: map[string]interface{}{
//   		"assignKey": assign,
//   	},
//   	Comment: jsii.String("comment"),
//   	InputPath: jsii.String("inputPath"),
//   	OutputPath: jsii.String("outputPath"),
//   	Outputs: outputs,
//   	Parameters: map[string]interface{}{
//   		"parametersKey": parameters,
//   	},
//   	PrefixStates: jsii.String("prefixStates"),
//   	QueryLanguage: awscdk.Aws_stepfunctions.QueryLanguage_JSON_PATH,
//   	ResultPath: jsii.String("resultPath"),
//   	ResultSelector: map[string]interface{}{
//   		"resultSelectorKey": resultSelector,
//   	},
//   	StateId: jsii.String("stateId"),
//   	StateName: jsii.String("stateName"),
//   }
//
type SingleStateOptions struct {
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
	// Parameters pass a collection of key-value pairs, either static values or JSONata expressions that select from the input.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/transforming-data.html
	//
	// Default: No arguments.
	//
	Arguments *map[string]interface{} `field:"optional" json:"arguments" yaml:"arguments"`
	// Parameters pass a collection of key-value pairs, either static values or JSONPath expressions that select from the input.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-parameters
	//
	// Default: No parameters.
	//
	Parameters *map[string]interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Default: $.
	//
	ResultPath *string `field:"optional" json:"resultPath" yaml:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Default: - None.
	//
	ResultSelector *map[string]interface{} `field:"optional" json:"resultSelector" yaml:"resultSelector"`
	// String to prefix all stateIds in the state machine with.
	// Default: stateId.
	//
	PrefixStates *string `field:"optional" json:"prefixStates" yaml:"prefixStates"`
	// ID of newly created containing state.
	// Default: Construct ID of the StateMachineFragment.
	//
	StateId *string `field:"optional" json:"stateId" yaml:"stateId"`
}

