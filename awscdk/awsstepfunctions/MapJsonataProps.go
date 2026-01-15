package awsstepfunctions


// Properties for defining a Map state that using JSONata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var assign interface{}
//   var itemSelector interface{}
//   var outputs interface{}
//   var parameters interface{}
//   var provideItems ProvideItems
//
//   mapJsonataProps := &MapJsonataProps{
//   	Assign: map[string]interface{}{
//   		"assignKey": assign,
//   	},
//   	Comment: jsii.String("comment"),
//   	Items: provideItems,
//   	ItemSelector: map[string]interface{}{
//   		"itemSelectorKey": itemSelector,
//   	},
//   	JsonataItemSelector: jsii.String("jsonataItemSelector"),
//   	JsonataMaxConcurrency: jsii.String("jsonataMaxConcurrency"),
//   	MaxConcurrency: jsii.Number(123),
//   	Outputs: outputs,
//   	Parameters: map[string]interface{}{
//   		"parametersKey": parameters,
//   	},
//   	QueryLanguage: awscdk.Aws_stepfunctions.QueryLanguage_JSON_PATH,
//   	StateName: jsii.String("stateName"),
//   }
//
type MapJsonataProps struct {
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
	// The JSON that you want to override your default iteration input (mutually exclusive  with `itemSelector`).
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-itemselector.html
	//
	// Default: $.
	//
	// Deprecated: Step Functions has deprecated the `parameters` field in favor of
	// the new `itemSelector` field.
	Parameters *map[string]interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

