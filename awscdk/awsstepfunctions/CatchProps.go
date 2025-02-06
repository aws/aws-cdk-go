package awsstepfunctions


// Error handler details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var assign interface{}
//   var outputs interface{}
//
//   catchProps := &CatchProps{
//   	Assign: map[string]interface{}{
//   		"assignKey": assign,
//   	},
//   	Errors: []*string{
//   		jsii.String("errors"),
//   	},
//   	Outputs: outputs,
//   	ResultPath: jsii.String("resultPath"),
//   }
//
type CatchProps struct {
	// Workflow variables to store in this step.
	//
	// Using workflow variables, you can store data in a step and retrieve that data in future steps.
	// See: https://docs.aws.amazon.com/ja_jp/step-functions/latest/dg/workflow-variables.html
	//
	// Default: - Not assign variables.
	//
	Assign *map[string]interface{} `field:"optional" json:"assign" yaml:"assign"`
	// Errors to recover from by going to the given state.
	//
	// A list of error strings to retry, which can be either predefined errors
	// (for example Errors.NoChoiceMatched) or a self-defined error.
	// Default: All errors.
	//
	Errors *[]*string `field:"optional" json:"errors" yaml:"errors"`
	// This option for JSONata only.
	//
	// When you use JSONPath, then the state ignores this property.
	// Used to specify and transform output from the state.
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
	// JSONPath expression to indicate where to inject the error data.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the error
	// data to be discarded.
	// Default: $.
	//
	ResultPath *string `field:"optional" json:"resultPath" yaml:"resultPath"`
}

