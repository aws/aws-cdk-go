package awsstepfunctions


// Properties for defining a Wait state that using JSONata.
//
// Example:
//   sfn.Wait_Jsonata(this, jsii.String("Wait"), &WaitJsonataProps{
//   	Time: sfn.WaitTime_Timestamp(jsii.String("{% $timestamp %}")),
//   	Outputs: map[string]interface{}{
//   		"stringArgument": jsii.String("inital-task"),
//   		"numberArgument": jsii.Number(123),
//   		"booleanArgument": jsii.Boolean(true),
//   		"arrayArgument": []interface{}{
//   			jsii.Number(1),
//   			jsii.String("{% $number %}"),
//   			jsii.Number(3),
//   		},
//   		"intrinsicFunctionsArgument": jsii.String("{% $join($each($obj, function($v) { $v }), ', ') %}"),
//   	},
//   })
//
type WaitJsonataProps struct {
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
	// See: https://docs.aws.amazon.com/ja_jp/step-functions/latest/dg/workflow-variables.html
	//
	// Default: - Not assign variables.
	//
	Assign *map[string]interface{} `field:"optional" json:"assign" yaml:"assign"`
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
	// Wait duration.
	Time WaitTime `field:"required" json:"time" yaml:"time"`
}

