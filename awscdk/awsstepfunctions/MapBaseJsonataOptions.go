package awsstepfunctions


// Base properties for defining a Map state that using JSONata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var outputs interface{}
//   var provideItems provideItems
//
//   mapBaseJsonataOptions := &MapBaseJsonataOptions{
//   	Items: provideItems,
//   	Outputs: outputs,
//   }
//
type MapBaseJsonataOptions struct {
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
	// The array that the Map state will iterate over.
	// Default: - The state input as is.
	//
	Items ProvideItems `field:"optional" json:"items" yaml:"items"`
}

