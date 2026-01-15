package awsstepfunctions


// Properties for defining a Map state that using JSONPath.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var assign interface{}
//   var itemSelector interface{}
//   var parameters interface{}
//   var resultSelector interface{}
//
//   mapJsonPathProps := &MapJsonPathProps{
//   	Assign: map[string]interface{}{
//   		"assignKey": assign,
//   	},
//   	Comment: jsii.String("comment"),
//   	InputPath: jsii.String("inputPath"),
//   	ItemSelector: map[string]interface{}{
//   		"itemSelectorKey": itemSelector,
//   	},
//   	ItemsPath: jsii.String("itemsPath"),
//   	JsonataItemSelector: jsii.String("jsonataItemSelector"),
//   	JsonataMaxConcurrency: jsii.String("jsonataMaxConcurrency"),
//   	MaxConcurrency: jsii.Number(123),
//   	MaxConcurrencyPath: jsii.String("maxConcurrencyPath"),
//   	OutputPath: jsii.String("outputPath"),
//   	Parameters: map[string]interface{}{
//   		"parametersKey": parameters,
//   	},
//   	QueryLanguage: awscdk.Aws_stepfunctions.QueryLanguage_JSON_PATH,
//   	ResultPath: jsii.String("resultPath"),
//   	ResultSelector: map[string]interface{}{
//   		"resultSelectorKey": resultSelector,
//   	},
//   	StateName: jsii.String("stateName"),
//   }
//
type MapJsonPathProps struct {
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
	// JSONPath expression to select the array to iterate over.
	// Default: $.
	//
	ItemsPath *string `field:"optional" json:"itemsPath" yaml:"itemsPath"`
	// MaxConcurrencyPath.
	//
	// A JsonPath that specifies the maximum concurrency dynamically from the state input.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-asl-use-map-state-inline.html#map-state-inline-additional-fields
	//
	// Default: - full concurrency.
	//
	MaxConcurrencyPath *string `field:"optional" json:"maxConcurrencyPath" yaml:"maxConcurrencyPath"`
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
	// The JSON that you want to override your default iteration input (mutually exclusive  with `itemSelector`).
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-itemselector.html
	//
	// Default: $.
	//
	// Deprecated: Step Functions has deprecated the `parameters` field in favor of
	// the new `itemSelector` field.
	Parameters *map[string]interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

