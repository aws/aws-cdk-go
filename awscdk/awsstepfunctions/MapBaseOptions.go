package awsstepfunctions


// Base properties for defining a Map state.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var assign interface{}
//   var itemSelector interface{}
//
//   mapBaseOptions := &MapBaseOptions{
//   	Assign: map[string]interface{}{
//   		"assignKey": assign,
//   	},
//   	ItemSelector: map[string]interface{}{
//   		"itemSelectorKey": itemSelector,
//   	},
//   	JsonataItemSelector: jsii.String("jsonataItemSelector"),
//   	JsonataMaxConcurrency: jsii.String("jsonataMaxConcurrency"),
//   	MaxConcurrency: jsii.Number(123),
//   }
//
type MapBaseOptions struct {
	// Workflow variables to store in this step.
	//
	// Using workflow variables, you can store data in a step and retrieve that data in future steps.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/workflow-variables.html
	//
	// Default: - Not assign variables.
	//
	Assign *map[string]interface{} `field:"optional" json:"assign" yaml:"assign"`
	// The JSON that you want to override your default iteration input (mutually exclusive  with `parameters` and `jsonataItemSelector`).
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-itemselector.html
	//
	// Default: $.
	//
	ItemSelector *map[string]interface{} `field:"optional" json:"itemSelector" yaml:"itemSelector"`
	// Jsonata expression that evaluates to a JSON array to override your default iteration input (mutually exclusive with `parameters` and `itemSelector`).
	//
	// Example value: `{% {\"foo\": \"foo\", \"input\": $states.input} %}`
	// Default: $.
	//
	JsonataItemSelector *string `field:"optional" json:"jsonataItemSelector" yaml:"jsonataItemSelector"`
	// JSONata expression for MaxConcurrency.
	//
	// A JSONata expression that evaluates to an integer, specifying the maximum
	// concurrency dynamically. Mutually exclusive with `maxConcurrency` and
	// `maxConcurrencyPath`.
	//
	// Example value: `{% $states.input.maxConcurrency %}`
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-asl-use-map-state-inline.html#map-state-inline-additional-fields
	//
	// Default: - full concurrency.
	//
	JsonataMaxConcurrency *string `field:"optional" json:"jsonataMaxConcurrency" yaml:"jsonataMaxConcurrency"`
	// MaxConcurrency.
	//
	// An upper bound on the number of iterations you want running at once.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-asl-use-map-state-inline.html#map-state-inline-additional-fields
	//
	// Default: - full concurrency.
	//
	MaxConcurrency *float64 `field:"optional" json:"maxConcurrency" yaml:"maxConcurrency"`
}

