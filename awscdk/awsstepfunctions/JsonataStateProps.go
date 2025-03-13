package awsstepfunctions


// Properties shared by all states that use JSONata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var arguments_ interface{}
//   var assign interface{}
//   var outputs interface{}
//
//   jsonataStateProps := &JsonataStateProps{
//   	Arguments: arguments_,
//   	Assign: map[string]interface{}{
//   		"assignKey": assign,
//   	},
//   	Comment: jsii.String("comment"),
//   	Outputs: outputs,
//   	QueryLanguage: awscdk.Aws_stepfunctions.QueryLanguage_JSON_PATH,
//   	StateName: jsii.String("stateName"),
//   }
//
type JsonataStateProps struct {
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
	// Default: - No arguments.
	//
	Arguments interface{} `field:"optional" json:"arguments" yaml:"arguments"`
	// Workflow variables to store in this step.
	//
	// Using workflow variables, you can store data in a step and retrieve that data in future steps.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/workflow-variables.html
	//
	// Default: - Not assign variables.
	//
	Assign *map[string]interface{} `field:"optional" json:"assign" yaml:"assign"`
}

