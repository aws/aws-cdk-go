package awsstepfunctions


// Option properties for JSONata task state.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var arguments_ interface{}
//   var outputs interface{}
//
//   jsonataStateOptions := &JsonataStateOptions{
//   	Arguments: arguments_,
//   	Outputs: outputs,
//   }
//
type JsonataStateOptions struct {
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
}

